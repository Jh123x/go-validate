package ttypes

// Validate returns true, nil if the validation passes, false, error otherwise.
// By default, the error returned is a Validate
type Validate func() error

// WithError changes the error returned by the validation.
func (v Validate) WithError(err error) Validate {
	return func() error {
		if oldErr := v(); oldErr != nil {
			return err
		}
		return nil
	}
}

// Not changes the error returned by the validation to the provided error.
func (v Validate) Not(err error) Validate {
	return func() error {
		if oldErr := v(); oldErr != nil {
			return nil
		}
		return err
	}
}
