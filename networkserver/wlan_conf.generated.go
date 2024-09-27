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

type WlanConf struct {
	ApGroupIds                  *[]string                       `json:"ap_group_ids,omitempty"`
	ApGroupMode                 *string                         `json:"ap_group_mode,omitempty"`
	AuthCache                   *bool                           `json:"auth_cache,omitempty"`
	BSupported                  *bool                           `json:"b_supported,omitempty"`
	BcFilterEnabled             *bool                           `json:"bc_filter_enabled,omitempty"`
	BcFilterList                *[]string                       `json:"bc_filter_list,omitempty"`
	BssTransition               *bool                           `json:"bss_transition,omitempty"`
	CountryBeacon               *bool                           `json:"country_beacon,omitempty"`
	DpiEnabled                  *bool                           `json:"dpi_enabled,omitempty"`
	DpigroupId                  *string                         `json:"dpigroup_id,omitempty"`
	Dtim6E                      *int64                          `json:"dtim_6e,omitempty"`
	DtimMode                    *string                         `json:"dtim_mode,omitempty"`
	DtimNa                      *int64                          `json:"dtim_na,omitempty"`
	DtimNg                      *int64                          `json:"dtim_ng,omitempty"`
	ElementAdopt                *bool                           `json:"element_adopt,omitempty"`
	Enabled                     *bool                           `json:"enabled,omitempty"`
	FastRoamingEnabled          *bool                           `json:"fast_roaming_enabled,omitempty"`
	GroupRekey                  *float64                        `json:"group_rekey,omitempty"`
	HideSsid                    *bool                           `json:"hide_ssid,omitempty"`
	Hotspot2                    *WlanConfHotspot2               `json:"hotspot2,omitempty"`
	Hotspot2ConfEnabled         *bool                           `json:"hotspot2conf_enabled,omitempty"`
	IappEnabled                 *bool                           `json:"iapp_enabled,omitempty"`
	IsGuest                     *bool                           `json:"is_guest,omitempty"`
	L2Isolation                 *bool                           `json:"l2_isolation,omitempty"`
	LogLevel                    *string                         `json:"log_level,omitempty"`
	MacFilterEnabled            *bool                           `json:"mac_filter_enabled,omitempty"`
	MacFilterList               *[]string                       `json:"mac_filter_list,omitempty"`
	MacFilterPolicy             *string                         `json:"mac_filter_policy,omitempty"`
	McastenhanceEnabled         *bool                           `json:"mcastenhance_enabled,omitempty"`
	MinrateNaAdvertisingRates   *bool                           `json:"minrate_na_advertising_rates,omitempty"`
	MinrateNaDataRateKbps       *int64                          `json:"minrate_na_data_rate_kbps,omitempty"`
	MinrateNaEnabled            *bool                           `json:"minrate_na_enabled,omitempty"`
	MinrateNgAdvertisingRates   *bool                           `json:"minrate_ng_advertising_rates,omitempty"`
	MinrateNgDataRateKbps       *int64                          `json:"minrate_ng_data_rate_kbps,omitempty"`
	MinrateNgEnabled            *bool                           `json:"minrate_ng_enabled,omitempty"`
	MinrateSettingPreference    *string                         `json:"minrate_setting_preference,omitempty"`
	MloEnabled                  *bool                           `json:"mlo_enabled,omitempty"`
	Name                        *string                         `json:"name,omitempty"`
	NameCombineEnabled          *bool                           `json:"name_combine_enabled,omitempty"`
	NameCombineSuffix           *string                         `json:"name_combine_suffix,omitempty"`
	NasIdentifier               *string                         `json:"nas_identifier,omitempty"`
	NasIdentifierType           *string                         `json:"nas_identifier_type,omitempty"`
	NetworkconfId               *string                         `json:"networkconf_id,omitempty"`
	No2GhzOui                   *bool                           `json:"no2ghz_oui,omitempty"`
	OptimizeIotWifiConnectivity *bool                           `json:"optimize_iot_wifi_connectivity,omitempty"`
	P2P                         *bool                           `json:"p2p,omitempty"`
	P2PCrossConnect             *bool                           `json:"p2p_cross_connect,omitempty"`
	PmfCipher                   *string                         `json:"pmf_cipher,omitempty"`
	PmfMode                     *string                         `json:"pmf_mode,omitempty"`
	Priority                    *string                         `json:"priority,omitempty"`
	PrivatePresharedKeys        *[]WlanConfPrivatePresharedKeys `json:"private_preshared_keys,omitempty"`
	PrivatePresharedKeysEnabled *bool                           `json:"private_preshared_keys_enabled,omitempty"`
	ProxyArp                    *bool                           `json:"proxy_arp,omitempty"`
	RadiusDasEnabled            *bool                           `json:"radius_das_enabled,omitempty"`
	RadiusMacAuthEnabled        *bool                           `json:"radius_mac_auth_enabled,omitempty"`
	RadiusMacaclEmptyPassword   *bool                           `json:"radius_macacl_empty_password,omitempty"`
	RadiusMacaclFormat          *string                         `json:"radius_macacl_format,omitempty"`
	RadiusprofileId             *string                         `json:"radiusprofile_id,omitempty"`
	RoamClusterId               *int64                          `json:"roam_cluster_id,omitempty"`
	RrmEnabled                  *bool                           `json:"rrm_enabled,omitempty"`
	SaeAntiClogging             *int64                          `json:"sae_anti_clogging,omitempty"`
	SaeGroups                   *[]int64                        `json:"sae_groups,omitempty"`
	SaePsk                      *[]WlanConfSaePsk               `json:"sae_psk,omitempty"`
	SaePskVlanRequired          *bool                           `json:"sae_psk_vlan_required,omitempty"`
	SaeSync                     *int64                          `json:"sae_sync,omitempty"`
	Schedule                    *[]string                       `json:"schedule,omitempty"`
	ScheduleEnabled             *bool                           `json:"schedule_enabled,omitempty"`
	ScheduleReversed            *bool                           `json:"schedule_reversed,omitempty"`
	ScheduleWithDuration        *[]WlanConfScheduleWithDuration `json:"schedule_with_duration,omitempty"`
	Security                    *string                         `json:"security,omitempty"`
	SettingPreference           *string                         `json:"setting_preference,omitempty"`
	TdlsProhibit                *bool                           `json:"tdls_prohibit,omitempty"`
	UapsdEnabled                *bool                           `json:"uapsd_enabled,omitempty"`
	UidWorkspaceUrl             *string                         `json:"uid_workspace_url,omitempty"`
	UsergroupId                 *string                         `json:"usergroup_id,omitempty"`
	Vlan                        *float64                        `json:"vlan,omitempty"`
	VlanEnabled                 *bool                           `json:"vlan_enabled,omitempty"`
	WepIdx                      *int64                          `json:"wep_idx,omitempty"`
	WlanBand                    *string                         `json:"wlan_band,omitempty"`
	WlanBands                   *[]string                       `json:"wlan_bands,omitempty"`
	Wpa3Enhanced192             *bool                           `json:"wpa3_enhanced_192,omitempty"`
	Wpa3FastRoaming             *bool                           `json:"wpa3_fast_roaming,omitempty"`
	Wpa3Support                 *bool                           `json:"wpa3_support,omitempty"`
	Wpa3Transition              *bool                           `json:"wpa3_transition,omitempty"`
	WpaEnc                      *string                         `json:"wpa_enc,omitempty"`
	WpaMode                     *string                         `json:"wpa_mode,omitempty"`
	WpaPskRadius                *string                         `json:"wpa_psk_radius,omitempty"`
	XIappKey                    *string                         `json:"x_iapp_key,omitempty"`
	XPassphrase                 *string                         `json:"x_passphrase,omitempty"`
	XWep                        *string                         `json:"x_wep,omitempty"`
}

