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

package membership

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.etcd.io/etcd/client/pkg/v3/types"
)

// TestUpdateRaftAttributesKeepsLearnerFlag checks that updating a learner's
// peer URLs does not clear IsLearner, since MemberUpdateRequest only
// supports changing peer URLs.
func TestUpdateRaftAttributesKeepsLearnerFlag(t *testing.T) {
	clientURLs := []string{"http://127.0.0.1:2379"}
	oldPeerURLs := []string{"http://127.0.0.1:2380"}
	newPeerURLs := []string{"http://127.0.0.1:2382"}

	c := newTestCluster(t, []*Member{
		newTestMemberAsLearner(2, oldPeerURLs, "2", clientURLs),
	})

	c.UpdateRaftAttributes(types.ID(2), RaftAttributes{PeerURLs: newPeerURLs, IsLearner: false}, false)

	m := c.Member(types.ID(2))
	require.NotNil(t, m)
	require.Equal(t, newPeerURLs, m.PeerURLs)
	require.True(t, m.IsLearner, "IsLearner must not be changed by UpdateRaftAttributes")
}
