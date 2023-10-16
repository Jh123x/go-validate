package ttypes

import "fmt"

var (
	errTest               = fmt.Errorf("test error")
	errTest2              = fmt.Errorf("test error 2")
	validateWErr Validate = func() error { return errTest }
	validateWNil Validate = func() error { return nil }
)
