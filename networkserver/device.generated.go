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

type Device struct {
	AtfEnabled                  *bool                      `json:"atf_enabled,omitempty"`
	BandsteeringMode            *string                    `json:"bandsteering_mode,omitempty"`
	BaresipAuthUser             *string                    `json:"baresip_auth_user,omitempty"`
	BaresipEnabled              *bool                      `json:"baresip_enabled,omitempty"`
	BaresipExtension            *string                    `json:"baresip_extension,omitempty"`
	ConfigNetwork               *DeviceConfigNetwork       `json:"config_network,omitempty"`
	Disabled                    *bool                      `json:"disabled,omitempty"`
	Dot1XFallbackNetworkconfId  *string                    `json:"dot1x_fallback_networkconf_id,omitempty"`
	Dot1XPortctrlEnabled        *bool                      `json:"dot1x_portctrl_enabled,omitempty"`
	DpiEnabled                  *bool                      `json:"dpi_enabled,omitempty"`
	EtherLighting               *DeviceEtherLighting       `json:"ether_lighting,omitempty"`
	EthernetOverrides           *[]DeviceEthernetOverrides `json:"ethernet_overrides,omitempty"`
	FlowctrlEnabled             *bool                      `json:"flowctrl_enabled,omitempty"`
	GatewayVrrpMode             *string                    `json:"gateway_vrrp_mode,omitempty"`
	GatewayVrrpPriority         *int64                     `json:"gateway_vrrp_priority,omitempty"`
	GreenApEnabled              *bool                      `json:"green_ap_enabled,omitempty"`
	HeightInMeters              *float64                   `json:"heightInMeters,omitempty"`
	Hostname                    *string                    `json:"hostname,omitempty"`
	JumboframeEnabled           *bool                      `json:"jumboframe_enabled,omitempty"`
	LcmBrightness               *int64                     `json:"lcm_brightness,omitempty"`
	LcmBrightnessOverride       *bool                      `json:"lcm_brightness_override,omitempty"`
	LcmIdleTimeout              *float64                   `json:"lcm_idle_timeout,omitempty"`
	LcmIdleTimeoutOverride      *bool                      `json:"lcm_idle_timeout_override,omitempty"`
	LcmNightModeBegins          *string                    `json:"lcm_night_mode_begins,omitempty"`
	LcmNightModeEnds            *string                    `json:"lcm_night_mode_ends,omitempty"`
	LcmOrientationOverride      *int64                     `json:"lcm_orientation_override,omitempty"`
	LcmSettingsRestrictedAccess *bool                      `json:"lcm_settings_restricted_access,omitempty"`
	LcmTrackerEnabled           *bool                      `json:"lcm_tracker_enabled,omitempty"`
	LcmTrackerSeed              *string                    `json:"lcm_tracker_seed,omitempty"`
	LedOverride                 *string                    `json:"led_override,omitempty"`
	LedOverrideColor            *string                    `json:"led_override_color,omitempty"`
	LedOverrideColorBrightness  *int64                     `json:"led_override_color_brightness,omitempty"`
	Locked                      *bool                      `json:"locked,omitempty"`
	LowpfmodeOverride           *bool                      `json:"lowpfmode_override,omitempty"`
	LteApn                      *string                    `json:"lte_apn,omitempty"`
	LteAuthType                 *string                    `json:"lte_auth_type,omitempty"`
	LteDataLimitEnabled         *bool                      `json:"lte_data_limit_enabled,omitempty"`
	LteDataWarningEnabled       *bool                      `json:"lte_data_warning_enabled,omitempty"`
	LteExtAnt                   *bool                      `json:"lte_ext_ant,omitempty"`
	LteHardLimit                *int64                     `json:"lte_hard_limit,omitempty"`
	LtePassword                 *string                    `json:"lte_password,omitempty"`
	LtePoe                      *bool                      `json:"lte_poe,omitempty"`
	LteRoamingAllowed           *bool                      `json:"lte_roaming_allowed,omitempty"`
	LteSimPin                   *int64                     `json:"lte_sim_pin,omitempty"`
	LteSoftLimit                *int64                     `json:"lte_soft_limit,omitempty"`
	LteUsername                 *string                    `json:"lte_username,omitempty"`
	MapId                       *string                    `json:"map_id,omitempty"`
	MeshStaVapEnabled           *bool                      `json:"mesh_sta_vap_enabled,omitempty"`
	MgmtNetworkId               *string                    `json:"mgmt_network_id,omitempty"`
	Name                        *string                    `json:"name,omitempty"`
	OutdoorModeOverride         *string                    `json:"outdoor_mode_override,omitempty"`
	OutletEnabled               *bool                      `json:"outlet_enabled,omitempty"`
	OutletOverrides             *[]DeviceOutletOverrides   `json:"outlet_overrides,omitempty"`
	OutletPowerCycleEnabled     *bool                      `json:"outlet_power_cycle_enabled,omitempty"`
	PeerToPeerMode              *string                    `json:"peer_to_peer_mode,omitempty"`
	PoeMode                     *string                    `json:"poe_mode,omitempty"`
	PortOverrides               *[]DevicePortOverrides     `json:"port_overrides,omitempty"`
	PowerSourceCtrl             *string                    `json:"power_source_ctrl,omitempty"`
	PowerSourceCtrlBudget       *int64                     `json:"power_source_ctrl_budget,omitempty"`
	PowerSourceCtrlEnabled      *bool                      `json:"power_source_ctrl_enabled,omitempty"`
	PtpApMac                    *string                    `json:"ptp_ap_mac,omitempty"`
	RadioTable                  *[]DeviceRadioTable        `json:"radio_table,omitempty"`
	RadiusprofileId             *string                    `json:"radiusprofile_id,omitempty"`
	ResetbtnEnabled             *string                    `json:"resetbtn_enabled,omitempty"`
	RpsOverride                 *DeviceRpsOverride         `json:"rps_override,omitempty"`
	SnmpContact                 *string                    `json:"snmp_contact,omitempty"`
	SnmpLocation                *string                    `json:"snmp_location,omitempty"`
	StationMode                 *string                    `json:"station_mode,omitempty"`
	StpPriority                 *float64                   `json:"stp_priority,omitempty"`
	StpVersion                  *string                    `json:"stp_version,omitempty"`
	SwitchVlanEnabled           *bool                      `json:"switch_vlan_enabled,omitempty"`
	UbbPairName                 *string                    `json:"ubb_pair_name,omitempty"`
	Volume                      *int64                     `json:"volume,omitempty"`
	X                           *string                    `json:"x,omitempty"`
	XBaresipPassword            *string                    `json:"x_baresip_password,omitempty"`
	Y                           *string                    `json:"y,omitempty"`
}

