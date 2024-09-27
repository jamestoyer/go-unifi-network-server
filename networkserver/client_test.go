// Copyright 2024 James Toyer
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

package networkserver

import (
	"context"
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Authenticate(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, "https://localhost:8443", "admin", "admin",
		WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)

	require.NoError(t, err)

	t.Run("successful authentication", func(t *testing.T) {
		got := client.Authenticate(ctx)
		assert.NoError(t, got)
	})

	t.Run("failed authentication", func(t *testing.T) {
		client.username = "invalid"
		got := client.Authenticate(ctx)
		assert.NoError(t, got)
	})
}