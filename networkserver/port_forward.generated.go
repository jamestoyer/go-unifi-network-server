// Code generated from Unifi Network Server API definitions.
// DO NOT EDIT.

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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

type PortForward struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	DestinationIp      *string                      `json:"destination_ip,omitempty"`
	DestinationIps     *[]PortForwardDestinationIps `json:"destination_ips,omitempty"`
	DstPort            *string                      `json:"dst_port,omitempty"`
	Enabled            *bool                        `json:"enabled,omitempty"`
	Fwd                *string                      `json:"fwd,omitempty"`
	FwdPort            *string                      `json:"fwd_port,omitempty"`
	Log                *bool                        `json:"log,omitempty"`
	Name               *string                      `json:"name,omitempty"`
	PfwdInterface      *string                      `json:"pfwd_interface,omitempty"`
	Proto              *string                      `json:"proto,omitempty"`
	Src                *string                      `json:"src,omitempty"`
	SrcFirewallGroupId *string                      `json:"src_firewall_group_id,omitempty"`
	SrcLimitingEnabled *bool                        `json:"src_limiting_enabled,omitempty"`
	SrcLimitingType    *string                      `json:"src_limiting_type,omitempty"`
}

func (s *PortForward) GetDestinationIp() string {
	if s == nil {
		return ""
	}

	return *s.DestinationIp
}

func (s *PortForward) GetDestinationIps() []PortForwardDestinationIps {
	if s == nil || s.DestinationIps == nil {
		return nil
	}

	return *s.DestinationIps
}

func (s *PortForward) GetDstPort() string {
	if s == nil {
		return ""
	}

	return *s.DstPort
}

func (s *PortForward) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *PortForward) GetFwd() string {
	if s == nil {
		return ""
	}

	return *s.Fwd
}

func (s *PortForward) GetFwdPort() string {
	if s == nil {
		return ""
	}

	return *s.FwdPort
}

func (s *PortForward) GetLog() bool {
	if s == nil {
		return false
	}

	return *s.Log
}

func (s *PortForward) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *PortForward) GetPfwdInterface() string {
	if s == nil {
		return ""
	}

	return *s.PfwdInterface
}

func (s *PortForward) GetProto() string {
	if s == nil {
		return ""
	}

	return *s.Proto
}

func (s *PortForward) GetSrc() string {
	if s == nil {
		return ""
	}

	return *s.Src
}

func (s *PortForward) GetSrcFirewallGroupId() string {
	if s == nil {
		return ""
	}

	return *s.SrcFirewallGroupId
}

func (s *PortForward) GetSrcLimitingEnabled() bool {
	if s == nil {
		return false
	}

	return *s.SrcLimitingEnabled
}

func (s *PortForward) GetSrcLimitingType() string {
	if s == nil {
		return ""
	}

	return *s.SrcLimitingType
}

type PortForwardDestinationIps struct {
	DestinationIp *string `json:"destination_ip,omitempty"`
	Interface     *string `json:"interface,omitempty"`
}

func (s *PortForwardDestinationIps) GetDestinationIp() string {
	if s == nil {
		return ""
	}

	return *s.DestinationIp
}

func (s *PortForwardDestinationIps) GetInterface() string {
	if s == nil {
		return ""
	}

	return *s.Interface
}

type responseBodyPortForward struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []PortForward   `json:"data"`
}

func (c *Client) CreatePortForward(ctx context.Context, site string, data *PortForward) (*PortForward, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "portforward")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyPortForward
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create portforward: %w`, err)
	}

	var item PortForward
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create PortForward`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeletePortForward(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "portforward", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyPortForward
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete PortForward: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetPortForward(ctx context.Context, site, id string) (*PortForward, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "portforward", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyPortForward
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get PortForward: %w`, err)
	}

	var item PortForward
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListPortForward(ctx context.Context, site string) ([]PortForward, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "portforward")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyPortForward
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get PortForward: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdatePortForward(ctx context.Context, site string, data *PortForward) (*PortForward, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "portforward", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyPortForward
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update portforward: %w`, err)
	}

	var item PortForward
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update PortForward`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