type DeviceConfigNetwork struct {
	BondingEnabled *bool   `json:"bonding_enabled,omitempty"`
	Dns1           *string `json:"dns1,omitempty"`
	Dns2           *string `json:"dns2,omitempty"`
	Dnssuffix      *string `json:"dnssuffix,omitempty"`
	Gateway        *string `json:"gateway,omitempty"`
	Ip             *string `json:"ip,omitempty"`
	Netmask        *string `json:"netmask,omitempty"`
	Type           *string `json:"type,omitempty"`
}

type DeviceEtherLighting struct {
	Behavior   *string `json:"behavior,omitempty"`
	Brightness *int64  `json:"brightness,omitempty"`
	Mode       *string `json:"mode,omitempty"`
}

type DeviceEthernetOverrides struct {
	Ifname       *string `json:"ifname,omitempty"`
	Networkgroup *string `json:"networkgroup,omitempty"`
}

type DeviceOutletOverrides struct {
	CycleEnabled *bool   `json:"cycle_enabled,omitempty"`
	Index        *int64  `json:"index,omitempty"`
	Name         *string `json:"name,omitempty"`
	RelayState   *bool   `json:"relay_state,omitempty"`
}

type DevicePortOverrides struct {
	AggregateNumPorts             *int64                         `json:"aggregate_num_ports,omitempty"`
	Autoneg                       *bool                          `json:"autoneg,omitempty"`
	Dot1XCtrl                     *string                        `json:"dot1x_ctrl,omitempty"`
	Dot1XIdleTimeout              *float64                       `json:"dot1x_idle_timeout,omitempty"`
	EgressRateLimitKbps           *int64                         `json:"egress_rate_limit_kbps,omitempty"`
	EgressRateLimitKbpsEnabled    *bool                          `json:"egress_rate_limit_kbps_enabled,omitempty"`
	ExcludedNetworkconfIds        *[]string                      `json:"excluded_networkconf_ids,omitempty"`
	FecMode                       *string                        `json:"fec_mode,omitempty"`
	Forward                       *string                        `json:"forward,omitempty"`
	FullDuplex                    *bool                          `json:"full_duplex,omitempty"`
	Isolation                     *bool                          `json:"isolation,omitempty"`
	LldpmedEnabled                *bool                          `json:"lldpmed_enabled,omitempty"`
	LldpmedNotifyEnabled          *bool                          `json:"lldpmed_notify_enabled,omitempty"`
	MirrorPortIdx                 *int64                         `json:"mirror_port_idx,omitempty"`
	MulticastRouterNetworkconfIds *[]string                      `json:"multicast_router_networkconf_ids,omitempty"`
	Name                          *string                        `json:"name,omitempty"`
	NativeNetworkconfId           *string                        `json:"native_networkconf_id,omitempty"`
	OpMode                        *string                        `json:"op_mode,omitempty"`
	PoeMode                       *string                        `json:"poe_mode,omitempty"`
	PortIdx                       *int64                         `json:"port_idx,omitempty"`
	PortKeepaliveEnabled          *bool                          `json:"port_keepalive_enabled,omitempty"`
	PortSecurityEnabled           *bool                          `json:"port_security_enabled,omitempty"`
	PortSecurityMacAddress        *[]string                      `json:"port_security_mac_address,omitempty"`
	PortconfId                    *string                        `json:"portconf_id,omitempty"`
	PriorityQueue1Level           *int64                         `json:"priority_queue1_level,omitempty"`
	PriorityQueue2Level           *int64                         `json:"priority_queue2_level,omitempty"`
	PriorityQueue3Level           *int64                         `json:"priority_queue3_level,omitempty"`
	PriorityQueue4Level           *int64                         `json:"priority_queue4_level,omitempty"`
	QosProfile                    *DevicePortOverridesQosProfile `json:"qos_profile,omitempty"`
	SettingPreference             *string                        `json:"setting_preference,omitempty"`
	Speed                         *float64                       `json:"speed,omitempty"`
	StormctrlBcastEnabled         *bool                          `json:"stormctrl_bcast_enabled,omitempty"`
	StormctrlBcastLevel           *int64                         `json:"stormctrl_bcast_level,omitempty"`
	StormctrlBcastRate            *float64                       `json:"stormctrl_bcast_rate,omitempty"`
	StormctrlMcastEnabled         *bool                          `json:"stormctrl_mcast_enabled,omitempty"`
	StormctrlMcastLevel           *int64                         `json:"stormctrl_mcast_level,omitempty"`
	StormctrlMcastRate            *float64                       `json:"stormctrl_mcast_rate,omitempty"`
	StormctrlType                 *string                        `json:"stormctrl_type,omitempty"`
	StormctrlUcastEnabled         *bool                          `json:"stormctrl_ucast_enabled,omitempty"`
	StormctrlUcastLevel           *int64                         `json:"stormctrl_ucast_level,omitempty"`
	StormctrlUcastRate            *float64                       `json:"stormctrl_ucast_rate,omitempty"`
	StpPortMode                   *bool                          `json:"stp_port_mode,omitempty"`
	TaggedVlanMgmt                *string                        `json:"tagged_vlan_mgmt,omitempty"`
	VoiceNetworkconfId            *string                        `json:"voice_networkconf_id,omitempty"`
}

