package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestDisplayName_Set(t *testing.T) {
	display := new(DisplayName)

	if err := display.Set(""); !errors.Is(err, exception.RequiredField) {
		t.Error("invalid display name validation, expected error")
	}

	// Populate fake data
	str := ""
	for i := 0; i < 257; i++ {
		str += "x"
	}

	if err := display.Set(str); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid display name validation, expected error")
	}

	if err := display.Set("Quantum Mechanics for people in a hurry"); err != nil {
		t.Error("invalid display name validation, expected nil error")
	}
}
