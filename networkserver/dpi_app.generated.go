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

type DpiApp struct {
	Apps           *[]int64 `json:"apps,omitempty"`
	Blocked        *bool    `json:"blocked,omitempty"`
	Cats           *[]int64 `json:"cats,omitempty"`
	Enabled        *bool    `json:"enabled,omitempty"`
	Log            *bool    `json:"log,omitempty"`
	Name           *string  `json:"name,omitempty"`
	QosRateMaxDown *float64 `json:"qos_rate_max_down,omitempty"`
	QosRateMaxUp   *float64 `json:"qos_rate_max_up,omitempty"`
}

func (s *DpiApp) GetApps() []int64 {
	if s == nil || s.Apps == nil {
		return nil
	}

	return *s.Apps
}

func (s *DpiApp) GetBlocked() bool {
	if s == nil {
		return false
	}

	return *s.Blocked
}

func (s *DpiApp) GetCats() []int64 {
	if s == nil || s.Cats == nil {
		return nil
	}

	return *s.Cats
}

func (s *DpiApp) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *DpiApp) GetLog() bool {
	if s == nil {
		return false
	}

	return *s.Log
}

func (s *DpiApp) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DpiApp) GetQosRateMaxDown() float64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxDown
}

func (s *DpiApp) GetQosRateMaxUp() float64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxUp
}
