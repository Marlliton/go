package fail

import (
	"errors"
	"testing"
)

func TestFail_ValidationError(t *testing.T) {
	err := WithValidationError("validation_error", "just fail with validation error")
	var expectedErr *ValidationError
	if !errors.As(err, &expectedErr) {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}
