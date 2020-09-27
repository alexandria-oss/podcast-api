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
)

// SocialMetadata social interaction metadata
type SocialMetadata struct {
	// totalViews total view counter
	totalViews uint64
	// totalPositiveReactions total positive social reactions
	totalPositiveReactions uint64
	// totalNegativeReactions total negative social reactions
	totalNegativeReactions uint64
	// likenessRatio social reaction percentage
	likenessRatio float64
	fieldName     string
}

// SetFieldName override field name for top-operations
func (s *SocialMetadata) SetFieldName(field string) {
	if field != "" {
		s.fieldName = field
	}
}

// GetFieldName get the url object field name
func (s SocialMetadata) GetFieldName() string {
	if s.fieldName != "" {
		return s.fieldName
	}

	// Default value
	return "social_metadata"
}

// SetTotalViews set total views counter
func (s *SocialMetadata) SetTotalViews(delta int64) error {
	if delta < 0 {
		return exception.NewFieldRange(s.GetFieldName()+"_total_views", "0", "n")
	}

	s.totalViews = uint64(delta)
	return nil
}

// IncrementViews increment total views by one
func (s *SocialMetadata) IncrementViews() {
	s.totalViews++
}

// DecreaseViews decrease total views by one
func (s *SocialMetadata) DecreaseViews() {
	if s.totalViews >= 1 {
		s.totalViews--
	}
}

// GetTotalViews get total views counter
func (s SocialMetadata) GetTotalViews() uint64 {
	return s.totalViews
}

// SetTotalPositiveReactions set total positive reactions counter
func (s *SocialMetadata) SetTotalPositiveReactions(delta int64) error {
	if delta < 0 {
		return exception.NewFieldRange(s.GetFieldName()+"_total_positive_reactions", "0", "n")
	}

	s.totalPositiveReactions = uint64(delta)
	return nil
}

// IncrementPositive increment total positive reactions by one
func (s *SocialMetadata) IncrementPositive() {
	s.totalPositiveReactions++
}

// DecreasePositive decrease total positive reactions by one
func (s *SocialMetadata) DecreasePositive() {
	if s.totalPositiveReactions >= 1 {
		s.totalPositiveReactions--
	}
}

// GetTotalPositiveReactions get total positive reactions counter
func (s SocialMetadata) GetTotalPositiveReactions() uint64 {
	return s.totalPositiveReactions
}

// SetTotalNegativeReactions set total negative reactions counter
func (s *SocialMetadata) SetTotalNegativeReactions(delta int64) error {
	if delta < 0 {
		return exception.NewFieldRange(s.GetFieldName()+"_total_negative_reactions", "0", "n")
	}

	s.totalNegativeReactions = uint64(delta)
	return nil
}

// IncrementNegative increment total negative reactions by one
func (s *SocialMetadata) IncrementNegative() {
	s.totalNegativeReactions++
}

// DecreaseNegative decrease total negative reactions by one
func (s *SocialMetadata) DecreaseNegative() {
	if s.totalNegativeReactions >= 1 {
		s.totalNegativeReactions--
	}
}

// GetTotalNegativeReactions get total negative reactions counter
func (s SocialMetadata) GetTotalNegativeReactions() uint64 {
	return s.totalNegativeReactions
}

// CalculateLikeness calculate and set likeness ratio value
func (s *SocialMetadata) CalculateLikeness() {
	totalReactions := float64(s.totalNegativeReactions + s.totalPositiveReactions)

	if totalReactions >= 1 {
		s.likenessRatio = (float64(s.totalPositiveReactions) / totalReactions) * 100 // Get percentage
	}
}

// GetLikenessRatio get likeness ratio value
func (s SocialMetadata) GetLikenessRatio() float64 {
	return s.likenessRatio
}

// SetLikenessRatio set likeness ratio
func (s *SocialMetadata) SetLikenessRatio(delta float64) error {
	if delta < 0 || delta > 100 {
		return exception.NewFieldRange(s.GetFieldName()+"_likeness_ratio", "0", "100")
	}

	s.likenessRatio = delta
	return nil
}
