package wrapper

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestValueWrapper(t *testing.T) {
	tests := map[string]struct {
		value       int
		options     ttypes.ValTest[int]
		expectedErr error
	}{
		"success": {
			value:       1,
			options:     options.VWithRequire(func(v int) bool { return true }, fmt.Errorf("value is empty")),
			expectedErr: nil,
		},
		"fail": {
			value:       0,
			options:     options.VWithRequire(func(v int) bool { return false }, fmt.Errorf("value is empty")),
			expectedErr: fmt.Errorf("value is empty"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value)
			err := valueWrapper.WithOptions(tc.options).Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestOptionCompatibility_SingleValues(t *testing.T) {
	tests := map[string]struct {
		value               int
		options             ttypes.ValTest[int]
		expectedValidateErr error
	}{
		"With Require success": {
			value:   1,
			options: options.VWithRequire(func(v int) bool { return true }, fmt.Errorf("value is error")),
		},
		"With Require fail": {
			value:               0,
			options:             options.VWithRequire(func(v int) bool { return false }, fmt.Errorf("value is error")),
			expectedValidateErr: fmt.Errorf("value is error"),
		},
		"IsNotEmpty success": {
			value:   1,
			options: options.VIsNotEmpty[int](),
		},
		"IsNotEmpty fail": {
			value:               0,
			options:             options.VIsNotEmpty[int](),
			expectedValidateErr: errs.IsNotEmptyErr,
		},
		"IsEmpty success": {
			value:   0,
			options: options.VIsEmpty[int](),
		},
		"IsEmpty fail": {
			value:               1,
			options:             options.VIsEmpty[int](),
			expectedValidateErr: errs.IsEmptyError,
		},
		"nil option should return err": {
			value:               1,
			options:             nil,
			expectedValidateErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value)
			valueWrapper = valueWrapper.WithOptions(tc.options)
			option := valueWrapper.ToOption()
			assert.Equal(t, tc.expectedValidateErr, option())
			assert.Equal(t, tc.expectedValidateErr, valueWrapper.Validate())
		})
	}
}

func TestOptionCompatibility_ArrValues(t *testing.T) {
	tests := map[string]struct {
		value               []int
		options             ttypes.ValTest[[]int]
		expectedValidateErr error
	}{

		"IsLength success": {
			value:   []int{1, 2, 3},
			options: options.VIsLength[int](1, 3),
		},
		"IsLength fail": {
			value:               []int{1, 2, 3},
			options:             options.VIsLength[int](4, 5),
			expectedValidateErr: errs.InvalidLengthError,
		},
		"Contains success": {
			value:   []int{1, 2, 3},
			options: options.VContains(1),
		},
		"Contains fail": {
			value:               []int{1, 2, 3},
			options:             options.VContains(4),
			expectedValidateErr: errs.ContainsError,
		},
		"and success": {
			value: []int{1, 2, 3},
			options: options.VAnd(
				options.VIsLength[int](1, 3),
				options.VContains(1),
			),
		},
		"and fail when 1 fails": {
			value: []int{1, 2, 3},
			options: options.VAnd(
				options.VIsLength[int](2, 3), // Success
				options.VContains(4),         // Fail
			),
			expectedValidateErr: errs.ContainsError,
		},
		"and fail when all fails": {
			value: []int{1, 2, 3},
			options: options.VAnd(
				options.VIsLength[int](4, 5), // Fail
				options.VContains(4),         // Fail
			),
			expectedValidateErr: errs.InvalidLengthError,
		},
		"or success": {
			value: []int{1, 2, 3},
			options: options.VOr(
				options.VIsLength[int](4, 5), // Fail
				options.VContains(1),         // Success
			),
		},
		"or success when all success": {
			value: []int{1, 2, 3, 4, 5},
			options: options.VOr(
				options.VIsLength[int](4, 5), // Success
				options.VContains(4),         // Success
			),
		},
		"or fail": {
			value: []int{1, 2, 3},
			options: options.VOr(
				options.VIsLength[int](4, 5), // Fail
				options.VContains(4),         // Fail
			),
			expectedValidateErr: errs.OrError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value)
			option := valueWrapper.WithOptions(tc.options).ToOption()
			assert.Equal(t, tc.expectedValidateErr, option())
		})
	}
}

func TestOptionCompatibility_StringOptions(t *testing.T) {
	tests := map[string]struct {
		value               string
		options             ttypes.ValTest[string]
		expectedValidateErr error
	}{
		"IsValidURI success": {
			value:   "https://www.google.com",
			options: options.VIsValidURI,
		},
		"IsValidURI fail": {
			value:               "invalid url",
			options:             options.VIsValidURI,
			expectedValidateErr: errs.InvalidURIError,
		},
		"IsValidEmail success": {
			value:   "test@test.com",
			options: options.VIsValidEmail,
		},
		"IsValidEmail fail": {
			value:               "invalid email",
			options:             options.VIsValidEmail,
			expectedValidateErr: errs.InvalidEmailError,
		},
		"IsValidJSON success": {
			value:   `{"key": "value"}`,
			options: options.VIsValidJson,
		},
		"IsValidJSON fail": {
			value:               "invalid json",
			options:             options.VIsValidJson,
			expectedValidateErr: errs.InvalidJsonError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value).WithOptions(tc.options)
			assert.Equal(t, tc.expectedValidateErr, valueWrapper.Validate())
		})
	}
}
