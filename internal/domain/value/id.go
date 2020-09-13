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

// ID unique identifier
type ID struct {
	value string
}

// Get get current unique identifier value
func (i ID) Get() string {
	return i.value
}

// Set override unique identifier value
func (i *ID) Set(id string) error {
	memo := i.value
	i.value = id

	if err := i.IsValid(); err != nil {
		i.value = memo
		return err
	}

	return nil
}

// IsValid verify unique identifier
func (i ID) IsValid() error {
	// Validation cases
	// - Required

	if i.value == "" {
		return exception.NewRequiredField("id")
	}

	return nil
}
