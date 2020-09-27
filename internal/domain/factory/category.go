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

package factory

import (
	"github.com/alexandria-oss/podcast-api/internal/domain/entity"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
)

// NewCategoryEntity create a new category entity
func NewCategoryEntity(categoryID int) (*entity.Category, error) {
	category := &entity.Category{
		ID:   new(value.ID),
		Name: new(value.Title),
	}
	category.SetFieldNames()
	if err := category.Set(categoryID); err != nil {
		return nil, err
	}

	if err := category.IsValid(); err != nil {
		return nil, err
	}

	return category, nil
}
