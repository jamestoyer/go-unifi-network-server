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

type ChannelPlan struct {
	ApBlacklistedChannels   *[]ChannelPlanApBlacklistedChannels   `json:"ap_blacklisted_channels,omitempty"`
	ConfSource              *string                               `json:"conf_source,omitempty"`
	Coupling                *[]ChannelPlanCoupling                `json:"coupling,omitempty"`
	Date                    *string                               `json:"date,omitempty"`
	Fitness                 *string                               `json:"fitness,omitempty"`
	Note                    *string                               `json:"note,omitempty"`
	Radio                   *string                               `json:"radio,omitempty"`
	RadioTable              *[]ChannelPlanRadioTable              `json:"radio_table,omitempty"`
	Satisfaction            *string                               `json:"satisfaction,omitempty"`
	SatisfactionTable       *[]ChannelPlanSatisfactionTable       `json:"satisfaction_table,omitempty"`
	SiteBlacklistedChannels *[]ChannelPlanSiteBlacklistedChannels `json:"site_blacklisted_channels,omitempty"`
}

type ChannelPlanApBlacklistedChannels struct {
	Channel   *float64 `json:"channel,omitempty"`
	Mac       *string  `json:"mac,omitempty"`
	Timestamp *int64   `json:"timestamp,omitempty"`
}

type ChannelPlanCoupling struct {
	Rssi   *int64  `json:"rssi,omitempty"`
	Source *string `json:"source,omitempty"`
	Target *string `json:"target,omitempty"`
}

type ChannelPlanRadioTable struct {
	BackupChannel *string `json:"backup_channel,omitempty"`
	Channel       *string `json:"channel,omitempty"`
	DeviceMac     *string `json:"device_mac,omitempty"`
	Name          *string `json:"name,omitempty"`
	TxPower       *string `json:"tx_power,omitempty"`
	TxPowerMode   *string `json:"tx_power_mode,omitempty"`
	Width         *int64  `json:"width,omitempty"`
}

type ChannelPlanSatisfactionTable struct {
	DeviceMac    *string `json:"device_mac,omitempty"`
	Satisfaction *string `json:"satisfaction,omitempty"`
}

type ChannelPlanSiteBlacklistedChannels struct {
	Channel   *float64 `json:"channel,omitempty"`
	Timestamp *int64   `json:"timestamp,omitempty"`
}
