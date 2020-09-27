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

package license

// Creative Commons Licenses
const (
	// (aka CC Zero) is a public dedication tool, which allows creators to give up their
	// copyright and put their works into the worldwide public domain.
	// CC0 allows reusers to distribute, remix, adapt, and build upon the material
	// in any medium or format, with no conditions.
	CC0 = iota
	// Allows reusers to distribute, remix, adapt, and build upon the material in any
	// medium or format, so long as attribution is given to the creator.
	// The license allows for commercial use.
	CCBY
	// Allows reusers to distribute, remix, adapt, and build upon the material in any
	// medium or format, so long as attribution is given to the creator.
	// The license allows for commercial use. If you remix, adapt, or build upon the
	// material, you must license the modified material under identical terms.
	CCBYSA
	// Allows reusers to distribute, remix, adapt, and build upon the material in any
	// medium or format for noncommercial purposes only, and only so long as attribution
	// is given to the creator.
	CCBYNC
	// Allows reusers to distribute, remix, adapt, and build upon the material in any
	// medium or format for noncommercial purposes only, and only so long as attribution
	// is given to the creator. If you remix, adapt, or build upon the material, you must
	// license the modified material under identical terms.
	CCBYNCSA
	// Allows reusers to copy and distribute the material in any medium or format in
	// unadapted form only, and only so long as attribution is given to the creator.
	// The license allows for commercial use.
	CCBYND
	// Allows reusers to copy and distribute the material in any medium or format in
	// unadapted form only, for noncommercial purposes only, and only so long as
	// attribution is given to the creator.
	CCBYNCND
)

// GetName get a complete name from a license ID
func GetName(licenseID int) string {
	switch licenseID {
	case 0:
		return "CC Zero"
	case 1:
		return "CC BY"
	case 2:
		return "CC BY-SA"
	case 3:
		return "CC BY-NC"
	case 4:
		return "CC BY-NC-SA"
	case 5:
		return "CC BY-ND"
	case 6:
		return "CC BY-NC-ND"
	default:
		return ""
	}
}
