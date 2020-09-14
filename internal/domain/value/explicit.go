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

// Explicit verify whether an entity contains explicit content or not
type Explicit struct {
	value bool
}

// IsExplicit get if it contains explicit content
func (e Explicit) IsExplicit() bool {
	return e.value
}

// IsExplicitString get if it contains explicit content as string
func (e Explicit) IsExplicitString() string {
	if e.value {
		return "yes"
	}

	return "no"
}

// Toggle set explicit state to opposite value
func (e *Explicit) Toggle() {
	if e.value {
		e.value = false
		return
	}

	e.value = true
}

// Set override current explicit state
func (e *Explicit) Set(state bool) {
	e.value = state
}
