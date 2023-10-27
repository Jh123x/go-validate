package ttypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWithRequire tests the WithRequire function.
func TestValidate_WithError(t *testing.T) {
	tests := map[string]struct {
		validate    Validate
		withError   error
		expectedErr error
	}{
		"WithError returns correct validate value with no error": {
			validate:    validateWErr,
			withError:   nil,
			expectedErr: nil,
		},
		"WithError returns correct validate value with error": {
			validate:    validateWErr,
			withError:   errTest2,
			expectedErr: errTest2,
		},
		"WithError returns no error withError": {
			validate:    validateWNil,
			withError:   errTest2,
			expectedErr: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := testCase.validate.WithError(testCase.expectedErr)
			assert.Equal(t, testCase.expectedErr, validateFn())
		})
	}
}

// TestIsEmpty tests the IsNotEmpty function.
func TestValidate_Not(t *testing.T) {
	tests := map[string]struct {
		validate    Validate
		err         error
		expectedErr error
	}{
		"Not on err validate should negate the error": {
			validate:    validateWErr,
			err:         errTest,
			expectedErr: nil,
		},
		"Not on nil validate should throw error": {
			validate:    validateWNil,
			err:         errTest,
			expectedErr: errTest,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := testCase.validate.Not(testCase.err)
			assert.Equal(t, testCase.expectedErr, validateFn())
		})
	}
}
