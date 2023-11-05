package options

import (
	"net/url"

	"github.com/Jh123x/go-validate/errs"
	types "github.com/Jh123x/go-validate/ttypes"
)

// IsValidURL validates that the provided string is a valid URL.
func IsValidURL(urlStr string) types.Validate {
	return WithRequire(func() bool {
		_, err := url.ParseRequestURI(urlStr)
		return err == nil
	}, errs.InvalidURLError)
}
