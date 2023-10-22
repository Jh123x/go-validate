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

## Installation

```sh
go get github.com/Jh123x/go-validate
```

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

Validation logic for this librarys' validators is as follows.

In this example we are using lazy validator but the logic is the same for all validators.

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
BenchmarkData/with_no_error_in_optional_for_TestValidator-12         	  958305	      1139 ns/op	    1056 B/op	      33 allocs/op
BenchmarkData/with_no_error_in_optional_for_TestLazyValidator-12     	  998826	      1072 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/with_no_error_in_optional_for_TestInvopop-12           	  726202	      1658 ns/op	    1913 B/op	      41 allocs/op
BenchmarkData/with_no_error_in_optional_for_TestParallelLazy-12      	  348286	      3317 ns/op	    1537 B/op	      43 allocs/op
BenchmarkData/with_err_in_setIfOptSet_for_TestParallelLazy-12        	  347270	      3483 ns/op	    1714 B/op	      47 allocs/op
BenchmarkData/with_err_in_setIfOptSet_for_TestValidator-12           	  958428	      1290 ns/op	    1232 B/op	      37 allocs/op
BenchmarkData/with_err_in_setIfOptSet_for_TestLazyValidator-12       	  939782	      1326 ns/op	    1248 B/op	      37 allocs/op
BenchmarkData/with_err_in_setIfOptSet_for_TestInvopop-12             	  574546	      2066 ns/op	    2353 B/op	      48 allocs/op
BenchmarkData/with_error_in_extras_for_TestValidator-12              	 1000000	      1090 ns/op	    1056 B/op	      33 allocs/op
BenchmarkData/with_error_in_extras_for_TestLazyValidator-12          	 1000000	      1107 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/with_error_in_extras_for_TestInvopop-12                	  570960	      2037 ns/op	    2329 B/op	      48 allocs/op
BenchmarkData/with_error_in_extras_for_TestParallelLazy-12           	  361460	      3362 ns/op	    1537 B/op	      43 allocs/op
BenchmarkData/no_errors_for_TestInvopop-12                           	  742675	      1671 ns/op	    1913 B/op	      41 allocs/op
BenchmarkData/no_errors_for_TestParallelLazy-12                      	  363316	      3303 ns/op	    1537 B/op	      43 allocs/op
BenchmarkData/no_errors_for_TestValidator-12                         	 1000000	      1070 ns/op	    1056 B/op	      33 allocs/op
BenchmarkData/no_errors_for_TestLazyValidator-12                     	 1000000	      1081 ns/op	    1072 B/op	      33 allocs/op
BenchmarkData/with_error_in_code_for_TestLazyValidator-12            	  959569	      1241 ns/op	    1216 B/op	      36 allocs/op
BenchmarkData/with_error_in_code_for_TestInvopop-12                  	  499682	      2063 ns/op	    2345 B/op	      48 allocs/op
BenchmarkData/with_error_in_code_for_TestParallelLazy-12             	  342500	      3410 ns/op	    1682 B/op	      46 allocs/op
BenchmarkData/with_error_in_code_for_TestValidator-12                	  959040	      1220 ns/op	    1200 B/op	      36 allocs/op
BenchmarkData/with_error_in_message_for_TestLazyValidator-12         	  959401	      1244 ns/op	    1216 B/op	      36 allocs/op
BenchmarkData/with_error_in_message_for_TestInvopop-12               	  571017	      2110 ns/op	    2393 B/op	      49 allocs/op
BenchmarkData/with_error_in_message_for_TestParallelLazy-12          	  337357	      3440 ns/op	    1682 B/op	      46 allocs/op
BenchmarkData/with_error_in_message_for_TestValidator-12             	  959062	      1221 ns/op	    1200 B/op	      36 allocs/op
BenchmarkData/with_error_in_optional_for_TestInvopop-12              	  599444	      2058 ns/op	    2353 B/op	      48 allocs/op
BenchmarkData/with_error_in_optional_for_TestParallelLazy-12         	  324042	      3492 ns/op	    1714 B/op	      47 allocs/op
BenchmarkData/with_error_in_optional_for_TestValidator-12            	  888217	      1317 ns/op	    1232 B/op	      37 allocs/op
BenchmarkData/with_error_in_optional_for_TestLazyValidator-12        	  799237	      1315 ns/op	    1248 B/op	      37 allocs/op
PASS
ok  	github.com/Jh123x/go-validate/validator	33.221s
```

#### Formatted Results (ns/op)

| No  | Test Case            | Input Value                                                                                   | Lazy Validator | Invopop    | Parallel Validator | Validator  | % improvement over Invopop |
| --- | -------------------- | --------------------------------------------------------------------------------------------- | -------------- | ---------- | ------------------ | ---------- | -------------------------- |
| 1   | No Errors            | `Response{Code: 200,Message: "OK",Extras: map[string]any{}}`                                  | 1081 ns/op     | 1671 ns/op | 3303 ns/op         | 1070 ns/op | 54.6%                      |
| 2   | Error in Code        | `Response{Code: 0,Message: "OK",Extras: map[string]any{}}`                                    | 1241 ns/op     | 2063 ns/op | 3410 ns/op         | 1220 ns/op | 66.2%                      |
| 3   | Error in Msg         | `Response{Code: 200,Message: "",Extras: map[string]any{}}`                                    | 1244 ns/op     | 2110 ns/op | 3440 ns/op         | 1221 ns/op | 69.6%                      |
| 4   | Error in Extras      | `Response{Code: 200,Message: "OK",Extras: nil}`                                               | 1107 ns/op     | 2037 ns/op | 3362 ns/op         | 1090 ns/op | 65.0%                      |
| 5   | Error in Opt         | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "test",SetIfOptSet: ""}` | 1315 ns/op     | 2058 ns/op | 3492 ns/op         | 1317 ns/op | 56.5%                      |
| 6   | No Error in Opt      | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: ""}`     | 1072 ns/op     | 1658 ns/op | 3317 ns/op         | 1139 ns/op | 100%                       |
| 7   | Error in setIfOptSet | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: "test"}` | 1326 ns/op     | 2066 ns/op | 3483 ns/op         | 1290 ns/op | 68.5%                      |

For short validations such as this one, it seems that the parallel validator is slower than any of the other validators.

#### Results

In general, the lazy validator is faster than the [Invopop](https://github.com/invopop/validation "validation") validator by a significant margin.
The lazy validator is also faster than the parallel validator in most cases unless there are a large number of validations to be done.

## Future tasks

- [ ] Add more validation options
- [ ] A more comprehensive benchmark
- [ ] Other types of validators for different use cases
