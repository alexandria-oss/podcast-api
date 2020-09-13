package entity

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"testing"
)

func TestCategory_Set(t *testing.T) {
	c := &Category{
		ID:   new(value.ID),
		Name: new(value.Title),
	}

	if err := c.Set(-1); !errors.Is(err, exception.NotFound) {
		t.Error("category", "invalid category validation, expected error")
	}

	if err := c.Set(13); err != nil {
		t.Error("category", "invalid category validation, expected nil error")
	}

	if c.ID.Get() != "13" {
		t.Error("category", "invalid category validation, expected id 13")
	}

	if c.Name.Get() != "Sociology" {
		t.Error("category", "invalid category validation, expected Sociology")
	}
}
