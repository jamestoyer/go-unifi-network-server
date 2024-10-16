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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *ClientIntegrationTestSuite) TestClient_Create() {
	t := suite.T()

	mac, _ := suite.macPool.MAC()

	t.Run("new client device", func(t *testing.T) {
		ctx := context.Background()

		name := "create client device" + time.Now().String()
		got, _, err := suite.client.ClientDevices.Create(ctx, &ClientDevice{
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

		got, _, err := suite.client.ClientDevices.Create(ctx, &ClientDevice{
			Name: String("existing client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_Delete() {
	t := suite.T()

	t.Run("delete client device", func(t *testing.T) {
		ctx := context.Background()
		mac, release := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("delete client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		// Validate the client device actually exists beforge deletion
		clientDevice, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.NotNil(t, clientDevice)

		_, err = suite.client.Delete(ctx, mac.String())
		assert.NoError(t, err)

		clientDevice, _, err = suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
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
		_, err := suite.client.Delete(ctx, mac.String())
		assert.NoError(t, err)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_Get() {
	t := suite.T()
	t.Run("get existing client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		wantName := "get client device" + time.Now().String()
		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String(wantName),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.GetID())
		assert.Equal(t, wantName, got.GetName())
		assert.Equal(t, mac.String(), got.GetMAC())
	})

	t.Run("get non existent client device", func(t *testing.T) {
		ctx := context.Background()

		got, _, err := suite.client.ClientDevices.Get(ctx, "fakeid")
		assert.NoError(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_List() {
	t := suite.T()
	t.Run("list client devices", func(t *testing.T) {
		ctx := context.Background()
		wantCount := 10
		// Ensure there are some clients for testing
		for i := 0; i < wantCount; i++ {
			mac, _ := suite.macPool.MAC()
			_ = suite.createClientDevice(ctx, t, &ClientDevice{
				Name: String(fmt.Sprintf("list client device %d", i)),
				MAC:  String(mac.String()),
			})
		}

		got, _, err := suite.client.ClientDevices.List(ctx)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(got), wantCount)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_Update() {
	t := suite.T()
	t.Run("update existing client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("update client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		updatedClientDevice := &ClientDevice{
			ID:   newClientDevice.ID,
			Name: String("updated"),
			MAC:  String(mac.String()),
			Note: String("a new note"),
		}
		got, _, err := suite.client.ClientDevices.Update(ctx, updatedClientDevice)
		assert.NoError(t, err)
		assert.NotNil(t, got)

		got, _, err = suite.client.ClientDevices.Get(ctx, updatedClientDevice.GetID())
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
		got, _, err := suite.client.ClientDevices.Update(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("change mac", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
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
			FixedIP:    String("192.168.1.2"),
		}
		got, _, err := suite.client.ClientDevices.Update(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)

		if !t.Failed() {
			release()
		}
	})

	t.Run("client device ID is not set", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		_ = suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("update ID is not set" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		updatedClientDevice := &ClientDevice{
			Name: String("updated"),
			MAC:  String(mac.String()),
		}
		got, _, err := suite.client.ClientDevices.Update(ctx, updatedClientDevice)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_GetByMAC() {
	t := suite.T()
	t.Run("get existing client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		wantName := "get client device" + time.Now().String()
		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String(wantName),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.ClientDevices.GetByMAC(ctx, mac.String())
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.GetID())
		assert.Equal(t, wantName, got.GetName())
		assert.Equal(t, mac.String(), got.GetMAC())
		assert.Equal(t, newClientDevice.GetID(), got.GetID())
	})

	t.Run("get non existent client device", func(t *testing.T) {
		ctx := context.Background()
		mac, release := suite.macPool.MAC()
		defer release()

		got, _, err := suite.client.ClientDevices.Get(ctx, mac.String())
		assert.NoError(t, err)
		assert.Nil(t, got)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_Block() {
	t := suite.T()
	t.Run("block client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("block client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.NotNil(t, got)

		_, err = suite.client.Block(ctx, mac.String())
		assert.NoError(t, err)

		got, _, err = suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.True(t, got.GetBlocked())
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_Unblock() {
	t := suite.T()
	t.Run("unblock client device", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("unblock client device" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.NotNil(t, got)

		_, err = suite.client.Block(ctx, mac.String())
		assert.NoError(t, err)

		got, _, err = suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		require.True(t, got.GetBlocked())

		_, err = suite.client.Unblock(ctx, mac.String())
		assert.NoError(t, err)

		got, _, err = suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.False(t, got.GetBlocked())
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_ForceReconnect() {
	t := suite.T()
	t.Run("force client device reconnect", func(t *testing.T) {
		t.Skip("Unable to test as a device needs to connect first")
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("force client device redirect" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.NotNil(t, got)

		_, err = suite.client.ForceReconnect(ctx, mac.String())
		assert.NoError(t, err)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_OverrideDeviceID() {
	t := suite.T()
	t.Run("override fingerprint", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name: String("force client device redirect" + time.Now().String()),
			MAC:  String(mac.String()),
		})

		fingerprint := 1096
		_, err := suite.client.OverrideDeviceID(ctx, mac.String(), fingerprint)
		assert.NoError(t, err)

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.Equal(t, Int64(int64(fingerprint)), got.DeviceIDOverride)
	})
}

func (suite *ClientIntegrationTestSuite) TestClient_RemoveDeviceIDOverride() {
	t := suite.T()
	t.Run("override fingerprint", func(t *testing.T) {
		ctx := context.Background()
		mac, _ := suite.macPool.MAC()
		fingerprint := 1096

		newClientDevice := suite.createClientDevice(ctx, t, &ClientDevice{
			Name:             String("force client device redirect" + time.Now().String()),
			MAC:              String(mac.String()),
			DeviceIDOverride: Int64(int64(fingerprint)),
		})

		_, err := suite.client.OverrideDeviceID(ctx, mac.String(), fingerprint)
		require.NoError(t, err)

		got, _, err := suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		require.NoError(t, err)
		require.Equal(t, Int64(int64(fingerprint)), got.DeviceIDOverride)

		_, err = suite.client.RemoveDeviceIDOverride(ctx, mac.String())
		assert.NoError(t, err)

		got, _, err = suite.client.ClientDevices.Get(ctx, newClientDevice.GetID())
		assert.NoError(t, err)
		assert.Nil(t, got.DeviceIDOverride)
	})
}

func (suite *ClientIntegrationTestSuite) createClientDevice(ctx context.Context, t *testing.T, device *ClientDevice) *ClientDevice {
	t.Helper()

	clientDevice, _, err := suite.client.ClientDevices.Create(ctx, device)
	require.NoError(t, err)
	require.NotNil(t, clientDevice)
	return clientDevice
}
