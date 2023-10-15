package govalidate

// Validate returns true, nil if the validation passes, false, error otherwise.
// By default, the error returned is a Validate
type Validate func() (bool, error)

// WithError changes the error returned by the validation.
func (v Validate) WithError(err error) Validate {
	return func() (bool, error) {
		ok, _ := v()
		return ok, err
	}
}
