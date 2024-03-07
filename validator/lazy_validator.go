package validator

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

// LazyValidator is a validator that lazily evaluates the options provided.
type LazyValidator struct {
	options ttypes.Validate
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
	newValidator := *l
	newValidator.options = options.And(l.options, options.And(opts...))
	return &newValidator
}

// Validate validates the options provided.
func (l *LazyValidator) Validate() error {
	if l == nil || l.options == nil {
		return nil
	}
	return l.options()
}
