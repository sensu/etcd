//go:build fips140
// +build fips140

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

package tlsutil

import (
	"crypto/tls"
	"fmt"
)

var fipsCipherSuites = map[uint16]struct{}{
	tls.TLS_AES_128_GCM_SHA256:                  {},
	tls.TLS_AES_256_GCM_SHA384:                  {},
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256: {},
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384: {},
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:   {},
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:   {},
}

// GetCipherSuite returns the corresponding cipher suite,
// and boolean value if it is supported.
func GetCipherSuite(s string) (uint16, bool) {
	for _, c := range tls.CipherSuites() {
		if s == c.Name {
			_, ok := fipsCipherSuites[c.ID]
			if !ok {
				return 0, false
			}
			return c.ID, true
		}
	}
	return 0, false
}

// GetCipherSuites returns list of corresponding cipher suite IDs.
func GetCipherSuites(ss []string) ([]uint16, error) {
	cs := make([]uint16, len(ss))
	for i, s := range ss {
		var ok bool
		cs[i], ok = GetCipherSuite(s)
		if !ok {
			return nil, fmt.Errorf("unexpected TLS cipher suite %q", s)
		}
	}

	return cs, nil
}
