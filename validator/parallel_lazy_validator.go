package validator

import "github.com/Jh123x/go-validate/ttypes"

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
	// Channel for parallel validation.
	errChan := make(chan error, len(l.options))

	// Run all validations in parallel.
	for _, opt := range l.options {
		go func(opt ttypes.Validate) {
			errChan <- opt()
		}(opt)
	}

	// Collect all errors.
	for i := 0; i < len(l.options); i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}
	return nil
}
