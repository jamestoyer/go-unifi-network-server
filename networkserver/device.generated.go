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

func (s *Device) GetAtfEnabled() bool {
	if s == nil {
		return false
	}

	return *s.AtfEnabled
}

func (s *Device) GetBandsteeringMode() string {
	if s == nil {
		return ""
	}

	return *s.BandsteeringMode
}

func (s *Device) GetBaresipAuthUser() string {
	if s == nil {
		return ""
	}

	return *s.BaresipAuthUser
}

func (s *Device) GetBaresipEnabled() bool {
	if s == nil {
		return false
	}

	return *s.BaresipEnabled
}

func (s *Device) GetBaresipExtension() string {
	if s == nil {
		return ""
	}

	return *s.BaresipExtension
}

func (s *Device) GetConfigNetwork() *DeviceConfigNetwork {
	if s == nil || s.ConfigNetwork == nil {
		return nil
	}

	return s.ConfigNetwork
}

func (s *Device) GetDisabled() bool {
	if s == nil {
		return false
	}

	return *s.Disabled
}

func (s *Device) GetDot1XFallbackNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.Dot1XFallbackNetworkconfId
}

func (s *Device) GetDot1XPortctrlEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Dot1XPortctrlEnabled
}

func (s *Device) GetDpiEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DpiEnabled
}

func (s *Device) GetEtherLighting() *DeviceEtherLighting {
	if s == nil || s.EtherLighting == nil {
		return nil
	}

	return s.EtherLighting
}

func (s *Device) GetEthernetOverrides() []DeviceEthernetOverrides {
	if s == nil || s.EthernetOverrides == nil {
		return nil
	}

	return *s.EthernetOverrides
}

func (s *Device) GetFlowctrlEnabled() bool {
	if s == nil {
		return false
	}

	return *s.FlowctrlEnabled
}

func (s *Device) GetGatewayVrrpMode() string {
	if s == nil {
		return ""
	}

	return *s.GatewayVrrpMode
}

func (s *Device) GetGatewayVrrpPriority() int64 {
	if s == nil {
		return 0
	}

	return *s.GatewayVrrpPriority
}

func (s *Device) GetGreenApEnabled() bool {
	if s == nil {
		return false
	}

	return *s.GreenApEnabled
}

func (s *Device) GetHeightInMeters() float64 {
	if s == nil {
		return 0
	}

	return *s.HeightInMeters
}

func (s *Device) GetHostname() string {
	if s == nil {
		return ""
	}

	return *s.Hostname
}

func (s *Device) GetJumboframeEnabled() bool {
	if s == nil {
		return false
	}

	return *s.JumboframeEnabled
}

func (s *Device) GetLcmBrightness() int64 {
	if s == nil {
		return 0
	}

	return *s.LcmBrightness
}

func (s *Device) GetLcmBrightnessOverride() bool {
	if s == nil {
		return false
	}

	return *s.LcmBrightnessOverride
}

func (s *Device) GetLcmIdleTimeout() float64 {
	if s == nil {
		return 0
	}

	return *s.LcmIdleTimeout
}

func (s *Device) GetLcmIdleTimeoutOverride() bool {
	if s == nil {
		return false
	}

	return *s.LcmIdleTimeoutOverride
}

func (s *Device) GetLcmNightModeBegins() string {
	if s == nil {
		return ""
	}

	return *s.LcmNightModeBegins
}

func (s *Device) GetLcmNightModeEnds() string {
	if s == nil {
		return ""
	}

	return *s.LcmNightModeEnds
}

func (s *Device) GetLcmOrientationOverride() int64 {
	if s == nil {
		return 0
	}

	return *s.LcmOrientationOverride
}

func (s *Device) GetLcmSettingsRestrictedAccess() bool {
	if s == nil {
		return false
	}

	return *s.LcmSettingsRestrictedAccess
}

func (s *Device) GetLcmTrackerEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LcmTrackerEnabled
}

func (s *Device) GetLcmTrackerSeed() string {
	if s == nil {
		return ""
	}

	return *s.LcmTrackerSeed
}

func (s *Device) GetLedOverride() string {
	if s == nil {
		return ""
	}

	return *s.LedOverride
}

func (s *Device) GetLedOverrideColor() string {
	if s == nil {
		return ""
	}

	return *s.LedOverrideColor
}

