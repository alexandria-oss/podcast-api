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
	"strings"
)

// Title heading of a document
type Title struct {
	value     string
	fieldName string
}

// SetFieldName override field name for top-operations
func (t *Title) SetFieldName(field string) {
	if field != "" {
		t.fieldName = field
	}
}

// GetFieldName get the url object field name
func (t Title) GetFieldName() string {
	if t.fieldName != "" {
		return t.fieldName
	}

	// Default value
	return "title"
}

// Set uniquely set title value (not mutable after first set)
func (t *Title) Set(title string) error {
	// Avoid mutations
	// TODO: Add NotMutable domain exception
	if t.value != "" {
		return exception.NewAlreadyExists(t.GetFieldName())
	}

	t.value = strings.Title(title)
	if err := t.IsValid(); err != nil {
		t.value = ""
		return err
	}

	return nil
}

// Get get title current value
func (t Title) Get() string {
	return t.value
}

// IsValid verify title
func (t Title) IsValid() error {
	// Validation cases
	// - Range from 1 to 512

	if t.value != "" {
		if len(t.value) > 512 {
			return exception.NewFieldRange(t.GetFieldName(), "1", "512")
		}
	}

	return nil
}
