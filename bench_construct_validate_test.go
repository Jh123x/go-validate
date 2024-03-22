package main

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/Jh123x/go-validate/validator"
	"github.com/Jh123x/go-validate/wrapper"
	"github.com/invopop/validation"
)

// validateResponseLazy is a benchmark for the Lazy Evaluator.
func validateResponseLazy(resp *Response) error {
	return validator.NewLazyValidator().WithOptions(
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

// validateResponseParallelLazy is a benchmark for the Parallel Lazy Evaluator.
func validateResponseParallelLazy(resp *Response) error {
	return validator.NewParallelLazyValidator().WithOptions(
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

// validateResponseValidator is a benchmark for the normal Validator.
func validateResponseValidator(resp *Response) error {
	return validator.NewValidator().WithOptions(
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

func validateResponseValueWrapperLong(resp *Response) error {
	return wrapper.NewValueWrapper[*Response]().WithOptions(
		options.VWithRequire(func(r *Response) bool { return r.Code != 0 }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool { return len(r.Message) > 0 }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool { return r.Extras != nil }, errs.IsEmptyError),
		options.VWithRequire(func(r *Response) bool {
			return len(r.Optional) > 0 && len(r.SetIfOptSet) > 0 || len(r.Optional) == 0 && len(r.SetIfOptSet) == 0
		}, errs.IsEmptyError),
	).Validate(resp)
}

func validateResponseValueWrapperShort(resp *Response) error {
	intValidator := wrapper.NewValueWrapper[int]().WithOptions(options.VIsNotDefault[int]())
	return wrapper.NewValueWrapper[*Response]().WithOptions(
		func(r *Response) error { return intValidator.Validate(r.Code) },
		func(r *Response) error { return intValidator.Validate(len(r.Message)) },
		options.VWithRequire(func(r *Response) bool { return r.Extras != nil }, errs.IsNotEmptyErr),
		func(r *Response) error {
			if len(r.Optional) > 0 && len(r.SetIfOptSet) > 0 || len(r.Optional) == 0 && len(r.SetIfOptSet) == 0 {
				return nil
			}
			return errs.IsEmptyError
		},
	).Validate(resp)
}

// BenchmarkConstructAndValidateData benchmarks the different validators.
func BenchmarkConstructAndValidateData(b *testing.B) {
	algorithms := map[string]ttypes.ValTest[*Response]{
		"TestLazyValidator":     validateResponseLazy,
		"TestInvopop":           validateResponseInvopop,
		"TestParallelLazy":      validateResponseParallelLazy,
		"TestValidator":         validateResponseValidator,
		"TestIfStmts":           validateIfImplementation,
		"TestValueWrapperLong":  validateResponseValueWrapperLong,
		"TestValueWrapperShort": validateResponseValueWrapperShort,
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
