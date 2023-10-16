package errs

const (
	ErrorFormat = "[validation error] error validating %s:%s"
)

var (
	IsEmptyError       = NewValidateError("IsEmpty", "value is not empty")
	IsNotEmptyErr      = NewValidateError("IsNotEmpty", "value is empty")
	InvalidLengthError = NewValidateError("IsLength", "invalid length")
)
