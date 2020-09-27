package entity

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"testing"
)

func TestVisibility_Set(t *testing.T) {
	v := &Visibility{
		ID:     new(value.ID),
		Policy: new(value.DisplayName),
	}

	if err := v.Set(-1); !errors.Is(err, exception.NotFound) {
		t.Error("visibility", "invalid visibility validation, expected error")
	}

	if err := v.Set(2); err != nil {
		t.Error("visibility", "invalid visibility validation, expected nil error")
	}

	if v.ID.Get() != "2" {
		t.Error("visibility", "invalid visibility validation, expected id 2")
	}

	if v.Policy.Get() != "PUBLIC" {
		t.Error("visibility", "invalid visibility validation, expected PUBLIC")
	}
}
