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

// Bool is a helper function that converts a bool into a point to the value.
func Bool(v bool) *bool { return &v }

// Float64 is a helper function that converts a float64 into a point to the value.
func Float64(v float64) *float64 { return &v }

// Int64 is a helper function that converts an int64 into a point to the value.
func Int64(v int64) *int64 { return &v }

// String is a helper function that converts a string into a point to the value.
func String(v string) *string { return &v }