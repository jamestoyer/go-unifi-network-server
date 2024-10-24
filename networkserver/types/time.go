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

package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type UnixEpochTime struct {
	time.Time
}

func (u *UnixEpochTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	if err := json.Unmarshal(b, &timestamp); err != nil {
		return err
	}

	u.Time = time.Unix(timestamp, 0)
	return nil
}

func (u *UnixEpochTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", u.Time.Unix())), nil
}
