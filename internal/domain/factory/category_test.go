package factory

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/shared/category"
	"testing"
)

func TestNewCategoryEntity(t *testing.T) {
	if _, err := NewCategoryEntity(-1); !errors.Is(err, exception.NotFound) {
		t.Error("category", "invalid category validation, expected not found error")
	}

	c, err := NewCategoryEntity(category.Misc)
	if err != nil {
		t.Error("category", "invalid category validation, expected nil error")
		return
	}

	if c.Name.Get() != "Miscellaneous" {
		t.Error("category", "invalid category validation, expected Miscellaneous name")
	}
}
