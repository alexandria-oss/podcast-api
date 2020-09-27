package factory

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestNewMediaEntity(t *testing.T) {
	if _, err := NewMediaEntity("", "", ""); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	if _, err := NewMediaEntity("Quantum Mechanics", "", ""); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	if _, err := NewMediaEntity("quantum mechanics",
		"quantum universe at scale", ""); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	media, err := NewMediaEntity("quantum mechanics",
		"quantum universe at scale", "en")
	if err != nil {
		t.Error("media", "invalid media validation, expected nil error")
		return
	}

	if media.Title.Get() != "Quantum Universe At Scale" {
		t.Error("media", "invalid media validation, expected Quantum Universe At Scale title")
	}
}
