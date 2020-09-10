package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestKeyword_SetKeyword(t *testing.T) {
	k := new(Keyword)

	if err := k.SetKeyword(""); !errors.Is(err, exception.RequiredField) {
		t.Error("invalid keyword validation, expected error")
	}

	if err := k.SetKeyword("quantum"); err != nil {
		t.Error("invalid keyword validation, expected nil error")
	}
}
