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

// Keyword words defined to properly filter data
type Keyword struct {
	value     string
	fieldName string
}

// SetFieldName override field name for top-operations
func (k *Keyword) SetFieldName(field string) {
	if field != "" {
		k.fieldName = field
	}
}

// GetFieldName get the url object field name
func (k Keyword) GetFieldName() string {
	if k.fieldName != "" {
		return k.fieldName
	}

	// Default value
	return "keyword"
}

// GetKeyword get the keyword value
func (k Keyword) GetKeyword() string {
	return k.value
}

// SetKeyword set the keyword value
func (k *Keyword) SetKeyword(keyword string) error {
	memo := k.value

	k.value = strings.ToLower(keyword)
	if err := k.IsKeywordValid(); err != nil {
		k.value = memo
		return err
	}

	return nil
}

// IsKeywordValid validate current keyword value
func (k Keyword) IsKeywordValid() error {
	// Validation cases
	// - Range from 1 to 128

	if k.value != "" {
		if len(k.value) > 128 {
			return exception.NewFieldRange(k.GetFieldName(), "1", "128")
		}
	}

	return nil
}
