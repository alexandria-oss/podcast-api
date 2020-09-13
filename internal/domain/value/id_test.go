package value

import (
	"testing"
)

func TestID_Set(t *testing.T) {
	id := new(ID)

	if err := id.Set("123"); err != nil {
		t.Error("invalid id validation, expected nil error")
	}
}