type DevicePortOverridesQosProfile struct {
	QosPolicies    *[]DevicePortOverridesQosProfileQosPolicies `json:"qos_policies,omitempty"`
	QosProfileMode *string                                     `json:"qos_profile_mode,omitempty"`
}

type DevicePortOverridesQosProfileQosPolicies struct {
	QosMarking  *DevicePortOverridesQosProfileQosPoliciesQosMarking  `json:"qos_marking,omitempty"`
	QosMatching *DevicePortOverridesQosProfileQosPoliciesQosMatching `json:"qos_matching,omitempty"`
}

type DevicePortOverridesQosProfileQosPoliciesQosMarking struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *float64 `json:"dscp_code,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Queue            *int64   `json:"queue,omitempty"`
}

type DevicePortOverridesQosProfileQosPoliciesQosMatching struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *int64   `json:"dscp_code,omitempty"`
	DstPort          *float64 `json:"dst_port,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Protocol         *string  `json:"protocol,omitempty"`
	SrcPort          *float64 `json:"src_port,omitempty"`
}

type DeviceRadioTable struct {
	AntennaGain                *int64                              `json:"antenna_gain,omitempty"`
	AntennaId                  *int64                              `json:"antenna_id,omitempty"`
	BackupChannel              *string                             `json:"backup_channel,omitempty"`
	Channel                    *string                             `json:"channel,omitempty"`
	ChannelOptimizationEnabled *bool                               `json:"channel_optimization_enabled,omitempty"`
	HardNoiseFloorEnabled      *bool                               `json:"hard_noise_floor_enabled,omitempty"`
	Ht                         *float64                            `json:"ht,omitempty"`
	LoadbalanceEnabled         *bool                               `json:"loadbalance_enabled,omitempty"`
	Maxsta                     *int64                              `json:"maxsta,omitempty"`
	MinRssi                    *int64                              `json:"min_rssi,omitempty"`
	MinRssiEnabled             *bool                               `json:"min_rssi_enabled,omitempty"`
	Name                       *string                             `json:"name,omitempty"`
	Radio                      *string                             `json:"radio,omitempty"`
	RadioIdentifiers           *[]DeviceRadioTableRadioIdentifiers `json:"radio_identifiers,omitempty"`
	SensLevel                  *int64                              `json:"sens_level,omitempty"`
	SensLevelEnabled           *bool                               `json:"sens_level_enabled,omitempty"`
	TxPower                    *string                             `json:"tx_power,omitempty"`
	TxPowerMode                *string                             `json:"tx_power_mode,omitempty"`
	VwireEnabled               *bool                               `json:"vwire_enabled,omitempty"`
}

type DeviceRadioTableRadioIdentifiers struct {
	DeviceId  *string `json:"device_id,omitempty"`
	RadioName *string `json:"radio_name,omitempty"`
}

type DeviceRpsOverride struct {
	PowerManagementMode *string                          `json:"power_management_mode,omitempty"`
	RpsPortTable        *[]DeviceRpsOverrideRpsPortTable `json:"rps_port_table,omitempty"`
}

type DeviceRpsOverrideRpsPortTable struct {
	Name     *string `json:"name,omitempty"`
	PortIdx  *int64  `json:"port_idx,omitempty"`
	PortMode *string `json:"port_mode,omitempty"`
}
