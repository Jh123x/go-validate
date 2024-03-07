package validator

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

// LazyValidator is a validator that lazily evaluates the options provided.
type LazyValidator struct {
	option ttypes.Validate
}

var _ ttypes.Validator[LazyValidator] = (*LazyValidator)(nil)

// NewLazyValidator returns a new LazyValidator.
func NewLazyValidator() *LazyValidator {
	return &LazyValidator{}
}

// WithOptions returns a new LazyValidator with the given options.
func (l *LazyValidator) WithOptions(opts ...ttypes.Validate) *LazyValidator {
	if l == nil {
		return nil
	}
	combineOptions := options.And(opts...)
	l.option = options.And(combineOptions, l.option)
	return l
}

// Validate validates the options provided.
func (l *LazyValidator) Validate() error {
	if l == nil || l.option == nil {
		return nil
	}
	return l.option()
}
