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
	"github.com/alexandria-oss/podcast-api/internal/domain/shared/category"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"strconv"
)

// Category classify an aggregate
type Category struct {
	// ID enum category unique identifier
	ID *value.ID
	// Name Category's display name
	Name *value.Title
}

// Start launch required operations
func (c *Category) Start() {
	c.SetFieldNames()
}

// SetFieldNames set required custom field name(s)
func (c *Category) SetFieldNames() {
	c.ID.SetFieldName("category_id")
	c.Name.SetFieldName("category_name")
}

// Set override category with an id
func (c *Category) Set(id int) error {
	if name := category.GetName(id); name != "" {
		if err := c.Name.Set(name); err != nil {
			return err
		}

		return c.ID.Set(strconv.FormatInt(int64(id), 10))
	}

	return exception.NewNotFound("category")
}

// MarshalString marshal category into string
func (c Category) MarshalString() string {
	return fmt.Sprintf("id: %s, name: %s", c.ID.Get(), c.Name.Get())
}

// IsValid verify category
func (c Category) IsValid() error {
	// Validation cases
	// - Required: ID, Name
	if c.ID.Get() == "" {
		return exception.NewRequiredField("category_id")
	} else if c.Name.Get() == "" {
		return exception.NewRequiredField("category_name")
	}

	// Value object
	if err := c.Name.IsValid(); err != nil {
		return err
	}

	return nil
}
