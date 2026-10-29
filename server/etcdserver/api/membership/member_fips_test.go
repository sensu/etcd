// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build fips140
// +build fips140

package membership

import (
	"net/url"
	"testing"

	"go.etcd.io/etcd/client/pkg/v3/types"
)

// TestMemberTimeFIPS checks member IDs derived with SHA-256 in the fips140 build.
// Expected values are computed independently of the implementation:
// first 8 bytes of SHA-256 over sorted peer URLs + cluster name + optional Unix time.
func TestMemberTimeFIPS(t *testing.T) {
	tests := []struct {
		mem *Member
		id  types.ID
	}{
		{NewMember("mem1", []url.URL{{Scheme: "http", Host: "10.0.0.8:2379"}}, "", nil), 4242955999852002778},
		// Same ID, different name (names shouldn't matter)
		{NewMember("memfoo", []url.URL{{Scheme: "http", Host: "10.0.0.8:2379"}}, "", nil), 4242955999852002778},
		// Same ID, different Time
		{NewMember("mem1", []url.URL{{Scheme: "http", Host: "10.0.0.8:2379"}}, "", timeParse("1984-12-23T15:04:05Z")), 6400357517120415902},
		// Different cluster name
		{NewMember("mcm1", []url.URL{{Scheme: "http", Host: "10.0.0.8:2379"}}, "etcd", timeParse("1984-12-23T15:04:05Z")), 14931056209725771900},
		{NewMember("mem1", []url.URL{{Scheme: "http", Host: "10.0.0.1:2379"}}, "", timeParse("1984-12-23T15:04:05Z")), 15614570837182403267},
		// Order shouldn't matter
		{NewMember("mem1", []url.URL{{Scheme: "http", Host: "10.0.0.1:2379"}, {Scheme: "http", Host: "10.0.0.2:2379"}}, "", nil), 13231417086386635835},
		{NewMember("mem1", []url.URL{{Scheme: "http", Host: "10.0.0.2:2379"}, {Scheme: "http", Host: "10.0.0.1:2379"}}, "", nil), 13231417086386635835},
	}
	for i, tt := range tests {
		if tt.mem.ID != tt.id {
			t.Errorf("#%d: mem.ID = %v, want %v", i, tt.mem.ID, tt.id)
		}
	}
}
