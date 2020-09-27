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

package factory

import (
	"github.com/alexandria-oss/podcast-api/internal/domain/entity"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"github.com/hashicorp/go-multierror"
	"time"
)

// NewMediaEntity create a new media entity
func NewMediaEntity(displayName, title, lang string) (*entity.Media, error) {
	id := new(value.NanoID)
	if err := id.Generate(); err != nil {
		return nil, err
	}

	var valuerErr error
	displayNameV := new(value.DisplayName)
	if err := displayNameV.Set(displayName); err != nil {
		valuerErr = multierror.Append(valuerErr, err)
	}

	titleV := new(value.Title)
	if err := titleV.Set(title); err != nil {
		valuerErr = multierror.Append(valuerErr, err)
	}

	langV := new(value.Language)
	if err := langV.Set(lang); err != nil {
		valuerErr = multierror.Append(valuerErr, err)
	}

	// Metadata auto-generation
	metadata := new(value.Metadata)
	metadata.SetFieldName("media")
	metadata.SetState(true)

	if err := metadata.SetCreateTime(time.Now()); err != nil {
		valuerErr = multierror.Append(valuerErr, err)
	}
	if err := metadata.SetUpdateTime(time.Now()); err != nil {
		valuerErr = multierror.Append(valuerErr, err)
	}

	social := new(value.SocialMetadata)
	social.SetFieldName("media")

	if valuerErr != nil {
		// Avoid create ref
		return nil, valuerErr
	}

	media := &entity.Media{
		InternalID:  new(value.ID),
		ID:          id,
		DisplayName: displayNameV,
		Title:       titleV,
		Description: new(value.Description),
		Explicit:    new(value.Explicit),
		Language:    langV,
		Keywords:    make([]*value.Keyword, 0),
		Cover:       new(value.ExternalMedia),
		Canvas:      new(value.ExternalMedia),
		Metadata:    metadata,
		Social:      social,
	}

	media.SetFieldNames()
	media.SetExternalMediaPolicy()

	if err := media.IsValid(); err != nil {
		return nil, err
	}

	return media, nil
}
