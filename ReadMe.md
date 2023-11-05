# Go Validate

![CI Badge](https://github.com/Jh123x/go-validate/actions/workflows/go.yml/badge.svg "CI Badge")

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

Validation logic for this library's validators is as follows.

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

The ideal results will be model as a series of if statements as shown in the code below.

```go
// validateIfImplementation is a benchmark for just pure if statements.
func validateIfImplementation(resp *Response) error {
	if resp.Code == 0 {
		return errs.IsEmptyError
	}

	if len(resp.Message) == 0 {
		return errs.IsEmptyError
	}

	if resp.Extras == nil {
		return errs.IsEmptyError
	}

	if len(resp.Optional) > 0 && len(resp.SetIfOptSet) == 0 {
		return errs.IsEmptyError
	}

	if len(resp.Optional) == 0 && len(resp.SetIfOptSet) > 0 {
		return errs.IsEmptyError
	}
	return nil
}
```

### Results

This test is conducted on my Desktop PC.
It is not a very powerful PC but it should be enough to test the performance of the library.

This is running on go version 1.21.3 windows/amd64. The results may vary on different versions of go and different operating systems.

#### Terms

| Term      | Meaning                                                                                         |
| --------- | ----------------------------------------------------------------------------------------------- |
| `ns`      | Nanoseconds                                                                                     |
| `op`      | Number of times the function is called                                                          |
| `ns/op` | Nanoseconds per operation. This is the average time taken for each operation. (Lower is better) |
| `allocs`  | Number of memory allocations. (Lower is better)                                                 |
| `B/op`  | Bytes per operation. This is the average memory allocated for each operation. (Lower is better) |

#### Raw Benchmark results

```bash
Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkData$ github.com/Jh123x/go-validate/validator

goos: windows
goarch: amd64
pkg: github.com/Jh123x/go-validate/validator
cpu: AMD Ryzen 5 7600 6-Core Processor
BenchmarkData/err_in_setIfOptSet_for_TestLazyValidator-12                 856053              1315 ns/op
BenchmarkData/err_in_setIfOptSet_for_TestInvopop-12                       558294              2068 ns/op
BenchmarkData/err_in_setIfOptSet_for_TestParallelLazy-12                  318944              3383 ns/op
BenchmarkData/err_in_setIfOptSet_for_TestValidator-12                    1566036               770.0 ns/op
BenchmarkData/err_in_setIfOptSet_for_TestIfStmts-12                      2353692               516.4 ns/op
BenchmarkData/err_in_extras_for_TestLazyValidator-12                      958228              1110 ns/op
BenchmarkData/err_in_extras_for_TestInvopop-12                            552584              2109 ns/op
BenchmarkData/err_in_extras_for_TestParallelLazy-12                       342315              3336 ns/op
BenchmarkData/err_in_extras_for_TestValidator-12                         2019872               611.2 ns/op
BenchmarkData/err_in_extras_for_TestIfStmts-12                           2325285               515.8 ns/op
BenchmarkData/no_err_for_TestInvopop-12                                   682748              1674 ns/op
BenchmarkData/no_err_for_TestParallelLazy-12                              307359              3378 ns/op
BenchmarkData/no_err_for_TestValidator-12                                2095874               587.1 ns/op
BenchmarkData/no_err_for_TestIfStmts-12                                  3606709               329.5 ns/op
BenchmarkData/no_err_for_TestLazyValidator-12                            1000000              1090 ns/op
BenchmarkData/err_in_code_for_TestLazyValidator-12                        921346              1232 ns/op
BenchmarkData/err_in_code_for_TestInvopop-12                              520557              2126 ns/op
BenchmarkData/err_in_code_for_TestParallelLazy-12                         320020              3489 ns/op
BenchmarkData/err_in_code_for_TestValidator-12                           1666033               730.3 ns/op
BenchmarkData/err_in_code_for_TestIfStmts-12                             2316396               521.0 ns/op
BenchmarkData/err_in_message_for_TestValidator-12                        1604588               744.0 ns/op
BenchmarkData/err_in_message_for_TestIfStmts-12                          2280768               600.6 ns/op
BenchmarkData/err_in_message_for_TestLazyValidator-12                     997928              1268 ns/op
BenchmarkData/err_in_message_for_TestInvopop-12                           544496              2260 ns/op
BenchmarkData/err_in_message_for_TestParallelLazy-12                      335959              3792 ns/op
BenchmarkData/err_in_optional_for_TestLazyValidator-12                    851504              1409 ns/op
BenchmarkData/err_in_optional_for_TestInvopop-12                          569956              2115 ns/op
BenchmarkData/err_in_optional_for_TestParallelLazy-12                     332600              3437 ns/op
BenchmarkData/err_in_optional_for_TestValidator-12                       1573495               776.4 ns/op
BenchmarkData/err_in_optional_for_TestIfStmts-12                         2266776               536.8 ns/op
BenchmarkData/no_err_in_optional_for_TestIfStmts-12                      3561025               337.3 ns/op
BenchmarkData/no_err_in_optional_for_TestLazyValidator-12                1000000              1094 ns/op
BenchmarkData/no_err_in_optional_for_TestInvopop-12                       715827              1678 ns/op
BenchmarkData/no_err_in_optional_for_TestParallelLazy-12                  356920              3485 ns/op
BenchmarkData/no_err_in_optional_for_TestValidator-12                    2071834               573.5 ns/op
PASS
ok      github.com/Jh123x/go-validate/validator 51.100s
```

#### Formatted Results (ns/op)

| No  | Test Case            | Input Value                                                                                   | If Stmts    | Lazy Validator | Parallel Validator | Validator   | Invopop    | Validator vs Invopop | Validator vs if |
| --- | -------------------- | --------------------------------------------------------------------------------------------- | ----------- | -------------- | ------------------ | ----------- | ---------- | -------------------- | --------------- |
| 1   | No Errors            | `Response{Code: 200,Message: "OK",Extras: map[string]any{}}`                                  | 329.5 ns/op | 1090 ns/op     | 3378 ns/op         | 587.1 ns/op | 1674 ns/op | 2.85x faster         | 0.561x as fast  |
| 2   | Error in Code        | `Response{Code: 0,Message: "OK",Extras: map[string]any{}}`                                    | 521.0 ns/op | 1232 ns/op     | 3489 ns/op         | 730.3 ns/op | 2126 ns/op | 6.70x faster         | 0.713x as fast  |
| 3   | Error in Msg         | `Response{Code: 200,Message: "",Extras: map[string]any{}}`                                    | 600.6 ns/op | 1268 ns/op     | 3792 ns/op         | 744.0 ns/op | 2260 ns/op | 3.04x faster         | 0.806x as fast  |
| 4   | Error in Extras      | `Response{Code: 200,Message: "OK",Extras: nil}`                                               | 515.8 ns/op | 1110 ns/op     | 3336 ns/op         | 611.2 ns/op | 2109 ns/op | 3.45x faster         | 0.844x as fast  |
| 5   | Error in Opt         | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "test",SetIfOptSet: ""}` | 536.8 ns/op | 1409 ns/op     | 3437 ns/op         | 776.4 ns/op | 2115 ns/op | 2.72x faster         | 0.691x as fast  |
| 6   | No Error in Opt      | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: ""}`     | 337.3 ns/op | 1094 ns/op     | 3485 ns/op         | 573.5 ns/op | 1678 ns/op | 2.93x faster         | 0.588x as fast  |
| 7   | Error in setIfOptSet | `Response{Code: 200,Message: "OK",Extras: map[string]any{},Optional: "",SetIfOptSet: "test"}` | 516.4 ns/op | 1315 ns/op     | 3383 ns/op         | 770.0 ns/op | 2068 ns/op | 2.69x faster         | 0.670x as fast  |

#### Results

For short validations such as this one, it seems that the parallel validator is slower than any of the other validators.

The normal validator performs best in most cases tested here.

In general, the validator is faster than the [Invopop](https://github.com/invopop/validation "validation") validator by a significant margin.
The validator is also faster than the parallel validator in most cases in the test suite here.

The if statements here are shown as a form of ideal performance. Given that the default evaluator performs at most 2x slower than the if statements, it is still a good performance improvement over the if statements if performance is not that tight a concern.

## Future tasks

- [ ] Add more validation options
- [ ] A more comprehensive benchmark
- [ ] Other types of validators for different use cases
