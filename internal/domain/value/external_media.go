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
	"fmt"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/service"
	"strings"
)

// ExternalMedia external link to resource
type ExternalMedia struct {
	url                string
	fieldName          string
	isFile             bool
	supportedExtension []string
}

// NewExternalMedia create a new ExternalMedia url object
func NewExternalMedia(field, url string, isFile bool, extension ...string) (*ExternalMedia, error) {
	u := new(ExternalMedia)

	u.SetFieldName(field)
	u.SetIsFile(isFile)
	if u.IsFile() {
		u.SetExtensions(extension)
	}

	if err := u.Set(url); err != nil {
		return nil, err
	}

	return u, nil
}

// SetFieldName override field name for top-operations
func (u *ExternalMedia) SetFieldName(field string) {
	if field != "" {
		u.fieldName = field
	}
}

// GetFieldName get the url object field name
func (u ExternalMedia) GetFieldName() string {
	if u.fieldName != "" {
		return u.fieldName
	}

	return "content_url"
}

// SetIsFile override ExternalMedia state isFile
func (u *ExternalMedia) SetIsFile(state bool) {
	u.isFile = state
}

// Toggle toggle to opposite isFile state
func (u *ExternalMedia) Toggle() {
	if u.isFile {
		u.isFile = false
		return
	}

	u.isFile = true
}

// IsFile return either ExternalMedia is referring to a file or not
func (u ExternalMedia) IsFile() bool {
	return u.isFile
}

// SetExtensions override supported extensions
func (u *ExternalMedia) SetExtensions(extensions []string) {
	if u.isFile {
		u.supportedExtension = extensions
	}
}

// AddExtension attach a new supported extension
func (u *ExternalMedia) AddExtension(extension ...string) {
	if u.isFile {
		for _, ext := range extension {
			u.supportedExtension = append(u.supportedExtension, ext)
		}
	}
}

// GetSupportedExtensions get all supported extensions by ExternalMedia
func (u ExternalMedia) GetSupportedExtensions() []string {
	return u.supportedExtension
}

// Get get content ExternalMedia
func (u ExternalMedia) Get() string {
	return u.url
}

// Set set content ExternalMedia
func (u *ExternalMedia) Set(url string) error {
	memo := u.url
	u.url = url

	if err := u.IsValid(); err != nil {
		u.url = memo
		return err
	}

	return nil
}

// IsValid validate content external ExternalMedia
func (u ExternalMedia) IsValid() error {
	// Validation cases
	// - HTTPS only
	// - ExternalMedia length standard (2-2048 characters)
	if u.url != "" {
		if err := service.ValidateURLLength(u.GetFieldName(), u.url); err != nil {
			return err
		}

		if u.isFile {
			for _, ext := range u.supportedExtension {
				if strings.HasSuffix(u.url, ext) {
					return nil
				}
			}

			return exception.NewFieldFormat(u.GetFieldName(),
				fmt.Sprintf("%v file extension(s)", u.supportedExtension))
		}
	}

	return nil
}
