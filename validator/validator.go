package validator

import "github.com/Jh123x/go-validate/ttypes"

type LazyValidator struct {
	options []ttypes.Validate
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
	newValidator.options = append(l.options, opts...)
	return &newValidator
}

// Validate validates the options provided.
func (l *LazyValidator) Validate() error {
	if l == nil {
		return nil
	}
	for _, opt := range l.options {
		if err := opt(); err != nil {
			return err
		}
	}
	return nil
}
