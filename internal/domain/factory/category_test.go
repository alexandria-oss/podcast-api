package factory

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestNewCategoryEntity(t *testing.T) {
	if _, err := NewCategoryEntity(-1); !errors.Is(err, exception.NotFound) {
		t.Error("category", "invalid category validation, expected not found error")
	}

	category, err := NewCategoryEntity(0)
	if err != nil {
		t.Error("category", "invalid category validation, expected nil error")
		return
	}

	if category.Name.Get() != "Miscellaneous" {
		t.Error("category", "invalid category validation, expected Miscellaneous name")
	}
}
