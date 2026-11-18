// Copyright 2026 The etcd Authors
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

package command

import (
	"bytes"
	"fmt"
	"testing"

	"go.etcd.io/etcd/api/v3/version"
)

func TestVersionCommandIncludesGitSHA(t *testing.T) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := NewVersionCommand()
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)

	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	want := fmt.Sprintf("etcdctl version: %s\nAPI version: %s\nGit SHA: %s\n", version.Version, version.APIVersion, version.GitSHA)
	if got := stdout.String(); got != want {
		t.Fatalf("stdout mismatch:\ngot:  %q\nwant: %q", got, want)
	}
	if got := stderr.String(); got != "" {
		t.Fatalf("unexpected stderr: %q", got)
	}
}
