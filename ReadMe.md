# Go Validate

A mini validation library in Pure Go.

## Usage

```go

import (
    validate "github.com/Jh123x/go-validate"
)

func main(){
    validator := NewLazyValidator()
    err := validator.WithOptions(
        IsNotEmpty("").WithError(fmt.Errorf("empty string")),               // Fails and returns error.
        IsLength([]string{}, 0, 3).WithError(fmt.Errorf("empty string")),   // Will not be evaluated.
    ).Validate()
    if err != nil {
        // handle error
        ...
    }
}
```
