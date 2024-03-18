package wrapper

import "github.com/Jh123x/go-validate/ttypes"

// ValueValidator is a wrapper for a value of type T.
// You can use repeated Tests on the wrapper to check for the same boolean.
type ValueValidator[T any] struct {
	Value T
	err   error
}

func NewValueWrapper[T any](value T) *ValueValidator[T] {
	return &ValueValidator[T]{Value: value}
}

func (v *ValueValidator[T]) WithOptions(options ...ttypes.ValTest[T]) *ValueValidator[T] {
	if v == nil {
		return nil
	}
	if v.err != nil {
		return v
	}
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(v.Value); err != nil {
			v.err = err
			break
		}
	}
	return v
}

func (v *ValueValidator[T]) Validate() error {
	if v == nil {
		return nil
	}
	return v.err
}

func (v *ValueValidator[T]) ToOption() ttypes.Validate {
	if v == nil {
		return func() error { return nil }
	}
	return func() error { return v.err }
}
