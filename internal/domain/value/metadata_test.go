package value

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"testing"
	"time"
)

func TestMetadata_SetCreateTime(t *testing.T) {
	m := new(Metadata)

	if err := m.SetCreateTime(time.Date(2018, time.December, 1, 0, 0,
		0, 0, time.UTC)); !errors.Is(err, exception.FieldRange) {
		t.Error("create_time", "invalid create time validation, expected error")
	}

	if err := m.SetCreateTime(time.Date(3000, time.January, 1, 0, 0,
		0, 0, time.UTC)); !errors.Is(err, exception.FieldRange) {
		t.Error("create_time", "invalid create time validation, expected error")
	}

	if err := m.SetCreateTime(time.Date(2020, time.January, 1, 0, 0,
		0, 0, time.UTC)); err != nil {
		t.Error("create_time", "invalid create time validation, expected nil error")
	}
}

func TestMetadata_SetUpdateTime(t *testing.T) {
	m := new(Metadata)

	if err := m.SetUpdateTime(time.Date(2018, time.December, 1, 0, 0,
		0, 0, time.UTC)); !errors.Is(err, exception.FieldRange) {
		t.Error("update_time", "invalid update time validation, expected error")
	}

	if err := m.SetUpdateTime(time.Date(3000, time.January, 1, 0, 0,
		0, 0, time.UTC)); !errors.Is(err, exception.FieldRange) {
		t.Error("update_time", "invalid update time validation, expected error")
	}

	if err := m.SetUpdateTime(time.Date(2020, time.January, 1, 0, 0,
		0, 0, time.UTC)); err != nil {
		t.Error("update_time", "invalid update time validation, expected nil error")
	}
}

func TestMetadata_IsActive(t *testing.T) {
	m := new(Metadata)

	if state := m.IsActive(); state {
		t.Error("active", "invalid active state, expected false")
	}

	m.SetState(false)
	if state := m.IsActive(); state {
		t.Error("active", "invalid active state, expected false")
	}

	m.ToggleState()
	if state := m.IsActive(); !state {
		t.Error("active", "invalid active state, expected true")
	}
}
