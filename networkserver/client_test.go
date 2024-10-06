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
	"log"
	"testing"

	"github.com/jamestoyer/go-unifi-network-server/networkserver/internal/testhelpers"
	"github.com/stretchr/testify/suite"
)

type ClientIntegrationTestSuite struct {
	suite.Suite

	client         *Client
	unifiContainer *testhelpers.UnifiNetworkServerContainer
}

func (suite *ClientIntegrationTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := testhelpers.NewUnifiNetworkServerContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.unifiContainer = container

	client, err := NewClient(ctx, suite.unifiContainer.Endpoint, "admin", "admin", "default",
		WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)
	if err != nil {
		log.Fatal(err)
	}

	suite.client = client
}

func (suite *ClientIntegrationTestSuite) TearDownSuite() {
	ctx := context.Background()
	if err := suite.unifiContainer.Terminate(ctx); err != nil {
		log.Fatalf("Error terminating Unifi container: %s", err)
	}
}

func TestClientIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ClientIntegrationTestSuite))
}
