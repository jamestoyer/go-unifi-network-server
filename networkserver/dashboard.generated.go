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

type Dashboard struct {
	ControllerVersion *string             `json:"controller_version,omitempty"`
	Desc              *string             `json:"desc,omitempty"`
	IsPublic          *bool               `json:"is_public,omitempty"`
	Modules           *[]DashboardModules `json:"modules,omitempty"`
	Name              *string             `json:"name,omitempty"`
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
