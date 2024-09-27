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

type Routing struct {
	ID                   *string `json:"_id,omitempty"`
	SiteID               *string `json:"site_id,omitempty"`
	Hidden               *bool   `json:"attr_hidden,omitempty"`
	HiddenID             *string `json:"attr_hidden_id,omitempty"`
	NoDelete             *bool   `json:"attr_no_delete,omitempty"`
	NoEdit               *bool   `json:"attr_no_edit,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty"`
	GatewayDevice        *string `json:"gateway_device,omitempty"`
	GatewayType          *string `json:"gateway_type,omitempty"`
	Name                 *string `json:"name,omitempty"`
	StaticRouteDistance  *int64  `json:"static-route_distance,omitempty"`
	StaticRouteInterface *string `json:"static-route_interface,omitempty"`
	StaticRouteNetwork   *string `json:"static-route_network,omitempty"`
	StaticRouteNexthop   *string `json:"static-route_nexthop,omitempty"`
	StaticRouteType      *string `json:"static-route_type,omitempty"`
	Type                 *string `json:"type,omitempty"`
}

func (s *Routing) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *Routing) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *Routing) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *Routing) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *Routing) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *Routing) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *Routing) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *Routing) GetGatewayDevice() string {
	if s == nil {
		return ""
	}

	return *s.GatewayDevice
}

func (s *Routing) GetGatewayType() string {
	if s == nil {
		return ""
	}

	return *s.GatewayType
}

func (s *Routing) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Routing) GetStaticRouteDistance() int64 {
	if s == nil {
		return 0
	}

	return *s.StaticRouteDistance
}

func (s *Routing) GetStaticRouteInterface() string {
	if s == nil {
		return ""
	}

	return *s.StaticRouteInterface
}

func (s *Routing) GetStaticRouteNetwork() string {
	if s == nil {
		return ""
	}

	return *s.StaticRouteNetwork
}

func (s *Routing) GetStaticRouteNexthop() string {
	if s == nil {
		return ""
	}

	return *s.StaticRouteNexthop
}

func (s *Routing) GetStaticRouteType() string {
	if s == nil {
		return ""
	}

	return *s.StaticRouteType
}

func (s *Routing) GetType() string {
	if s == nil {
		return ""
	}

	return *s.Type
}

type responseBodyRouting struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []Routing       `json:"data"`
}

func (c *Client) CreateRouting(ctx context.Context, site string, data *Routing) (*Routing, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "routing")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyRouting
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create routing: %w`, err)
	}

	var item Routing
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create Routing`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteRouting(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "routing", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyRouting
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete Routing: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetRouting(ctx context.Context, site, id string) (*Routing, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "routing", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyRouting
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Routing: %w`, err)
	}

	var item Routing
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListRouting(ctx context.Context, site string) ([]Routing, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "routing")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyRouting
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Routing: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateRouting(ctx context.Context, site string, data *Routing) (*Routing, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "routing", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyRouting
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update routing: %w`, err)
	}

	var item Routing
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update Routing`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
