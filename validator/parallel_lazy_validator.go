package validator

import (
	"github.com/Jh123x/go-validate/ttypes"
	lop "github.com/gozelle/lo/parallel"
)

var (
	mapperFn = func(opt ttypes.Validate, _ int) error { return opt() }
)

type ParallelLazyValidator struct {
	options []ttypes.Validate
}

var _ ttypes.Validator[ParallelLazyValidator] = (*ParallelLazyValidator)(nil)

// NewParallelLazyValidator returns a new ParallelLazyValidator.
func NewParallelLazyValidator() *ParallelLazyValidator {
	return &ParallelLazyValidator{}
}

// WithOptions returns a new ParallelLazyValidator with the given options.
func (l *ParallelLazyValidator) WithOptions(opts ...ttypes.Validate) *ParallelLazyValidator {
	if l == nil {
		return nil
	}
	newValidator := *l
	newValidator.options = append(l.options, opts...)
	return &newValidator
}

// Validate validates the options provided.
func (l *ParallelLazyValidator) Validate() error {
	if l == nil {
		return nil
	}

	for _, err := range lop.Map(l.options, mapperFn) {
		if err != nil {
			return err
		}
	}

	return nil
}
