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
	"time"
)

// File binary extension
type File struct {
	// extension file extension
	extension string
	// byteLength file byte length
	byteLength uint64
	// uploadTime file upload time
	uploadTime time.Time
}

// GetExtension get file extension
func (f File) GetExtension() string {
	return f.extension
}

// GetByteLength get file length in bytes
func (f File) GetByteLength() uint64 {
	return f.byteLength
}

// GetUploadTime get file upload time
func (f File) GetUploadTime() time.Time {
	return f.uploadTime
}

// SetExtension set file extension
func (f *File) SetExtension(extension string) error {
	memo := f.extension
	f.extension = strings.ToLower(extension)

	if err := f.IsExtensionValid(); err != nil {
		f.extension = memo
		return err
	}

	return nil
}

// SetByteLength set file length in bytes
func (f *File) SetByteLength(length int64) error {
	// Avoid uint overflow
	if length < 0 {
		return exception.NewFieldRange("file_length", "1 MB", "500 MB")
	}

	memo := f.byteLength
	f.byteLength = uint64(length)

	if err := f.IsLengthValid(); err != nil {
		f.byteLength = memo
		return err
	}

	return nil
}

// SetUploadTime set file upload time
func (f *File) SetUploadTime(uploadTime time.Time) error {
	memo := f.uploadTime
	f.uploadTime = uploadTime

	if err := f.IsUploadTimeValid(); err != nil {
		f.uploadTime = memo
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
	// - Byte length between 1 and 500 Mebibyte

	// Mebibyte limit in bytes (500 MB)
	byteLimit := uint64(524288000)
	if f.byteLength < 0 || f.byteLength > byteLimit {
		return exception.NewFieldRange("file_length", "1 MB", "500 MB")
	}

	return nil
}

func (f File) IsUploadTimeValid() error {
	// Validation cases
	// - Must be above 19th century (accept old radio programs as podcasts)

	if f.uploadTime.Year() < 1900 {
		return exception.NewFieldRange("upload_time", "1900 year", "n year")
	}

	return nil
}