func (s *Device) GetLedOverrideColorBrightness() int64 {
	if s == nil {
		return 0
	}

	return *s.LedOverrideColorBrightness
}

func (s *Device) GetLocked() bool {
	if s == nil {
		return false
	}

	return *s.Locked
}

func (s *Device) GetLowpfmodeOverride() bool {
	if s == nil {
		return false
	}

	return *s.LowpfmodeOverride
}

func (s *Device) GetLteApn() string {
	if s == nil {
		return ""
	}

	return *s.LteApn
}

func (s *Device) GetLteAuthType() string {
	if s == nil {
		return ""
	}

	return *s.LteAuthType
}

func (s *Device) GetLteDataLimitEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LteDataLimitEnabled
}

func (s *Device) GetLteDataWarningEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LteDataWarningEnabled
}

func (s *Device) GetLteExtAnt() bool {
	if s == nil {
		return false
	}

	return *s.LteExtAnt
}

func (s *Device) GetLteHardLimit() int64 {
	if s == nil {
		return 0
	}

	return *s.LteHardLimit
}

func (s *Device) GetLtePassword() string {
	if s == nil {
		return ""
	}

	return *s.LtePassword
}

func (s *Device) GetLtePoe() bool {
	if s == nil {
		return false
	}

	return *s.LtePoe
}

func (s *Device) GetLteRoamingAllowed() bool {
	if s == nil {
		return false
	}

	return *s.LteRoamingAllowed
}

func (s *Device) GetLteSimPin() int64 {
	if s == nil {
		return 0
	}

	return *s.LteSimPin
}

func (s *Device) GetLteSoftLimit() int64 {
	if s == nil {
		return 0
	}

	return *s.LteSoftLimit
}

func (s *Device) GetLteUsername() string {
	if s == nil {
		return ""
	}

	return *s.LteUsername
}

func (s *Device) GetMapId() string {
	if s == nil {
		return ""
	}

	return *s.MapId
}

func (s *Device) GetMeshStaVapEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MeshStaVapEnabled
}

func (s *Device) GetMgmtNetworkId() string {
	if s == nil {
		return ""
	}

	return *s.MgmtNetworkId
}

func (s *Device) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Device) GetOutdoorModeOverride() string {
	if s == nil {
		return ""
	}

	return *s.OutdoorModeOverride
}

func (s *Device) GetOutletEnabled() bool {
	if s == nil {
		return false
	}

	return *s.OutletEnabled
}

func (s *Device) GetOutletOverrides() []DeviceOutletOverrides {
	if s == nil || s.OutletOverrides == nil {
		return nil
	}

	return *s.OutletOverrides
}

func (s *Device) GetOutletPowerCycleEnabled() bool {
	if s == nil {
		return false
	}

	return *s.OutletPowerCycleEnabled
}

func (s *Device) GetPeerToPeerMode() string {
	if s == nil {
		return ""
	}

	return *s.PeerToPeerMode
}

func (s *Device) GetPoeMode() string {
	if s == nil {
		return ""
	}

	return *s.PoeMode
}

func (s *Device) GetPortOverrides() []DevicePortOverrides {
	if s == nil || s.PortOverrides == nil {
		return nil
	}

	return *s.PortOverrides
}

func (s *Device) GetPowerSourceCtrl() string {
	if s == nil {
		return ""
	}

	return *s.PowerSourceCtrl
}

func (s *Device) GetPowerSourceCtrlBudget() int64 {
	if s == nil {
		return 0
	}

	return *s.PowerSourceCtrlBudget
}

func (s *Device) GetPowerSourceCtrlEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PowerSourceCtrlEnabled
}

func (s *Device) GetPtpApMac() string {
	if s == nil {
		return ""
	}

	return *s.PtpApMac
}

func (s *Device) GetRadioTable() []DeviceRadioTable {
	if s == nil || s.RadioTable == nil {
		return nil
	}

	return *s.RadioTable
}

func (s *Device) GetRadiusprofileId() string {
	if s == nil {
		return ""
	}

	return *s.RadiusprofileId
}

func (s *Device) GetResetbtnEnabled() string {
	if s == nil {
		return ""
	}

	return *s.ResetbtnEnabled
}

func (s *Device) GetRpsOverride() *DeviceRpsOverride {
	if s == nil || s.RpsOverride == nil {
		return nil
	}

	return s.RpsOverride
}