func (s *WlanConf) GetApGroupIds() []string {
	if s == nil || s.ApGroupIds == nil {
		return nil
	}

	return *s.ApGroupIds
}

func (s *WlanConf) GetApGroupMode() string {
	if s == nil {
		return ""
	}

	return *s.ApGroupMode
}

func (s *WlanConf) GetAuthCache() bool {
	if s == nil {
		return false
	}

	return *s.AuthCache
}

func (s *WlanConf) GetBSupported() bool {
	if s == nil {
		return false
	}

	return *s.BSupported
}

func (s *WlanConf) GetBcFilterEnabled() bool {
	if s == nil {
		return false
	}

	return *s.BcFilterEnabled
}

func (s *WlanConf) GetBcFilterList() []string {
	if s == nil || s.BcFilterList == nil {
		return nil
	}

	return *s.BcFilterList
}

func (s *WlanConf) GetBssTransition() bool {
	if s == nil {
		return false
	}

	return *s.BssTransition
}

func (s *WlanConf) GetCountryBeacon() bool {
	if s == nil {
		return false
	}

	return *s.CountryBeacon
}

func (s *WlanConf) GetDpiEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DpiEnabled
}

func (s *WlanConf) GetDpigroupId() string {
	if s == nil {
		return ""
	}

	return *s.DpigroupId
}

