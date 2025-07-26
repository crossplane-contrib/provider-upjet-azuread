#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(turkenf): We are getting "Resource '5f5fdd0e-e056-4c4d-a407-43d80e5c11b5' does not exist or 
# one of its queried reference-property objects are not present." exception if member for
# the group got deleted before the member resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete member.groups.azuread.upbound.io -l testing.upbound.io/example-name=example
