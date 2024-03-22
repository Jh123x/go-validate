package wrapper

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

// ValueValidator is a wrapper for a value of type T.
// You can use repeated Tests on the wrapper to check for the same boolean.
type ValueValidator[T any] struct {
	option ttypes.ValTest[T]
}

func NewValueWrapper[T any]() *ValueValidator[T] {
	return &ValueValidator[T]{
		option: options.VWithRequire[T](func(T) bool { return true }, nil),
	}
}

func (v *ValueValidator[T]) WithOptions(valOptions ...ttypes.ValTest[T]) *ValueValidator[T] {
	if v == (*ValueValidator[T])(nil) {
		return nil
	}
	v.option = options.VAnd(v.option, options.VAnd(valOptions...))
	return v
}

func (v *ValueValidator[T]) Validate(val T) error {
	if v == nil {
		return nil
	}
	return v.option(val)
}

func (v *ValueValidator[T]) ToOption(val T) ttypes.Validate {
	if v == nil {
		return func() error { return nil }
	}
	return func() error { return v.option(val) }
}
