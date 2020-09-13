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

package service

import (
	"github.com/alexandria-oss/common-go/exception"
	"strings"
)

// ValidateURLLength verify if the given URL complies with URL GET maximum length
func ValidateURLLength(field, value string) error {
	// Validation cases
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	if !strings.HasPrefix(value, "https://") {
		return exception.NewFieldFormat(field, "HTTPS link")
	} else if len(value) <= 2 || len(value) >= 2048 {
		return exception.NewFieldRange(field, "2 characters", "2048 characters")
	}

	return nil
}
