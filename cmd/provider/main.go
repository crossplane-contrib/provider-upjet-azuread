// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/certificates"
	xpcontroller "github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/feature"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	"gopkg.in/alecthomas/kingpin.v2"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/upbound/provider-azuread/apis"
	"github.com/upbound/provider-azuread/apis/v1alpha1"
	"github.com/upbound/provider-azuread/config"
	"github.com/upbound/provider-azuread/internal/clients"
	"github.com/upbound/provider-azuread/internal/controller"
	"github.com/upbound/provider-azuread/internal/features"
)

func main() {
	var (
		app               = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Azuread").DefaultEnvars()
		deprecationAction = func(flagName string) kingpin.Action {
			return func(c *kingpin.ParseContext) error {
				_, err := fmt.Fprintf(os.Stderr, "warning: Command-line flag %q is deprecated and no longer used. It will be removed in a future release. Please remove it from all of your configurations (ControllerConfigs, etc.).\n", flagName)
				kingpin.FatalIfError(err, "Failed to print the deprecation notice.")
				return nil
			}
		}
		debug            = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod       = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval     = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		leaderElection   = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("10").Int()
		// now deprecated command-line arguments with the Terraform SDK-based upjet architecture
		_ = app.Flag("terraform-version", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform version.").Envar("TERRAFORM_VERSION").Hidden().Action(deprecationAction("terraform-version")).String()
		_ = app.Flag("terraform-provider-version", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform provider version.").Envar("TERRAFORM_PROVIDER_VERSION").Hidden().Action(deprecationAction("terraform-provider-version")).String()
		_ = app.Flag("terraform-native-provider-path", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform native provider path for shared execution.").Default("").Envar("TERRAFORM_NATIVE_PROVIDER_PATH").Hidden().Action(deprecationAction("terraform-native-provider-path")).String()
		_ = app.Flag("terraform-provider-source", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Terraform provider source.").Envar("TERRAFORM_PROVIDER_SOURCE").Hidden().Action(deprecationAction("terraform-provider-source")).String()
		_ = app.Flag("provider-ttl", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] TTL for the native plugin processes before they are replaced. Changing the default may increase memory consumption.").Default("100").Hidden().Action(deprecationAction("provider-ttl")).Int()

		namespace                  = app.Flag("namespace", "Namespace used to set as default scope in default secret store config.").Default("crossplane-system").Envar("POD_NAMESPACE").String()
		enableExternalSecretStores = app.Flag("enable-external-secret-stores", "Enable support for ExternalSecretStores.").Default("false").Envar("ENABLE_EXTERNAL_SECRET_STORES").Bool()
		essTLSCertsPath            = app.Flag("ess-tls-cert-dir", "Path of ESS TLS certificates.").Envar("ESS_TLS_CERTS_DIR").String()
		enableManagementPolicies   = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-azuread"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	// currently, we configure the jitter to be the 5% of the poll interval
	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	log.Debug("Starting", "sync-period", syncPeriod.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-azuread",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add Azuread APIs to scheme")

	ctx := context.Background()
	provider, err := config.GetProvider(ctx, false)
	kingpin.FatalIfError(err, "Cannot initialize the provider configuration")
	o := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
		},
		Provider: provider,
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		SetupFn:               clients.TerraformSetupBuilder(provider.TerraformProvider),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}

	if *enableManagementPolicies {
		o.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableExternalSecretStores {
		o.SecretStoreConfigGVK = &v1alpha1.StoreConfigGroupVersionKind
		log.Info("Alpha feature enabled", "flag", features.EnableAlphaExternalSecretStores)

		o.ESSOptions = &tjcontroller.ESSOptions{}
		if *essTLSCertsPath != "" {
			log.Info("ESS TLS certificates path is set. Loading mTLS configuration.")
			tCfg, err := certificates.LoadMTLSConfig(filepath.Join(*essTLSCertsPath, "ca.crt"), filepath.Join(*essTLSCertsPath, "tls.crt"), filepath.Join(*essTLSCertsPath, "tls.key"), false)
			kingpin.FatalIfError(err, "Cannot load ESS TLS config.")

			o.ESSOptions.TLSConfig = tCfg
		}

		// Ensure default store config exists.
		kingpin.FatalIfError(resource.Ignore(kerrors.IsAlreadyExists, mgr.GetClient().Create(ctx, &v1alpha1.StoreConfig{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
			Spec: v1alpha1.StoreConfigSpec{
				// NOTE(turkenh): We only set required spec and expect optional
				// ones to properly be initialized with CRD level default values.
				SecretStoreConfig: xpv1.SecretStoreConfig{
					DefaultScope: *namespace,
				},
			},
		})), "cannot create default store config")
	}

	kingpin.FatalIfError(controller.Setup(mgr, o), "Cannot setup Azuread controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
