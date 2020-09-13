// Copyright 2020 The Alexandria Foundation
//
// Licensed under the GNU Affero General Public License, Version 3.0 (the "License");
// you may not use this FileMetadata except in compliance with the License.
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

// FileMetadata binary extension
type FileMetadata struct {
	// extension FileMetadata extension
	extension string
	// byteLength FileMetadata byte length
	byteLength uint64
	// uploadTime FileMetadata upload time
	uploadTime time.Time
}

// GetExtension get FileMetadata extension
func (f FileMetadata) GetExtension() string {
	return f.extension
}

// GetByteLength get FileMetadata length in bytes
func (f FileMetadata) GetByteLength() uint64 {
	return f.byteLength
}

// GetUploadTime get FileMetadata upload time
func (f FileMetadata) GetUploadTime() time.Time {
	return f.uploadTime
}

// SetExtension set FileMetadata extension
func (f *FileMetadata) SetExtension(extension string) error {
	memo := f.extension
	f.extension = strings.ToLower(extension)

	if err := f.IsExtensionValid(); err != nil {
		f.extension = memo
		return err
	}

	return nil
}

// SetByteLength set FileMetadata length in bytes
func (f *FileMetadata) SetByteLength(length int64) error {
	// Avoid uint overflow
	if length < 0 {
		return exception.NewFieldRange("FileMetadata_length", "1 MB", "500 MB")
	}

	memo := f.byteLength
	f.byteLength = uint64(length)

	if err := f.IsLengthValid(); err != nil {
		f.byteLength = memo
		return err
	}

	return nil
}

// SetUploadTime set FileMetadata upload time
func (f *FileMetadata) SetUploadTime(uploadTime time.Time) error {
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
func (f FileMetadata) IsValid() error {
	if err := f.IsExtensionValid(); err != nil {
		return err
	} else if err := f.IsLengthValid(); err != nil {
		return err
	}

	return nil
}

// IsExtensionValid validate FileMetadata extension
func (f FileMetadata) IsExtensionValid() error {
	// Validation cases
	// - Audio extension .mp3 and .flac only
	if f.extension != "" {
		if f.extension != "mp3" && f.extension != "flac" {
			return exception.NewFieldFormat("FileMetadata_extension", "mp3 or flac")
		}

		return nil
	}

	return nil
}

// IsLengthValid validate FileMetadata byte length
func (f FileMetadata) IsLengthValid() error {
	// Validation cases
	// - Byte length between 1 and 500 Mebibyte

	// Mebibyte limit in bytes (500 MB)
	byteLimit := uint64(524288000)
	if f.byteLength < 0 || f.byteLength > byteLimit {
		return exception.NewFieldRange("FileMetadata_length", "1 MB", "500 MB")
	}

	return nil
}

func (f FileMetadata) IsUploadTimeValid() error {
	// Validation cases
	// - Must be above 19th century (accept old radio programs as podcasts)

	if f.uploadTime.Year() < 1900 {
		return exception.NewFieldRange("upload_time", "1900 year", "n year")
	}

	return nil
}
