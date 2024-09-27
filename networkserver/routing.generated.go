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

type Routing struct {
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
