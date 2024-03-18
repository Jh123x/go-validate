package validator

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/Jh123x/go-validate/validator"
	"github.com/Jh123x/go-validate/wrapper"
	"github.com/invopop/validation"
	"github.com/stretchr/testify/assert"
)

var (
	errTest = fmt.Errorf("test error")
)

/* Test Scenario for Benchmarking */
type Response struct {
	Code        int            // Must be non-zero.
	Message     string         // Must be non-empty.
	Extras      map[string]any // Must be non-nil.
	Optional    string         // Optional.
	SetIfOptSet string         // Set if Optional is set, empty otherwise.
}

// benchmarkValidator benchmarks the validateFn.
func benchmarkValidator(b *testing.B, response *Response, validateFn ttypes.ValTest[*Response], hasErr bool) {
	for i := 0; i < b.N; i++ {
		err := validateFn(response)
		assert.Equal(b, hasErr, err != nil, fmt.Sprintf("expected error: %v, got: %v", hasErr, err))
	}
}

// validateOnlyResponseLazy is a benchmark for the Lazy Evaluator.
func validateOnlyResponseLazy() ttypes.ValTest[*Response] {
	validator := validator.NewLazyValidator()
	return func(resp *Response) error {
		return validator.WithOptions(
			options.IsNotEmpty(resp.Code),
			options.IsNotEmpty(resp.Message),
			options.WithRequire(func() bool { return resp.Extras != nil }, errTest),
			options.Or(
				options.And(
					options.IsEmpty(resp.Optional),
					options.IsEmpty(resp.SetIfOptSet),
				),
				options.And(
					options.IsNotEmpty(resp.Optional),
					options.IsNotEmpty(resp.SetIfOptSet),
				),
			),
		).Validate()
	}
}

// validateOnlyResponseParallelLazy is a benchmark for the Parallel Lazy Evaluator.
func validateOnlyResponseParallelLazy() ttypes.ValTest[*Response] {
	parallelValidator := validator.NewParallelLazyValidator()
	return func(resp *Response) error {
		return parallelValidator.WithOptions(
			options.IsNotEmpty(resp.Code),
			options.IsNotEmpty(resp.Message),
			options.WithRequire(func() bool { return resp.Extras != nil }, errTest),
			options.Or(
				options.And(
					options.IsEmpty(resp.Optional),
					options.IsEmpty(resp.SetIfOptSet),
				),
				options.And(
					options.IsNotEmpty(resp.Optional),
					options.IsNotEmpty(resp.SetIfOptSet),
				),
			),
		).Validate()
	}
}

// validateResponseValidator is a benchmark for the normal Validator.
func validateOnlyResponseValidator() ttypes.ValTest[*Response] {
	validator := validator.NewValidator()
	return func(resp *Response) error {
		return validator.WithOptions(
			options.IsNotEmpty(resp.Code),
			options.IsNotEmpty(resp.Message),
			options.WithRequire(func() bool { return resp.Extras != nil }, errTest),
			options.Or(
				options.And(
					options.IsEmpty(resp.Optional),
					options.IsEmpty(resp.SetIfOptSet),
				),
				options.And(
					options.IsNotEmpty(resp.Optional),
					options.IsNotEmpty(resp.SetIfOptSet),
				),
			),
		).Validate()
	}
}

// validateIfImplementation is a benchmark for just pure if statements.
func validateIfImplementation(resp *Response) error {
	if resp.Code == 0 {
		return errs.IsEmptyError
	}

	if len(resp.Message) == 0 {
		return errs.IsEmptyError
	}

	if resp.Extras == nil {
		return errs.IsEmptyError
	}

	if len(resp.Optional) > 0 && len(resp.SetIfOptSet) == 0 {
		return errs.IsEmptyError
	}

	if len(resp.Optional) == 0 && len(resp.SetIfOptSet) > 0 {
		return errs.IsEmptyError
	}
	return nil
}

