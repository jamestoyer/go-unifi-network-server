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

type DhcpOption struct {
	Code   *string `json:"code,omitempty"`
	Name   *string `json:"name,omitempty"`
	Signed *bool   `json:"signed,omitempty"`
	Type   *string `json:"type,omitempty"`
	Width  *int64  `json:"width,omitempty"`
}

func (s *DhcpOption) GetCode() string {
	if s == nil {
		return ""
	}

	return *s.Code
}

func (s *DhcpOption) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DhcpOption) GetSigned() bool {
	if s == nil {
		return false
	}

	return *s.Signed
}

func (s *DhcpOption) GetType() string {
	if s == nil {
		return ""
	}

	return *s.Type
}

func (s *DhcpOption) GetWidth() int64 {
	if s == nil {
		return 0
	}

	return *s.Width
}
