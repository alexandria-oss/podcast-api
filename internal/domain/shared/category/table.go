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

package category

const (
	Misc = iota
	Formal
	Empirical

	Social
	Natural
	Physical
	Life

	Logic
	Mathematics
	Statistics
	Computer

	Economics
	Politics
	Sociology
	Psychology

	Chemistry
	Physics
	Biology
	Earth
	Space

	Biochemistry
	Microbiology
	Botany
	Zoology
	Ecology

	Engineering
	Agricultural
	Medicine
	Dentistry
	Pharmacy

	Business
	Jurisprudence
	Pedagogy
)

// GetName get a complete description from a category ID
func GetName(categoryID int) string {
	switch categoryID {
	case 0:
		return "Miscellaneous"
	case 1:
		return "Formal Science"
	case 2:
		return "Empirical Sciences"
	case 3:
		return "Social Science"
	case 4:
		return "Natural Science"
	case 5:
		return "Physical Science"
	case 6:
		return "Life Science"
	case 7:
		return "Logic"
	case 8:
		return "Mathematics"
	case 9:
		return "Statistics"
	case 10:
		return "Computer Science"
	case 11:
		return "Economics"
	case 12:
		return "Politics"
	case 13:
		return "Sociology"
	case 14:
		return "Psychology"
	case 15:
		return "Chemistry"
	case 16:
		return "Physics"
	case 17:
		return "Biology"
	case 18:
		return "Earth Science"
	case 19:
		return "Space Science"
	case 20:
		return "Biochemistry"
	case 21:
		return "Microbiology"
	case 22:
		return "Botany"
	case 23:
		return "Zoology"
	case 24:
		return "Ecology"
	case 25:
		return "Engineering"
	case 26:
		return "Agricultural"
	case 27:
		return "Medicine"
	case 28:
		return "Dentistry"
	case 29:
		return "Pharmacy"
	case 30:
		return "Business"
	case 31:
		return "Jurisprudence"
	case 32:
		return "Pedagogy"
	default:
		return ""
	}
}
