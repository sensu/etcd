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

//go:build !fips140
// +build !fips140

package auth

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// encodePassword returns the stored form of a password for the default build:
// a bcrypt hash, base64-encoded.
func encodePassword(s string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	return base64.StdEncoding.EncodeToString(hashedPassword)
}
