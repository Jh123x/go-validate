# Go Validate

A mini validation library in Pure Go.

Please create an issue if you have any:

1. Questions
2. Suggestions
3. Find any bugs
4. Want to point out any mistakes

## Features

1. Customized error messages for your validation rules
2. Easy to write custom rules
3. Define your own validation method
4. Fast

## Example

```go
package main

import (
    "github.com/Jh123x/go-validate/options"
    "github.com/Jh123x/go-validate/validator"
)

func main(){
    lazyValidator := validator.NewLazyValidator()
	err := lazyValidator.WithOptions(
		options.IsNotEmpty("").WithError(fmt.Errorf("empty string")),             // Fails and returns error.
		options.IsLength([]string{}, 0, 3).WithError(fmt.Errorf("empty string")), // Will not be evaluated.
	).Validate()
    if err != nil {
        // handle error
        ...
    }
}
```

## Benchmark (Needs improvement)

### Methodology

We will be using a fake response struct to test the benchmark the performance of the library.
This can be further extended to test the performance of the library on a real world scenarios.

```go
type Response struct {
	Code    int            // Must be non-zero.
	Message string         // Must be non-empty.
	Extras  map[string]any // Must be non-nil.
}
```

### Validation logic

Validation logic for this library is as follows.

```go
func validateResponseLazy(resp *Response) error {
	return NewLazyValidator().WithOptions(
		options.IsNotEmpty(resp.Code),
		options.IsNotEmpty(resp.Message),
		options.WithRequire(func() bool { return resp.Extras != nil }, errTest),
	).Validate()
}

```

Validation logic for the Ozzo library is as follows.

```go
func validateResponseOzzo(resp *Response) error {
	return validation.ValidateStruct(
		resp,
		validation.Field(&resp.Code, validation.NilOrNotEmpty),
		validation.Field(&resp.Message, validation.NilOrNotEmpty),
		validation.Field(&resp.Extras, validation.NotNil),
	)
}
```

### Results

This is the test results on my PC.

```
goos: windows
goarch: amd64
pkg: github.com/Jh123x/go-validate/validator
cpu: AMD Ryzen 5 7600 6-Core Processor

BenchmarkData/with_error_in_message_for_TestLazyValidator-12         	 2721825	       431.0 ns/op	     280 B/op	       9 allocs/op
BenchmarkData/with_error_in_message_for_TestOzzo-12                  	 1291412	       928.2 ns/op	    1144 B/op	      20 allocs/op
BenchmarkData/with_error_in_extras_for_TestLazyValidator-12          	 2808309	       430.3 ns/op	     280 B/op	       9 allocs/op
BenchmarkData/with_error_in_extras_for_TestOzzo-12                   	 1307173	       913.8 ns/op	    1144 B/op	      20 allocs/op
BenchmarkData/no_errors_for_TestOzzo-12                              	 1448179	       817.0 ns/op	     840 B/op	      18 allocs/op
BenchmarkData/no_errors_for_TestLazyValidator-12                     	 2804840	       430.8 ns/op	     280 B/op	       9 allocs/op
BenchmarkData/with_error_in_code_for_TestLazyValidator-12            	 2858064	       418.4 ns/op	     280 B/op	       9 allocs/op
BenchmarkData/with_error_in_code_for_TestOzzo-12                     	 1288452	       924.6 ns/op	    1144 B/op	      20 allocs/op
PASS

ok  	github.com/Jh123x/go-validate/validator	15.144s
```

## Future tasks

- [ ] Add more validation options
- [ ] A more comprehensive benchmark
- [ ] Other types of validators for different use cases