func (s *Device) GetSnmpContact() string {
	if s == nil {
		return ""
	}

	return *s.SnmpContact
}

func (s *Device) GetSnmpLocation() string {
	if s == nil {
		return ""
	}

	return *s.SnmpLocation
}

func (s *Device) GetStationMode() string {
	if s == nil {
		return ""
	}

	return *s.StationMode
}

func (s *Device) GetStpPriority() float64 {
	if s == nil {
		return 0
	}

	return *s.StpPriority
}

func (s *Device) GetStpVersion() string {
	if s == nil {
		return ""
	}

	return *s.StpVersion
}

func (s *Device) GetSwitchVlanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.SwitchVlanEnabled
}

func (s *Device) GetUbbPairName() string {
	if s == nil {
		return ""
	}

	return *s.UbbPairName
}

func (s *Device) GetVolume() int64 {
	if s == nil {
		return 0
	}

	return *s.Volume
}

func (s *Device) GetX() string {
	if s == nil {
		return ""
	}

	return *s.X
}

func (s *Device) GetXBaresipPassword() string {
	if s == nil {
		return ""
	}

	return *s.XBaresipPassword
}

func (s *Device) GetY() string {
	if s == nil {
		return ""
	}

	return *s.Y
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

func (s *DeviceConfigNetwork) GetBondingEnabled() bool {
	if s == nil {
		return false
	}

	return *s.BondingEnabled
}

func (s *DeviceConfigNetwork) GetDns1() string {
	if s == nil {
		return ""
	}

	return *s.Dns1
}

func (s *DeviceConfigNetwork) GetDns2() string {
	if s == nil {
		return ""
	}

	return *s.Dns2
}

func (s *DeviceConfigNetwork) GetDnssuffix() string {
	if s == nil {
		return ""
	}

	return *s.Dnssuffix
}

func (s *DeviceConfigNetwork) GetGateway() string {
	if s == nil {
		return ""
	}

	return *s.Gateway
}

func (s *DeviceConfigNetwork) GetIp() string {
	if s == nil {
		return ""
	}

	return *s.Ip
}

func (s *DeviceConfigNetwork) GetNetmask() string {
	if s == nil {
		return ""
	}

	return *s.Netmask
}

func (s *DeviceConfigNetwork) GetType() string {
	if s == nil {
		return ""
	}

	return *s.Type
}

type DeviceEtherLighting struct {
	Behavior   *string `json:"behavior,omitempty"`
	Brightness *int64  `json:"brightness,omitempty"`
	Mode       *string `json:"mode,omitempty"`
}

func (s *DeviceEtherLighting) GetBehavior() string {
	if s == nil {
		return ""
	}

	return *s.Behavior
}

func (s *DeviceEtherLighting) GetBrightness() int64 {
	if s == nil {
		return 0
	}

	return *s.Brightness
}

func (s *DeviceEtherLighting) GetMode() string {
	if s == nil {
		return ""
	}

	return *s.Mode
}

type DeviceEthernetOverrides struct {
	Ifname       *string `json:"ifname,omitempty"`
	Networkgroup *string `json:"networkgroup,omitempty"`
}

func (s *DeviceEthernetOverrides) GetIfname() string {
	if s == nil {
		return ""
	}

	return *s.Ifname
}

func (s *DeviceEthernetOverrides) GetNetworkgroup() string {
	if s == nil {
		return ""
	}

	return *s.Networkgroup
}

type DeviceOutletOverrides struct {
	CycleEnabled *bool   `json:"cycle_enabled,omitempty"`
	Index        *int64  `json:"index,omitempty"`
	Name         *string `json:"name,omitempty"`
	RelayState   *bool   `json:"relay_state,omitempty"`
}

func (s *DeviceOutletOverrides) GetCycleEnabled() bool {
	if s == nil {
		return false
	}

	return *s.CycleEnabled
}

func (s *DeviceOutletOverrides) GetIndex() int64 {
	if s == nil {
		return 0
	}

	return *s.Index
}

func (s *DeviceOutletOverrides) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DeviceOutletOverrides) GetRelayState() bool {
	if s == nil {
		return false
	}

	return *s.RelayState
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

func (s *DevicePortOverrides) GetAggregateNumPorts() int64 {
	if s == nil {
		return 0
	}

	return *s.AggregateNumPorts
}

func (s *DevicePortOverrides) GetAutoneg() bool {
	if s == nil {
		return false
	}

	return *s.Autoneg
}

func (s *DevicePortOverrides) GetDot1XCtrl() string {
	if s == nil {
		return ""
	}

	return *s.Dot1XCtrl
}

func (s *DevicePortOverrides) GetDot1XIdleTimeout() float64 {
	if s == nil {
		return 0
	}

	return *s.Dot1XIdleTimeout
}

func (s *DevicePortOverrides) GetEgressRateLimitKbps() int64 {
	if s == nil {
		return 0
	}

	return *s.EgressRateLimitKbps
}

func (s *DevicePortOverrides) GetEgressRateLimitKbpsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.EgressRateLimitKbpsEnabled
}

