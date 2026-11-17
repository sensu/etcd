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

//go:build fips140

package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/metadata"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

// TestAuthInfoFromCtxBearerFIPS checks that a "Bearer "-prefixed token is
// stripped before lookup on the FIPS path.
func TestAuthInfoFromCtxBearerFIPS(t *testing.T) {
	tp, err := NewTokenProvider(zaptest.NewLogger(t), testJWTOpts(), dummyIndexWaiter, simpleTokenTTLDefault)
	require.NoError(t, err)
	as := NewAuthStore(zaptest.NewLogger(t), newBackendMock(), tp, 4)
	defer as.Close()
	require.NoError(t, enableAuthAndCreateRoot(as))

	resp, err := as.Authenticate(t.Context(), "root", "123")
	require.NoError(t, err)

	md := metadata.New(map[string]string{rpctypes.TokenFieldNameGRPC: "Bearer " + resp.Token})
	inCtx := metadata.NewIncomingContext(t.Context(), md)

	ai, err := as.AuthInfoFromCtx(inCtx)
	require.NoError(t, err)
	require.NotNil(t, ai)
	require.Equal(t, "root", ai.Username)
}

// TestHashPasswordRoundTripFIPS checks that a password hashed by the
// store verifies with the right password and fails with a wrong one.
func TestHashPasswordRoundTripFIPS(t *testing.T) {
	as := NewAuthStore(zaptest.NewLogger(t), newBackendMock(), nil, 4)
	defer as.Close()

	hashed, err := as.HashPassword("correct-horse")
	require.NoError(t, err)
	require.NoError(t, compareFIPSPasswordHash(hashed, "correct-horse"))
	require.Error(t, compareFIPSPasswordHash(hashed, "wrong-password"))
}
