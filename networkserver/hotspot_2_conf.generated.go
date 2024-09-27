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

type Hotspot2Conf struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

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

func (s *Hotspot2Conf) GetAnqpDomainId() float64 {
	if s == nil {
		return 0
	}

	return *s.AnqpDomainId
}

func (s *Hotspot2Conf) GetCapab() []Hotspot2ConfCapab {
	if s == nil || s.Capab == nil {
		return nil
	}

	return *s.Capab
}

func (s *Hotspot2Conf) GetCellularNetworkList() []Hotspot2ConfCellularNetworkList {
	if s == nil || s.CellularNetworkList == nil {
		return nil
	}

	return *s.CellularNetworkList
}

func (s *Hotspot2Conf) GetDeauthReqTimeout() float64 {
	if s == nil {
		return 0
	}

	return *s.DeauthReqTimeout
}

func (s *Hotspot2Conf) GetDisableDgaf() bool {
	if s == nil {
		return false
	}

	return *s.DisableDgaf
}

func (s *Hotspot2Conf) GetDomainNameList() []string {
	if s == nil || s.DomainNameList == nil {
		return nil
	}

	return *s.DomainNameList
}

func (s *Hotspot2Conf) GetFriendlyName() []Hotspot2ConfFriendlyName {
	if s == nil || s.FriendlyName == nil {
		return nil
	}

	return *s.FriendlyName
}

func (s *Hotspot2Conf) GetGasAdvanced() bool {
	if s == nil {
		return false
	}

	return *s.GasAdvanced
}

func (s *Hotspot2Conf) GetGasComebackDelay() int64 {
	if s == nil {
		return 0
	}

	return *s.GasComebackDelay
}

func (s *Hotspot2Conf) GetGasFragLimit() int64 {
	if s == nil {
		return 0
	}

	return *s.GasFragLimit
}

func (s *Hotspot2Conf) GetHessid() string {
	if s == nil {
		return ""
	}

	return *s.Hessid
}

func (s *Hotspot2Conf) GetHessidUsed() bool {
	if s == nil {
		return false
	}

	return *s.HessidUsed
}

func (s *Hotspot2Conf) GetIcons() []Hotspot2ConfIcons {
	if s == nil || s.Icons == nil {
		return nil
	}

	return *s.Icons
}

func (s *Hotspot2Conf) GetIpaddrTypeAvailV4() int64 {
	if s == nil {
		return 0
	}

	return *s.IpaddrTypeAvailV4
}

func (s *Hotspot2Conf) GetIpaddrTypeAvailV6() int64 {
	if s == nil {
		return 0
	}

	return *s.IpaddrTypeAvailV6
}

func (s *Hotspot2Conf) GetMetricsDownlinkLoad() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsDownlinkLoad
}

func (s *Hotspot2Conf) GetMetricsDownlinkLoadSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsDownlinkLoadSet
}

func (s *Hotspot2Conf) GetMetricsDownlinkSpeed() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsDownlinkSpeed
}

func (s *Hotspot2Conf) GetMetricsDownlinkSpeedSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsDownlinkSpeedSet
}

func (s *Hotspot2Conf) GetMetricsInfoAtCapacity() bool {
	if s == nil {
		return false
	}

	return *s.MetricsInfoAtCapacity
}

func (s *Hotspot2Conf) GetMetricsInfoLinkStatus() string {
	if s == nil {
		return ""
	}

	return *s.MetricsInfoLinkStatus
}

func (s *Hotspot2Conf) GetMetricsInfoSymmetric() bool {
	if s == nil {
		return false
	}

	return *s.MetricsInfoSymmetric
}

func (s *Hotspot2Conf) GetMetricsMeasurement() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsMeasurement
}

func (s *Hotspot2Conf) GetMetricsMeasurementSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsMeasurementSet
}

func (s *Hotspot2Conf) GetMetricsStatus() bool {
	if s == nil {
		return false
	}

	return *s.MetricsStatus
}

func (s *Hotspot2Conf) GetMetricsUplinkLoad() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsUplinkLoad
}

func (s *Hotspot2Conf) GetMetricsUplinkLoadSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsUplinkLoadSet
}

