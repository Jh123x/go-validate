package options

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/ttypes"

	"github.com/stretchr/testify/assert"
)

// TestWithRequire tests the WithRequire function.
func TestWithRequire(t *testing.T) {
	tests := map[string]struct {
		name        string
		testFn      ttypes.Test
		err         error
		expectedErr error
	}{
		"WithRequire returns correct validate value with no error": {
			testFn:      func() bool { return true },
			err:         nil,
			expectedErr: nil,
		},
		"WithRequire returns correct validate value with error": {
			testFn:      func() bool { return false },
			err:         errTest,
			expectedErr: errTest,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := WithRequire(testCase.testFn, testCase.err)
			assert.Equal(t, testCase.expectedErr, validateFn())
		})
	}
}

// TestIsEmpty tests the IsNotEmpty function.
func TestIsEmpty_string(t *testing.T) {
	tests := map[string]struct {
		value       string
		expectedErr error
	}{
		"empty string should not return error": {
			value:       "",
			expectedErr: nil,
		},
		"non-empty string should return error": {
			value:       "test",
			expectedErr: errs.IsEmptyError,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := IsEmpty(testCase.value)
			assert.Equal(t, testCase.expectedErr, validateFn(), fmt.Sprintf("%v != %v", testCase.expectedErr, validateFn()))
		})
	}
}

// TestIsNotEmpty tests the IsNotEmpty function.
func TestIsNotEmpty_string(t *testing.T) {
	tests := map[string]struct {
		value       string
		expectedErr error
	}{
		"empty string should not return error": {
			value:       "",
			expectedErr: errs.IsNotEmptyErr,
		},
		"non-empty string should return error": {
			value:       "test",
			expectedErr: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := IsNotEmpty(testCase.value)
			assert.Equal(t, testCase.expectedErr, validateFn(), fmt.Sprintf("%v != %v", testCase.expectedErr, validateFn()))
		})
	}
}

// TestIsLength test if the IsLength function works as expected.
func TestIsLength(t *testing.T) {
	tests := map[string]struct {
		arr         []any
		start       int
		end         int
		expectedErr error
	}{
		"array is too short": {
			arr:         testArr,
			start:       len(testArr) + 1,
			end:         len(testArr) + 2,
			expectedErr: errs.InvalidLengthError,
		},
		"array is too long": {
			arr:         testArr,
			start:       len(testArr) - 2,
			end:         len(testArr) - 1,
			expectedErr: errs.InvalidLengthError,
		},
		"array is correct length": {
			arr:         testArr,
			start:       len(testArr) - 2,
			end:         len(testArr) + 2,
			expectedErr: nil,
		},
		"array is at lower boundary": {
			arr:         testArr,
			start:       len(testArr),
			end:         len(testArr) + 1,
			expectedErr: nil,
		},
		"array is at upper boundary": {
			arr:         testArr,
			start:       len(testArr) - 1,
			end:         len(testArr),
			expectedErr: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := IsLength(testCase.arr, testCase.start, testCase.end)
			assert.Equal(t, testCase.expectedErr, validateFn(), fmt.Sprintf("%v != %v", testCase.expectedErr, validateFn()))
		})
	}
}