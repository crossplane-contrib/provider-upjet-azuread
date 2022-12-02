/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"

	"github.com/upbound/provider-azuread/config/applications"
	"github.com/upbound/provider-azuread/config/groups"
	"github.com/upbound/provider-azuread/config/invitations"
)

const (
	resourcePrefix = "azuread"
	modulePath     = "github.com/upbound/provider-azuread"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		invitations.Configure,
		applications.Configure,
		groups.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
