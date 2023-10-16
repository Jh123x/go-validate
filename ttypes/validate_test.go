package ttypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_WithError(t *testing.T) {
	tests := map[string]struct {
		withError   error
		expectedErr error
	}{
		"WithError returns correct validate value with no error": {
			withError:   nil,
			expectedErr: nil,
		},
		"WithError returns correct validate value with error": {
			withError:   errTest2,
			expectedErr: errTest2,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := validateWErr.WithError(testCase.expectedErr)
			assert.Equal(t, testCase.expectedErr, validateFn())
		})
	}
}

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
