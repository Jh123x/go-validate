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
		arr         []int
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

// TestOr tests if the Or function works as expected.
func TestOr(t *testing.T) {
	tests := map[string]struct {
		options     []ttypes.Validate
		expectedErr error
	}{
		"all options are valid": {
			options: []ttypes.Validate{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			expectedErr: nil,
		},
		"one option is valid": {
			options: []ttypes.Validate{
				func() error { return errs.OrError },
				func() error { return nil },
				func() error { return errs.OrError },
			},
			expectedErr: nil,
		},
		"no options are valid": {
			options: []ttypes.Validate{
				func() error { return errs.IsNotEmptyErr },
				func() error { return errs.IsEmptyError },
				func() error { return errs.InvalidLengthError },
			},
			expectedErr: errs.OrError,
		},
		"nil options will be skipped": {
			options: []ttypes.Validate{
				nil,
				func() error { return nil },
			},
			expectedErr: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateOr := Or(testCase.options...)
			assert.Equal(t, testCase.expectedErr, validateOr())
		})
	}
}

// TestAnd tests if the Or function works as expected.
func TestAnd(t *testing.T) {
	tests := map[string]struct {
		options     []ttypes.Validate
		expectedErr error
	}{
		"all options are valid": {
			options: []ttypes.Validate{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			expectedErr: nil,
		},
		"one option is valid": {
			options: []ttypes.Validate{
				func() error { return errs.OrError },
				func() error { return nil },
				func() error { return errs.OrError },
			},
			expectedErr: errs.OrError,
		},
		"no options are valid": {
			options: []ttypes.Validate{
				func() error { return errs.IsNotEmptyErr },
				func() error { return errs.IsEmptyError },
				func() error { return errs.InvalidLengthError },
			},
			expectedErr: errs.IsNotEmptyErr,
		},
		"nil options will be skipped": {
			options: []ttypes.Validate{
				nil,
				func() error { return nil },
			},
			expectedErr: nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateOr := And(testCase.options...)
			assert.Equal(t, testCase.expectedErr, validateOr())
		})
	}
}

// TestContains tests if the Contains function works as expected.
func TestContains(t *testing.T) {
	tests := map[string]struct {
		arr         []int
		elem        int
		expectedErr error
	}{
		"array contains element": {
			arr:         testArr,
			elem:        testArr[2],
			expectedErr: nil,
		},
		"array does not contain element": {
			arr:         testArr,
			elem:        4,
			expectedErr: errs.ContainsError,
		},
		"empty array": {
			arr:         []int{},
			elem:        4,
			expectedErr: errs.ContainsError,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := Contains(testCase.arr, testCase.elem)
			assert.Equal(t, testCase.expectedErr, validateFn())
		})
	}
}

func TestVWithRequire(t *testing.T) {
	tests := map[string]struct {
		val         string
		testFn      func(string) bool
		err         error
		expectedErr error
	}{
		"testFn returns true": {
			val:         "test",
			testFn:      func(_ string) bool { return true },
			err:         errTest,
			expectedErr: nil,
		},
		"testFn returns false": {
			val:         "test",
			testFn:      func(_ string) bool { return false },
			err:         errTest,
			expectedErr: errTest,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			validateFn := VWithRequire(tc.testFn, tc.err)
			assert.Equal(t, tc.expectedErr, validateFn(tc.val))
		})
	}
}

func TestVIsNotDefault(t *testing.T) {
	tests := map[string]struct {
		val         string
		expectedErr error
	}{
		"is default string": {
			val:         "",
			expectedErr: errs.IsNotDefaultErr,
		},
		"not default string": {
			val:         "test",
			expectedErr: nil,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VIsNotDefault[string]()(tc.val))
		})
	}
}

func TestVIsDefault(t *testing.T) {
	tests := map[string]struct {
		val         string
		expectedErr error
	}{
		"is default string": {
			val:         "",
			expectedErr: nil,
		},
		"not default string": {
			val:         "test",
			expectedErr: errs.IsDefaultErr,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VIsDefault[string]()(tc.val))
		})
	}
}

