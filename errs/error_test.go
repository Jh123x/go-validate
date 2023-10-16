package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateError_Error tests the NewValidateError and Error method of ValidateError.
func TestValidateError_Error(t *testing.T) {
	tests := map[string]struct {
		fieldName   string
		errMsg      string
		expectedMsg string
	}{
		"error is correct": {
			fieldName:   "test",
			errMsg:      "test error",
			expectedMsg: "[validation error] error validating test:test error",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			ve := NewValidateError(testCase.fieldName, testCase.errMsg)
			assert.Equal(t, testCase.expectedMsg, ve.Error())
		})
	}
}
