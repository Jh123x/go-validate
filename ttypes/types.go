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
