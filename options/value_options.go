package options

import (
	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/ttypes"
)

func VWithRequire[T any](t ttypes.VTest[T], err error) ttypes.ValTest[T] {
	return func(val T) error {
		if !t(val) {
			return err
		}
		return nil
	}
}

func VIsNotEmpty[T comparable]() ttypes.VTest[T] {
	var defaultVal T
	return func(t T) bool {
		return defaultVal != t
	}
}

func VIsEmpty[T comparable]() ttypes.VTest[T] {
	var defaultVal T
	return func(val T) bool {
		return defaultVal == val
	}
}

func VIsLength[T any](minLen, maxLen int) ttypes.ValTest[[]T] {
	return func(val []T) error {
		if len(val) >= minLen && len(val) <= maxLen {
			return nil
		}
		return errs.InvalidLengthError
	}
}

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

func VOr[T comparable](options ...ttypes.ValTest[T]) ttypes.ValTest[T] {
	return func(val T) error {
		for _, option := range options {
			if err := option(val); err == nil {
				return nil
			}
		}
		return errs.OrError
	}
}

func VAnd[T comparable](options ...ttypes.ValTest[T]) ttypes.ValTest[T] {
	return func(val T) error {
		for _, option := range options {
			if err := option(val); err != nil {
				return err
			}
		}
		return nil
	}
}
