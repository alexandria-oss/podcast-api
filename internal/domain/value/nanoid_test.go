package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestNanoID_Set(t *testing.T) {
	id := new(NanoID)

	// Populate fake id
	str := ""
	for i := 0; i < 65; i++ {
		str += "x"
	}

	if err := id.Set(str); !errors.Is(err, exception.FieldRange) {
		t.Error("invalid id validation, expected error")
	}

	if err := id.Set("123"); err != nil {
		t.Error("invalid id validation, expected nil error")
	}
}
