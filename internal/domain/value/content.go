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
