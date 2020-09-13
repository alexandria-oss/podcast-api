package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
	"time"
)

func TestFileMetadata_SetByteLength(t *testing.T) {
	f := new(FileMetadata)

	err := f.SetByteLength(-10)
	if !errors.Is(err, exception.FieldRange) {
		t.Error("invalid byte length validation, expected error")
	}

	err = f.SetByteLength(534288000)
	if !errors.Is(err, exception.FieldRange) {
		t.Error("invalid byte length validation, expected error")
	}

	err = f.SetByteLength(524288000)
	if err != nil {
		t.Error("invalid byte length validation, expected nil error")
	}
}

func TestFileMetadata_SetExtension(t *testing.T) {
	f := new(FileMetadata)

	err := f.SetExtension("")
	if !errors.Is(err, exception.RequiredField) {
		t.Error("invalid FileMetadata extension validation, expected error")
	}

	err = f.SetExtension("wav")
	if !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid FileMetadata extension validation, expected error")
	}

	err = f.SetExtension(".mp3")
	if !errors.Is(err, exception.FieldFormat) {
		t.Error("invalid FileMetadata extension validation, expected error")
	}

	err = f.SetExtension("mp3")
	if err != nil {
		t.Error("invalid FileMetadata extension validation, expected nil error")
	}

	err = f.SetExtension("flac")
	if err != nil {
		t.Error("invalid FileMetadata extension validation, expected nil error")
	}
}

func TestFileMetadata_SetUploadTime(t *testing.T) {
	f := new(FileMetadata)

	err := f.SetUploadTime(time.Time{})
	if !errors.Is(err, exception.FieldRange) {
		t.Error("invalid FileMetadata upload time validation, expected error")
	}

	// 1899-31-12
	err = f.SetUploadTime(time.Time{}.AddDate(1898, 11, 30))
	if !errors.Is(err, exception.FieldRange) {
		t.Error("invalid FileMetadata upload time validation, expected error")
	}

	err = f.SetUploadTime(time.Time{}.AddDate(1958, 11, 30))
	if err != nil {
		t.Error("invalid FileMetadata upload time validation, expected nil error")
	}

	err = f.SetUploadTime(time.Now())
	if err != nil {
		t.Error("invalid FileMetadata upload time validation, expected nil error")
	}

	err = f.SetUploadTime(time.Time{}.AddDate(2999, 12, 31))
	if err != nil {
		t.Error("invalid FileMetadata upload time validation, expected nil error")
	}
}
