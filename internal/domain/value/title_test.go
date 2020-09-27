package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestTitle_Set(t *testing.T) {
	title := new(Title)

	// Populate fake title
	str := ""
	for i := 0; i < 513; i++ {
		str += "x"
	}

	if err := title.Set(str); !errors.Is(err, exception.FieldRange) {
		t.Error("title", "invalid title validation, expected error")
	}

	if err := title.Set("Newtonian Mechanics"); err != nil {
		t.Error("title", "invalid title validation, expected nil error")
	}

	if err := title.Set("Euclidean Vectors"); !errors.Is(err, exception.AlreadyExists) {
		t.Error("title", "invalid title validation, expected error")
	}
}
