package models

import (
	"testing"
)

func TestNewField(t *testing.T) {
	t.Log("A user with empty name or empty fieldType ")
	field, err := NewField("", "")
	if err == nil {
		t.Error("error should be raised if name or field is empty")
	}
	t.Log("Test to see if a new field can be created properly")
	name := "textbox"
	fieldType := "textbox"
	field, err = NewField(name, fieldType)
	if len(field.ID) <= 0 || field.Name != name || field.FieldType != fieldType {
		t.Error("Failure when creating new field")
	}
}
