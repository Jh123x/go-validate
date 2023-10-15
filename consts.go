package govalidate

const (
	ErrorFormat = "[validation error] error validating %s:%s"
)

var (
	IsNotEmptyErr      = ValidateErrGenerator("IsNotEmpty", "value is empty")
	InvalidLengthError = ValidateErrGenerator("IsLength", "invalid length")
)
