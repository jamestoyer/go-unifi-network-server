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

type PortConf struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Autoneg                       *bool               `json:"autoneg,omitempty"`
	Dot1XCtrl                     *string             `json:"dot1x_ctrl,omitempty"`
	Dot1XIdleTimeout              *float64            `json:"dot1x_idle_timeout,omitempty"`
	EgressRateLimitKbps           *int64              `json:"egress_rate_limit_kbps,omitempty"`
	EgressRateLimitKbpsEnabled    *bool               `json:"egress_rate_limit_kbps_enabled,omitempty"`
	ExcludedNetworkconfIds        *[]string           `json:"excluded_networkconf_ids,omitempty"`
	FecMode                       *string             `json:"fec_mode,omitempty"`
	Forward                       *string             `json:"forward,omitempty"`
	FullDuplex                    *bool               `json:"full_duplex,omitempty"`
	Isolation                     *bool               `json:"isolation,omitempty"`
	LldpmedEnabled                *bool               `json:"lldpmed_enabled,omitempty"`
	LldpmedNotifyEnabled          *bool               `json:"lldpmed_notify_enabled,omitempty"`
	MulticastRouterNetworkconfIds *[]string           `json:"multicast_router_networkconf_ids,omitempty"`
	Name                          *string             `json:"name,omitempty"`
	NativeNetworkconfId           *string             `json:"native_networkconf_id,omitempty"`
	OpMode                        *string             `json:"op_mode,omitempty"`
	PoeMode                       *string             `json:"poe_mode,omitempty"`
	PortKeepaliveEnabled          *bool               `json:"port_keepalive_enabled,omitempty"`
	PortSecurityEnabled           *bool               `json:"port_security_enabled,omitempty"`
	PortSecurityMacAddress        *[]string           `json:"port_security_mac_address,omitempty"`
	PriorityQueue1Level           *int64              `json:"priority_queue1_level,omitempty"`
	PriorityQueue2Level           *int64              `json:"priority_queue2_level,omitempty"`
	PriorityQueue3Level           *int64              `json:"priority_queue3_level,omitempty"`
	PriorityQueue4Level           *int64              `json:"priority_queue4_level,omitempty"`
	QosProfile                    *PortConfQosProfile `json:"qos_profile,omitempty"`
	SettingPreference             *string             `json:"setting_preference,omitempty"`
	Speed                         *float64            `json:"speed,omitempty"`
	StormctrlBcastEnabled         *bool               `json:"stormctrl_bcast_enabled,omitempty"`
	StormctrlBcastLevel           *int64              `json:"stormctrl_bcast_level,omitempty"`
	StormctrlBcastRate            *float64            `json:"stormctrl_bcast_rate,omitempty"`
	StormctrlMcastEnabled         *bool               `json:"stormctrl_mcast_enabled,omitempty"`
	StormctrlMcastLevel           *int64              `json:"stormctrl_mcast_level,omitempty"`
	StormctrlMcastRate            *float64            `json:"stormctrl_mcast_rate,omitempty"`
	StormctrlType                 *string             `json:"stormctrl_type,omitempty"`
	StormctrlUcastEnabled         *bool               `json:"stormctrl_ucast_enabled,omitempty"`
	StormctrlUcastLevel           *int64              `json:"stormctrl_ucast_level,omitempty"`
	StormctrlUcastRate            *float64            `json:"stormctrl_ucast_rate,omitempty"`
	StpPortMode                   *bool               `json:"stp_port_mode,omitempty"`
	TaggedVlanMgmt                *string             `json:"tagged_vlan_mgmt,omitempty"`
	VoiceNetworkconfId            *string             `json:"voice_networkconf_id,omitempty"`
}

func (s *PortConf) GetAutoneg() bool {
	if s == nil {
		return false
	}

	return *s.Autoneg
}

func (s *PortConf) GetDot1XCtrl() string {
	if s == nil {
		return ""
	}

	return *s.Dot1XCtrl
}

func (s *PortConf) GetDot1XIdleTimeout() float64 {
	if s == nil {
		return 0
	}

	return *s.Dot1XIdleTimeout
}

func (s *PortConf) GetEgressRateLimitKbps() int64 {
	if s == nil {
		return 0
	}

	return *s.EgressRateLimitKbps
}

func (s *PortConf) GetEgressRateLimitKbpsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.EgressRateLimitKbpsEnabled
}

func (s *PortConf) GetExcludedNetworkconfIds() []string {
	if s == nil || s.ExcludedNetworkconfIds == nil {
		return nil
	}

	return *s.ExcludedNetworkconfIds
}

func (s *PortConf) GetFecMode() string {
	if s == nil {
		return ""
	}

	return *s.FecMode
}

func (s *PortConf) GetForward() string {
	if s == nil {
		return ""
	}

	return *s.Forward
}

func (s *PortConf) GetFullDuplex() bool {
	if s == nil {
		return false
	}

	return *s.FullDuplex
}

func (s *PortConf) GetIsolation() bool {
	if s == nil {
		return false
	}

	return *s.Isolation
}

func (s *PortConf) GetLldpmedEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LldpmedEnabled
}

func (s *PortConf) GetLldpmedNotifyEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LldpmedNotifyEnabled
}

func (s *PortConf) GetMulticastRouterNetworkconfIds() []string {
	if s == nil || s.MulticastRouterNetworkconfIds == nil {
		return nil
	}

	return *s.MulticastRouterNetworkconfIds
}

func (s *PortConf) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *PortConf) GetNativeNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.NativeNetworkconfId
}

