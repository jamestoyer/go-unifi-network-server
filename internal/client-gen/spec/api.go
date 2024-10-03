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

package spec

import (
	"fmt"
	"maps"
	"slices"
)

type API struct {
	endpoints map[string]*Endpoint
	version   string
}

func NewAPI(version string, endpoints ...*Endpoint) (*API, error) {
	e := make(map[string]*Endpoint, len(endpoints))
	for _, endpoint := range endpoints {
		if _, ok := e[endpoint.Name]; ok {
			return nil, fmt.Errorf("duplicate endpoint: %s", endpoint.Name)
		}

		e[endpoint.Name] = endpoint
	}

	return &API{
		endpoints: e,
		version:   version,
	}, nil
}

func (a *API) AddEndpoint(endpoint *Endpoint) error {
	if _, ok := a.endpoints[endpoint.Name]; ok {
		return fmt.Errorf("duplicate endpoint: %s", endpoint.Name)
	}

	a.endpoints[endpoint.Name] = endpoint
	return nil
}

func (a *API) RemoveEndpoint(name string) {
	delete(a.endpoints, name)
}

func (a *API) Endpoint(name string) *Endpoint {
	if ep, ok := a.endpoints[name]; ok {
		return ep
	}

	return nil
}

func (a *API) Endpoints() []*Endpoint {
	sortedKeys := slices.Sorted(maps.Keys(a.endpoints))
	var endpoints []*Endpoint
	for _, key := range sortedKeys {
		endpoints = append(endpoints, a.endpoints[key])
	}

	return endpoints
}

func (a *API) Version() string {
	return a.version
}