func (s *WlanConf) GetDtim6E() int64 {
	if s == nil {
		return 0
	}

	return *s.Dtim6E
}

func (s *WlanConf) GetDtimMode() string {
	if s == nil {
		return ""
	}

	return *s.DtimMode
}

func (s *WlanConf) GetDtimNa() int64 {
	if s == nil {
		return 0
	}

	return *s.DtimNa
}

func (s *WlanConf) GetDtimNg() int64 {
	if s == nil {
		return 0
	}

	return *s.DtimNg
}

func (s *WlanConf) GetElementAdopt() bool {
	if s == nil {
		return false
	}

	return *s.ElementAdopt
}

func (s *WlanConf) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *WlanConf) GetFastRoamingEnabled() bool {
	if s == nil {
		return false
	}

	return *s.FastRoamingEnabled
}

func (s *WlanConf) GetGroupRekey() float64 {
	if s == nil {
		return 0
	}

	return *s.GroupRekey
}

func (s *WlanConf) GetHideSsid() bool {
	if s == nil {
		return false
	}

	return *s.HideSsid
}

func (s *WlanConf) GetHotspot2() *WlanConfHotspot2 {
	if s == nil || s.Hotspot2 == nil {
		return nil
	}

	return s.Hotspot2
}

func (s *WlanConf) GetHotspot2ConfEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Hotspot2ConfEnabled
}

func (s *WlanConf) GetIappEnabled() bool {
	if s == nil {
		return false
	}

	return *s.IappEnabled
}

func (s *WlanConf) GetIsGuest() bool {
	if s == nil {
		return false
	}

	return *s.IsGuest
}

func (s *WlanConf) GetL2Isolation() bool {
	if s == nil {
		return false
	}

	return *s.L2Isolation
}

func (s *WlanConf) GetLogLevel() string {
	if s == nil {
		return ""
	}

	return *s.LogLevel
}

func (s *WlanConf) GetMacFilterEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MacFilterEnabled
}

func (s *WlanConf) GetMacFilterList() []string {
	if s == nil || s.MacFilterList == nil {
		return nil
	}

	return *s.MacFilterList
}

func (s *WlanConf) GetMacFilterPolicy() string {
	if s == nil {
		return ""
	}

	return *s.MacFilterPolicy
}

func (s *WlanConf) GetMcastenhanceEnabled() bool {
	if s == nil {
		return false
	}

	return *s.McastenhanceEnabled
}

func (s *WlanConf) GetMinrateNaAdvertisingRates() bool {
	if s == nil {
		return false
	}

	return *s.MinrateNaAdvertisingRates
}

func (s *WlanConf) GetMinrateNaDataRateKbps() int64 {
	if s == nil {
		return 0
	}

	return *s.MinrateNaDataRateKbps
}

func (s *WlanConf) GetMinrateNaEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MinrateNaEnabled
}

func (s *WlanConf) GetMinrateNgAdvertisingRates() bool {
	if s == nil {
		return false
	}

	return *s.MinrateNgAdvertisingRates
}