func (s *PortConf) GetOpMode() string {
	if s == nil {
		return ""
	}

	return *s.OpMode
}

func (s *PortConf) GetPoeMode() string {
	if s == nil {
		return ""
	}

	return *s.PoeMode
}

func (s *PortConf) GetPortKeepaliveEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PortKeepaliveEnabled
}

func (s *PortConf) GetPortSecurityEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PortSecurityEnabled
}

func (s *PortConf) GetPortSecurityMacAddress() []string {
	if s == nil || s.PortSecurityMacAddress == nil {
		return nil
	}

	return *s.PortSecurityMacAddress
}

func (s *PortConf) GetPriorityQueue1Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue1Level
}

func (s *PortConf) GetPriorityQueue2Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue2Level
}

func (s *PortConf) GetPriorityQueue3Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue3Level
}

func (s *PortConf) GetPriorityQueue4Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue4Level
}

func (s *PortConf) GetQosProfile() *PortConfQosProfile {
	if s == nil || s.QosProfile == nil {
		return nil
	}

	return s.QosProfile
}

func (s *PortConf) GetSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.SettingPreference
}

func (s *PortConf) GetSpeed() float64 {
	if s == nil {
		return 0
	}

	return *s.Speed
}

func (s *PortConf) GetStormctrlBcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlBcastEnabled
}

func (s *PortConf) GetStormctrlBcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlBcastLevel
}

func (s *PortConf) GetStormctrlBcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlBcastRate
}

func (s *PortConf) GetStormctrlMcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlMcastEnabled
}

func (s *PortConf) GetStormctrlMcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlMcastLevel
}

func (s *PortConf) GetStormctrlMcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlMcastRate
}

func (s *PortConf) GetStormctrlType() string {
	if s == nil {
		return ""
	}

	return *s.StormctrlType
}

func (s *PortConf) GetStormctrlUcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlUcastEnabled
}

func (s *PortConf) GetStormctrlUcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlUcastLevel
}

func (s *PortConf) GetStormctrlUcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlUcastRate
}

func (s *PortConf) GetStpPortMode() bool {
	if s == nil {
		return false
	}

	return *s.StpPortMode
}

func (s *PortConf) GetTaggedVlanMgmt() string {
	if s == nil {
		return ""
	}

	return *s.TaggedVlanMgmt
}

func (s *PortConf) GetVoiceNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.VoiceNetworkconfId
}

type PortConfQosProfile struct {
	QosPolicies    *[]PortConfQosProfileQosPolicies `json:"qos_policies,omitempty"`
	QosProfileMode *string                          `json:"qos_profile_mode,omitempty"`
}

func (s *PortConfQosProfile) GetQosPolicies() []PortConfQosProfileQosPolicies {
	if s == nil || s.QosPolicies == nil {
		return nil
	}

	return *s.QosPolicies
}

func (s *PortConfQosProfile) GetQosProfileMode() string {
	if s == nil {
		return ""
	}

	return *s.QosProfileMode
}

type PortConfQosProfileQosPolicies struct {
	QosMarking  *PortConfQosProfileQosPoliciesQosMarking  `json:"qos_marking,omitempty"`
	QosMatching *PortConfQosProfileQosPoliciesQosMatching `json:"qos_matching,omitempty"`
}

func (s *PortConfQosProfileQosPolicies) GetQosMarking() *PortConfQosProfileQosPoliciesQosMarking {
	if s == nil || s.QosMarking == nil {
		return nil
	}

	return s.QosMarking
}

func (s *PortConfQosProfileQosPolicies) GetQosMatching() *PortConfQosProfileQosPoliciesQosMatching {
	if s == nil || s.QosMatching == nil {
		return nil
	}

	return s.QosMatching
}

type PortConfQosProfileQosPoliciesQosMarking struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *float64 `json:"dscp_code,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Queue            *int64   `json:"queue,omitempty"`
}

func (s *PortConfQosProfileQosPoliciesQosMarking) GetCosCode() int64 {
	if s == nil {
		return 0
	}

	return *s.CosCode
}

func (s *PortConfQosProfileQosPoliciesQosMarking) GetDscpCode() float64 {
	if s == nil {
		return 0
	}

	return *s.DscpCode
}

func (s *PortConfQosProfileQosPoliciesQosMarking) GetIpPrecedenceCode() int64 {
	if s == nil {
		return 0
	}

	return *s.IpPrecedenceCode
}

func (s *PortConfQosProfileQosPoliciesQosMarking) GetQueue() int64 {
	if s == nil {
		return 0
	}

	return *s.Queue
}

type PortConfQosProfileQosPoliciesQosMatching struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *int64   `json:"dscp_code,omitempty"`
	DstPort          *float64 `json:"dst_port,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Protocol         *string  `json:"protocol,omitempty"`
	SrcPort          *float64 `json:"src_port,omitempty"`
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetCosCode() int64 {
	if s == nil {
		return 0
	}

	return *s.CosCode
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetDscpCode() int64 {
	if s == nil {
		return 0
	}

	return *s.DscpCode
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetDstPort() float64 {
	if s == nil {
		return 0
	}

	return *s.DstPort
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetIpPrecedenceCode() int64 {
	if s == nil {
		return 0
	}

	return *s.IpPrecedenceCode
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetProtocol() string {
	if s == nil {
		return ""
	}

	return *s.Protocol
}

func (s *PortConfQosProfileQosPoliciesQosMatching) GetSrcPort() float64 {
	if s == nil {
		return 0
	}

	return *s.SrcPort
}
