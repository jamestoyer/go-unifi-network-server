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

type DynamicDNS struct {
	CustomService *string   `json:"custom_service,omitempty"`
	HostName      *string   `json:"host_name,omitempty"`
	Interface     *string   `json:"interface,omitempty"`
	Login         *string   `json:"login,omitempty"`
	Options       *[]string `json:"options,omitempty"`
	Server        *string   `json:"server,omitempty"`
	Service       *string   `json:"service,omitempty"`
	XPassword     *string   `json:"x_password,omitempty"`
}

func (s *DynamicDNS) GetCustomService() string {
	if s == nil {
		return ""
	}

	return *s.CustomService
}

func (s *DynamicDNS) GetHostName() string {
	if s == nil {
		return ""
	}

	return *s.HostName
}

func (s *DynamicDNS) GetInterface() string {
	if s == nil {
		return ""
	}

	return *s.Interface
}

func (s *DynamicDNS) GetLogin() string {
	if s == nil {
		return ""
	}

	return *s.Login
}

func (s *DynamicDNS) GetOptions() []string {
	if s == nil || s.Options == nil {
		return nil
	}

	return *s.Options
}

func (s *DynamicDNS) GetServer() string {
	if s == nil {
		return ""
	}

	return *s.Server
}

func (s *DynamicDNS) GetService() string {
	if s == nil {
		return ""
	}

	return *s.Service
}

func (s *DynamicDNS) GetXPassword() string {
	if s == nil {
		return ""
	}

	return *s.XPassword
}
