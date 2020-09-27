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

package model

import (
	"time"
)

// Podcast holds Alexandria standard Podcast read-only primitive attributes
type Podcast struct {
	// ID domain unique identifier
	ID string `json:"id"`
	// DisplayName mutable title
	DisplayName string `json:"display_name"`
	// Title immutable title
	Title string `json:"title"`
	// Description detailed description
	Description string `json:"description"`
	// UploadBy user's unique identifier who uploaded Podcast
	UploadBy string `json:"upload_by"`
	// Explicit contains explicit content
	Explicit bool `json:"explicit"`
	// CreateTime create time timestamp
	CreateTime time.Time `json:"create_time"`
	// UpdateTime update time timestamp
	UpdateTime time.Time `json:"update_time"`
	// Active content is available
	Active bool `json:"active"`
	// Lang language code in ISO 639-1
	Lang string `json:"lang"`
	// Keywords list of keywords
	Keywords []string `json:"keywords"`
	// Visibility user-access
	Visibility string `json:"visibility"`
	// License content license like CC
	License string `json:"license"`
	// Cover image url
	Cover string `json:"cover"`
	// Canvas image or small loop-video background url
	Canvas string `json:"canvas"`
	// ContentURL content url
	ContentURL string `json:"content_url"`
	// RssID RSS feed item GUID
	RssID string `json:"rss_id"`
	// DurationSec total duration in seconds
	DurationSec int64 `json:"duration_sec"`
	// Duration total duration in hh:mm:ss format
	Duration string `json:"duration"`
	// Episode episode number
	Episode int32 `json:"episode"`
	// EpisodeType kind of episode
	EpisodeType string `json:"episode_type"`
	// Transcript content's transcript url
	Transcript string `json:"transcript"`
	// Category category association
	Category struct {
		// ID foreign Category correlation unique identifier
		ID string `json:"id"`
		// Name denormalized Category name field
		Name string `json:"name"`
	} `json:"category"`
	// File binary file
	File struct {
		// Extension file extension
		Extension string `json:"extension"`
		// Length file byte length properties
		Length int64 `json:"length"`
		// UploadTime file upload time
		UploadTime time.Time `json:"upload_time"`
	} `json:"file"`
	// SocialMetadata social interaction(s)
	SocialMetadata struct {
		// TotalViews total view counter
		TotalViews int64 `json:"total_views"`
		// TotalPositiveReactions total positive social reactions
		TotalPositiveReactions int64 `json:"total_positive_reactions"`
		// TotalNegativeReactions total negative social reactions
		TotalNegativeReactions int64 `json:"total_negative_reactions"`
		// LikenessRatio positive reaction percentage
		LikenessRatio float64 `json:"likeness_ratio"`
	} `json:"social_metadata"`
}