func (s *DevicePortOverrides) GetExcludedNetworkconfIds() []string {
	if s == nil || s.ExcludedNetworkconfIds == nil {
		return nil
	}

	return *s.ExcludedNetworkconfIds
}

func (s *DevicePortOverrides) GetFecMode() string {
	if s == nil {
		return ""
	}

	return *s.FecMode
}

func (s *DevicePortOverrides) GetForward() string {
	if s == nil {
		return ""
	}

	return *s.Forward
}

func (s *DevicePortOverrides) GetFullDuplex() bool {
	if s == nil {
		return false
	}

	return *s.FullDuplex
}

func (s *DevicePortOverrides) GetIsolation() bool {
	if s == nil {
		return false
	}

	return *s.Isolation
}

func (s *DevicePortOverrides) GetLldpmedEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LldpmedEnabled
}

func (s *DevicePortOverrides) GetLldpmedNotifyEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LldpmedNotifyEnabled
}

func (s *DevicePortOverrides) GetMirrorPortIdx() int64 {
	if s == nil {
		return 0
	}

	return *s.MirrorPortIdx
}

func (s *DevicePortOverrides) GetMulticastRouterNetworkconfIds() []string {
	if s == nil || s.MulticastRouterNetworkconfIds == nil {
		return nil
	}

	return *s.MulticastRouterNetworkconfIds
}

func (s *DevicePortOverrides) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DevicePortOverrides) GetNativeNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.NativeNetworkconfId
}

func (s *DevicePortOverrides) GetOpMode() string {
	if s == nil {
		return ""
	}

	return *s.OpMode
}

func (s *DevicePortOverrides) GetPoeMode() string {
	if s == nil {
		return ""
	}

	return *s.PoeMode
}

func (s *DevicePortOverrides) GetPortIdx() int64 {
	if s == nil {
		return 0
	}

	return *s.PortIdx
}

func (s *DevicePortOverrides) GetPortKeepaliveEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PortKeepaliveEnabled
}

func (s *DevicePortOverrides) GetPortSecurityEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PortSecurityEnabled
}

func (s *DevicePortOverrides) GetPortSecurityMacAddress() []string {
	if s == nil || s.PortSecurityMacAddress == nil {
		return nil
	}

	return *s.PortSecurityMacAddress
}

func (s *DevicePortOverrides) GetPortconfId() string {
	if s == nil {
		return ""
	}

	return *s.PortconfId
}

func (s *DevicePortOverrides) GetPriorityQueue1Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue1Level
}

func (s *DevicePortOverrides) GetPriorityQueue2Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue2Level
}

func (s *DevicePortOverrides) GetPriorityQueue3Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue3Level
}

func (s *DevicePortOverrides) GetPriorityQueue4Level() int64 {
	if s == nil {
		return 0
	}

	return *s.PriorityQueue4Level
}

func (s *DevicePortOverrides) GetQosProfile() *DevicePortOverridesQosProfile {
	if s == nil || s.QosProfile == nil {
		return nil
	}

	return s.QosProfile
}

func (s *DevicePortOverrides) GetSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.SettingPreference
}

func (s *DevicePortOverrides) GetSpeed() float64 {
	if s == nil {
		return 0
	}

	return *s.Speed
}

func (s *DevicePortOverrides) GetStormctrlBcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlBcastEnabled
}

func (s *DevicePortOverrides) GetStormctrlBcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlBcastLevel
}

func (s *DevicePortOverrides) GetStormctrlBcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlBcastRate
}