func (s *WlanConf) GetMinrateNgDataRateKbps() int64 {
	if s == nil {
		return 0
	}

	return *s.MinrateNgDataRateKbps
}

func (s *WlanConf) GetMinrateNgEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MinrateNgEnabled
}

func (s *WlanConf) GetMinrateSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.MinrateSettingPreference
}

func (s *WlanConf) GetMloEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MloEnabled
}

func (s *WlanConf) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *WlanConf) GetNameCombineEnabled() bool {
	if s == nil {
		return false
	}

	return *s.NameCombineEnabled
}

func (s *WlanConf) GetNameCombineSuffix() string {
	if s == nil {
		return ""
	}

	return *s.NameCombineSuffix
}

func (s *WlanConf) GetNasIdentifier() string {
	if s == nil {
		return ""
	}

	return *s.NasIdentifier
}

func (s *WlanConf) GetNasIdentifierType() string {
	if s == nil {
		return ""
	}

	return *s.NasIdentifierType
}

func (s *WlanConf) GetNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.NetworkconfId
}

func (s *WlanConf) GetNo2GhzOui() bool {
	if s == nil {
		return false
	}

	return *s.No2GhzOui
}

func (s *WlanConf) GetOptimizeIotWifiConnectivity() bool {
	if s == nil {
		return false
	}

	return *s.OptimizeIotWifiConnectivity
}

func (s *WlanConf) GetP2P() bool {
	if s == nil {
		return false
	}

	return *s.P2P
}

func (s *WlanConf) GetP2PCrossConnect() bool {
	if s == nil {
		return false
	}

	return *s.P2PCrossConnect
}

func (s *WlanConf) GetPmfCipher() string {
	if s == nil {
		return ""
	}

	return *s.PmfCipher
}

func (s *WlanConf) GetPmfMode() string {
	if s == nil {
		return ""
	}

	return *s.PmfMode
}

func (s *WlanConf) GetPriority() string {
	if s == nil {
		return ""
	}

	return *s.Priority
}

func (s *WlanConf) GetPrivatePresharedKeys() []WlanConfPrivatePresharedKeys {
	if s == nil || s.PrivatePresharedKeys == nil {
		return nil
	}

	return *s.PrivatePresharedKeys
}

func (s *WlanConf) GetPrivatePresharedKeysEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PrivatePresharedKeysEnabled
}

func (s *WlanConf) GetProxyArp() bool {
	if s == nil {
		return false
	}

	return *s.ProxyArp
}

func (s *WlanConf) GetRadiusDasEnabled() bool {
	if s == nil {
		return false
	}

	return *s.RadiusDasEnabled
}

func (s *WlanConf) GetRadiusMacAuthEnabled() bool {
	if s == nil {
		return false
	}

	return *s.RadiusMacAuthEnabled
}

func (s *WlanConf) GetRadiusMacaclEmptyPassword() bool {
	if s == nil {
		return false
	}

	return *s.RadiusMacaclEmptyPassword
}

func (s *WlanConf) GetRadiusMacaclFormat() string {
	if s == nil {
		return ""
	}

	return *s.RadiusMacaclFormat
}

func (s *WlanConf) GetRadiusprofileId() string {
	if s == nil {
		return ""
	}

	return *s.RadiusprofileId
}

func (s *WlanConf) GetRoamClusterId() int64 {
	if s == nil {
		return 0
	}

	return *s.RoamClusterId
}

func (s *WlanConf) GetRrmEnabled() bool {
	if s == nil {
		return false
	}

	return *s.RrmEnabled
}

func (s *WlanConf) GetSaeAntiClogging() int64 {
	if s == nil {
		return 0
	}

	return *s.SaeAntiClogging
}

func (s *WlanConf) GetSaeGroups() []int64 {
	if s == nil || s.SaeGroups == nil {
		return nil
	}

	return *s.SaeGroups
}

func (s *WlanConf) GetSaePsk() []WlanConfSaePsk {
	if s == nil || s.SaePsk == nil {
		return nil
	}

	return *s.SaePsk
}

