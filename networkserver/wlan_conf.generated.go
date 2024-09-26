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

type WlanConfHotspot2Capab struct {
	Port     *float64 `json:"port,omitempty"`
	Protocol *string  `json:"protocol,omitempty"`
	Status   *string  `json:"status,omitempty"`
}

type WlanConfHotspot2CellularNetworkList struct {
	CountryCode *int64  `json:"country_code,omitempty"`
	Mcc         *int64  `json:"mcc,omitempty"`
	Mnc         *int64  `json:"mnc,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type WlanConfHotspot2FriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type WlanConfHotspot2NaiRealmList struct {
	AuthIds   *[]int64 `json:"auth_ids,omitempty"`
	AuthVals  *[]int64 `json:"auth_vals,omitempty"`
	EapMethod *int64   `json:"eap_method,omitempty"`
	Encoding  *int64   `json:"encoding,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Status    *bool    `json:"status,omitempty"`
}

type WlanConfHotspot2RoamingConsortiumList struct {
	Name *string `json:"name,omitempty"`
	Oid  *string `json:"oid,omitempty"`
}

type WlanConfHotspot2VenueName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}

type WlanConfPrivatePresharedKeys struct {
	NetworkconfId *string `json:"networkconf_id,omitempty"`
	Password      *string `json:"password,omitempty"`
}

type WlanConfSaePsk struct {
	Id   *string  `json:"id,omitempty"`
	Mac  *string  `json:"mac,omitempty"`
	Psk  *string  `json:"psk,omitempty"`
	Vlan *float64 `json:"vlan,omitempty"`
}

type WlanConfScheduleWithDuration struct {
	DurationMinutes *int64    `json:"duration_minutes,omitempty"`
	Name            *string   `json:"name,omitempty"`
	StartDaysOfWeek *[]string `json:"start_days_of_week,omitempty"`
	StartHour       *int64    `json:"start_hour,omitempty"`
	StartMinute     *int64    `json:"start_minute,omitempty"`
}
