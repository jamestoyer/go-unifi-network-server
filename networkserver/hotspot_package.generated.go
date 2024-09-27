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

type HotspotPackage struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Amount                         *float64 `json:"amount,omitempty"`
	ChargedAs                      *string  `json:"charged_as,omitempty"`
	Currency                       *string  `json:"currency,omitempty"`
	CustomPaymentFieldsEnabled     *bool    `json:"custom_payment_fields_enabled,omitempty"`
	Hours                          *int64   `json:"hours,omitempty"`
	Index                          *int64   `json:"index,omitempty"`
	LimitDown                      *int64   `json:"limit_down,omitempty"`
	LimitOverwrite                 *bool    `json:"limit_overwrite,omitempty"`
	LimitQuota                     *int64   `json:"limit_quota,omitempty"`
	LimitUp                        *int64   `json:"limit_up,omitempty"`
	Name                           *string  `json:"name,omitempty"`
	PaymentFieldsAddressEnabled    *bool    `json:"payment_fields_address_enabled,omitempty"`
	PaymentFieldsAddressRequired   *bool    `json:"payment_fields_address_required,omitempty"`
	PaymentFieldsCityEnabled       *bool    `json:"payment_fields_city_enabled,omitempty"`
	PaymentFieldsCityRequired      *bool    `json:"payment_fields_city_required,omitempty"`
	PaymentFieldsCountryEnabled    *bool    `json:"payment_fields_country_enabled,omitempty"`
	PaymentFieldsCountryRequired   *bool    `json:"payment_fields_country_required,omitempty"`
	PaymentFieldsEmailEnabled      *bool    `json:"payment_fields_email_enabled,omitempty"`
	PaymentFieldsEmailRequired     *bool    `json:"payment_fields_email_required,omitempty"`
	PaymentFieldsFirstNameEnabled  *bool    `json:"payment_fields_first_name_enabled,omitempty"`
	PaymentFieldsFirstNameRequired *bool    `json:"payment_fields_first_name_required,omitempty"`
	PaymentFieldsLastNameEnabled   *bool    `json:"payment_fields_last_name_enabled,omitempty"`
	PaymentFieldsLastNameRequired  *bool    `json:"payment_fields_last_name_required,omitempty"`
	PaymentFieldsStateEnabled      *bool    `json:"payment_fields_state_enabled,omitempty"`
	PaymentFieldsStateRequired     *bool    `json:"payment_fields_state_required,omitempty"`
	PaymentFieldsZipEnabled        *bool    `json:"payment_fields_zip_enabled,omitempty"`
	PaymentFieldsZipRequired       *bool    `json:"payment_fields_zip_required,omitempty"`
	TrialDurationMinutes           *int64   `json:"trial_duration_minutes,omitempty"`
	TrialReset                     *float64 `json:"trial_reset,omitempty"`
}

func (s *HotspotPackage) GetAmount() float64 {
	if s == nil {
		return 0
	}

	return *s.Amount
}

func (s *HotspotPackage) GetChargedAs() string {
	if s == nil {
		return ""
	}

	return *s.ChargedAs
}

func (s *HotspotPackage) GetCurrency() string {
	if s == nil {
		return ""
	}

	return *s.Currency
}

func (s *HotspotPackage) GetCustomPaymentFieldsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.CustomPaymentFieldsEnabled
}

func (s *HotspotPackage) GetHours() int64 {
	if s == nil {
		return 0
	}

	return *s.Hours
}

func (s *HotspotPackage) GetIndex() int64 {
	if s == nil {
		return 0
	}

	return *s.Index
}

func (s *HotspotPackage) GetLimitDown() int64 {
	if s == nil {
		return 0
	}

	return *s.LimitDown
}

func (s *HotspotPackage) GetLimitOverwrite() bool {
	if s == nil {
		return false
	}

	return *s.LimitOverwrite
}

func (s *HotspotPackage) GetLimitQuota() int64 {
	if s == nil {
		return 0
	}

	return *s.LimitQuota
}

func (s *HotspotPackage) GetLimitUp() int64 {
	if s == nil {
		return 0
	}

	return *s.LimitUp
}

func (s *HotspotPackage) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *HotspotPackage) GetPaymentFieldsAddressEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsAddressEnabled
}

func (s *HotspotPackage) GetPaymentFieldsAddressRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsAddressRequired
}

func (s *HotspotPackage) GetPaymentFieldsCityEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsCityEnabled
}

func (s *HotspotPackage) GetPaymentFieldsCityRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsCityRequired
}

func (s *HotspotPackage) GetPaymentFieldsCountryEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsCountryEnabled
}

func (s *HotspotPackage) GetPaymentFieldsCountryRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsCountryRequired
}

func (s *HotspotPackage) GetPaymentFieldsEmailEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsEmailEnabled
}

func (s *HotspotPackage) GetPaymentFieldsEmailRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsEmailRequired
}

func (s *HotspotPackage) GetPaymentFieldsFirstNameEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsFirstNameEnabled
}

func (s *HotspotPackage) GetPaymentFieldsFirstNameRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsFirstNameRequired
}

func (s *HotspotPackage) GetPaymentFieldsLastNameEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsLastNameEnabled
}

func (s *HotspotPackage) GetPaymentFieldsLastNameRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsLastNameRequired
}

func (s *HotspotPackage) GetPaymentFieldsStateEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsStateEnabled
}

func (s *HotspotPackage) GetPaymentFieldsStateRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsStateRequired
}

func (s *HotspotPackage) GetPaymentFieldsZipEnabled() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsZipEnabled
}

func (s *HotspotPackage) GetPaymentFieldsZipRequired() bool {
	if s == nil {
		return false
	}

	return *s.PaymentFieldsZipRequired
}

func (s *HotspotPackage) GetTrialDurationMinutes() int64 {
	if s == nil {
		return 0
	}

	return *s.TrialDurationMinutes
}

func (s *HotspotPackage) GetTrialReset() float64 {
	if s == nil {
		return 0
	}

	return *s.TrialReset
}
