package value

import "testing"

func TestExplicit_Toggle(t *testing.T) {
	e := new(Explicit)

	if e.IsExplicit() {
		t.Error("explicit", "invalid explicit validation, expected false")
	}

	e.Toggle()
	if !e.IsExplicit() {
		t.Error("explicit", "invalid explicit validation, expected true")
	}
}

func TestExplicit_Set(t *testing.T) {
	e := new(Explicit)

	if e.IsExplicit() {
		t.Error("explicit", "invalid explicit validation, expected false")
	}

	e.Set(true)
	if !e.IsExplicit() {
		t.Error("explicit", "invalid explicit validation, expected true")
	}
}
