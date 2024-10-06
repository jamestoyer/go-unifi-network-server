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
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func (suite *ClientIntegrationTestSuite) TestClient_CreateClientDevice() {
	t := suite.T()

	mac, _ := suite.macPool.MAC()

	t.Run("new client device", func(t *testing.T) {
		ctx := context.Background()

		name := "create client device" + time.Now().String()
		got, _, err := suite.client.CreateClientDevice(ctx, &ClientDevice{
			Name: String(name),
			MAC:  String(mac.String()),
		})
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.GetID())
		assert.Equal(t, name, got.GetName())
		assert.Equal(t, mac.String(), got.GetMAC())
	})

	t.Run("existing client device", func(t *testing.T) {
		ctx := context.Background()

		got, _, err := suite.client.CreateClientDevice(ctx, &ClientDevice{
			Name: String("existing client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_DeleteClientDevice() {
	t := suite.T()

	t.Run("delete client device", func(t *testing.T) {
		ctx := context.Background()
		mac, release := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(t, ctx, &ClientDevice{
			Name: String("delete client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		// Validate the client device actually exists before deletion
		clientDevice, _, err := suite.client.GetClientDevice(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.NotNil(t, clientDevice)

		_, err = suite.client.DeleteClientDevice(ctx, newClientDevice.GetID())
		assert.NoError(t, err)

		clientDevice, _, err = suite.client.GetClientDevice(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.Nil(t, clientDevice)

		if !t.Failed() {
			// Only release the MAC if deletion was a success. This ensures that it will not affect other tests
			release()
		}
	})

	t.Run("delete non existent client device", func(t *testing.T) {
		ctx := context.Background()
		mac, release := suite.macPool.MAC()
		defer release()
		_, err := suite.client.DeleteClientDevice(ctx, mac.String())
		assert.NoError(t, err)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_GetClientDevice() {
	t := suite.T()
	t.Run("get existing client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		wantName := "get client device" + time.Now().String()
		newClientDevice := suite.createClientDevice(t, ctx, &ClientDevice{
			Name: String(wantName),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.GetClientDevice(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.GetID())
		assert.Equal(t, wantName, got.GetName())
		assert.Equal(t, mac.String(), got.GetMAC())
	})

	t.Run("get non existent client device", func(t *testing.T) {
		ctx := context.Background()

		got, _, err := suite.client.GetClientDevice(ctx, "fakeid")
		assert.NoError(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_ListClientDevice() {
	t := suite.T()
	t.Run("list client devices", func(t *testing.T) {
		ctx := context.Background()
		wantCount := 10
		// Ensure there are some clients for testing
		for i := 0; i < wantCount; i++ {
			mac, _ := suite.macPool.MAC()
			_ = suite.createClientDevice(t, ctx, &ClientDevice{
				Name: String(fmt.Sprintf("list client device %d", i)),
				MAC:  String(mac.String()),
			})
		}

		got, _, err := suite.client.ListClientDevice(ctx)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, wantCount, len(got))
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_UpdateClientDevice() {
	t := suite.T()
	t.Run("update existing client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(t, ctx, &ClientDevice{
			Name: String("update client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		updatedClientDevice := &ClientDevice{
			ID:         newClientDevice.ID,
			Name:       String("updated"),
			MAC:        String(mac.String()),
			Note:       String("a new note"),
			UseFixedIP: Bool(true),
			FixedIP:    String("192.168.1.25"),
		}
		got, _, err := suite.client.UpdateClientDevice(ctx, updatedClientDevice)
		assert.NoError(t, err)
		assert.NotNil(t, got)

		got, _, err = suite.client.GetClientDevice(ctx, updatedClientDevice.GetID())
		assert.NoError(t, err)

		assert.Equal(t, updatedClientDevice.ID, got.ID)
		assert.Equal(t, updatedClientDevice.Name, got.Name)
		assert.Equal(t, updatedClientDevice.MAC, got.MAC)
		assert.Equal(t, updatedClientDevice.Note, got.Note)
		assert.Equal(t, updatedClientDevice.UseFixedIP, got.UseFixedIP)
		assert.Equal(t, updatedClientDevice.FixedIP, got.FixedIP)
	})

	t.Run("update non existent client device", func(t *testing.T) {
		ctx := context.Background()
		mac, release := suite.macPool.MAC()
		defer release()

		updatedClientDevice := &ClientDevice{
			ID:   String("fakeid"),
			Name: String("updated"),
			MAC:  String(mac.String()),
		}
		got, _, err := suite.client.UpdateClientDevice(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("change mac", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(t, ctx, &ClientDevice{
			Name: String("change mac" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		newMAC, release := suite.macPool.MAC()
		updatedClientDevice := &ClientDevice{
			ID:         newClientDevice.ID,
			Name:       String("updated"),
			MAC:        String(newMAC.String()),
			Note:       String("a new note"),
			UseFixedIP: Bool(true),
			FixedIP:    String("192.168.1.25"),
		}
		got, _, err := suite.client.UpdateClientDevice(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)

		if !t.Failed() {
			release()
		}
	})

	t.Run("client device ID is not set", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		_ = suite.createClientDevice(t, ctx, &ClientDevice{
			Name: String("update ID is not set" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		updatedClientDevice := &ClientDevice{
			Name: String("updated"),
			MAC:  String(mac.String()),
		}
		got, _, err := suite.client.UpdateClientDevice(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) createClientDevice(t *testing.T, ctx context.Context, device *ClientDevice) *ClientDevice {
	t.Helper()

	clientDevice, _, err := suite.client.CreateClientDevice(ctx, device)
	require.NoError(t, err)
	require.NotNil(t, clientDevice)
	return clientDevice
}
