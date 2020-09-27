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

// ID generic unique identifier
type ID struct {
	value     string
	fieldName string
}

// SetFieldName override field name for top-operations
func (i *ID) SetFieldName(field string) {
	if field != "" {
		i.fieldName = field
	}
}

// GetFieldName get the url object field name
func (i ID) GetFieldName() string {
	if i.fieldName != "" {
		return i.fieldName
	}

	// Default value
	return "id"
}

// Get get current unique identifier value
func (i ID) Get() string {
	return i.value
}

// Set override unique identifier value
func (i *ID) Set(id string) error {
	i.value = id
	return nil
}
