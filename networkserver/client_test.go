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
	"github.com/jamestoyer/go-unifi-network-server/networkserver/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type ClientIntegrationTestSuite struct {
	suite.Suite

	unifiContainer *testhelpers.UnifiNetworkServerContainer
}

func (suite *ClientIntegrationTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := testhelpers.NewUnifiNetworkServerContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.unifiContainer = container
}

func (suite *ClientIntegrationTestSuite) TearDownSuite() {
	ctx := context.Background()
	if err := suite.unifiContainer.Terminate(ctx); err != nil {
		log.Fatalf("Error terminating Unifi container: %s", err)
	}
}

func (suite *ClientIntegrationTestSuite) TestNewClient() {
	t := suite.T()
	t.Run("new client", func(t *testing.T) {
		ctx := context.Background()
		client, err := NewClient(ctx, suite.unifiContainer.Endpoint, "admin", "admin",
			WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		)
		require.NoError(t, err)
		assert.NoError(t, client.Authenticate(ctx))
	})
}

func TestClientIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ClientIntegrationTestSuite))
}