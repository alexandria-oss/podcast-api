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
	// cover image url
	cover string
	// canvas image or small loop-video background url
	canvas string
	// contentURL media url
	contentURL string
}

// GetCover get cover url
func (c Content) GetCover() string {
	return c.cover
}

// GetCanvas get canvas url
func (c Content) GetCanvas() string {
	return c.canvas
}

// GetContentURL get content url
func (c Content) GetContentURL() string {
	return c.contentURL
}

// SetCover set cover url
func (c *Content) SetCover(cover string) error {
	memo := c.cover
	c.cover = cover

	if err := c.IsCoverValid(); err != nil {
		c.cover = memo
		return err
	}

	return nil
}

// SetCanvas set canvas url
func (c *Content) SetCanvas(canvas string) error {
	memo := c.canvas
	c.canvas = canvas

	if err := c.IsCanvasValid(); err != nil {
		c.canvas = memo
		return err
	}

	return nil
}

// SetContentURL set content url
func (c *Content) SetContentURL(content string) error {
	memo := c.contentURL
	c.contentURL = content

	if err := c.IsContentURLValid(); err != nil {
		c.contentURL = memo
		return err
	}

	return nil
}

// Validators

// IsCoverValid validate cover external URL
func (c Content) IsCoverValid() error {
	// Validation cases
	// - Required
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	// - Image extension .jpg, .jpeg, .png and .webp only
	field := "cover"
	if c.cover != "" {
		if err := c.validateLink(field, c.cover); err != nil {
			return err
		} else if !strings.HasSuffix(c.cover, ".jpg") &&
			!strings.HasSuffix(c.cover, ".jpeg") && !strings.HasSuffix(c.cover, ".png") &&
			!strings.HasSuffix(c.cover, ".webp") {
			return exception.NewFieldFormat(field, ".jpg, .jpeg, .png or .webp file extension")
		}

		return nil
	}

	return exception.NewRequiredField(field)
}

// IsCanvasValid validate canvas external URL
func (c Content) IsCanvasValid() error {
	// Validation cases
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	// - Image/Video extension .jpg, .jpeg, .png, .webp and .mp4 only
	field := "canvas"
	if c.canvas != "" {
		if err := c.validateLink(field, c.canvas); err != nil {
			return err
		} else if !strings.HasSuffix(c.canvas, ".jpg") &&
			!strings.HasSuffix(c.canvas, ".jpeg") &&
			!strings.HasSuffix(c.canvas, ".png") && !strings.HasSuffix(c.canvas, ".webp") &&
			!strings.HasSuffix(c.canvas, ".mp4") {
			return exception.NewFieldFormat(field, ".jpg, .jpeg, .png or .webp file extension")
		}
	}

	return nil
}

// IsContentURLValid validate content external URL
func (c Content) IsContentURLValid() error {
	// Validation cases
	// - Required
	// - HTTPS only
	// - URL length standard (2-2048 characters)
	// - Audio extension .mp3 and .flac only
	field := "content_url"
	if c.contentURL != "" {
		if err := c.validateLink(field, c.contentURL); err != nil {
			return err
		} else if !strings.HasSuffix(c.contentURL, ".mp3") &&
			!strings.HasSuffix(c.contentURL, ".flac") {
			return exception.NewFieldFormat(field, ".mp3 or .flac file extension")
		}

		return nil
	}

	return exception.NewRequiredField(field)
}

func (Content) validateLink(field, value string) error {
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