func (s *DevicePortOverrides) GetStormctrlMcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlMcastEnabled
}

func (s *DevicePortOverrides) GetStormctrlMcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlMcastLevel
}

func (s *DevicePortOverrides) GetStormctrlMcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlMcastRate
}

func (s *DevicePortOverrides) GetStormctrlType() string {
	if s == nil {
		return ""
	}

	return *s.StormctrlType
}

func (s *DevicePortOverrides) GetStormctrlUcastEnabled() bool {
	if s == nil {
		return false
	}

	return *s.StormctrlUcastEnabled
}

func (s *DevicePortOverrides) GetStormctrlUcastLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlUcastLevel
}

func (s *DevicePortOverrides) GetStormctrlUcastRate() float64 {
	if s == nil {
		return 0
	}

	return *s.StormctrlUcastRate
}

func (s *DevicePortOverrides) GetStpPortMode() bool {
	if s == nil {
		return false
	}

	return *s.StpPortMode
}

func (s *DevicePortOverrides) GetTaggedVlanMgmt() string {
	if s == nil {
		return ""
	}

	return *s.TaggedVlanMgmt
}

func (s *DevicePortOverrides) GetVoiceNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.VoiceNetworkconfId
}

type DevicePortOverridesQosProfile struct {
	QosPolicies    *[]DevicePortOverridesQosProfileQosPolicies `json:"qos_policies,omitempty"`
	QosProfileMode *string                                     `json:"qos_profile_mode,omitempty"`
}

func (s *DevicePortOverridesQosProfile) GetQosPolicies() []DevicePortOverridesQosProfileQosPolicies {
	if s == nil || s.QosPolicies == nil {
		return nil
	}

	return *s.QosPolicies
}

func (s *DevicePortOverridesQosProfile) GetQosProfileMode() string {
	if s == nil {
		return ""
	}

	return *s.QosProfileMode
}

type DevicePortOverridesQosProfileQosPolicies struct {
	QosMarking  *DevicePortOverridesQosProfileQosPoliciesQosMarking  `json:"qos_marking,omitempty"`
	QosMatching *DevicePortOverridesQosProfileQosPoliciesQosMatching `json:"qos_matching,omitempty"`
}

func (s *DevicePortOverridesQosProfileQosPolicies) GetQosMarking() *DevicePortOverridesQosProfileQosPoliciesQosMarking {
	if s == nil || s.QosMarking == nil {
		return nil
	}

	return s.QosMarking
}

func (s *DevicePortOverridesQosProfileQosPolicies) GetQosMatching() *DevicePortOverridesQosProfileQosPoliciesQosMatching {
	if s == nil || s.QosMatching == nil {
		return nil
	}

	return s.QosMatching
}