func (s *WlanConf) GetSaePskVlanRequired() bool {
	if s == nil {
		return false
	}

	return *s.SaePskVlanRequired
}

func (s *WlanConf) GetSaeSync() int64 {
	if s == nil {
		return 0
	}

	return *s.SaeSync
}

func (s *WlanConf) GetSchedule() []string {
	if s == nil || s.Schedule == nil {
		return nil
	}

	return *s.Schedule
}

func (s *WlanConf) GetScheduleEnabled() bool {
	if s == nil {
		return false
	}

	return *s.ScheduleEnabled
}

func (s *WlanConf) GetScheduleReversed() bool {
	if s == nil {
		return false
	}

	return *s.ScheduleReversed
}

func (s *WlanConf) GetScheduleWithDuration() []WlanConfScheduleWithDuration {
	if s == nil || s.ScheduleWithDuration == nil {
		return nil
	}

	return *s.ScheduleWithDuration
}

func (s *WlanConf) GetSecurity() string {
	if s == nil {
		return ""
	}

	return *s.Security
}

func (s *WlanConf) GetSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.SettingPreference
}

func (s *WlanConf) GetTdlsProhibit() bool {
	if s == nil {
		return false
	}

	return *s.TdlsProhibit
}

func (s *WlanConf) GetUapsdEnabled() bool {
	if s == nil {
		return false
	}

	return *s.UapsdEnabled
}

func (s *WlanConf) GetUidWorkspaceUrl() string {
	if s == nil {
		return ""
	}

	return *s.UidWorkspaceUrl
}

func (s *WlanConf) GetUsergroupId() string {
	if s == nil {
		return ""
	}

	return *s.UsergroupId
}

func (s *WlanConf) GetVlan() float64 {
	if s == nil {
		return 0
	}

	return *s.Vlan
}

func (s *WlanConf) GetVlanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VlanEnabled
}

func (s *WlanConf) GetWepIdx() int64 {
	if s == nil {
		return 0
	}

	return *s.WepIdx
}

func (s *WlanConf) GetWlanBand() string {
	if s == nil {
		return ""
	}

	return *s.WlanBand
}

func (s *WlanConf) GetWlanBands() []string {
	if s == nil || s.WlanBands == nil {
		return nil
	}

	return *s.WlanBands
}

func (s *WlanConf) GetWpa3Enhanced192() bool {
	if s == nil {
		return false
	}

	return *s.Wpa3Enhanced192
}

func (s *WlanConf) GetWpa3FastRoaming() bool {
	if s == nil {
		return false
	}

	return *s.Wpa3FastRoaming
}

func (s *WlanConf) GetWpa3Support() bool {
	if s == nil {
		return false
	}

	return *s.Wpa3Support
}

func (s *WlanConf) GetWpa3Transition() bool {
	if s == nil {
		return false
	}

	return *s.Wpa3Transition
}

func (s *WlanConf) GetWpaEnc() string {
	if s == nil {
		return ""
	}

	return *s.WpaEnc
}

func (s *WlanConf) GetWpaMode() string {
	if s == nil {
		return ""
	}

	return *s.WpaMode
}

func (s *WlanConf) GetWpaPskRadius() string {
	if s == nil {
		return ""
	}

	return *s.WpaPskRadius
}

func (s *WlanConf) GetXIappKey() string {
	if s == nil {
		return ""
	}

	return *s.XIappKey
}

func (s *WlanConf) GetXPassphrase() string {
	if s == nil {
		return ""
	}

	return *s.XPassphrase
}

func (s *WlanConf) GetXWep() string {
	if s == nil {
		return ""
	}

	return *s.XWep
}

