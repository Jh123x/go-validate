package options

import (
	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/ttypes"
)

// VWithRequire returns a new ValTest that will be evaluated.
func VWithRequire[T any](t ttypes.VTest[T], err error) ttypes.ValTest[T] {
	return func(val T) error {
		if !t(val) {
			return err
		}
		return nil
	}
}

// VIsNotDefault validates that the provided value is not the empty/default value.
func VIsNotDefault[T comparable]() ttypes.ValTest[T] {
	var defaultVal T
	return func(t T) error {
		if defaultVal == t {
			return errs.IsNotDefaultErr
		}
		return nil
	}
}

// VIsDefault validates that the provided value is equals to the empty/default value.
func VIsDefault[T comparable]() ttypes.ValTest[T] {
	var defaultVal T
	return func(val T) error {
		if defaultVal != val {
			return errs.IsDefaultErr
		}
		return nil
	}
}

// VIsEmpty validates that the provided value is empty.
func VIsEmpty[T any](val []T) error {
	if len(val) == 0 {
		return nil
	}
	return errs.IsEmptyError
}

// VIsNotEmpty validates that the provided value is not empty.
func VIsNotEmpty[T any](val []T) error {
	if len(val) != 0 {
		return nil
	}
	return errs.IsNotEmptyErr
}

// VIsLength validates the the provided value is between, inclusive, the start and end values.
func VIsLength[T any](minLen, maxLen int) ttypes.ValTest[[]T] {
	return func(val []T) error {
		if len(val) >= minLen && len(val) <= maxLen {
			return nil
		}
		return errs.InvalidLengthError
	}
}

// VContains validates that the provided array contains the provided element.
func VContains[T comparable](elem T) ttypes.ValTest[[]T] {
	return func(arr []T) error {
		for _, v := range arr {
			if v == elem {
				return nil
			}
		}
		return errs.ContainsError
	}
}

// VOr validates that at least one of the provided options is valid.
func VOr[T any](options ...ttypes.ValTest[T]) ttypes.ValTest[T] {
	return func(val T) error {
		for _, option := range options {
			if option == nil {
				continue
			}
			if err := option(val); err == nil {
				return nil
			}
		}
		return errs.OrError
	}
}

// VAnd validates that all of the provided options are valid.
func VAnd[T any](options ...ttypes.ValTest[T]) ttypes.ValTest[T] {
	return func(val T) error {
		for _, option := range options {
			if option == nil {
				continue
			}
			if err := option(val); err != nil {
				return err
			}
		}
		return nil
	}
}
