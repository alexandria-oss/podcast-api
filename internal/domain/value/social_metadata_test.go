package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestSocialMetadata_SetTotalViews(t *testing.T) {
	s := new(SocialMetadata)

	if err := s.SetTotalViews(-1); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid social metadata views validation, expected error")
	}

	if err := s.SetTotalViews(0); err != nil {
		t.Error("invalid social metadata views validation, expected nil error")
	}

	if err := s.SetTotalViews(5099345); err != nil {
		t.Error("invalid social metadata views validation, expected nil error")
	}
}

func TestSocialMetadata_IncrementViews(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalViews(100)

	s.IncrementViews()

	if s.GetTotalViews() != 101 {
		t.Error("invalid social metadata views validation, expected 101")
	}
}

func TestSocialMetadata_DecreaseViews(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalViews(100)

	s.DecreaseViews()

	if s.GetTotalViews() != 99 {
		t.Error("invalid social metadata views validation, expected 99")
	}
}

func TestSocialMetadata_SetTotalPositiveReactions(t *testing.T) {
	s := new(SocialMetadata)

	if err := s.SetTotalPositiveReactions(-1); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid social metadata positive reaction validation, expected error")
	}

	if err := s.SetTotalPositiveReactions(0); err != nil {
		t.Error("invalid social metadata positive reaction validation, expected nil error")
	}

	if err := s.SetTotalPositiveReactions(148765); err != nil {
		t.Error("invalid social metadata positive reaction validation, expected nil error")
	}
}

func TestSocialMetadata_IncrementPositive(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalPositiveReactions(1024)

	s.IncrementPositive()

	if s.GetTotalPositiveReactions() != 1025 {
		t.Error("invalid social metadata positive reaction validation, expected 1025")
	}
}

func TestSocialMetadata_DecreasePositive(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalPositiveReactions(1024)

	s.DecreasePositive()

	if s.GetTotalPositiveReactions() != 1023 {
		t.Error("invalid social metadata positive reaction validation, expected 1023")
	}
}

func TestSocialMetadata_SetTotalNegativeReactions(t *testing.T) {
	s := new(SocialMetadata)

	if err := s.SetTotalNegativeReactions(-1); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid social metadata negative reaction validation, expected error")
	}

	if err := s.SetTotalNegativeReactions(0); err != nil {
		t.Error("invalid social metadata negative reaction validation, expected nil error")
	}

	if err := s.SetTotalNegativeReactions(23654); err != nil {
		t.Error("invalid social metadata negative reaction validation, expected nil error")
	}
}

func TestSocialMetadata_IncrementNegative(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalNegativeReactions(1024)

	s.IncrementNegative()

	if s.GetTotalNegativeReactions() != 1025 {
		t.Error("invalid social metadata negative reaction validation, expected 1025")
	}
}

func TestSocialMetadata_DecreaseNegative(t *testing.T) {
	s := new(SocialMetadata)

	_ = s.SetTotalNegativeReactions(1024)

	s.DecreaseNegative()

	if s.GetTotalNegativeReactions() != 1023 {
		t.Error("invalid social metadata negative reaction validation, expected 1023")
	}
}

func TestSocialMetadata_CalculateLikeness(t *testing.T) {
	s := new(SocialMetadata)
	s.CalculateLikeness()

	if s.GetLikenessRatio() != 0 {
		t.Error("invalid social metadata likeness ratio validation, expected 0")
	}

	_ = s.SetTotalPositiveReactions(50)
	_ = s.SetTotalNegativeReactions(50)

	s.CalculateLikeness()
	if s.GetLikenessRatio() != 50 {
		t.Error("invalid social metadata likeness ratio validation, expected 50%")
	}
}

func TestSocialMetadata_SetLikenessRatio(t *testing.T) {
	s := new(SocialMetadata)

	if err := s.SetLikenessRatio(-1); !errors.Is(err, exception.FieldRange) {
		t.Error("likeness_ratio", "invalid social metadata likeness ratio validation, expected error")
	}

	if err := s.SetLikenessRatio(101); !errors.Is(err, exception.FieldRange) {
		t.Error("likeness_ratio", "invalid social metadata likeness ratio validation, expected error")
	}

	if err := s.SetLikenessRatio(99); err != nil {
		t.Error("likeness_ratio", "invalid social metadata likeness ratio validation, expected nil error")
	}
}
