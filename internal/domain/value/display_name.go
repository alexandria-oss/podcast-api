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

// DisplayName UI display text for title
type DisplayName struct {
	value string
}

// Get get display name current value
func (d DisplayName) Get() string {
	return d.value
}

// Set override display name value
func (d *DisplayName) Set(displayName string) error {
	memo := d.value
	d.value = displayName

	if err := d.IsValid(); err != nil {
		d.value = memo
		return err
	}

	return nil
}

// IsValid verify display name
func (d DisplayName) IsValid() error {
	// Validation cases
	// - Required
	// - Range from 1 to 256

	if d.value != "" {
		if len(d.value) > 256 {
			return exception.NewFieldRange("display_name", "1", "256")
		}

		return nil
	}

	return exception.NewRequiredField("display_name")
}
