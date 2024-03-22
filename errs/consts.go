package errs

const (
	ErrorFormat = "[validation error] error validating %s:%s"
)

var (
	IsEmptyError       = NewValidateError("IsEmpty", "value is not empty")
	IsNotEmptyErr      = NewValidateError("IsNotEmpty", "value is empty")
	IsDefaultErr       = NewValidateError("IsDefault", "value is not default")
	IsNotDefaultErr    = NewValidateError("IsNotDefault", "value is default")
	InvalidLengthError = NewValidateError("IsLength", "invalid length")
	OrError            = NewValidateError("Or", "no options passed")
	ContainsError      = NewValidateError("Contains", "value not found in array")
	InvalidURIError    = NewValidateError("IsValidURL", "invalid url")
	InvalidJsonError   = NewValidateError("IsValidJson", "invalid json")
	InvalidEmailError  = NewValidateError("IsValidEmail", "invalid email")
)
