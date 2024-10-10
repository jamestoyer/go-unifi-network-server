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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	// Adapted from https://github.com/unpoller/unifi/blob/5679712eac71f81836fcc44ce843b00fbfda6dc9/types.go#L122-L161

	// apiLoginPath is Unifi Controller Login API Path.
	apiLoginPath string = "/api/login"
	// apiLoginPathNew is how we log into UDM 5.12.55+.
	apiLoginPathNew string = "/api/auth/login"

	// apiPrefixNew is the prefix added to the new API paths; except login. duh.
	apiPrefixNew string = "/proxy/network"

	// apiV1Path is the v1 API prefix.
	apiV1Path = "api"
	// apiV2Path is the v2 API prefix.
	apiV2Path = "/v2/api"

	stamgrCommandBlock   = "block-sta"
	stamgrCommandForget  = "forget-sta"
	stamgrCommandKick    = "kick-sta"
	stamgrCommandUnblock = "unblock-sta"
)

var (
	ErrAuthenticationFailed = errors.New("authentication failed")

	defaultTransport = &http.Transport{
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
)

type service struct {
	*Client
}

// NewClient creates a new Unifi Network Server client.
func NewClient(ctx context.Context, endpoint, username, password, site string, options ...ClientOption) (*Client, error) {
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
		site:     site,
	}

	if err = client.initialise(ctx); err != nil {
		return nil, fmt.Errorf("initialise client failed: %w", err)
	}

	return client, nil
}

// ResourceAPIPath generates the correct API path for a given rest resource that can be used with NewRequest.
func (c *Client) ResourceAPIPath(resource string) string {
	return path.Join(apiV1Path, "s", c.site, "rest", resource)
}

// StatAPIPath generates the correct API path for a given stat resource that can be used with NewRequest.
func (c *Client) StatAPIPath(resource string) string {
	return path.Join(apiV1Path, "s", c.site, "stat", resource)
}

func (c *Client) NewRequest(ctx context.Context, method, urlPath string, body interface{}) (*http.Request, error) {
	u, err := c.endpoint.Parse(c.path(urlPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create API request URL: %w", err)
	}

	var bodyBuffer io.ReadWriter
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}

		bodyBuffer = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	}

	req.Header.Set("User-Agent", c.config.UserAgent)
	req.Header.Add("Accept", "application/json")

	if c.csrf != "" {
		c.csrfLock.RLock()
		req.Header.Set("X-Csrf-Token", c.csrf)
		c.csrfLock.RUnlock()
	}

	return req, nil
}

// Authenticate is predominantly a helper method that should only be called to refresh the authentication cookies if
// they have gone stale.
func (c *Client) Authenticate(ctx context.Context) error {
	req, err := c.NewRequest(ctx, http.MethodPost, apiLoginPath, struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: c.username, Password: c.password})
	if err != nil {
		return err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("authenticate request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return ErrAuthenticationFailed
	}

	return nil
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
	c.initialiseServices()

	httpClient, err := c.newHTTPClient()
	if err != nil {
		return fmt.Errorf("error creating http client: %w", err)
	}

	c.httpClient = httpClient
	if err = c.setAddAPIPrefix(ctx); err != nil {
		return err
	}

	return c.Authenticate(ctx)
}

// setAddAPIPrefix checks to see which type of Network Server it is running against and sets the addAPIPrefix field
// accordingly.
//
// Adapted from https://github.com/unpoller/unifi/blob/5679712eac71f81836fcc44ce843b00fbfda6dc9/unifi.go#L156-L201
func (c *Client) setAddAPIPrefix(ctx context.Context) error {
	req, err := c.NewRequest(ctx, http.MethodGet, "", nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

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

// path sets the correct API endpoint path based on the addAPIPrefix property.
//
// Adapted from https://github.com/unpoller/unifi/blob/5679712eac71f81836fcc44ce843b00fbfda6dc9/types.go#L163-L177
func (c *Client) path(p string) string {
	if c.addAPIPrefix {
		if p == apiLoginPath {
			return apiLoginPathNew
		}

		if !strings.HasPrefix(p, apiPrefixNew) && p != apiLoginPathNew {
			return apiPrefixNew + p
		}
	}

	return p
}

// Do will execute a HTTP request to the Unifi Network Server. When v implements an io.Writer the raw response will be
// copied, otherwise the response body will be JSON decoded to any non nil value of v.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context not set")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failure: %w", err)
	}

	defer resp.Body.Close()

	if csrf := resp.Header.Get("X-Csrf-Token"); csrf != "" {
		c.csrfLock.Lock()
		c.csrf = csrf
		c.csrfLock.Unlock()
	}

	err = checkResponseForError(resp)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return resp, nil
	}

	switch v := v.(type) {
	case io.Writer:
		if _, err = io.Copy(v, resp.Body); err != nil {
			return nil, fmt.Errorf("failed to copy response body: %w", err)
		}
	default:
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil && !errors.Is(err, io.EOF) {
			return resp, fmt.Errorf("json decode failure: %w", err)
		}
	}

	return resp, nil
}

type stamgrCommand string

// stamgr will execute station manager commands. These are generally used for manipulating a ClientDevice or a guest.
//
// Many of the commands are either exposed via the external references or discovered by using the UI.
//
// External References:
//
//   - https://dl.ui.com/unifi/8.4.62/unifi_sh_api
func (c *Client) stamgr(ctx context.Context, command stamgrCommand, argument map[string]interface{}) (*http.Response, error) {
	endpointPath := path.Join(apiV1Path, "s", c.site, "cmd/stamgr")

	strmgrRequest := map[string]interface{}{
		"cmd": string(command),
	}

	for name, value := range argument {
		strmgrRequest[name] = value
	}

	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, strmgrRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return resp, fmt.Errorf(`failed to execute stamgr command: %w`, err)
	}

	return resp, nil
}

type ErrResponse struct {
	Meta struct {
		ResponseCode string `json:"rc"`
		Message      string `json:"msg"`
	} `json:"meta"`
	Data     interface{} `json:"data"`
	response *http.Response
}

func (e *ErrResponse) Error() string {
	if e.response != nil && e.response.Request != nil {
		return fmt.Sprintf("%v %v: %d %s %+v", e.response.Request.Method, e.response.Request.URL, e.response.StatusCode, e.Meta.Message, e.Data)
	}

	if e.response != nil {
		return fmt.Sprintf("%d %s %+v", e.response.StatusCode, e.Meta.Message, e.Data)
	}

	return fmt.Sprintf("%s %+v", e.Meta.Message, e.Data)
}

func checkResponseForError(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	errResp := &ErrResponse{response: resp}
	data, err := io.ReadAll(resp.Body)
	if err == nil && data != nil {
		if err = json.Unmarshal(data, errResp); err != nil {
			// Reset the response to ensure we don't leave it in an inconsistent state after the failed unmarshalling
			errResp = &ErrResponse{}
		}
	}

	// Repopulate the body as it's unclear what the formal spec for the error response is. This allows for the user
	// to interrogate this later.
	resp.Body = io.NopCloser(bytes.NewBuffer(data))
	return errResp
}
