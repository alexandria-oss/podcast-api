package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestKeyword_SetKeyword(t *testing.T) {
	k := new(Keyword)

	// Populate fake keyword
	str := ""
	for i := 0; i < 129; i++ {
		str += "x"
	}

	if err := k.SetKeyword(str); !errors.Is(err, exception.FieldRange) {
		t.Error("keyword", "invalid keyword validation, expected error")
	}

	if err := k.SetKeyword("quantum"); err != nil {
		t.Error("keyword", "invalid keyword validation, expected nil error")
	}
}
