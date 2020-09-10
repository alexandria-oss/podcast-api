package entity

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
)

func TestCategory_SetName(t *testing.T) {
	c := new(Category)

	err := c.SetName("")
	if !errors.Is(err, exception.RequiredField) {
		t.Error("invalid category set name validation, expected error")
	}

	err = c.SetName("quantum mechanics")
	expStr := "Quantum Mechanics"
	if err != nil {
		t.Error("invalid category set name validation, expected nil error")
	} else if c.GetName() != expStr {
		t.Error("invalid category set name sanitization, expected " + expStr)
	}
}
