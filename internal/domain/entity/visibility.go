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
	"github.com/alexandria-oss/podcast-api/internal/domain/shared/visibility"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"strconv"
)

// Visibility brings policy access to users
type Visibility struct {
	ID     *value.ID
	Policy *value.DisplayName
}

// SetFieldNames set required custom field name(s)
func (v *Visibility) SetFieldNames() {
	v.ID.SetFieldName("visibility_id")
	v.Policy.SetFieldName("policy")
}

// Set override visibility policy with an id
func (v *Visibility) Set(id int) error {
	if policy := visibility.GetName(id); policy != "" {
		if err := v.Policy.Set(policy); err != nil {
			return err
		}

		return v.ID.Set(strconv.FormatInt(int64(id), 10))
	}

	return exception.NewNotFound("visibility")
}

// MarshalString marshal visibility into string
func (v Visibility) MarshalString() string {
	return fmt.Sprintf("id: %s, policy: %s", v.ID.Get(), v.Policy.Get())
}

// IsValid verify visibility
func (v Visibility) IsValid() error {
	// Validation cases
	// - Required: ID, Policy
	if v.ID.Get() == "" {
		return exception.NewRequiredField("visibility_id")
	} else if v.Policy.Get() == "" {
		return exception.NewRequiredField("visibility_policy")
	}

	// Value object
	if err := v.Policy.IsValid(); err != nil {
		return err
	}

	return nil
}