func (s *Hotspot2Conf) GetMetricsUplinkSpeed() int64 {
	if s == nil {
		return 0
	}

	return *s.MetricsUplinkSpeed
}

func (s *Hotspot2Conf) GetMetricsUplinkSpeedSet() bool {
	if s == nil {
		return false
	}

	return *s.MetricsUplinkSpeedSet
}

func (s *Hotspot2Conf) GetNaiRealmList() []Hotspot2ConfNaiRealmList {
	if s == nil || s.NaiRealmList == nil {
		return nil
	}

	return *s.NaiRealmList
}

func (s *Hotspot2Conf) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Hotspot2Conf) GetNetworkAccessAsra() bool {
	if s == nil {
		return false
	}

	return *s.NetworkAccessAsra
}

func (s *Hotspot2Conf) GetNetworkAccessEsr() bool {
	if s == nil {
		return false
	}

	return *s.NetworkAccessEsr
}

func (s *Hotspot2Conf) GetNetworkAccessInternet() bool {
	if s == nil {
		return false
	}

	return *s.NetworkAccessInternet
}

func (s *Hotspot2Conf) GetNetworkAccessUesa() bool {
	if s == nil {
		return false
	}

	return *s.NetworkAccessUesa
}

func (s *Hotspot2Conf) GetNetworkAuthType() int64 {
	if s == nil {
		return 0
	}

	return *s.NetworkAuthType
}

func (s *Hotspot2Conf) GetNetworkAuthUrl() string {
	if s == nil {
		return ""
	}

	return *s.NetworkAuthUrl
}

func (s *Hotspot2Conf) GetNetworkType() int64 {
	if s == nil {
		return 0
	}

	return *s.NetworkType
}

func (s *Hotspot2Conf) GetOsu() []Hotspot2ConfOsu {
	if s == nil || s.Osu == nil {
		return nil
	}

	return *s.Osu
}

func (s *Hotspot2Conf) GetOsuSsid() string {
	if s == nil {
		return ""
	}

	return *s.OsuSsid
}

func (s *Hotspot2Conf) GetQosMapDcsp() []Hotspot2ConfQosMapDcsp {
	if s == nil || s.QosMapDcsp == nil {
		return nil
	}

	return *s.QosMapDcsp
}

func (s *Hotspot2Conf) GetQosMapExceptions() []Hotspot2ConfQosMapExceptions {
	if s == nil || s.QosMapExceptions == nil {
		return nil
	}

	return *s.QosMapExceptions
}

func (s *Hotspot2Conf) GetQosMapStatus() bool {
	if s == nil {
		return false
	}

	return *s.QosMapStatus
}

func (s *Hotspot2Conf) GetRoamingConsortiumList() []Hotspot2ConfRoamingConsortiumList {
	if s == nil || s.RoamingConsortiumList == nil {
		return nil
	}

	return *s.RoamingConsortiumList
}

func (s *Hotspot2Conf) GetSaveTimestamp() string {
	if s == nil {
		return ""
	}

	return *s.SaveTimestamp
}

func (s *Hotspot2Conf) GetTCFilename() string {
	if s == nil {
		return ""
	}

	return *s.TCFilename
}

func (s *Hotspot2Conf) GetTCTimestamp() int64 {
	if s == nil {
		return 0
	}

	return *s.TCTimestamp
}

func (s *Hotspot2Conf) GetVenueGroup() int64 {
	if s == nil {
		return 0
	}

	return *s.VenueGroup
}

func (s *Hotspot2Conf) GetVenueName() []Hotspot2ConfVenueName {
	if s == nil || s.VenueName == nil {
		return nil
	}

	return *s.VenueName
}

func (s *Hotspot2Conf) GetVenueType() float64 {
	if s == nil {
		return 0
	}

	return *s.VenueType
}

type Hotspot2ConfCapab struct {
	Port     *float64 `json:"port,omitempty"`
	Protocol *string  `json:"protocol,omitempty"`
	Status   *string  `json:"status,omitempty"`
}

