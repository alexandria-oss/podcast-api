package factory

import (
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/aggregate"
	"github.com/alexandria-oss/podcast-api/internal/domain/entity"
)

// NewPodcast generate a new podcast aggregate root
func NewPodcast(pod *entity.Podcast, media *entity.Media, category *entity.Category) (*aggregate.Podcast, error) {
	// Validations
	if pod == nil {
		return nil, exception.NewRequiredField("podcast_entity")
	} else if media == nil {
		return nil, exception.NewRequiredField("media_entity")
	} else if category == nil {
		return nil, exception.NewRequiredField("category_entity")
	}

	return &aggregate.Podcast{
		Self:     pod,
		Media:    media,
		Category: category,
	}, nil
}
