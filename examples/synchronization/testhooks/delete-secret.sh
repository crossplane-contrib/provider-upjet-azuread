#!/usr/bin/env bash
set -aeuo pipefail

# Note(turkenf): We are getting "Retrieving synchronization secrets for service principal" exception if secret for
# the principal got deleted before the secret resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete secret.synchronization.azuread.upbound.io -l testing.upbound.io/example-name=example