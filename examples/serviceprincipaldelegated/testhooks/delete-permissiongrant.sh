#!/usr/bin/env bash
set -aeuo pipefail

# Note(turkenf): We are getting "Retrieving Delegated Permission Grant" exception if permissiongrant for
# the principal got deleted before the permissiongrant resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete permissiongrant.serviceprincipaldelegated.azuread.upbound.io -l testing.upbound.io/example-name=example