// validateResponseInvopop is a benchmark for the Invopop Validation Library.
func validateResponseInvopop(resp *Response) error {
	return validation.ValidateStruct(
		resp,
		validation.Field(&resp.Code, validation.Required),
		validation.Field(&resp.Message, validation.Required),
		validation.Field(&resp.Extras, validation.NotNil),
		validation.Field(
			&resp.SetIfOptSet,
			validation.Required.When(
				len(resp.Optional) > 0,
			),
			validation.Empty.When(
				len(resp.Optional) == 0,
			),
		),
	)
}

func validateResponseValueWrapperLong() ttypes.ValTest[*Response] {
	validator := wrapper.NewValueWrapper[*Response]().WithOptions(
		options.VWithRequire(func(r *Response) bool { return r.Code != 0 }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool { return len(r.Message) > 0 }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool { return r.Extras != nil }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool {
			return len(r.Optional) > 0 && len(r.SetIfOptSet) > 0 || len(r.Optional) == 0 && len(r.SetIfOptSet) == 0
		}, errs.IsEmptyError),
	)
	return func(resp *Response) error { return validator.Validate(resp) }
}

func validateResponseValueWrapperShort() ttypes.ValTest[*Response] {
	intValidator := wrapper.NewValueWrapper[int]().WithOptions(options.VIsNotDefault[int]())
	validator := wrapper.NewValueWrapper[*Response]().WithOptions(
		func(r *Response) error { return intValidator.Validate(r.Code) },
		func(r *Response) error { return intValidator.Validate(len(r.Message)) },
		options.VWithRequire(func(r *Response) bool { return r.Extras != nil }, errs.IsNotEmptyErr),
		func(r *Response) error {
			if len(r.Optional) > 0 && len(r.SetIfOptSet) > 0 || len(r.Optional) == 0 && len(r.SetIfOptSet) == 0 {
				return nil
			}
			return errs.IsEmptyError
		},
	)
	return func(resp *Response) error { return validator.Validate(resp) }
}

// BenchmarkOnlyValidate Data benchmarks the different validators only for their validation cost.
func BenchmarkOnlyValidateData(b *testing.B) {
	algorithms := map[string]ttypes.ValTest[*Response]{
		"TestLazyValidator":     validateOnlyResponseLazy(),
		"TestInvopop":           validateResponseInvopop, // No Option to create a validator first.
		"TestParallelLazy":      validateOnlyResponseParallelLazy(),
		"TestValidator":         validateOnlyResponseValidator(),
		"TestIfStmts":           validateIfImplementation,
		"TestValueWrapperLong":  validateResponseValueWrapperLong(),
		"TestValueWrapperShort": validateResponseValueWrapperShort(),
	}
	tests := map[string]struct {
		resp   Response
		hasErr bool
	}{
		"no err": {
			resp: Response{
				Code:    200,
				Message: "OK",
				Extras:  map[string]any{},
			},
			hasErr: false,
		},
		"err in code": {
			resp: Response{
				Code:    0,
				Message: "OK",
				Extras:  map[string]any{},
			},
			hasErr: true,
		},
		"err in message": {
			resp: Response{
				Code:    200,
				Message: "",
				Extras:  map[string]any{},
			},
			hasErr: true,
		},
		"err in extras": {
			resp: Response{
				Code:    200,
				Message: "OK",
				Extras:  nil,
			},
			hasErr: true,
		},
		"err in optional": {
			resp: Response{
				Code:        200,
				Message:     "OK",
				Extras:      map[string]any{},
				Optional:    "optional",
				SetIfOptSet: "",
			},
			hasErr: true,
		},
		"no err in optional": {
			resp: Response{
				Code:        200,
				Message:     "OK",
				Extras:      map[string]any{},
				Optional:    "optional",
				SetIfOptSet: "set",
			},
			hasErr: false,
		},
		"err in setIfOptSet": {
			resp: Response{
				Code:        200,
				Message:     "OK",
				Extras:      map[string]any{},
				Optional:    "",
				SetIfOptSet: "set",
			},
			hasErr: true,
		},
	}

	for testName, testCase := range tests {
		for name, algo := range algorithms {
			tcName := fmt.Sprintf("%s for %s", testName, name)
			b.Run(tcName, func(b *testing.B) {
				benchmarkValidator(b, &testCase.resp, algo, testCase.hasErr)
			})
		}
	}
}
