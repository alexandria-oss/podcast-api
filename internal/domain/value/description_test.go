package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestDescription_Set(t *testing.T) {
	d := new(Description)

	// Populate fake description
	str := ""
	for i := 0; i < 2049; i++ {
		str += "x"
	}
	if err := d.Set(str); !errors.Is(err, exception.FieldRange) {
		t.Error("description", "invalid description validation, expected error")
	}

	if err := d.Set("Enjoy a simple briefing about the quantum mechanics universe."); err != nil {
		t.Error("description", "invalid description validation, expected nil error")
	}
}
