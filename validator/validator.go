package validator

import "github.com/Jh123x/go-validate/ttypes"

type Validator struct {
	currErr error
}

var _ ttypes.Validator[Validator] = (*Validator)(nil)

// NewLazyValidator returns a new LazyValidator.
func NewValidator() *Validator {
	return &Validator{currErr: nil}
}

// WithOptions returns a new LazyValidator with the given options.
func (l *Validator) WithOptions(opts ...ttypes.Validate) *Validator {
	if l == nil {
		return nil
	}
	for _, opt := range opts {
		if l.currErr != nil {
			return &Validator{currErr: l.currErr}
		}
		if err := opt(); err != nil {
			return &Validator{currErr: err}
		}
	}
	return NewValidator()
}

// Validate validates the options provided.
func (l *Validator) Validate() error {
	if l == nil {
		return nil
	}
	return l.currErr
}
