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

// This package is meant to be used whenever a value object is needed.
//
// 	An object that contains attributes but has no conceptual identity. They should be treated as immutable
package value

// ValueObject an object that contains attributes but has no conceptual identity.
// They should be treated as immutable
//	This could be an abstract class with generics
type ValueObject interface {
	Get() interface{}
	Set(value interface{}) error
	IsValid() error

	SetFieldName(field string)
	GetFieldName() string
}
