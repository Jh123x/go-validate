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
			valueWrapper := NewValueWrapper[int]()
			err := valueWrapper.WithOptions(tc.options).Validate(tc.value)
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
		"IsNotDefault success": {
			value:   1,
			options: options.VIsNotDefault[int](),
		},
		"IsNotDefault fail": {
			value:               0,
			options:             options.VIsNotDefault[int](),
			expectedValidateErr: errs.IsNotDefaultErr,
		},
		"IsDefault success": {
			value:   0,
			options: options.VIsDefault[int](),
		},
		"IsDefault fail": {
			value:               1,
			options:             options.VIsDefault[int](),
			expectedValidateErr: errs.IsDefaultErr,
		},
		"nil option should return no err": {
			value:               1,
			options:             nil,
			expectedValidateErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper[int]()
			valueWrapper = valueWrapper.WithOptions(tc.options)
			option := valueWrapper.ToOption(tc.value)
			assert.Equal(t, tc.expectedValidateErr, option())
			assert.Equal(t, tc.expectedValidateErr, valueWrapper.Validate(tc.value))
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
		"is empty success": {
			value:   []int{},
			options: options.VIsEmpty[int],
		},
		"is empty fail": {
			value:               []int{1, 2, 3},
			options:             options.VIsEmpty[int],
			expectedValidateErr: errs.IsEmptyError,
		},
		"is not empty success": {
			value:   []int{1, 2, 3},
			options: options.VIsNotEmpty[int],
		},
		"is not empty fail": {
			value:               []int{},
			options:             options.VIsNotEmpty[int],
			expectedValidateErr: errs.IsNotEmptyErr,
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
			valueWrapper := NewValueWrapper[[]int]()
			option := valueWrapper.WithOptions(tc.options).ToOption(tc.value)
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
			valueWrapper := NewValueWrapper[string]().WithOptions(tc.options)
			assert.Equal(t, tc.expectedValidateErr, valueWrapper.Validate(tc.value))
		})
	}
}

func TestValueWrapper_NilBehaviour(t *testing.T) {
	var valueWrapper *ValueValidator[int]
	assert.Equal(t, (*ValueValidator[int])(nil), valueWrapper)
	assert.Nil(t, valueWrapper.Validate(1))
	assert.Nil(t, valueWrapper.ToOption(1)())
	assert.Nil(t, valueWrapper.WithOptions(
		options.VIsDefault[int](),
		options.VIsNotDefault[int](),
	).Validate(1))
	assert.Nil(t, valueWrapper.WithOptions(nil).ToOption(1)())
}

func TestValueWrapper_DoubleWrap(t *testing.T) {
	valueWrapper := NewValueWrapper[int]()
	valueWrapper = valueWrapper.WithOptions(options.VIsDefault[int]())
	valueWrapper = valueWrapper.WithOptions(options.VIsNotDefault[int]()) // Should not be used
	assert.Equal(t, errs.IsDefaultErr, valueWrapper.Validate(1))
}
