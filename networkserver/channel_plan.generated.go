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
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

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

func (s *ChannelPlan) GetApBlacklistedChannels() []ChannelPlanApBlacklistedChannels {
	if s == nil || s.ApBlacklistedChannels == nil {
		return nil
	}

	return *s.ApBlacklistedChannels
}

func (s *ChannelPlan) GetConfSource() string {
	if s == nil {
		return ""
	}

	return *s.ConfSource
}

func (s *ChannelPlan) GetCoupling() []ChannelPlanCoupling {
	if s == nil || s.Coupling == nil {
		return nil
	}

	return *s.Coupling
}

func (s *ChannelPlan) GetDate() string {
	if s == nil {
		return ""
	}

	return *s.Date
}

func (s *ChannelPlan) GetFitness() string {
	if s == nil {
		return ""
	}

	return *s.Fitness
}

func (s *ChannelPlan) GetNote() string {
	if s == nil {
		return ""
	}

	return *s.Note
}

func (s *ChannelPlan) GetRadio() string {
	if s == nil {
		return ""
	}

	return *s.Radio
}

func (s *ChannelPlan) GetRadioTable() []ChannelPlanRadioTable {
	if s == nil || s.RadioTable == nil {
		return nil
	}

	return *s.RadioTable
}

func (s *ChannelPlan) GetSatisfaction() string {
	if s == nil {
		return ""
	}

	return *s.Satisfaction
}

func (s *ChannelPlan) GetSatisfactionTable() []ChannelPlanSatisfactionTable {
	if s == nil || s.SatisfactionTable == nil {
		return nil
	}

	return *s.SatisfactionTable
}

func (s *ChannelPlan) GetSiteBlacklistedChannels() []ChannelPlanSiteBlacklistedChannels {
	if s == nil || s.SiteBlacklistedChannels == nil {
		return nil
	}

	return *s.SiteBlacklistedChannels
}

type ChannelPlanApBlacklistedChannels struct {
	Channel   *float64 `json:"channel,omitempty"`
	Mac       *string  `json:"mac,omitempty"`
	Timestamp *int64   `json:"timestamp,omitempty"`
}

func (s *ChannelPlanApBlacklistedChannels) GetChannel() float64 {
	if s == nil {
		return 0
	}

	return *s.Channel
}

func (s *ChannelPlanApBlacklistedChannels) GetMac() string {
	if s == nil {
		return ""
	}

	return *s.Mac
}

func (s *ChannelPlanApBlacklistedChannels) GetTimestamp() int64 {
	if s == nil {
		return 0
	}

	return *s.Timestamp
}

type ChannelPlanCoupling struct {
	Rssi   *int64  `json:"rssi,omitempty"`
	Source *string `json:"source,omitempty"`
	Target *string `json:"target,omitempty"`
}

func (s *ChannelPlanCoupling) GetRssi() int64 {
	if s == nil {
		return 0
	}

	return *s.Rssi
}

func (s *ChannelPlanCoupling) GetSource() string {
	if s == nil {
		return ""
	}

	return *s.Source
}

func (s *ChannelPlanCoupling) GetTarget() string {
	if s == nil {
		return ""
	}

	return *s.Target
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

func (s *ChannelPlanRadioTable) GetBackupChannel() string {
	if s == nil {
		return ""
	}

	return *s.BackupChannel
}

func (s *ChannelPlanRadioTable) GetChannel() string {
	if s == nil {
		return ""
	}

	return *s.Channel
}

func (s *ChannelPlanRadioTable) GetDeviceMac() string {
	if s == nil {
		return ""
	}

	return *s.DeviceMac
}

func (s *ChannelPlanRadioTable) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *ChannelPlanRadioTable) GetTxPower() string {
	if s == nil {
		return ""
	}

	return *s.TxPower
}

func (s *ChannelPlanRadioTable) GetTxPowerMode() string {
	if s == nil {
		return ""
	}

	return *s.TxPowerMode
}

func (s *ChannelPlanRadioTable) GetWidth() int64 {
	if s == nil {
		return 0
	}

	return *s.Width
}

type ChannelPlanSatisfactionTable struct {
	DeviceMac    *string `json:"device_mac,omitempty"`
	Satisfaction *string `json:"satisfaction,omitempty"`
}

func (s *ChannelPlanSatisfactionTable) GetDeviceMac() string {
	if s == nil {
		return ""
	}

	return *s.DeviceMac
}

func (s *ChannelPlanSatisfactionTable) GetSatisfaction() string {
	if s == nil {
		return ""
	}

	return *s.Satisfaction
}

type ChannelPlanSiteBlacklistedChannels struct {
	Channel   *float64 `json:"channel,omitempty"`
	Timestamp *int64   `json:"timestamp,omitempty"`
}

func (s *ChannelPlanSiteBlacklistedChannels) GetChannel() float64 {
	if s == nil {
		return 0
	}

	return *s.Channel
}

func (s *ChannelPlanSiteBlacklistedChannels) GetTimestamp() int64 {
	if s == nil {
		return 0
	}

	return *s.Timestamp
}
