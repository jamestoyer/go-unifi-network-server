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

package testhelpers

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type UnifiNetworkServerContainer struct {
	testcontainers.Container
	Endpoint string
}

func NewUnifiNetworkServerContainer(ctx context.Context) (*UnifiNetworkServerContainer, error) {
	// When set to "true" debug logs from the server will be output
	unifiStdOut := ""
	var logConsumers []testcontainers.LogConsumer
	if stdOut := os.Getenv("UNIFI_STDOUT"); stdOut != "" {
		unifiStdOut = stdOut
	}

	logConsumers = append(logConsumers, &testcontainers.StdoutLogConsumer{})

	// TODO: (jtoyer) This should be loaded from the rendered details version, once that is available
	unifiVersion := "v8.4.62"
	if version := os.Getenv("UNIFI_VERSION"); version != "" {
		unifiVersion = version
	}

	demoModePath, err := filepath.Abs("../dev/unifi/demo-mode")
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path for demo-mode: %w", err)
	}

	r, err := os.Open(demoModePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open Unifi Network Server configuration file: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Name:  fmt.Sprintf("unifi-%v", time.Now().Unix()),
		Image: "jacobalberty/unifi:" + unifiVersion,
		Env: map[string]string{
			"PKGURL":       "",
			"UNIFI_STDOUT": unifiStdOut,
		},
		ExposedPorts: []string{"8443/tcp"},
		Files: []testcontainers.ContainerFile{
			{
				Reader:            r,
				HostFilePath:      demoModePath,
				ContainerFilePath: "/usr/local/unifi/init.d/demo-mode",
				FileMode:          0o700,
			},
		},
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Consumers: logConsumers,
		},
		WaitingFor: wait.ForHealthCheck(),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to start Unifi Network Server container: %w", err)
	}

	endpoint, err := container.PortEndpoint(ctx, "8443/tcp", "https")
	if err != nil {
		return nil, fmt.Errorf("unable to get Unifi Network Server endpoint: %w", err)
	}

	return &UnifiNetworkServerContainer{
		Container: container,
		Endpoint:  endpoint,
	}, nil
}
