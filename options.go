package govalidate

// WithRequire returns a new Validate that will be evaluated.
func WithRequire(t Test, err error) Validate {
	return func() (bool, error) { return t(), err }
}

// IsNotEmpty validates that the provided value is not empty/default value.
func IsNotEmpty[T comparable](val T) Validate {
	var defaultVal T
	return WithRequire(func() bool { return val == defaultVal }, IsNotEmptyErr)
}

// IsLength validates the the provided value is between, inclusive, the start and end values.
func IsLength[T any](arr []T, start, end int) Validate {
	return WithRequire(func() bool { return len(arr) >= start && len(arr) <= end }, InvalidLengthError)
}
