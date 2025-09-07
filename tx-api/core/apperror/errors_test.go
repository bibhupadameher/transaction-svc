package apperror

import (
	"errors"
	"testing"
)

func TestAppError_Error_WithDetails(t *testing.T) {
	e := AppError{
		Code:    "123",
		Message: "Something went wrong",
		Type:    "Validation",
		Details: "field X is required",
	}
	got := e.Error()
	wantSubstring := "[123] Something went wrong (Validation): field X is required"

	if got != wantSubstring {
		t.Errorf("expected %q, got %q", wantSubstring, got)
	}
}

func TestAppError_Error_WithoutDetails(t *testing.T) {
	e := AppError{
		Code:    "123",
		Message: "Something went wrong",
		Type:    "Validation",
	}
	got := e.Error()
	wantSubstring := "[123] Something went wrong (Validation)"

	if got != wantSubstring {
		t.Errorf("expected %q, got %q", wantSubstring, got)
	}
}

func TestIs(t *testing.T) {
	e := New("401", "Auth", "Unauthorized", 401)
	err := &e

	if !Is(err, "401") {
		t.Errorf("expected Is() true for code 401")
	}
	if Is(err, "999") {
		t.Errorf("expected Is() false for code 999")
	}
	if Is(errors.New("some error"), "401") {
		t.Errorf("expected Is() false for non-AppError")
	}
}