func (s *Hotspot2ConfCapab) GetPort() float64 {
	if s == nil {
		return 0
	}

	return *s.Port
}

func (s *Hotspot2ConfCapab) GetProtocol() string {
	if s == nil {
		return ""
	}

	return *s.Protocol
}

func (s *Hotspot2ConfCapab) GetStatus() string {
	if s == nil {
		return ""
	}

	return *s.Status
}

type Hotspot2ConfCellularNetworkList struct {
	Mcc  *int64  `json:"mcc,omitempty"`
	Mnc  *int64  `json:"mnc,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (s *Hotspot2ConfCellularNetworkList) GetMcc() int64 {
	if s == nil {
		return 0
	}

	return *s.Mcc
}

func (s *Hotspot2ConfCellularNetworkList) GetMnc() int64 {
	if s == nil {
		return 0
	}

	return *s.Mnc
}

func (s *Hotspot2ConfCellularNetworkList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type Hotspot2ConfFriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

func (s *Hotspot2ConfFriendlyName) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *Hotspot2ConfFriendlyName) GetText() string {
	if s == nil {
		return ""
	}

	return *s.Text
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

func (s *Hotspot2ConfIcons) GetData() string {
	if s == nil {
		return ""
	}

	return *s.Data
}

func (s *Hotspot2ConfIcons) GetFilename() string {
	if s == nil {
		return ""
	}

	return *s.Filename
}

func (s *Hotspot2ConfIcons) GetHeight() int64 {
	if s == nil {
		return 0
	}

	return *s.Height
}

func (s *Hotspot2ConfIcons) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *Hotspot2ConfIcons) GetMedia() string {
	if s == nil {
		return ""
	}

	return *s.Media
}

func (s *Hotspot2ConfIcons) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Hotspot2ConfIcons) GetSize() int64 {
	if s == nil {
		return 0
	}

	return *s.Size
}

func (s *Hotspot2ConfIcons) GetWidth() int64 {
	if s == nil {
		return 0
	}

	return *s.Width
}

type Hotspot2ConfNaiRealmList struct {
	AuthIds   *string `json:"auth_ids,omitempty"`
	AuthVals  *string `json:"auth_vals,omitempty"`
	EapMethod *int64  `json:"eap_method,omitempty"`
	Encoding  *int64  `json:"encoding,omitempty"`
	Name      *string `json:"name,omitempty"`
	Status    *bool   `json:"status,omitempty"`
}

func (s *Hotspot2ConfNaiRealmList) GetAuthIds() string {
	if s == nil {
		return ""
	}

	return *s.AuthIds
}

func (s *Hotspot2ConfNaiRealmList) GetAuthVals() string {
	if s == nil {
		return ""
	}

	return *s.AuthVals
}

func (s *Hotspot2ConfNaiRealmList) GetEapMethod() int64 {
	if s == nil {
		return 0
	}

	return *s.EapMethod
}

func (s *Hotspot2ConfNaiRealmList) GetEncoding() int64 {
	if s == nil {
		return 0
	}

	return *s.Encoding
}

func (s *Hotspot2ConfNaiRealmList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Hotspot2ConfNaiRealmList) GetStatus() bool {
	if s == nil {
		return false
	}

	return *s.Status
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

func (s *Hotspot2ConfOsu) GetDescription() []Hotspot2ConfOsuDescription {
	if s == nil || s.Description == nil {
		return nil
	}

	return *s.Description
}

func (s *Hotspot2ConfOsu) GetFriendlyName() []Hotspot2ConfOsuFriendlyName {
	if s == nil || s.FriendlyName == nil {
		return nil
	}

	return *s.FriendlyName
}

func (s *Hotspot2ConfOsu) GetIcon() []Hotspot2ConfOsuIcon {
	if s == nil || s.Icon == nil {
		return nil
	}

	return *s.Icon
}

func (s *Hotspot2ConfOsu) GetMethodOmaDm() bool {
	if s == nil {
		return false
	}

	return *s.MethodOmaDm
}

func (s *Hotspot2ConfOsu) GetMethodSoapXmlSpp() bool {
	if s == nil {
		return false
	}

	return *s.MethodSoapXmlSpp
}

func (s *Hotspot2ConfOsu) GetNai() string {
	if s == nil {
		return ""
	}

	return *s.Nai
}

func (s *Hotspot2ConfOsu) GetNai2() string {
	if s == nil {
		return ""
	}

	return *s.Nai2
}

func (s *Hotspot2ConfOsu) GetOperatingClass() string {
	if s == nil {
		return ""
	}

	return *s.OperatingClass
}

func (s *Hotspot2ConfOsu) GetServerUri() string {
	if s == nil {
		return ""
	}

	return *s.ServerUri
}

type Hotspot2ConfOsuDescription struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

func (s *Hotspot2ConfOsuDescription) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *Hotspot2ConfOsuDescription) GetText() string {
	if s == nil {
		return ""
	}

	return *s.Text
}

type Hotspot2ConfOsuFriendlyName struct {
	Language *string `json:"language,omitempty"`
	Text     *string `json:"text,omitempty"`
}

func (s *Hotspot2ConfOsuFriendlyName) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *Hotspot2ConfOsuFriendlyName) GetText() string {
	if s == nil {
		return ""
	}

	return *s.Text
}

type Hotspot2ConfOsuIcon struct {
	Name *string `json:"name,omitempty"`
}

func (s *Hotspot2ConfOsuIcon) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type Hotspot2ConfQosMapDcsp struct {
	High *int64 `json:"high,omitempty"`
	Low  *int64 `json:"low,omitempty"`
}

func (s *Hotspot2ConfQosMapDcsp) GetHigh() int64 {
	if s == nil {
		return 0
	}

	return *s.High
}

func (s *Hotspot2ConfQosMapDcsp) GetLow() int64 {
	if s == nil {
		return 0
	}

	return *s.Low
}

type Hotspot2ConfQosMapExceptions struct {
	Dcsp *int64 `json:"dcsp,omitempty"`
	Up   *int64 `json:"up,omitempty"`
}

func (s *Hotspot2ConfQosMapExceptions) GetDcsp() int64 {
	if s == nil {
		return 0
	}

	return *s.Dcsp
}

func (s *Hotspot2ConfQosMapExceptions) GetUp() int64 {
	if s == nil {
		return 0
	}

	return *s.Up
}

type Hotspot2ConfRoamingConsortiumList struct {
	Name *string `json:"name,omitempty"`
	Oid  *string `json:"oid,omitempty"`
}

func (s *Hotspot2ConfRoamingConsortiumList) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Hotspot2ConfRoamingConsortiumList) GetOid() string {
	if s == nil {
		return ""
	}

	return *s.Oid
}

type Hotspot2ConfVenueName struct {
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}

func (s *Hotspot2ConfVenueName) GetLanguage() string {
	if s == nil {
		return ""
	}

	return *s.Language
}

func (s *Hotspot2ConfVenueName) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Hotspot2ConfVenueName) GetUrl() string {
	if s == nil {
		return ""
	}

	return *s.Url
}

type responseBodyHotspot2Conf struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []Hotspot2Conf  `json:"data"`
}

func (c *Client) CreateHotspot2Conf(ctx context.Context, site string, data *Hotspot2Conf) (*Hotspot2Conf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspot2conf")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspot2Conf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create hotspot2conf: %w`, err)
	}

	var item Hotspot2Conf
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create Hotspot2Conf`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteHotspot2Conf(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspot2conf", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyHotspot2Conf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete Hotspot2Conf: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetHotspot2Conf(ctx context.Context, site, id string) (*Hotspot2Conf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspot2conf", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspot2Conf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Hotspot2Conf: %w`, err)
	}

	var item Hotspot2Conf
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListHotspot2Conf(ctx context.Context, site string) ([]Hotspot2Conf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspot2conf")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspot2Conf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get Hotspot2Conf: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateHotspot2Conf(ctx context.Context, site string, data *Hotspot2Conf) (*Hotspot2Conf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "hotspot2conf", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyHotspot2Conf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update hotspot2conf: %w`, err)
	}

	var item Hotspot2Conf
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update Hotspot2Conf`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
