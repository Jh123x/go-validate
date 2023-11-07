package options

import (
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
