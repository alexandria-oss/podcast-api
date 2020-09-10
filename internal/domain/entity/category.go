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
	"github.com/alexandria-oss/common-go/exception"
	"strings"
)

type Category struct {
	// id foreign unique identifier
	id string
	// name denormalized category's name field
	name string
}

// GetID get category unique identifier
func (c Category) GetID() string {
	return c.id
}

// GetName get category name
func (c Category) GetName() string {
	return c.name
}

// SetID set category unique identifier
func (c *Category) SetID(id string) error {
	c.id = id
	return nil
}

// SetName set category name
func (c *Category) SetName(name string) error {
	mem := c.name

	c.name = strings.Title(name)
	if err := c.IsNameValid(); err != nil {
		c.name = mem
		return err
	}

	return nil
}

// Validators

// IsNameValid validate category name
func (c Category) IsNameValid() error {
	// Validation cases
	// * Foreign data
	// - Required
	if c.name == "" {
		return exception.NewRequiredField("category_name")
	}

	return nil
}
