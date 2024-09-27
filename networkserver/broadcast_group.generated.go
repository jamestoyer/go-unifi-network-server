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

type BroadcastGroup struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	MemberTable *[]string `json:"member_table,omitempty"`
	Name        *string   `json:"name,omitempty"`
}

func (s *BroadcastGroup) GetMemberTable() []string {
	if s == nil || s.MemberTable == nil {
		return nil
	}

	return *s.MemberTable
}

func (s *BroadcastGroup) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}
