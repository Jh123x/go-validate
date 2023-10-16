package options

import (
	"github.com/Jh123x/go-validate/errs"
	types "github.com/Jh123x/go-validate/ttypes"
)

// WithRequire returns a new Validate that will be evaluated.
func WithRequire(t types.Test, err error) types.Validate {
	return func() error {
		if !t() {
			return err
		}
		return nil
	}
}

// IsNotEmpty validates that the provided value is not the empty/default value.
func IsNotEmpty[T comparable](val T) types.Validate {
	var defaultVal T
	return WithRequire(func() bool { return val != defaultVal }, errs.IsNotEmptyErr)
}

// IsEmpty validates that the provided value is equals to the empty/default value.
func IsEmpty[T comparable](val T) types.Validate {
	return IsNotEmpty(val).Not(errs.IsEmptyError)
}

// IsLength validates the the provided value is between, inclusive, the start and end values.
func IsLength[T any](arr []T, start, end int) types.Validate {
	return WithRequire(func() bool { return len(arr) >= start && len(arr) <= end }, errs.InvalidLengthError)
}
