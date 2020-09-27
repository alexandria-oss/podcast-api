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

import "github.com/alexandria-oss/common-go/exception"

// Description large/mid-text to describe an entity
type Description struct {
	value     string
	fieldName string
}

// SetFieldName override field name for top-operations
func (d *Description) SetFieldName(field string) {
	if field != "" {
		d.fieldName = field
	}
}

// GetFieldName get the url object field name
func (d Description) GetFieldName() string {
	if d.fieldName != "" {
		return d.fieldName
	}

	// Default value
	return "description"
}

// Get get description value
func (d Description) Get() string {
	return d.value
}

// Set override description value
func (d *Description) Set(description string) error {
	memo := d.value
	d.value = description

	if err := d.IsValid(); err != nil {
		d.value = memo
		return err
	}

	return nil
}

// IsValid verify description value
func (d Description) IsValid() error {
	// Validation cases
	// - Range from 1 to 2048 characters
	if d.value != "" {
		if len(d.value) > 2048 {
			return exception.NewFieldRange(d.GetFieldName(), "1", "2048")
		}
	}

	return nil
}
