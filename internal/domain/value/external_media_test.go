package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestExternalMedia_SetExtensions(t *testing.T) {
	u := new(ExternalMedia)
	u.SetFieldName("")
	u.SetIsFile(true)
	u.SetExtensions([]string{".mp3", ".flac"})

	if err := u.Set("https://domain.com/media/foo.mp4"); !errors.Is(err, exception.FieldFormat) {
		t.Error("ExternalMedia", "invalid external media validation, expected error")
	}

	if err := u.Set("https://domain.com/media/bar.mp3"); err != nil {
		t.Error("ExternalMedia", "invalid external media validation, expected nil error")
	}
}

func TestExternalMedia_SetFieldName(t *testing.T) {
	u := new(ExternalMedia)
	u.SetFieldName("picture")
	u.SetIsFile(true)
	u.SetExtensions([]string{".png", ".jpeg"})

	if err := u.Set("https://domain.com/media/foo.mp4"); !errors.Is(err, exception.FieldFormat) {
		t.Error("ExternalMedia", "invalid external media validation, expected error")
	}

	if err := u.Set("https://domain.com/media/bar.jpeg"); err != nil {
		t.Error("ExternalMedia", "invalid external media validation, expected nil error")
	}

	if u.GetFieldName() != "picture" {
		t.Error("ExternalMedia", "invalid external media validation, expected picture field name")
	}
}

func TestExternalMedia_AddExtension(t *testing.T) {
	u := new(ExternalMedia)
	u.SetFieldName("canvas")
	u.SetIsFile(true)
	u.SetExtensions([]string{".png"})

	if err := u.Set("https://domain.com/media/foo.mp4"); !errors.Is(err, exception.FieldFormat) {
		t.Error("ExternalMedia", "invalid external media validation, expected error")
	}

	if err := u.Set("https://domain.com/media/bar.jpeg"); !errors.Is(err, exception.FieldFormat) {
		t.Error("ExternalMedia", "invalid external media validation, expected error")
	}

	u.AddExtension("jpeg")

	if len(u.GetSupportedExtensions()) < 2 {
		t.Error("ExternalMedia", "invalid external media validation, expected 2 supported extensions")
	}

	if err := u.Set("https://domain.com/media/foo.png"); err != nil {
		t.Error("ExternalMedia", "invalid external media validation, expected nil error")
	}

	if err := u.Set("https://domain.com/media/bar.jpeg"); err != nil {
		t.Error("ExternalMedia", "invalid external media validation, expected nil error")
	}
}

func TestNewExternalMedia(t *testing.T) {
	if _, err := NewExternalMedia("cover", "http://domain.com/media/canvas.jpeg",
		true, ".jpeg"); !errors.Is(err, exception.FieldFormat) {
		t.Error("ExternalMedia", "invalid external media validation, expected error")
	}

	if _, err := NewExternalMedia("canvas", "https://domain.com/media/canvas.mp4",
		true, ".mp4"); err != nil {
		t.Error("ExternalMedia", "invalid external media validation, expected nil error")
	}
}
