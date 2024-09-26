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

type Account struct {
	Ip               *string `json:"ip,omitempty"`
	Name             *string `json:"name,omitempty"`
	NetworkconfId    *string `json:"networkconf_id,omitempty"`
	TunnelConfigType *string `json:"tunnel_config_type,omitempty"`
	TunnelMediumType *int64  `json:"tunnel_medium_type,omitempty"`
	TunnelType       *int64  `json:"tunnel_type,omitempty"`
	Vlan             *int64  `json:"vlan,omitempty"`
	XPassword        *string `json:"x_password,omitempty"`
}
