package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestContent_SetCanvas(t *testing.T) {
	c := new(Content)

	// Non-SSL secured site
	if err := c.SetCanvas("http://domain.com/pics/123.jpg"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set canvas validation, expected error")
	}

	// Link more than 2048
	str := ""
	for i := 0; i < 260; i++ {
		str += "https://"
	}

	if err := c.SetCanvas(str); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid content set canvas validation, expected error")
	}

	// Non-valid image/video format
	if err := c.SetCanvas("https://domain.com/pics/123.gif"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set canvas validation, expected error")
	}

	// Valid
	if err := c.SetCanvas("https://domain.com/pics/123.jpg"); err != nil {
		t.Error("invalid content set canvas validation, expected nil error")
	}
}

func TestContent_SetCover(t *testing.T) {
	c := new(Content)

	// Non-SSL secured site
	if err := c.SetCover("http://domain.com/pics/123.jpg"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set cover validation, expected error")
	}

	// Link more than 2048
	str := ""
	for i := 0; i < 260; i++ {
		str += "https://"
	}

	if err := c.SetCover(str); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid content set cover validation, expected error")
	}

	// Non-valid image/video format
	if err := c.SetCover("https://domain.com/video/123.mp4"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set cover validation, expected error")
	}

	// Valid
	if err := c.SetCover("https://domain.com/pics/123.jpg"); err != nil {
		t.Error("invalid content set cover validation, expected nil error")
	}
}

func TestContent_SetContentURL(t *testing.T) {
	c := new(Content)

	// Non-SSL secured site
	if err := c.SetContentURL("http://domain.com/pics/123.jpg"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set content_url validation, expected error")
	}

	// Link more than 2048
	str := ""
	for i := 0; i < 260; i++ {
		str += "https://"
	}

	if err := c.SetContentURL(str); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid content set content_url validation, expected error")
	}

	// Non-valid image/video format
	if err := c.SetContentURL("https://domain.com/audio/123.wav"); !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid content set content_url validation, expected error")
	}

	// Valid
	if err := c.SetContentURL("https://domain.com/audio/123.flac"); err != nil {
		t.Error("invalid content set content_url validation, expected nil error")
	}
}
