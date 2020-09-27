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

package entity

import (
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/shared/license"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"strconv"
)

// License brings license to an aggregate
type License struct {
	ID   *value.ID
	Name *value.DisplayName
}

// Start launch required operations
func (l *License) Start() {
	l.SetFieldNames()
}

// SetFieldNames set required custom field name(s)
func (l *License) SetFieldNames() {
	l.ID.SetFieldName("License_id")
	l.Name.SetFieldName("license_name")
}

// Set override License with an id
func (l *License) Set(id int) error {
	if policy := license.GetName(id); policy != "" {
		if err := l.Name.Set(policy); err != nil {
			return err
		}

		return l.ID.Set(strconv.FormatInt(int64(id), 10))
	}

	return exception.NewNotFound("license")
}

// MarshalString marshal License into string
func (l License) MarshalString() string {
	return fmt.Sprintf("id: %s, name: %s", l.ID.Get(), l.Name.Get())
}

// IsValid verify License
func (l License) IsValid() error {
	// Validation cases
	// - Required: ID, Name
	if l.ID.Get() == "" {
		return exception.NewRequiredField("license_id")
	} else if l.Name.Get() == "" {
		return exception.NewRequiredField("license_name")
	}

	// Value object
	if err := l.Name.IsValid(); err != nil {
		return err
	}

	return nil
}
