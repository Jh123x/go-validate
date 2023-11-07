package errs

const (
	ErrorFormat = "[validation error] error validating %s:%s"
)

var (
	IsEmptyError       = NewValidateError("IsEmpty", "value is not empty")
	IsNotEmptyErr      = NewValidateError("IsNotEmpty", "value is empty")
	InvalidLengthError = NewValidateError("IsLength", "invalid length")
	OrError            = NewValidateError("Or", "no options passed")
	ContainsError      = NewValidateError("Contains", "value not found in array")
	InvalidURIError    = NewValidateError("IsValidURL", "invalid url")
)
