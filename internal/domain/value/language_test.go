package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestLanguage_Set(t *testing.T) {
	lang := new(Language)

	if err := lang.Set("a"); !errors.Is(err, exception.FieldRange) {
		t.Error("language", "invalid lang validation, expected error")
	}

	if err := lang.Set("abcdefgh"); !errors.Is(err, exception.FieldRange) {
		t.Error("language", "invalid lang validation, expected error")
	}

	if err := lang.Set("ek"); !errors.Is(err, exception.FieldFormat) {
		t.Error("language", "invalid lang validation, expected error")
	}

	if err := lang.Set("en"); err != nil {
		t.Error("language", "invalid lang validation, expected nil error")
	}

	if err := lang.Set("en-US"); err != nil {
		t.Error("language", "invalid lang validation, expected nil error")
	}
}
