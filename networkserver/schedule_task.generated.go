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

type ScheduleTask struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Action          *string                       `json:"action,omitempty"`
	CronExpr        *string                       `json:"cron_expr,omitempty"`
	ExecuteOnlyOnce *bool                         `json:"execute_only_once,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	UpgradeTargets  *[]ScheduleTaskUpgradeTargets `json:"upgrade_targets,omitempty"`
}

func (s *ScheduleTask) GetAction() string {
	if s == nil {
		return ""
	}

	return *s.Action
}

func (s *ScheduleTask) GetCronExpr() string {
	if s == nil {
		return ""
	}

	return *s.CronExpr
}

func (s *ScheduleTask) GetExecuteOnlyOnce() bool {
	if s == nil {
		return false
	}

	return *s.ExecuteOnlyOnce
}

func (s *ScheduleTask) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *ScheduleTask) GetUpgradeTargets() []ScheduleTaskUpgradeTargets {
	if s == nil || s.UpgradeTargets == nil {
		return nil
	}

	return *s.UpgradeTargets
}

type ScheduleTaskUpgradeTargets struct {
	Mac *string `json:"mac,omitempty"`
}

func (s *ScheduleTaskUpgradeTargets) GetMac() string {
	if s == nil {
		return ""
	}

	return *s.Mac
}

type responseBodyScheduleTask struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []ScheduleTask  `json:"data"`
}

func (c *Client) CreateScheduleTask(ctx context.Context, site string, data *ScheduleTask) (*ScheduleTask, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "scheduletask")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyScheduleTask
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create scheduletask: %w`, err)
	}

	var item ScheduleTask
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create ScheduleTask`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteScheduleTask(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "scheduletask", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyScheduleTask
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete ScheduleTask: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetScheduleTask(ctx context.Context, site, id string) (*ScheduleTask, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "scheduletask", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyScheduleTask
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get ScheduleTask: %w`, err)
	}

	var item ScheduleTask
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListScheduleTask(ctx context.Context, site string) ([]ScheduleTask, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "scheduletask")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyScheduleTask
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get ScheduleTask: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateScheduleTask(ctx context.Context, site string, data *ScheduleTask) (*ScheduleTask, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "scheduletask", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyScheduleTask
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update scheduletask: %w`, err)
	}

	var item ScheduleTask
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update ScheduleTask`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
