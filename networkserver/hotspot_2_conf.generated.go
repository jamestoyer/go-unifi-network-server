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

type Hotspot2Conf struct {
	AnqpDomainId            *float64                             `json:"anqp_domain_id,omitempty"`
	Capab                   *[]Hotspot2ConfCapab                 `json:"capab,omitempty"`
	CellularNetworkList     *[]Hotspot2ConfCellularNetworkList   `json:"cellular_network_list,omitempty"`
	DeauthReqTimeout        *float64                             `json:"deauth_req_timeout,omitempty"`
	DisableDgaf             *bool                                `json:"disable_dgaf,omitempty"`
	DomainNameList          *[]string                            `json:"domain_name_list,omitempty"`
	FriendlyName            *[]Hotspot2ConfFriendlyName          `json:"friendly_name,omitempty"`
	GasAdvanced             *bool                                `json:"gas_advanced,omitempty"`
	GasComebackDelay        *int64                               `json:"gas_comeback_delay,omitempty"`
	GasFragLimit            *int64                               `json:"gas_frag_limit,omitempty"`
	Hessid                  *string                              `json:"hessid,omitempty"`
	HessidUsed              *bool                                `json:"hessid_used,omitempty"`
	Icons                   *[]Hotspot2ConfIcons                 `json:"icons,omitempty"`
	IpaddrTypeAvailV4       *int64                               `json:"ipaddr_type_avail_v4,omitempty"`
	IpaddrTypeAvailV6       *int64                               `json:"ipaddr_type_avail_v6,omitempty"`
	MetricsDownlinkLoad     *int64                               `json:"metrics_downlink_load,omitempty"`
	MetricsDownlinkLoadSet  *bool                                `json:"metrics_downlink_load_set,omitempty"`
	MetricsDownlinkSpeed    *int64                               `json:"metrics_downlink_speed,omitempty"`
	MetricsDownlinkSpeedSet *bool                                `json:"metrics_downlink_speed_set,omitempty"`
	MetricsInfoAtCapacity   *bool                                `json:"metrics_info_at_capacity,omitempty"`
	MetricsInfoLinkStatus   *string                              `json:"metrics_info_link_status,omitempty"`
	MetricsInfoSymmetric    *bool                                `json:"metrics_info_symmetric,omitempty"`
	MetricsMeasurement      *int64                               `json:"metrics_measurement,omitempty"`
	MetricsMeasurementSet   *bool                                `json:"metrics_measurement_set,omitempty"`
	MetricsStatus           *bool                                `json:"metrics_status,omitempty"`
	MetricsUplinkLoad       *int64                               `json:"metrics_uplink_load,omitempty"`
	MetricsUplinkLoadSet    *bool                                `json:"metrics_uplink_load_set,omitempty"`
	MetricsUplinkSpeed      *int64                               `json:"metrics_uplink_speed,omitempty"`
	MetricsUplinkSpeedSet   *bool                                `json:"metrics_uplink_speed_set,omitempty"`
	NaiRealmList            *[]Hotspot2ConfNaiRealmList          `json:"nai_realm_list,omitempty"`
	Name                    *string                              `json:"name,omitempty"`
	NetworkAccessAsra       *bool                                `json:"network_access_asra,omitempty"`
	NetworkAccessEsr        *bool                                `json:"network_access_esr,omitempty"`
	NetworkAccessInternet   *bool                                `json:"network_access_internet,omitempty"`
	NetworkAccessUesa       *bool                                `json:"network_access_uesa,omitempty"`
	NetworkAuthType         *int64                               `json:"network_auth_type,omitempty"`
	NetworkAuthUrl          *string                              `json:"network_auth_url,omitempty"`
	NetworkType             *int64                               `json:"network_type,omitempty"`
	Osu                     *[]Hotspot2ConfOsu                   `json:"osu,omitempty"`
	OsuSsid                 *string                              `json:"osu_ssid,omitempty"`
	QosMapDcsp              *[]Hotspot2ConfQosMapDcsp            `json:"qos_map_dcsp,omitempty"`
	QosMapExceptions        *[]Hotspot2ConfQosMapExceptions      `json:"qos_map_exceptions,omitempty"`
	QosMapStatus            *bool                                `json:"qos_map_status,omitempty"`
	RoamingConsortiumList   *[]Hotspot2ConfRoamingConsortiumList `json:"roaming_consortium_list,omitempty"`
	SaveTimestamp           *string                              `json:"save_timestamp,omitempty"`
	TCFilename              *string                              `json:"t_c_filename,omitempty"`
	TCTimestamp             *int64                               `json:"t_c_timestamp,omitempty"`
	VenueGroup              *int64                               `json:"venue_group,omitempty"`
	VenueName               *[]Hotspot2ConfVenueName             `json:"venue_name,omitempty"`
	VenueType               *float64                             `json:"venue_type,omitempty"`
}

type Hotspot2ConfCapab struct {
	Port     *float64 `json:"port,omitempty"`
	Protocol *string  `json:"protocol,omitempty"`
	Status   *string  `json:"status,omitempty"`
}

type Hotspot2ConfCellularNetworkList struct {
	Mcc  *int64  `json:"mcc,omitempty"`
	Mnc  *int64  `json:"mnc,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Hotspot2ConfFriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type Hotspot2ConfIcons struct {
	Data     *string `json:"data,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Height   *int64  `json:"height,omitempty"`
	Language *string `json:"language,omitempty"`
	Media    *string `json:"media,omitempty"`
	Name     *string `json:"name,omitempty"`
	Size     *int64  `json:"size,omitempty"`
	Width    *int64  `json:"width,omitempty"`
}

type Hotspot2ConfNaiRealmList struct {
	AuthIds   *string `json:"auth_ids,omitempty"`
	AuthVals  *string `json:"auth_vals,omitempty"`
	EapMethod *int64  `json:"eap_method,omitempty"`
	Encoding  *int64  `json:"encoding,omitempty"`
	Name      *string `json:"name,omitempty"`
	Status    *bool   `json:"status,omitempty"`
}

type Hotspot2ConfOsu struct {
	Description      *[]Hotspot2ConfOsuDescription  `json:"description,omitempty"`
	FriendlyName     *[]Hotspot2ConfOsuFriendlyName `json:"friendly_name,omitempty"`
	Icon             *[]Hotspot2ConfOsuIcon         `json:"icon,omitempty"`
	MethodOmaDm      *bool                          `json:"method_oma_dm,omitempty"`
	MethodSoapXmlSpp *bool                          `json:"method_soap_xml_spp,omitempty"`
	Nai              *string                        `json:"nai,omitempty"`
	Nai2             *string                        `json:"nai2,omitempty"`
	OperatingClass   *string                        `json:"operating_class,omitempty"`
	ServerUri        *string                        `json:"server_uri,omitempty"`
}

type Hotspot2ConfOsuDescription struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type Hotspot2ConfOsuFriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

type Hotspot2ConfOsuIcon struct {
	Name *string `json:"name,omitempty"`
}

type Hotspot2ConfQosMapDcsp struct {
	High *int64 `json:"high,omitempty"`
	Low  *int64 `json:"low,omitempty"`
}

type Hotspot2ConfQosMapExceptions struct {
	Dcsp *int64 `json:"dcsp,omitempty"`
	Up   *int64 `json:"up,omitempty"`
}

type Hotspot2ConfRoamingConsortiumList struct {
	Name *string `json:"name,omitempty"`
	Oid  *string `json:"oid,omitempty"`
}

type Hotspot2ConfVenueName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}
