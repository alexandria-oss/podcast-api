// Copyright 2020 The Alexandria Foundation
//
// Licensed under the GNU Affero General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package value

import (
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/shared"
	"time"
)

// Metadata struct default metadata
type Metadata struct {
	// createTime create time timestamp
	createTime time.Time
	// updateTime update time timestamp
	updateTime time.Time
	// active content is available
	active bool
}

// GetCreateTime get current create time
func (m Metadata) GetCreateTime() time.Time {
	return m.createTime
}

// SetCreateTime override create time value
func (m *Metadata) SetCreateTime(createTime time.Time) error {
	memo := m.createTime
	m.createTime = createTime

	if err := m.IsCreateTimeValid(); err != nil {
		m.createTime = memo
		return err
	}
	return nil
}

// IsCreateTimeValid verify create time value
func (m Metadata) IsCreateTimeValid() error {
	// Validation cases
	// - Range from 01-01-2020
	// - Future not valid
	if m.createTime.Before(time.Date(2020, time.January, 1, 0, 0,
		0, 0, time.UTC)) || m.createTime.After(time.Now().UTC()) {
		return exception.NewFieldRange("create_time", "2020-01-01",
			time.Now().UTC().Format(shared.RFC3339Micro))
	}

	return nil
}

// GetUpdateTime get current create time
func (m Metadata) GetUpdateTime() time.Time {
	return m.updateTime
}

// SetUpdateTime override create time value
func (m *Metadata) SetUpdateTime(updateTime time.Time) error {
	memo := m.updateTime
	m.updateTime = updateTime

	if err := m.IsUpdateTimeValid(); err != nil {
		m.updateTime = memo
		return err
	}
	return nil
}

// IsUpdateTimeValid verify create time value
func (m Metadata) IsUpdateTimeValid() error {
	// Validation cases
	// - Range from 01-01-2020
	// - Future not valid
	if m.updateTime.Before(time.Date(2020, time.January, 1, 0, 0,
		0, 0, time.UTC)) || m.updateTime.After(time.Now().UTC()) {
		return exception.NewFieldRange("update_time", "2020-01-01",
			time.Now().UTC().Format(shared.RFC3339Micro))
	}

	return nil
}

// IsActive get current state
func (m Metadata) IsActive() bool {
	return m.active
}

// SetState override state value
func (m *Metadata) SetState(state bool) {
	m.active = state
}

// ToggleState set state to opposite
func (m *Metadata) ToggleState() {
	if m.active {
		m.active = false
		return
	}

	m.active = true
}
