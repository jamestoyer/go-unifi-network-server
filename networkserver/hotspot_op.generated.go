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

type HotspotOp struct {
	ID        *string `json:"_id,omitempty"`
	SiteID    *string `json:"site_id,omitempty"`
	Hidden    *bool   `json:"attr_hidden,omitempty"`
	HiddenID  *string `json:"attr_hidden_id,omitempty"`
	NoDelete  *bool   `json:"attr_no_delete,omitempty"`
	NoEdit    *bool   `json:"attr_no_edit,omitempty"`
	Name      *string `json:"name,omitempty"`
	Note      *string `json:"note,omitempty"`
	XPassword *string `json:"x_password,omitempty"`
}

func (s *HotspotOp) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *HotspotOp) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *HotspotOp) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *HotspotOp) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *HotspotOp) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *HotspotOp) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *HotspotOp) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *HotspotOp) GetNote() string {
	if s == nil {
		return ""
	}

	return *s.Note
}

func (s *HotspotOp) GetXPassword() string {
	if s == nil {
		return ""
	}

	return *s.XPassword
}

type responseBodyHotspotOp struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []HotspotOp     `json:"data"`
}

func (c *Client) CreateHotspotOp(ctx context.Context, site string, data *HotspotOp) (*HotspotOp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspotop")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspotOp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create hotspotop: %w`, err)
	}

	var item HotspotOp
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create HotspotOp`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteHotspotOp(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspotop", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyHotspotOp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete HotspotOp: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetHotspotOp(ctx context.Context, site, id string) (*HotspotOp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspotop", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspotOp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get HotspotOp: %w`, err)
	}

	var item HotspotOp
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListHotspotOp(ctx context.Context, site string) ([]HotspotOp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspotop")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspotOp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get HotspotOp: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateHotspotOp(ctx context.Context, site string, data *HotspotOp) (*HotspotOp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspotop", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspotOp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update hotspotop: %w`, err)
	}

	var item HotspotOp
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update HotspotOp`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