type DevicePortOverridesQosProfileQosPoliciesQosMarking struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *float64 `json:"dscp_code,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Queue            *int64   `json:"queue,omitempty"`
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMarking) GetCosCode() int64 {
	if s == nil {
		return 0
	}

	return *s.CosCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMarking) GetDscpCode() float64 {
	if s == nil {
		return 0
	}

	return *s.DscpCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMarking) GetIpPrecedenceCode() int64 {
	if s == nil {
		return 0
	}

	return *s.IpPrecedenceCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMarking) GetQueue() int64 {
	if s == nil {
		return 0
	}

	return *s.Queue
}

type DevicePortOverridesQosProfileQosPoliciesQosMatching struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *int64   `json:"dscp_code,omitempty"`
	DstPort          *float64 `json:"dst_port,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Protocol         *string  `json:"protocol,omitempty"`
	SrcPort          *float64 `json:"src_port,omitempty"`
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetCosCode() int64 {
	if s == nil {
		return 0
	}

	return *s.CosCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetDscpCode() int64 {
	if s == nil {
		return 0
	}

	return *s.DscpCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetDstPort() float64 {
	if s == nil {
		return 0
	}

	return *s.DstPort
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetIpPrecedenceCode() int64 {
	if s == nil {
		return 0
	}

	return *s.IpPrecedenceCode
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetProtocol() string {
	if s == nil {
		return ""
	}

	return *s.Protocol
}

func (s *DevicePortOverridesQosProfileQosPoliciesQosMatching) GetSrcPort() float64 {
	if s == nil {
		return 0
	}

	return *s.SrcPort
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

func (s *DeviceRadioTable) GetAntennaGain() int64 {
	if s == nil {
		return 0
	}

	return *s.AntennaGain
}

func (s *DeviceRadioTable) GetAntennaId() int64 {
	if s == nil {
		return 0
	}

	return *s.AntennaId
}

func (s *DeviceRadioTable) GetBackupChannel() string {
	if s == nil {
		return ""
	}

	return *s.BackupChannel
}

func (s *DeviceRadioTable) GetChannel() string {
	if s == nil {
		return ""
	}

	return *s.Channel
}

func (s *DeviceRadioTable) GetChannelOptimizationEnabled() bool {
	if s == nil {
		return false
	}

	return *s.ChannelOptimizationEnabled
}

func (s *DeviceRadioTable) GetHardNoiseFloorEnabled() bool {
	if s == nil {
		return false
	}

	return *s.HardNoiseFloorEnabled
}

func (s *DeviceRadioTable) GetHt() float64 {
	if s == nil {
		return 0
	}

	return *s.Ht
}

func (s *DeviceRadioTable) GetLoadbalanceEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LoadbalanceEnabled
}

func (s *DeviceRadioTable) GetMaxsta() int64 {
	if s == nil {
		return 0
	}

	return *s.Maxsta
}

func (s *DeviceRadioTable) GetMinRssi() int64 {
	if s == nil {
		return 0
	}

	return *s.MinRssi
}

func (s *DeviceRadioTable) GetMinRssiEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MinRssiEnabled
}

func (s *DeviceRadioTable) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DeviceRadioTable) GetRadio() string {
	if s == nil {
		return ""
	}

	return *s.Radio
}

func (s *DeviceRadioTable) GetRadioIdentifiers() []DeviceRadioTableRadioIdentifiers {
	if s == nil || s.RadioIdentifiers == nil {
		return nil
	}

	return *s.RadioIdentifiers
}

func (s *DeviceRadioTable) GetSensLevel() int64 {
	if s == nil {
		return 0
	}

	return *s.SensLevel
}

func (s *DeviceRadioTable) GetSensLevelEnabled() bool {
	if s == nil {
		return false
	}

	return *s.SensLevelEnabled
}

func (s *DeviceRadioTable) GetTxPower() string {
	if s == nil {
		return ""
	}

	return *s.TxPower
}

func (s *DeviceRadioTable) GetTxPowerMode() string {
	if s == nil {
		return ""
	}

	return *s.TxPowerMode
}

func (s *DeviceRadioTable) GetVwireEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VwireEnabled
}

type DeviceRadioTableRadioIdentifiers struct {
	DeviceId  *string `json:"device_id,omitempty"`
	RadioName *string `json:"radio_name,omitempty"`
}

func (s *DeviceRadioTableRadioIdentifiers) GetDeviceId() string {
	if s == nil {
		return ""
	}

	return *s.DeviceId
}

func (s *DeviceRadioTableRadioIdentifiers) GetRadioName() string {
	if s == nil {
		return ""
	}

	return *s.RadioName
}

type DeviceRpsOverride struct {
	PowerManagementMode *string                          `json:"power_management_mode,omitempty"`
	RpsPortTable        *[]DeviceRpsOverrideRpsPortTable `json:"rps_port_table,omitempty"`
}

func (s *DeviceRpsOverride) GetPowerManagementMode() string {
	if s == nil {
		return ""
	}

	return *s.PowerManagementMode
}

func (s *DeviceRpsOverride) GetRpsPortTable() []DeviceRpsOverrideRpsPortTable {
	if s == nil || s.RpsPortTable == nil {
		return nil
	}

	return *s.RpsPortTable
}

type DeviceRpsOverrideRpsPortTable struct {
	Name     *string `json:"name,omitempty"`
	PortIdx  *int64  `json:"port_idx,omitempty"`
	PortMode *string `json:"port_mode,omitempty"`
}

func (s *DeviceRpsOverrideRpsPortTable) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DeviceRpsOverrideRpsPortTable) GetPortIdx() int64 {
	if s == nil {
		return 0
	}

	return *s.PortIdx
}

func (s *DeviceRpsOverrideRpsPortTable) GetPortMode() string {
	if s == nil {
		return ""
	}

	return *s.PortMode
}
