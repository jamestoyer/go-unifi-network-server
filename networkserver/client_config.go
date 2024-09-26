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
	"crypto/tls"
)

type ClientConfig struct {
	TLSConfig *tls.Config
	UserAgent string
}

type ClientOption interface {
	apply(ClientConfig) ClientConfig
}

type clientOptionFunc func(ClientConfig) ClientConfig

func (fn clientOptionFunc) apply(config ClientConfig) ClientConfig {
	return fn(config)
}

func withDefaultUserAgent() ClientOption {
	// TODO: (jtoyer) Set the version in the user agent
	return WithUserAgent("jamestoyer/go-unifi-network-server")
}

func WithTLSConfig(tlsConfig *tls.Config) ClientOption {
	return clientOptionFunc(func(config ClientConfig) ClientConfig {
		config.TLSConfig = tlsConfig
		return config
	})
}

func WithUserAgent(userAgent string) ClientOption {
	return clientOptionFunc(func(config ClientConfig) ClientConfig {
		config.UserAgent = userAgent
		return config
	})
}
