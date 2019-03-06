package errors_test

import (
	"testing"

	"github.com/payfazz/authfazz/errors"
)

func TestValidationError(t *testing.T) {
	e := errors.NewValidationError()
	// test add field error
	e.FieldError("f1", "field1 error")
	if len(e.Fields) == 0 {
		t.Fatal("expected fields len greater than 0")
	}
	fe := e.GetFieldError("f1")
	if fe == nil {
		t.Fatal("expected field error f1")
	}
	if fe.Message != "field1 error" {
		t.Fatalf("expected field error message to be '%s'", "field1 error")
	}
	// test update field error
	e.FieldError("f1", "field1 error updated")
	fe = e.GetFieldError("f1")
	if fe == nil {
		t.Fatal("expected field error f1")
	}
	if fe.Message != "field1 error updated" {
		t.Fatalf("expected field error message to be '%s'", "field1 error updated")
	}
	// test clear field errors
	e.ClearFieldErrors()
	if len(e.Fields) != 0 {
		t.Fatal("expected fields len 0")
	}

	var ge error
	ge = e
	if ge.Error() != "Validation error" {
		t.Fatal("expected Error is 'Validation error'")
	}
}
