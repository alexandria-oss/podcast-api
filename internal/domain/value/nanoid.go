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
	gonanoid "github.com/matoous/go-nanoid"
)

// NanoID unique identifier using 16-bit nano ID strategy
type NanoID struct {
	value string
}

// Get get current unique identifier value
func (i NanoID) Get() string {
	return i.value
}

// Set override unique identifier value
func (i *NanoID) Set(id string) error {
	memo := i.value
	i.value = id

	if err := i.IsValid(); err != nil {
		i.value = memo
		return err
	}

	return nil
}

// Generate generate a new unique identifier
func (i NanoID) Generate() error {
	id, err := gonanoid.Nanoid(16)
	if err != nil {
		return err
	}

	i.value = id
	return nil
}

// IsValid verify unique identifier
func (i NanoID) IsValid() error {
	// Validation cases
	// - Range from 1 to 64

	if i.value != "" {
		if len(i.value) > 64 {
			return exception.NewFieldRange("id", "1", "64")
		}
	}

	return nil
}
