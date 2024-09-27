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

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

type FirewallRule struct {
	ID                    *string   `json:"_id,omitempty"`
	SiteID                *string   `json:"site_id,omitempty"`
	Hidden                *bool     `json:"attr_hidden,omitempty"`
	HiddenID              *string   `json:"attr_hidden_id,omitempty"`
	NoDelete              *bool     `json:"attr_no_delete,omitempty"`
	NoEdit                *bool     `json:"attr_no_edit,omitempty"`
	Action                *string   `json:"action,omitempty"`
	Contiguous            *bool     `json:"contiguous,omitempty"`
	DstAddress            *string   `json:"dst_address,omitempty"`
	DstAddressIpv6        *string   `json:"dst_address_ipv6,omitempty"`
	DstFirewallgroupIds   *[]string `json:"dst_firewallgroup_ids,omitempty"`
	DstNetworkconfId      *string   `json:"dst_networkconf_id,omitempty"`
	DstNetworkconfType    *string   `json:"dst_networkconf_type,omitempty"`
	DstPort               *string   `json:"dst_port,omitempty"`
	Enabled               *bool     `json:"enabled,omitempty"`
	IcmpTypename          *string   `json:"icmp_typename,omitempty"`
	Icmpv6Typename        *string   `json:"icmpv6_typename,omitempty"`
	Ipsec                 *string   `json:"ipsec,omitempty"`
	Logging               *bool     `json:"logging,omitempty"`
	Monthdays             *string   `json:"monthdays,omitempty"`
	MonthdaysNegate       *bool     `json:"monthdays_negate,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	Protocol              *string   `json:"protocol,omitempty"`
	ProtocolMatchExcepted *bool     `json:"protocol_match_excepted,omitempty"`
	ProtocolV6            *string   `json:"protocol_v6,omitempty"`
	RuleIndex             *int64    `json:"rule_index,omitempty"`
	Ruleset               *string   `json:"ruleset,omitempty"`
	SettingPreference     *string   `json:"setting_preference,omitempty"`
	SrcAddress            *string   `json:"src_address,omitempty"`
	SrcAddressIpv6        *string   `json:"src_address_ipv6,omitempty"`
	SrcFirewallgroupIds   *[]string `json:"src_firewallgroup_ids,omitempty"`
	SrcMacAddress         *string   `json:"src_mac_address,omitempty"`
	SrcNetworkconfId      *string   `json:"src_networkconf_id,omitempty"`
	SrcNetworkconfType    *string   `json:"src_networkconf_type,omitempty"`
	SrcPort               *string   `json:"src_port,omitempty"`
	Startdate             *string   `json:"startdate,omitempty"`
	Starttime             *string   `json:"starttime,omitempty"`
	StateEstablished      *bool     `json:"state_established,omitempty"`
	StateInvalid          *bool     `json:"state_invalid,omitempty"`
	StateNew              *bool     `json:"state_new,omitempty"`
	StateRelated          *bool     `json:"state_related,omitempty"`
	Stopdate              *string   `json:"stopdate,omitempty"`
	Stoptime              *string   `json:"stoptime,omitempty"`
	Utc                   *bool     `json:"utc,omitempty"`
	Weekdays              *string   `json:"weekdays,omitempty"`
	WeekdaysNegate        *bool     `json:"weekdays_negate,omitempty"`
}

func (s *FirewallRule) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *FirewallRule) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *FirewallRule) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *FirewallRule) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *FirewallRule) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *FirewallRule) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *FirewallRule) GetAction() string {
	if s == nil {
		return ""
	}

	return *s.Action
}

func (s *FirewallRule) GetContiguous() bool {
	if s == nil {
		return false
	}

	return *s.Contiguous
}

func (s *FirewallRule) GetDstAddress() string {
	if s == nil {
		return ""
	}

	return *s.DstAddress
}

func (s *FirewallRule) GetDstAddressIpv6() string {
	if s == nil {
		return ""
	}

	return *s.DstAddressIpv6
}

func (s *FirewallRule) GetDstFirewallgroupIds() []string {
	if s == nil || s.DstFirewallgroupIds == nil {
		return nil
	}

	return *s.DstFirewallgroupIds
}

func (s *FirewallRule) GetDstNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.DstNetworkconfId
}

func (s *FirewallRule) GetDstNetworkconfType() string {
	if s == nil {
		return ""
	}

	return *s.DstNetworkconfType
}

func (s *FirewallRule) GetDstPort() string {
	if s == nil {
		return ""
	}

	return *s.DstPort
}

func (s *FirewallRule) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *FirewallRule) GetIcmpTypename() string {
	if s == nil {
		return ""
	}

	return *s.IcmpTypename
}

func (s *FirewallRule) GetIcmpv6Typename() string {
	if s == nil {
		return ""
	}

	return *s.Icmpv6Typename
}

func (s *FirewallRule) GetIpsec() string {
	if s == nil {
		return ""
	}

	return *s.Ipsec
}

func (s *FirewallRule) GetLogging() bool {
	if s == nil {
		return false
	}

	return *s.Logging
}

func (s *FirewallRule) GetMonthdays() string {
	if s == nil {
		return ""
	}

	return *s.Monthdays
}

func (s *FirewallRule) GetMonthdaysNegate() bool {
	if s == nil {
		return false
	}

	return *s.MonthdaysNegate
}

func (s *FirewallRule) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *FirewallRule) GetProtocol() string {
	if s == nil {
		return ""
	}

	return *s.Protocol
}

func (s *FirewallRule) GetProtocolMatchExcepted() bool {
	if s == nil {
		return false
	}

	return *s.ProtocolMatchExcepted
}

func (s *FirewallRule) GetProtocolV6() string {
	if s == nil {
		return ""
	}

	return *s.ProtocolV6
}

func (s *FirewallRule) GetRuleIndex() int64 {
	if s == nil {
		return 0
	}

	return *s.RuleIndex
}

func (s *FirewallRule) GetRuleset() string {
	if s == nil {
		return ""
	}

	return *s.Ruleset
}

func (s *FirewallRule) GetSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.SettingPreference
}

func (s *FirewallRule) GetSrcAddress() string {
	if s == nil {
		return ""
	}

	return *s.SrcAddress
}

func (s *FirewallRule) GetSrcAddressIpv6() string {
	if s == nil {
		return ""
	}

	return *s.SrcAddressIpv6
}

func (s *FirewallRule) GetSrcFirewallgroupIds() []string {
	if s == nil || s.SrcFirewallgroupIds == nil {
		return nil
	}

	return *s.SrcFirewallgroupIds
}

func (s *FirewallRule) GetSrcMacAddress() string {
	if s == nil {
		return ""
	}

	return *s.SrcMacAddress
}

func (s *FirewallRule) GetSrcNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.SrcNetworkconfId
}

func (s *FirewallRule) GetSrcNetworkconfType() string {
	if s == nil {
		return ""
	}

	return *s.SrcNetworkconfType
}

func (s *FirewallRule) GetSrcPort() string {
	if s == nil {
		return ""
	}

	return *s.SrcPort
}

func (s *FirewallRule) GetStartdate() string {
	if s == nil {
		return ""
	}

	return *s.Startdate
}

func (s *FirewallRule) GetStarttime() string {
	if s == nil {
		return ""
	}

	return *s.Starttime
}

func (s *FirewallRule) GetStateEstablished() bool {
	if s == nil {
		return false
	}

	return *s.StateEstablished
}

func (s *FirewallRule) GetStateInvalid() bool {
	if s == nil {
		return false
	}

	return *s.StateInvalid
}

func (s *FirewallRule) GetStateNew() bool {
	if s == nil {
		return false
	}

	return *s.StateNew
}

func (s *FirewallRule) GetStateRelated() bool {
	if s == nil {
		return false
	}

	return *s.StateRelated
}

func (s *FirewallRule) GetStopdate() string {
	if s == nil {
		return ""
	}

	return *s.Stopdate
}

func (s *FirewallRule) GetStoptime() string {
	if s == nil {
		return ""
	}

	return *s.Stoptime
}

func (s *FirewallRule) GetUtc() bool {
	if s == nil {
		return false
	}

	return *s.Utc
}

func (s *FirewallRule) GetWeekdays() string {
	if s == nil {
		return ""
	}

	return *s.Weekdays
}

func (s *FirewallRule) GetWeekdaysNegate() bool {
	if s == nil {
		return false
	}

	return *s.WeekdaysNegate
}

type responseBodyFirewallRule struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []FirewallRule  `json:"data"`
}

func (c *Client) CreateFirewallRule(ctx context.Context, site string, data *FirewallRule) (*FirewallRule, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallrule")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallRule
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create firewallrule: %w`, err)
	}

	var item FirewallRule
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create FirewallRule`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteFirewallRule(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallrule", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyFirewallRule
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete FirewallRule: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetFirewallRule(ctx context.Context, site, id string) (*FirewallRule, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallrule", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallRule
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get FirewallRule: %w`, err)
	}

	var item FirewallRule
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListFirewallRule(ctx context.Context, site string) ([]FirewallRule, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallrule")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallRule
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get FirewallRule: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateFirewallRule(ctx context.Context, site string, data *FirewallRule) (*FirewallRule, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallrule", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallRule
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update firewallrule: %w`, err)
	}

	var item FirewallRule
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update FirewallRule`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
