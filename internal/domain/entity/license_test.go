package entity

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"testing"
)

func TestLicense_Set(t *testing.T) {
	l := &License{
		ID:   new(value.ID),
		Name: new(value.DisplayName),
	}

	if err := l.Set(-1); !errors.Is(err, exception.NotFound) {
		t.Error("license", "invalid license validation, expected error")
	}

	if err := l.Set(0); err != nil {
		t.Error("license", "invalid license validation, expected nil error")
	}

	if l.ID.Get() != "0" {
		t.Error("license", "invalid license validation, expected id 0")
	}

	if l.Name.Get() != "CC Zero" {
		t.Error("license", "invalid license validation, expected CC Zero")
	}
}