type WlanConfHotspot2 struct {
	Capab                   *[]WlanConfHotspot2Capab                 `json:"capab,omitempty"`
	CellularNetworkList     *[]WlanConfHotspot2CellularNetworkList   `json:"cellular_network_list,omitempty"`
	DomainNameList          *[]string                                `json:"domain_name_list,omitempty"`
	FriendlyName            *[]WlanConfHotspot2FriendlyName          `json:"friendly_name,omitempty"`
	IpaddrTypeAvailV4       *int64                                   `json:"ipaddr_type_avail_v4,omitempty"`
	IpaddrTypeAvailV6       *int64                                   `json:"ipaddr_type_avail_v6,omitempty"`
	MetricsDownlinkLoad     *int64                                   `json:"metrics_downlink_load,omitempty"`
	MetricsDownlinkLoadSet  *bool                                    `json:"metrics_downlink_load_set,omitempty"`
	MetricsDownlinkSpeed    *int64                                   `json:"metrics_downlink_speed,omitempty"`
	MetricsDownlinkSpeedSet *bool                                    `json:"metrics_downlink_speed_set,omitempty"`
	MetricsInfoAtCapacity   *bool                                    `json:"metrics_info_at_capacity,omitempty"`
	MetricsInfoLinkStatus   *string                                  `json:"metrics_info_link_status,omitempty"`
	MetricsInfoSymmetric    *bool                                    `json:"metrics_info_symmetric,omitempty"`
	MetricsMeasurement      *int64                                   `json:"metrics_measurement,omitempty"`
	MetricsMeasurementSet   *bool                                    `json:"metrics_measurement_set,omitempty"`
	MetricsStatus           *bool                                    `json:"metrics_status,omitempty"`
	MetricsUplinkLoad       *int64                                   `json:"metrics_uplink_load,omitempty"`
	MetricsUplinkLoadSet    *bool                                    `json:"metrics_uplink_load_set,omitempty"`
	MetricsUplinkSpeed      *int64                                   `json:"metrics_uplink_speed,omitempty"`
	MetricsUplinkSpeedSet   *bool                                    `json:"metrics_uplink_speed_set,omitempty"`
	NaiRealmList            *[]WlanConfHotspot2NaiRealmList          `json:"nai_realm_list,omitempty"`
	NetworkType             *int64                                   `json:"network_type,omitempty"`
	RoamingConsortiumList   *[]WlanConfHotspot2RoamingConsortiumList `json:"roaming_consortium_list,omitempty"`
	VenueGroup              *int64                                   `json:"venue_group,omitempty"`
	VenueName               *[]WlanConfHotspot2VenueName             `json:"venue_name,omitempty"`
	VenueType               *float64                                 `json:"venue_type,omitempty"`
}

func (s *WlanConfHotspot2) GetCapab() []WlanConfHotspot2Capab {
	if s == nil || s.Capab == nil {
		return nil
	}

	return *s.Capab
}

func (s *WlanConfHotspot2) GetCellularNetworkList() []WlanConfHotspot2CellularNetworkList {
	if s == nil || s.CellularNetworkList == nil {
		return nil
	}

	return *s.CellularNetworkList
}

func (s *WlanConfHotspot2) GetDomainNameList() []string {
	if s == nil || s.DomainNameList == nil {
		return nil
	}

	return *s.DomainNameList
}

func (s *WlanConfHotspot2) GetFriendlyName() []WlanConfHotspot2FriendlyName {
	if s == nil || s.FriendlyName == nil {
		return nil
	}

	return *s.FriendlyName
}

func (s *WlanConfHotspot2) GetIpaddrTypeAvailV4() int64 {
	if s == nil {
		return 0
	}

	return *s.IpaddrTypeAvailV4
}

func (s *WlanConfHotspot2) GetIpaddrTypeAvailV6() int64 {
	if s == nil {
		return 0
	}

	return *s.IpaddrTypeAvailV6
}

func (s *WlanConfHotspot2) GetMetricsDownlinkLoad() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsDownlinkLoad
}

func (s *WlanConfHotspot2) GetMetricsDownlinkLoadSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsDownlinkLoadSet
}

func (s *WlanConfHotspot2) GetMetricsDownlinkSpeed() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsDownlinkSpeed
}

func (s *WlanConfHotspot2) GetMetricsDownlinkSpeedSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsDownlinkSpeedSet
}

func (s *WlanConfHotspot2) GetMetricsInfoAtCapacity() bool {
	if s == nil {
		return false
	}

	return *s.MetricsInfoAtCapacity
}

