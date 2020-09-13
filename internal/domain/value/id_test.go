package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestID_Set(t *testing.T) {
	id := new(ID)
	if err := id.Set(""); !errors.Is(err, exception.RequiredField) {
		t.Error("invalid id validation, expected error")
	}

	if err := id.Set("123"); err != nil {
		t.Error("invalid id validation, expected nil error")
	}
}
