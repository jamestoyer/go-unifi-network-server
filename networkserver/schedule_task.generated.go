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

type ScheduleTask struct {
	Action          *string                       `json:"action,omitempty"`
	CronExpr        *string                       `json:"cron_expr,omitempty"`
	ExecuteOnlyOnce *bool                         `json:"execute_only_once,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	UpgradeTargets  *[]ScheduleTaskUpgradeTargets `json:"upgrade_targets,omitempty"`
}

func (s *ScheduleTask) GetAction() string {
	if s == nil {
		return ""
	}

	return *s.Action
}

func (s *ScheduleTask) GetCronExpr() string {
	if s == nil {
		return ""
	}

	return *s.CronExpr
}

func (s *ScheduleTask) GetExecuteOnlyOnce() bool {
	if s == nil {
		return false
	}

	return *s.ExecuteOnlyOnce
}

func (s *ScheduleTask) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *ScheduleTask) GetUpgradeTargets() []ScheduleTaskUpgradeTargets {
	if s == nil || s.UpgradeTargets == nil {
		return nil
	}

	return *s.UpgradeTargets
}

type ScheduleTaskUpgradeTargets struct {
	Mac *string `json:"mac,omitempty"`
}

func (s *ScheduleTaskUpgradeTargets) GetMac() string {
	if s == nil {
		return ""
	}

	return *s.Mac
}
