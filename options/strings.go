package options

import (
	"encoding/json"
	"net/url"

	"github.com/Jh123x/go-validate/errs"
	types "github.com/Jh123x/go-validate/ttypes"
)

// IsValidURI validates that the provided string is a valid URL.
func IsValidURI(uriStr string) types.Validate {
	return WithRequire(func() bool {
		_, err := url.ParseRequestURI(uriStr)
		return err == nil
	}, errs.InvalidURIError)
}

// IsValidJson validates that the provided string is a valid JSON.
func IsValidJson(jsonStr string) types.Validate {
	return WithRequire(func() bool {
		return json.Valid([]byte(jsonStr))
	}, errs.InvalidJsonError)
}

// IsValidEmail validates the provided string is a valid email address.
func IsValidEmail(email string) types.Validate {
	return WithRequire(func() bool {
		return emailRegex.MatchString(email)
	}, errs.InvalidEmailError)
}
