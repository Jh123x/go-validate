package validator

import "fmt"

var (
	errTest      = fmt.Errorf("test error")
	validateWNil = func() error { return nil }
	validateWErr = func() error { return errTest }
)
