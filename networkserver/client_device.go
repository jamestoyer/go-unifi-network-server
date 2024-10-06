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
	"net/http"
)

// DeleteClientDeviceByMAC will remove a given ClientDevice based on its MAC address.
//
// Client devices cannot be removed using their IDs instead the `stamgr` must be used, as discovered from the UI.
func (c *Client) DeleteClientDeviceByMAC(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandForget, map[string]interface{}{
		"macs": []string{mac},
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to delete client device: %w`, err)
	}

	return resp, nil
}

func (c *Client) BlockClientDevice(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandBlock, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to block client device: %w`, err)
	}

	return resp, nil
}

func (c *Client) UnblockClientDevice(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandUnblock, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to unblock client device: %w`, err)
	}

	return resp, nil
}

// ForceClientDeviceReconnect will kick a ClientDevice forcing them to reconnect to a site.
//
// References:
//   - https://dl.ui.com/unifi/8.4.62/unifi_sh_api
//   - https://lists.freeradius.org/pipermail/freeradius-users/2017-February/086467.html
func (c *Client) ForceClientDeviceReconnect(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandKick, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to force client device reconnect: %w`, err)
	}

	return resp, nil
}
