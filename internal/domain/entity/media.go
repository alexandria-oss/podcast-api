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

package entity

import (
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"time"
)

// Media holds Alexandria standard media attributes
type Media struct {
	// ID domain unique identifier
	ID string `json:"id"`
	// DisplayName mutable media title
	DisplayName string `json:"display_name"`
	// Title immutable media title
	Title string `json:"title"`
	// Description complete media description
	Description string `json:"description"`
	// UploadBy user's unique identifier who uploaded media
	UploadBy string `json:"upload_by"`
	// Explicit podcasts contains explicit content
	Explicit bool `json:"explicit"`
	// CreateTime create time timestamp
	CreateTime time.Time `json:"create_time"`
	// UpdateTime update time timestamp
	UpdateTime time.Time `json:"update_time"`
	// Active content is available
	Active bool `json:"active"`
	// Lang language code in ISO 639-1
	Lang string `json:"lang"`
	// Visibility media user-access
	Visibility string `json:"visibility"`
	// License content license like CC
	License string `json:"license"`
	// TotalViews total view counter
	TotalViews int64 `json:"total_views"`
	// File binary properties
	File value.File `json:"file"`
	// Content external content URLs
	Content value.Content `json:"content"`
}
