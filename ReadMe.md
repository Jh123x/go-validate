# Go Validate

A mini validation library in Pure Go.

## Usage

```go

import (
    validate "github.com/Jh123x/go-validate"
)

func main(){
    validator := NewLazyValidator()
    validator.Validate(
        IsNotEmpty("").WithError()
    )
}

```
