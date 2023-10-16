# Go Validate

A mini validation library in Pure Go.

Please create an issue if you have any suggestions or find any bugs / problems.

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

## Benchmark

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
