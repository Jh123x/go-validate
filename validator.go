package govalidate

type LazyValidator struct {
	options []Validate
}

// NewLazyValidator returns a new LazyValidator.
func NewLazyValidator() LazyValidator {
	return LazyValidator{}
}

// WithOptions returns a new LazyValidator with the given options.
func (l *LazyValidator) WithOptions(opts ...Validate) LazyValidator {
	l.options = append(l.options, opts...)
	return *l
}

// Validate validates the options provided.
func (l *LazyValidator) Validate() (bool, error) {
	for _, opt := range l.options {
		if ok, err := opt(); !ok {
			return false, err
		}
	}
	return true, nil
}
