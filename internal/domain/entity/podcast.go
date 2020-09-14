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

// Podcast Podcasts entity
type Podcast struct {
	// RssID RSS feed item GUID
	RssID string
	// DurationSec total duration in seconds
	DurationSec int64
	// Duration total duration in hh:mm:ss format
	Duration string
	// Episode episode number
	Episode int32
	// EpisodeType kind of episode
	EpisodeType string
	// Transcript content's transcript url
	Transcript string
}
