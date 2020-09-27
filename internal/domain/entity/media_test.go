package entity

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"github.com/alexandria-oss/podcast-api/internal/domain/value"
	"testing"
)

func TestMedia_IsValid(t *testing.T) {
	// init required data
	media := new(Media)
	media.InternalID = new(value.ID)
	media.Cover = new(value.ExternalMedia)
	media.Canvas = new(value.ExternalMedia)

	media.Start()

	// Init required data
	media.ID = new(value.NanoID)
	media.DisplayName = new(value.DisplayName)
	media.Title = new(value.Title)
	media.Language = new(value.Language)
	media.Metadata = value.NewMetadata()
	if err := media.IsValid(); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	_ = media.ID.Generate()
	if err := media.IsValid(); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	_ = media.DisplayName.Set("Quantum Mechanics")
	if err := media.IsValid(); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	_ = media.Title.Set("quantum mechanics universe: ed 2")
	if err := media.IsValid(); !errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected required field error")
	}

	_ = media.Language.Set("es")
	if err := media.IsValid(); errors.Is(err, exception.RequiredField) {
		t.Error("media", "invalid media validation, expected not a required field error")
	}
}
