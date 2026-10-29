// Copyright 2018 The etcd Authors
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

package tlsutil

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/require"
)

// The fips140 build only accepts AEAD (GCM) cipher suites. CBC, ChaCha20
// and RC4 suites must be rejected.
func TestGetCipherSuiteFIPSAllowed(t *testing.T) {
	CipherSuiteExpectedToExist(t, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256)
	CipherSuiteExpectedToExist(t, "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384)
	CipherSuiteExpectedToExist(t, "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256)
}

func TestGetCipherSuiteFIPSRejected(t *testing.T) {
	for _, name := range []string{
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
		"TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	} {
		_, ok := GetCipherSuite(name)
		require.Falsef(t, ok, "expected %s to be rejected in fips140 build", name)
	}
}
