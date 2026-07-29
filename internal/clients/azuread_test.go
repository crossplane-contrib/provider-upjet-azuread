// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_graphServiceFromPath(t *testing.T) {
	cases := map[string]struct {
		path string
		want string
	}{
		"v1_groups_collection": {
			path: "/v1.0/groups",
			want: "groups",
		},
		"v1_groups_item": {
			path: "/v1.0/groups/abc-123",
			want: "groups",
		},
		"not_valid_graph_path": {
			path: "/some",
			want: "unknown",
		},
		"empty_path": {
			path: "",
			want: "unknown",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := graphServiceFromPath(tc.path)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("graphServiceFromPath(%q) mismatch (-want +got):\n%s", tc.path, diff)
			}
		})
	}
}
