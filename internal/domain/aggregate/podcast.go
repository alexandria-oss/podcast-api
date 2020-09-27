package aggregate

import "github.com/alexandria-oss/podcast-api/internal/domain/entity"

// Podcast digital audio media content known as podcast
type Podcast struct {
	// Self Podcast attributes
	Self *entity.Podcast `json:"root"`
	// Media Alexandria standard media attributes
	Media *entity.Media `json:"media"`
	// CreatedBy user who created the aggregate
	CreatedBy *entity.User `json:"upload_by"`
	// TODO: Add contrib pool entity
	// Contributors authors who made a contribution
	Contributors []interface{} `json:"contributors"`
	// Category classify aggregate
	Category *entity.Category `json:"category"`
	// Visibility media user-access
	Visibility *entity.Visibility `json:"visibility"`
	// License content creative commons (CC) license
	License *entity.License
}
