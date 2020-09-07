package value

import (
	"github.com/alexandria-oss/common-go/exception"
	"strings"
)

// File binary extension
type File struct {
	// extension file extension
	extension string
	// byteLength file byte length
	byteLength int64
}

// GetExtension get file extension
func (f File) GetExtension() string {
	return f.extension
}

// GetByteLength get file length in bytes
func (f File) GetByteLength() int64 {
	return f.byteLength
}

// SetExtension set file extension
func (f *File) SetExtension(extension string) error {
	f.extension = strings.ToLower(extension)

	if err := f.IsExtensionValid(); err != nil {
		return err
	}

	return nil
}

// SetByteLength set file length in bytes
func (f *File) SetByteLength(length int64) error {
	f.byteLength = length

	if err := f.IsLengthValid(); err != nil {
		return err
	}

	return nil
}

// Validators

// IsValid validate overall values
func (f File) IsValid() error {
	if err := f.IsExtensionValid(); err != nil {
		return err
	} else if err := f.IsLengthValid(); err != nil {
		return err
	}

	return nil
}

// IsExtensionValid validate file extension
func (f File) IsExtensionValid() error {
	// Validation cases
	// - Required
	// - Audio extension .mp3 and .flac only
	if f.extension != "" {
		if f.extension != "mp3" && f.extension != "flac" {
			return exception.NewFieldFormat("file_extension", "mp3 or flac")
		}

		return nil
	}

	return exception.NewRequiredField("file_extension")
}

// IsLengthValid validate file byte length
func (f File) IsLengthValid() error {
	// Validation cases
	// - Required
	// - Byte length between 1 and 500 Mebibyte

	// Mebibyte limit in bytes (500 MB)
	byteLimit := int64(500 * 1048576)
	if f.byteLength < 0 && f.byteLength > byteLimit {
		return exception.NewFieldRange("file_length", "1 MB", "500 MB")
	}

	return nil
}