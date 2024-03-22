package ttypes

// Test is a function that returns true if the test succeeds.
type Test func() bool

// Require is a validation that will be evaluated.
type Require func(Test, error) Validate

// Validator is a type that can be validated.
type Validator[T any] interface {
	WithOptions(...Validate) *T
	Validate() error
}

// VTest
type VTest[T any] func(T) bool

// ValTest is a test for Type T.
// Returns true if the test succeeds.
type ValTest[T any] func(T) error

// ValValidator is a type that can be validated
type ValValidator[T any] interface {
	WithOptions(...ValTest[T]) *ValValidator[T]
	Validate() error
	ToOption() Validate
}