func (s *WlanConfHotspot2) GetMetricsInfoLinkStatus() string {
	if s == nil {
		return ""
	}

	return *s.MetricsInfoLinkStatus
}

func (s *WlanConfHotspot2) GetMetricsInfoSymmetric() bool {
	if s == nil {
		return false
	}

	return *s.MetricsInfoSymmetric
}

func (s *WlanConfHotspot2) GetMetricsMeasurement() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsMeasurement
}

func (s *WlanConfHotspot2) GetMetricsMeasurementSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsMeasurementSet
}

func (s *WlanConfHotspot2) GetMetricsStatus() bool {
	if s == nil {
		return false
	}

	return *s.MetricsStatus
}

func (s *WlanConfHotspot2) GetMetricsUplinkLoad() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsUplinkLoad
}

func (s *WlanConfHotspot2) GetMetricsUplinkLoadSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsUplinkLoadSet
}

func (s *WlanConfHotspot2) GetMetricsUplinkSpeed() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsUplinkSpeed
}

func (s *WlanConfHotspot2) GetMetricsUplinkSpeedSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsUplinkSpeedSet
}

func (s *WlanConfHotspot2) GetNaiRealmList() []WlanConfHotspot2NaiRealmList {
	if s == nil || s.NaiRealmList == nil {
		return nil
	}

	return *s.NaiRealmList
}

func (s *WlanConfHotspot2) GetNetworkType() int64 {
	if s == nil {
		return 0
	}

	return *s.NetworkType
}

func (s *WlanConfHotspot2) GetRoamingConsortiumList() []WlanConfHotspot2RoamingConsortiumList {
	if s == nil || s.RoamingConsortiumList == nil {
		return nil
	}

	return *s.RoamingConsortiumList
}

func (s *WlanConfHotspot2) GetVenueGroup() int64 {
	if s == nil {
		return 0
	}

	return *s.VenueGroup
}

func (s *WlanConfHotspot2) GetVenueName() []WlanConfHotspot2VenueName {
	if s == nil || s.VenueName == nil {
		return nil
	}

	return *s.VenueName
}

func (s *WlanConfHotspot2) GetVenueType() float64 {
	if s == nil {
		return 0
	}

	return *s.VenueType
}

type WlanConfHotspot2Capab struct {
	Port     *float64 `json:"port,omitempty"`
	Protocol *string  `json:"protocol,omitempty"`
	Status   *string  `json:"status,omitempty"`
}

func (s *WlanConfHotspot2Capab) GetPort() float64 {
	if s == nil {
		return 0
	}

	return *s.Port
}

func (s *WlanConfHotspot2Capab) GetProtocol() string {
	if s == nil {
		return ""
	}

	return *s.Protocol
}

func (s *WlanConfHotspot2Capab) GetStatus() string {
	if s == nil {
		return ""
	}

	return *s.Status
}

type WlanConfHotspot2CellularNetworkList struct {
	CountryCode *int64  `json:"country_code,omitempty"`
	Mcc         *int64  `json:"mcc,omitempty"`
	Mnc         *int64  `json:"mnc,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (s *WlanConfHotspot2CellularNetworkList) GetCountryCode() int64 {
	if s == nil {
		return 0
	}

	return *s.CountryCode
}

func (s *WlanConfHotspot2CellularNetworkList) GetMcc() int64 {
	if s == nil {
		return 0
	}

	return *s.Mcc
}

func (s *WlanConfHotspot2CellularNetworkList) GetMnc() int64 {
	if s == nil {
		return 0
	}

	return *s.Mnc
}

func (s *WlanConfHotspot2CellularNetworkList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type WlanConfHotspot2FriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

func (s *WlanConfHotspot2FriendlyName) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *WlanConfHotspot2FriendlyName) GetText() string {
	if s == nil {
		return ""
	}

	return *s.Text
}

type WlanConfHotspot2NaiRealmList struct {
	AuthIds   *[]int64 `json:"auth_ids,omitempty"`
	AuthVals  *[]int64 `json:"auth_vals,omitempty"`
	EapMethod *int64   `json:"eap_method,omitempty"`
	Encoding  *int64   `json:"encoding,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Status    *bool    `json:"status,omitempty"`
}

