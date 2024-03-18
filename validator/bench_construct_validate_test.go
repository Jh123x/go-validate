package validator

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/invopop/validation"
	"github.com/stretchr/testify/assert"
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

// validateResponseLazy is a benchmark for the Lazy Evaluator.
func validateResponseLazy(resp *Response) error {
	return NewLazyValidator().WithOptions(
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
	return NewParallelLazyValidator().WithOptions(
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
	return NewValidator().WithOptions(
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

// BenchmarkConstructAndValidateData benchmarks the different validators.
func BenchmarkConstructAndValidateData(b *testing.B) {
	algorithms := map[string]ttypes.ValTest[*Response]{
		"TestLazyValidator": validateResponseLazy,
		"TestInvopop":       validateResponseInvopop,
		"TestParallelLazy":  validateResponseParallelLazy,
		"TestValidator":     validateResponseValidator,
		"TestIfStmts":       validateIfImplementation,
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
