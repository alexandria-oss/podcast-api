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

import (
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/shared"
)

type Language struct {
	value string
}

// Get get language current value
func (l Language) Get() string {
	return l.value
}

// Set override language
func (l *Language) Set(lang string) error {
	memo := l.value
	l.value = lang

	if err := l.IsValid(); err != nil {
		l.value = memo
		return err
	}

	return nil
}

// IsValid verify language
func (l Language) IsValid() error {
	// Validation cases
	// - Range 2 to 5 digit ("de" - "en-US")
	// - Is supported and complies with ISO 639-2 language code standard

	if l.value != "" {
		if len(l.value) < 2 || len(l.value) > 5 {
			return exception.NewFieldRange("language", "2", "5")
		}

		if shared.SupportedLangMap[l.value].String() != "und" {
			return nil
		}

		return exception.NewFieldFormat("language", "ISO 639-2 (de, en-US)")
	}

	return nil
}
