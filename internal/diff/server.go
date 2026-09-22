// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Package diff implements the provider's diff gRPC services.
package diff

import (
	"context"
	"net"
	"os"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	diffv1alpha1 "github.com/crossplane/upjet/v2/proto/diff/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	errListen = "cannot listen on %s address %q"
	errServe  = "cannot serve the diff gRPC services"
)

// PlanService implements the upjet.diff.v1alpha1.PlanService gRPC service.
type PlanService struct {
	diffv1alpha1.UnimplementedPlanServiceServer

	log logging.Logger
}

// NewPlanService returns a new PlanService.
func NewPlanService(log logging.Logger) *PlanService {
	return &PlanService{log: log}
}

// Plan computes a diff between the desired and the live resources supplied in
// the request.
func (s *PlanService) Plan(_ context.Context, _ *diffv1alpha1.PlanRequest) (*diffv1alpha1.PlanResponse, error) {
	return &diffv1alpha1.PlanResponse{}, nil
}

// Serve starts a gRPC server serving the diff services on the given network
// and address, and blocks until ctx is done or the server fails. A "unix"
// network address is removed before it's bound so that a socket left behind by
// a previous run does not prevent the server from starting.
func Serve(ctx context.Context, network, address string, log logging.Logger) error {
	if network == "unix" {
		if err := os.Remove(address); err != nil && !os.IsNotExist(err) {
			return errors.Wrapf(err, "cannot remove the existing socket at %q", address)
		}
	}

	var lc net.ListenConfig
	l, err := lc.Listen(ctx, network, address)
	if err != nil {
		return errors.Wrapf(err, errListen, network, address)
	}

	s := grpc.NewServer()
	diffv1alpha1.RegisterPlanServiceServer(s, NewPlanService(log))
	// Reflection lets clients such as grpcurl discover the served services.
	reflection.Register(s)

	// GracefulStop closes the listener and waits for the in-flight RPCs to
	// complete, which also unblocks the Serve call below.
	go func() {
		<-ctx.Done()
		log.Info("Stopping the diff gRPC server")
		s.GracefulStop()
	}()

	log.Info("Starting the diff gRPC server", "network", network, "address", address)
	return errors.Wrap(s.Serve(l), errServe)
}
