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

type Dashboard struct {
	ID                *string             `json:"_id,omitempty"`
	SiteID            *string             `json:"site_id,omitempty"`
	Hidden            *bool               `json:"attr_hidden,omitempty"`
	HiddenID          *string             `json:"attr_hidden_id,omitempty"`
	NoDelete          *bool               `json:"attr_no_delete,omitempty"`
	NoEdit            *bool               `json:"attr_no_edit,omitempty"`
	ControllerVersion *string             `json:"controller_version,omitempty"`
	Desc              *string             `json:"desc,omitempty"`
	IsPublic          *bool               `json:"is_public,omitempty"`
	Modules           *[]DashboardModules `json:"modules,omitempty"`
	Name              *string             `json:"name,omitempty"`
}

func (s *Dashboard) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *Dashboard) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *Dashboard) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *Dashboard) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *Dashboard) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *Dashboard) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *Dashboard) GetControllerVersion() string {
	if s == nil {
		return ""
	}

	return *s.ControllerVersion
}

func (s *Dashboard) GetDesc() string {
	if s == nil {
		return ""
	}

	return *s.Desc
}

func (s *Dashboard) GetIsPublic() bool {
	if s == nil {
		return false
	}

	return *s.IsPublic
}

func (s *Dashboard) GetModules() []DashboardModules {
	if s == nil || s.Modules == nil {
		return nil
	}

	return *s.Modules
}

func (s *Dashboard) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type DashboardModules struct {
	Config       *string `json:"config,omitempty"`
	Id           *string `json:"id,omitempty"`
	ModuleId     *string `json:"module_id,omitempty"`
	Restrictions *string `json:"restrictions,omitempty"`
}

func (s *DashboardModules) GetConfig() string {
	if s == nil {
		return ""
	}

	return *s.Config
}

func (s *DashboardModules) GetId() string {
	if s == nil {
		return ""
	}

	return *s.Id
}

func (s *DashboardModules) GetModuleId() string {
	if s == nil {
		return ""
	}

	return *s.ModuleId
}

func (s *DashboardModules) GetRestrictions() string {
	if s == nil {
		return ""
	}

	return *s.Restrictions
}

type responseBodyDashboard struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []Dashboard     `json:"data"`
}

func (c *Client) CreateDashboard(ctx context.Context, site string, data *Dashboard) (*Dashboard, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dashboard")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDashboard
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create dashboard: %w`, err)
	}

	var item Dashboard
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create Dashboard`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteDashboard(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dashboard", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyDashboard
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete Dashboard: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetDashboard(ctx context.Context, site, id string) (*Dashboard, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dashboard", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDashboard
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Dashboard: %w`, err)
	}

	var item Dashboard
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListDashboard(ctx context.Context, site string) ([]Dashboard, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dashboard")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDashboard
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Dashboard: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateDashboard(ctx context.Context, site string, data *Dashboard) (*Dashboard, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dashboard", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDashboard
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update dashboard: %w`, err)
	}

	var item Dashboard
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update Dashboard`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
