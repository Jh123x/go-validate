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
/* Test Scenario for Benchmarking */
type Response struct {
	Code        int            // Must be non-zero.
	Message     string         // Must be non-empty.
	Extras      map[string]any // Must be non-nil.
	Optional    string         // Optional
	SetIfOptSet string         // Set if Optional is set
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
		options.Or(
			options.And(
				options.IsEmpty(resp.Optional),
				options.IsEmpty(resp.SetIfOptSet),
			),
			options.And(
				options.IsNotEmpty(resp.Optional),
				options.IsNotEmpty(resp.SetIfOptSet),
			),
		),
	).Validate()
}

```

Validation logic for the [~~Ozzo~~](https://github.com/go-ozzo/ozzo-validation/ "ozzo") [Invopop](https://github.com/invopop/validation "validation") library is as follows.

Invopop is an updated fork of Ozzo and will be used instead.

```go
func validateResponseInvopop(resp *Response) error {
	return validation.ValidateStruct(
		resp,
		validation.Field(&resp.Code, validation.Required),
		validation.Field(&resp.Message, validation.Required),
		validation.Field(&resp.Extras, validation.NotNil),
		validation.Field(
			&resp.SetIfOptSet,
			validation.Required.When(
				len(resp.Optional) > 0,
			),
			validation.Empty.When(
				len(resp.Optional) == 0,
			),
		),
	)
}
```

### Results

This test is conducted on my Desktop PC.
It is not a very powerful PC but it should be enough to test the performance of the library.

#### Terms

| Term    | Meaning                                                                                         |
| ------- | ----------------------------------------------------------------------------------------------- |
| ns      | Nanoseconds                                                                                     |
| op      | Number of times the function is called                                                          |
| ns / op | Nanoseconds per operation. This is the average time taken for each operation. (Lower is better) |
| allocs  | Number of memory allocations. (Lower is better)                                                 |
| B / op  | Bytes per operation. This is the average memory allocated for each operation. (Lower is better) |

#### Raw Benchmark results

```
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkData$ github.com/Jh123x/go-validate/validator

goos: windows
goarch: amd64
pkg: github.com/Jh123x/go-validate/validator
cpu: AMD Ryzen 5 7600 6-Core Processor
BenchmarkData/with_err_in_setIfOptSet_for_TestLazyValidator-12         	  826958	      1300 ns/op	    1248 B/op	      37 allocs/op
BenchmarkData/with_err_in_setIfOptSet_for_TestInvopop-12               	  570406	      2112 ns/op	    2353 B/op	      48 allocs/op
BenchmarkData/with_error_in_extras_for_TestLazyValidator-12            	  994496	      1092 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/with_error_in_extras_for_TestInvopop-12                  	  570580	      2047 ns/op	    2329 B/op	      48 allocs/op
BenchmarkData/no_errors_for_TestLazyValidator-12                       	 1000000	      1075 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/no_errors_for_TestInvopop-12                             	  773055	      1682 ns/op	    1913 B/op	      41 allocs/op
BenchmarkData/with_error_in_code_for_TestLazyValidator-12              	  932284	      1231 ns/op	    1216 B/op	      36 allocs/op
BenchmarkData/with_error_in_code_for_TestInvopop-12                    	  557136	      2064 ns/op	    2345 B/op	      48 allocs/op
BenchmarkData/with_error_in_message_for_TestLazyValidator-12           	  955764	      1229 ns/op	    1216 B/op	      36 allocs/op
BenchmarkData/with_error_in_message_for_TestInvopop-12                 	  567538	      2127 ns/op	    2393 B/op	      49 allocs/op
BenchmarkData/with_error_in_optional_for_TestInvopop-12                	  596101	      2049 ns/op	    2353 B/op	      48 allocs/op
BenchmarkData/with_error_in_optional_for_TestLazyValidator-12          	  957708	      1284 ns/op	    1248 B/op	      37 allocs/op
BenchmarkData/with_no_error_in_optional_for_TestLazyValidator-12       	 1000000	      1074 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/with_no_error_in_optional_for_TestInvopop-12             	  749296	      1666 ns/op	    1913 B/op	      41 allocs/op
PASS
ok  	github.com/Jh123x/go-validate/validator	16.762s
```

#### Formatted Results (ns/op)

| No  | Test Case            | Input Value                                                                                   | Lazy Validator | Invopop    |
| --- | -------------------- | --------------------------------------------------------------------------------------------- | -------------- | ---------- |
| 1   | No Errors            | `Response{Code: 200,Message: "OK",Extras: map[string]any{}}`                                  | 1075 ns/op     | 1682 ns/op |
| 2   | Error in Code        | `Response{Code: 0,Message: "OK",Extras: map[string]any{}}`                                    | 1231 ns/op     | 2064 ns/op |
| 3   | Error in Msg         | `Response{Code: 200,Message: "",Extras: map[string]any{}}`                                    | 1229 ns/op     | 2127 ns/op |
| 4   | Error in Extras      | `Response{Code: 200,Message: "OK",Extras: nil}`                                               | 1092 ns/op     | 2047 ns/op |
| 5   | Error in Opt         | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "test",SetIfOptSet: ""}` | 1284 ns/op     | 2049 ns/op |
| 6   | No Error in Opt      | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: ""}`     | 1074 ns/op     | 1666 ns/op |
| 7   | Error in setIfOptSet | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: "test"}` | 1300 ns/op     | 2112 ns/op |

## Future tasks

- [ ] Add more validation options
- [ ] A more comprehensive benchmark
- [ ] Other types of validators for different use cases
