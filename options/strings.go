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

func VIsValidURI(uriStr string) types.ValTest[string] {
	return func(uriStr string) error {
		_, err := url.ParseRequestURI(uriStr)
		if err != nil {
			return errs.InvalidURIError
		}
		return nil
	}
}

func VIsValidJson(jsonStr string) types.ValTest[string] {
	return func(jsonStr string) error {
		if !json.Valid([]byte(jsonStr)) {
			return errs.InvalidJsonError
		}
		return nil
	}
}

func VIsValidEmail(email string) types.ValTest[string] {
	return func(email string) error {
		if !emailRegex.MatchString(email) {
			return errs.InvalidEmailError
		}
		return nil
	}
}
