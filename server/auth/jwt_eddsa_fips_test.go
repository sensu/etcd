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

//go:build fips140
// +build fips140

package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// The fips140 build must reject EdDSA JWT signing methods outright.
func TestJWTEdDSARejectedInFIPS(t *testing.T) {
	for k, opts := range map[string]map[string]string{
		"Ed25519-priv": {"priv-key": jwtEdPrivKey, "sign-method": "EdDSA"},
		"Ed25519":      {"pub-key": jwtEdPubKey, "priv-key": jwtEdPrivKey, "sign-method": "EdDSA"},
	} {
		t.Run(k, func(tt *testing.T) {
			_, err := newTokenProviderJWT(zap.NewNop(), opts)
			require.Error(tt, err, "expected EdDSA to be rejected in fips140 build")
		})
	}
}
