package entity

import (
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"github.com/hashicorp/go-multierror"
)

// User identity
type User struct {
	// ID foreign unique identifier
	ID *value.ID
	// DisplayName denormalized User's display name field
	DisplayName *value.DisplayName
	// Image denormalized User's image field
	Image *value.ExternalMedia
}

// Validators

// IsValid validate User
func (c User) IsValid() error {
	// Validation cases
	// * Foreign data
	// - Required: ID, Display Name
	if c.DisplayName.Get() == "" {
		return exception.NewRequiredField("user_display_name")
	} else if c.ID.Get() == "" {
		return exception.NewRequiredField("user_id")
	}

	// Value Object
	var valueErr error
	if err := c.DisplayName.IsValid(); err != nil {
		valueErr = multierror.Append(valueErr, err)
	}
	if c.Image != nil {
		if err := c.Image.IsValid(); err != nil {
			valueErr = multierror.Append(valueErr, err)
		}
	}

	return valueErr
}
