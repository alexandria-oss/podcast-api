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
	"strings"
)

// Content holds external content attributes
type Content struct {
	// Cover image url
	cover string
	// Canvas image or small loop-video background url
	canvas string
	// ContentURL media url
	contentURL string
}

// Validators

// IsCoverValid validate cover external URL
func (c Content) IsCoverValid() error {
	// Validation cases
	// - Required
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	// - Image extension .jpg, .jpeg, .png and .webp only
	if c.cover != "" {
		if !strings.HasPrefix(c.cover, "https://") {
			return exception.NewFieldFormat("cover", "HTTPS link")
		} else if len(c.cover) < 2 && len(c.cover) > 2048 {
			return exception.NewFieldRange("cover", "2 characters", "2048 characters")
		} else if !strings.HasSuffix(c.cover, ".jpg") && !strings.HasSuffix(c.cover, ".jpeg") &&
			!strings.HasSuffix(c.cover, ".png") && !strings.HasSuffix(c.cover, ".webp") {
			return exception.NewFieldFormat("cover", ".jpg, .jpeg, .png or .webp file extension")
		}

		return nil
	}

	return exception.NewRequiredField("cover")
}

// IsCanvasValid validate canvas external URL
func (c Content) IsCanvasValid() error {
	// Validation cases
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	// - Image/Video extension .jpg, .jpeg, .png, .webp and .mp4 only
	if c.cover != "" {
		if !strings.HasPrefix(c.cover, "https://") {
			return exception.NewFieldFormat("canvas", "HTTPS link")
		} else if len(c.cover) < 2 && len(c.cover) > 2048 {
			return exception.NewFieldRange("canvas", "2 characters", "2048 characters")
		} else if !strings.HasSuffix(c.cover, ".jpg") && !strings.HasSuffix(c.cover, ".jpeg") &&
			!strings.HasSuffix(c.cover, ".png") && !strings.HasSuffix(c.cover, ".webp") {
			return exception.NewFieldFormat("canvas", ".jpg, .jpeg, .png or .webp file extension")
		}

		return nil
	}

	return exception.NewRequiredField("canvas")
}
