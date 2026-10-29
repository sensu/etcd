// Copyright 2016 The etcd Authors
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

//go:build !fips140
// +build !fips140

package auth

import (
	"testing"
)

// Ed25519 (EdDSA) JWT signing is only available in the default build.
// The fips140 build restricts JWT signing to RSA, ECDSA and HMAC.
func jwtEdDSATestOptions() map[string]map[string]string {
	return map[string]map[string]string{
		"Ed25519-priv": {
			"priv-key":    jwtEdPrivKey,
			"sign-method": "EdDSA",
		},
		"Ed25519": {
			"pub-key":     jwtEdPubKey,
			"priv-key":    jwtEdPrivKey,
			"sign-method": "EdDSA",
		},
	}
}

func TestJWTInfoEdDSA(t *testing.T) {
	for k, opts := range jwtEdDSATestOptions() {
		t.Run(k, func(tt *testing.T) {
			testJWTInfo(tt, opts)
		})
	}
}
