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

package visibility

const (
	Private = iota
	Unlisted
	Public
)

// GetName get visibility state name
func GetName(visibilityID int) string {
	switch visibilityID {
	case 0:
		return "PRIVATE"
	case 1:
		return "UNLISTED"
	case 2:
		return "PUBLIC"
	default:
		return ""
	}
}
