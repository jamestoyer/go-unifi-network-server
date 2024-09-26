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
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

var defaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

type Client struct {
	httpClient *http.Client

	addAPIPrefix bool
	config       ClientConfig
	endpoint     *url.URL
	username     string
	password     string
}

// NewClient creates a new Unifi Network Server client.
func NewClient(ctx context.Context, endpoint, username, password string, options ...ClientOption) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}

	defaultOpts := []ClientOption{
		withDefaultUserAgent(),
	}
	options = append(defaultOpts, options...)

	var config ClientConfig
	for _, option := range options {
		config = option.apply(config)
	}

	client := &Client{
		config:   config,
		endpoint: u,
		username: username,
		password: password,
	}

	if err = client.initialise(ctx); err != nil {
		return nil, fmt.Errorf("initialise client failed: %w", err)
	}

	return client, nil
}

func (c *Client) newHTTPClient() (*http.Client, error) {
	// https://github.com/unpoller/unifi/blob/5679712eac71f81836fcc44ce843b00fbfda6dc9/unifi.go#L38
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("error creating default cookie jar: %w", err)
	}

	httpClient := &http.Client{
		Jar:       jar,
		Transport: defaultTransport,
	}

	if c.config.TLSConfig != nil {
		clonedTransport := defaultTransport.Clone()
		httpClient.Transport = clonedTransport
		clonedTransport.TLSClientConfig = c.config.TLSConfig
	}

	return httpClient, nil
}

func (c *Client) initialise(ctx context.Context) error {
	httpClient, err := c.newHTTPClient()
	if err != nil {
		return fmt.Errorf("error creating http client: %w", err)
	}

	c.httpClient = httpClient
	return c.setAddAPIPrefix(ctx)
}

func (c *Client) setAddAPIPrefix(ctx context.Context) error {
	// Adapted from https://github.com/unpoller/unifi/blob/5679712eac71f81836fcc44ce843b00fbfda6dc9/unifi.go#L156-L201
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Set("User-Agent", c.config.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// We can't share these cookies with other requests, so make a new client.
	client, err := c.newHTTPClient()
	if err != nil {
		return err
	}

	// Checking the return code on the first request so don't follow a redirect.
	client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to determine API type: %w", err)
	}
	_ = resp.Body.Close()

	// The proxy prefixes need to be added when the status code is a 200. The older style controller endpoints, without
	// prefixes, typically return 302s
	c.addAPIPrefix = resp.StatusCode == http.StatusOK
	return nil
}
