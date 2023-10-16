package validator

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

// TestLazyValidator tests the LazyValidator.
func TestLazyValidator(t *testing.T) {
	tests := map[string]struct {
		options     []ttypes.Validate
		expectedErr error
	}{
		"default case with options should return nil": {
			options:     []ttypes.Validate{},
			expectedErr: nil,
		},
		"with options with no errors should not return an error": {
			options:     []ttypes.Validate{validateWNil},
			expectedErr: nil,
		},
		"with options with errors should return an error": {
			options:     []ttypes.Validate{validateWErr},
			expectedErr: errTest,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validator := NewLazyValidator()
			assert.Equal(t, testCase.expectedErr, validator.WithOptions(testCase.options...).Validate())
		})
	}
}

// TestLazyValidator_cache ensure that LazyValidator can be cached.
func TestLazyValidator_Caching(t *testing.T) {
	validator := NewLazyValidator()
	validator2 := validator.WithOptions(validateWErr)
	assert.Nil(t, validator.Validate())
	assert.Equal(t, errTest, validator2.Validate())
}

// TestLazyValidator_ReadMeExample tests the example in the README.
func TestLazyValidator_ReadMeExample(t *testing.T) {
	validator := NewLazyValidator()
	err := validator.WithOptions(
		options.IsNotEmpty("").WithError(fmt.Errorf("empty string")),             // Fails and returns error.
		options.IsLength([]string{}, 0, 3).WithError(fmt.Errorf("empty string")), // Will not be evaluated.
	).Validate()
	assert.Equal(t, fmt.Errorf("empty string"), err)
}

/* Test Scenario for Benchmarking */
type Response struct {
	Code    int            // Must be non-zero.
	Message string         // Must be non-empty.
	Extras  map[string]any // Must be non-nil.
}

type testFn func(*Response) error

func benchmarkValidator(b *testing.B, response *Response, validateFn testFn, hasErr bool) {
	for i := 0; i < b.N; i++ {
		err := validateFn(response)
		assert.Equal(b, hasErr, err != nil, err)
	}
}

func validateResponseLazy(resp *Response) error {
	return NewLazyValidator().WithOptions(
		options.IsNotEmpty(resp.Code),
		options.IsNotEmpty(resp.Message),
		options.WithRequire(func() bool { return resp.Extras != nil }, errTest),
	).Validate()
}

func validateResponseOzzo(resp *Response) error {
	return validation.ValidateStruct(
		resp,
		validation.Field(&resp.Code, validation.NilOrNotEmpty),
		validation.Field(&resp.Message, validation.NilOrNotEmpty),
		validation.Field(&resp.Extras, validation.NotNil),
	)
}

func BenchmarkData(b *testing.B) {
	algorithms := map[string]testFn{
		"TestLazyValidator": validateResponseLazy,
		"TestOzzo":          validateResponseOzzo,
	}
	tests := map[string]struct {
		resp   Response
		hasErr bool
	}{
		"no errors": {
			resp: Response{
				Code:    200,
				Message: "OK",
				Extras:  map[string]any{},
			},
			hasErr: false,
		},
		"with error in code": {
			resp: Response{
				Code:    0,
				Message: "OK",
				Extras:  map[string]any{},
			},
			hasErr: true,
		},
		"with error in message": {
			resp: Response{
				Code:    200,
				Message: "",
				Extras:  map[string]any{},
			},
			hasErr: true,
		},
		"with error in extras": {
			resp: Response{
				Code:    200,
				Message: "OK",
				Extras:  nil,
			},
			hasErr: true,
		},
	}

	for testName, testCase := range tests {
		for name, algo := range algorithms {
			b.Run(fmt.Sprintf("%s for %s", testName, name), func(b *testing.B) {
				benchmarkValidator(b, &testCase.resp, algo, testCase.hasErr)
			})
		}
	}
}
