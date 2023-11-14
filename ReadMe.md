# Go Validate

![CI Badge](https://github.com/Jh123x/go-validate/actions/workflows/go.yml/badge.svg "CI Badge")

A validation library in Pure Go.

## Features

1. Customized error messages for your validation rules
2. Easy to write custom rules
3. Define your own validation method
4. Fast (Visit our [benchmark](docs/benchmark.md) page for more details)

To see the list of options, you can refer to the [options page](docs/options.md).

## Installation

```sh
go get github.com/Jh123x/go-validate
```

## Usage

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

An example usage is shown in the code snippet above. To see a full list of options, you can refer to the [options page](docs/options.md).

## Issues

Please create an issue if you have any:

1. Questions
2. Suggestions
3. Find any bugs
4. Want to point out any mistakes

## Contribution Guide

1. Code a new feature or fix a bug.
2. Add tests for the new feature or the bug you fixed.
3. Fill up the [pull request template](.github/pull_request_template.md).
4. Create a pull request.
5. Wait for the CI to pass.
6. Wait for the maintainer to review your code.
7. Wait for the maintainer to merge your code.

## Future tasks

- [ ] Add other types of validators for different use cases.
- [ ] Add more comprehensive documentations.