func TestVIsEmpty(t *testing.T) {
	tests := map[string]struct {
		val         []rune
		expectedErr error
	}{
		"empty string": {
			val:         []rune(""),
			expectedErr: nil,
		},
		"non-empty string": {
			val:         []rune("test"),
			expectedErr: errs.IsEmptyError,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VIsEmpty[rune](tc.val))
		})
	}
}

func TestVIsNotEmpty(t *testing.T) {
	tests := map[string]struct {
		val         []rune
		expectedErr error
	}{
		"empty string": {
			val:         []rune(""),
			expectedErr: errs.IsNotEmptyErr,
		},
		"non-empty string": {
			val:         []rune("test"),
			expectedErr: nil,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VIsNotEmpty[rune](tc.val))
		})
	}
}

func TestVIsLength(t *testing.T) {
	tests := map[string]struct {
		val         []rune
		minLen      int
		maxLen      int
		expectedErr error
	}{
		"valid length": {
			val:         []rune("test"),
			minLen:      1,
			maxLen:      5,
			expectedErr: nil,
		},
		"invalid length": {
			val:         []rune("test"),
			minLen:      5,
			maxLen:      10,
			expectedErr: errs.InvalidLengthError,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VIsLength[rune](tc.minLen, tc.maxLen)(tc.val))
		})
	}
}

func TestVContains(t *testing.T) {
	tests := map[string]struct {
		val         []rune
		elem        rune
		expectedErr error
	}{
		"contains element": {
			val:         []rune("test"),
			elem:        't',
			expectedErr: nil,
		},
		"does not contain element": {
			val:         []rune("test"),
			elem:        'z',
			expectedErr: errs.ContainsError,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VContains[rune](tc.elem)(tc.val))
		})
	}
}

func TestVOr(t *testing.T) {
	tests := map[string]struct {
		val         []string
		orOptions   []ttypes.ValTest[[]string]
		expectedErr error
	}{
		"all options are valid": {
			val: []string{"test", "test2"},
			orOptions: []ttypes.ValTest[[]string]{
				VIsNotEmpty[string],      // Success
				VIsLength[string](1, 10), // Success
			},
			expectedErr: nil,
		},
		"one option is valid": {
			val: []string{"test", "test2"},
			orOptions: []ttypes.ValTest[[]string]{
				VIsEmpty[string],         // Fail
				VIsLength[string](1, 10), // Success
			},
			expectedErr: nil,
		},
		"no options are valid": {
			val: []string{"test", "test2"},
			orOptions: []ttypes.ValTest[[]string]{
				VIsEmpty[string],          // Fail
				VIsLength[string](10, 20), // Fail
			},
			expectedErr: errs.OrError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VOr[[]string](tc.orOptions...)(tc.val))
		})
	}
}

func TestVAnd(t *testing.T) {
	tests := map[string]struct {
		val         []string
		andOptions  []ttypes.ValTest[[]string]
		expectedErr error
	}{
		"all options are valid": {
			val: []string{"test", "test2"},
			andOptions: []ttypes.ValTest[[]string]{
				VIsNotEmpty[string],      // Success
				VIsLength[string](1, 10), // Success
			},
			expectedErr: nil,
		},
		"one option is valid": {
			val: []string{"test", "test2"},
			andOptions: []ttypes.ValTest[[]string]{
				VIsEmpty[string],         // Fail
				VIsLength[string](1, 10), // Success
			},
			expectedErr: errs.IsEmptyError,
		},
		"no options are valid": {
			val: []string{"test", "test2"},
			andOptions: []ttypes.ValTest[[]string]{
				VIsEmpty[string],          // Fail
				VIsLength[string](10, 20), // Fail
			},
			expectedErr: errs.IsEmptyError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErr, VAnd[[]string](tc.andOptions...)(tc.val))
		})
	}
}
