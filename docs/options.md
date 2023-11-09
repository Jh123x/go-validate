# Options

In this section, we will describe all the options in the validator.

## Usages

To use an option, follow the example below.

```go
package main

import (
    "github.com/Jh123x/go-validate/options"
    "github.com/Jh123x/go-validate/validator"
)

func main(){
    validator := validator.NewValidator()
	err := validator.WithOptions(
        // Insert options here when required.
		options.IsNotEmpty(""),
		options.IsLength([]string{}, 0, 3),
        ...
	).Validate()
    if err != nil {
        // handle error
        ...
    }
}
```

## IsNotEmpty

Takes in any value and returns `errs.IsNotEmptyErr` if the value is empty.

### Usage

```go
// Returns error
validator.WithOptions(
    options.IsNotEmpty(""),
).Validate()

// Returns nil
validator.WithOptions(
    options.IsNotEmpty("not empty"), // Returns error
).Validate()
```

## Basic Options

### IsEmpty

The opposite of [`IsNotEmpty`](#isnotempty).
Takes in any value and returns `errs.IsEmptyErr` if the value is not default/empty.

#### Usage

```go
// Returns error
validator.WithOptions(
    options.IsEmpty("not empty"),
).Validate()

// Returns nil
validator.WithOptions(
    options.IsEmpty(""),
).Validate()
```

### IsLength

Takes in an array of any type, the start length and the end length (Both inclusive).
Returns `errs.IsLengthErr` if the length of the value is not within the specified range.

#### Usage

```go
// Returns error
validator.WithOptions(
    options.IsLength([]string{}, 1, 5),
).Validate()

// Returns nil
validator.WithOptions(
    options.IsLength([]string{"123", "test"}, 1, 5),
).Validate()
```

### Contains

Takes in an array of any type and a value. Returns `errs.ContainsErr` if the value is not in the array.

#### Usage

```go
// No error
validator.WithOptions(
    options.Contains([]string{"123", "test"}, "123"),
).Validate()

// returns error
validator.WithOptions(
    options.Contains([]string{"123", "test"}, "bcd"),
).Validate()
```

## String Operations

### IsValidURI

Takes in a string and returns `errs.IsValidURLErr` if the string is not a valid URL.

#### Usage

```go
// No error
validator.WithOptions(
    options.IsValidURI("https://github.com/Jh123x/go-validate"),
).Validate()

// returns error
validator.WithOptions(
    options.IsValidURI("invalid url"),
).Validate()
```

### IsValidJson

Takes in a string and returns `errs.InvalidJsonError` if the string is not a valid JSON.

#### Usage

```go
// No error
validator.WithOptions(
    options.IsValidJson(`{"name":"jh123x"}`),
).Validate()

// returns error
validator.WithOptions(
    options.IsValidJson(`{"name":"jh123x"`),
).Validate()
```

### IsValidEmail

Takes in a string and returns `errs.IsValidEmailErr` if the string is not a valid email.

#### Usage

```go
// No error
validator.WithOptions(
    options.IsValidEmail("test@test.com"),
).Validate()

// returns error
validator.WithOptions(
    options.IsValidEmail("test.com"),
).Validate()
```

## Option Composition

### Or

This is a special option that takes in multiple options and returns nil if any of the options returns nil, otherwise, it returns an `errs.OrError` if all of the options return an error.
You can use this in conjunction with [`And`](#and) to create complex validation rules.

```go
// Returns no error
validator.WithOptions(
    options.Or(
        options.IsEmpty(""), // No error
        options.IsLength([]string{}, 1, 5), // Error
    ),
).Validate()

// Returns error
validator.WithOptions(
    options.Or(
        options.IsEmpty("not empty"), // Error
        options.IsLength([]string{}, 1, 5), // Error
    ),
).Validate()
```

### And

This is a special option that takes in multiple options and returns an error if any of the option errors.
You can use this in conjunction with [`Or`](#or) to create complex validation rules.

```go
// Returns no error
validator.WithOptions(
    options.And(
        options.IsEmpty(""), // No error
        options.IsLength([]string{"test1"}, 1, 5), // No error
    ),
).Validate()

// Returns error (errs.InvalidLengthError)
validator.WithOptions(
    options.And(
        options.IsEmpty(""), // No error
        options.IsLength([]string{}, 1, 5), // Error
    ),
).Validate()
```

## Custom Options

### WithRequire

If the above options are not enough for you, you can use `WithRequire` to define your own validation method.
It takes in a [`Test`](../ttypes/types.go) type (`func() bool`) and a `error` type as parameters. If the function returns `true`, `WithRequire` will return nil, otherwise it will return the error.

#### Usage

```go
// My Custom Function
func IsSpace(value string) types.Validate {
	return WithRequire(func() bool { return value == " " }, fmt.Errorf("value is not space"))
}
```
