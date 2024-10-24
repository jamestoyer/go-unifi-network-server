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
	"path"
)

// Delete will remove a given ClientDevice based on its MAC address.
//
// Client devices cannot be removed using their IDs instead the `stamgr` must be used, as discovered from the UI.
func (c *Client) Delete(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandForget, map[string]interface{}{
		"macs": []string{mac},
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to delete client device: %w`, err)
	}

	return resp, nil
}

// GetByMAC will look up a ClientDevice by its MAC address rather than the internal ID used by Get.
//
// Importantly, the information returned by these to functions is not actually the same as they return different
// variations of the data for a client device. The primary difference is that this function will populate the current IP
// unlike Get which doesn't include this.
func (s *ClientDeviceService) GetByMAC(ctx context.Context, mac string) (*ClientDevice, *http.Response, error) {
	endpointPath := path.Join(s.StatAPIPath("user"), mac)
	req, err := s.NewRequest(http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyClientDevice
	resp, err := s.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get ClientDevice: %w`, err)
	}

	var item *ClientDevice
	switch len(body.Payload) {
	case 0:
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}

func (c *Client) Block(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandBlock, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to block client device: %w`, err)
	}

	return resp, nil
}

func (c *Client) Unblock(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandUnblock, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to unblock client device: %w`, err)
	}

	return resp, nil
}

// ForceReconnect will kick a ClientDevice forcing them to reconnect to a site.
//
// References:
//   - https://dl.ui.com/unifi/8.4.62/unifi_sh_api
//   - https://lists.freeradius.org/pipermail/freeradius-users/2017-February/086467.html
func (c *Client) ForceReconnect(ctx context.Context, mac string) (*http.Response, error) {
	resp, err := c.stamgr(ctx, stamgrCommandKick, map[string]interface{}{
		"mac": mac,
	})
	if err != nil {
		return resp, fmt.Errorf(`unable to force client device reconnect: %w`, err)
	}

	return resp, nil
}

type fingerprintOverrideRequest struct {
	DeviceIDOverride int    `json:"dev_id_override"`
	MAC              string `json:"mac"`
}

func (c *Client) OverrideDeviceID(ctx context.Context, mac string, fingerprint int) (*http.Response, error) {
	endpointPath := path.Join(apiV2Path, "site", c.site, "station", mac, "fingerprint_override")
	f := fingerprintOverrideRequest{
		DeviceIDOverride: fingerprint,
		MAC:              mac,
	}

	req, err := c.NewRequest(http.MethodPut, endpointPath, f)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return resp, fmt.Errorf(`failed to override the device ID: %w`, err)
	}

	return resp, nil
}

func (c *Client) RemoveDeviceIDOverride(ctx context.Context, mac string) (*http.Response, error) {
	endpointPath := path.Join(apiV2Path, "site", c.site, "station", mac, "fingerprint_override")
	req, err := c.NewRequest(http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return resp, fmt.Errorf(`failed to remove the device ID override: %w`, err)
	}

	return resp, nil
}