func (s *WlanConfHotspot2NaiRealmList) GetAuthIds() []int64 {
	if s == nil || s.AuthIds == nil {
		return nil
	}

	return *s.AuthIds
}

func (s *WlanConfHotspot2NaiRealmList) GetAuthVals() []int64 {
	if s == nil || s.AuthVals == nil {
		return nil
	}

	return *s.AuthVals
}

func (s *WlanConfHotspot2NaiRealmList) GetEapMethod() int64 {
	if s == nil {
		return 0
	}

	return *s.EapMethod
}

func (s *WlanConfHotspot2NaiRealmList) GetEncoding() int64 {
	if s == nil {
		return 0
	}

	return *s.Encoding
}

func (s *WlanConfHotspot2NaiRealmList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *WlanConfHotspot2NaiRealmList) GetStatus() bool {
	if s == nil {
		return false
	}

	return *s.Status
}

type WlanConfHotspot2RoamingConsortiumList struct {
	Name *string `json:"name,omitempty"`
	Oid  *string `json:"oid,omitempty"`
}

func (s *WlanConfHotspot2RoamingConsortiumList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *WlanConfHotspot2RoamingConsortiumList) GetOid() string {
	if s == nil {
		return ""
	}

	return *s.Oid
}

type WlanConfHotspot2VenueName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}

func (s *WlanConfHotspot2VenueName) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *WlanConfHotspot2VenueName) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *WlanConfHotspot2VenueName) GetUrl() string {
	if s == nil {
		return ""
	}

	return *s.Url
}

type WlanConfPrivatePresharedKeys struct {
	NetworkconfId *string `json:"networkconf_id,omitempty"`
	Password      *string `json:"password,omitempty"`
}

func (s *WlanConfPrivatePresharedKeys) GetNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.NetworkconfId
}

func (s *WlanConfPrivatePresharedKeys) GetPassword() string {
	if s == nil {
		return ""
	}

	return *s.Password
}

type WlanConfSaePsk struct {
	Id   *string  `json:"id,omitempty"`
	Mac  *string  `json:"mac,omitempty"`
	Psk  *string  `json:"psk,omitempty"`
	Vlan *float64 `json:"vlan,omitempty"`
}

func (s *WlanConfSaePsk) GetId() string {
	if s == nil {
		return ""
	}

	return *s.Id
}

func (s *WlanConfSaePsk) GetMac() string {
	if s == nil {
		return ""
	}

	return *s.Mac
}

func (s *WlanConfSaePsk) GetPsk() string {
	if s == nil {
		return ""
	}

	return *s.Psk
}

func (s *WlanConfSaePsk) GetVlan() float64 {
	if s == nil {
		return 0
	}

	return *s.Vlan
}

type WlanConfScheduleWithDuration struct {
	DurationMinutes *int64    `json:"duration_minutes,omitempty"`
	Name            *string   `json:"name,omitempty"`
	StartDaysOfWeek *[]string `json:"start_days_of_week,omitempty"`
	StartHour       *int64    `json:"start_hour,omitempty"`
	StartMinute     *int64    `json:"start_minute,omitempty"`
}

func (s *WlanConfScheduleWithDuration) GetDurationMinutes() int64 {
	if s == nil {
		return 0
	}

	return *s.DurationMinutes
}

func (s *WlanConfScheduleWithDuration) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *WlanConfScheduleWithDuration) GetStartDaysOfWeek() []string {
	if s == nil || s.StartDaysOfWeek == nil {
		return nil
	}

	return *s.StartDaysOfWeek
}

func (s *WlanConfScheduleWithDuration) GetStartHour() int64 {
	if s == nil {
		return 0
	}

	return *s.StartHour
}

func (s *WlanConfScheduleWithDuration) GetStartMinute() int64 {
	if s == nil {
		return 0
	}

	return *s.StartMinute
}
