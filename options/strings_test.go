package options

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/stretchr/testify/assert"
)

func TestIsValidURI(t *testing.T) {
	tests := map[string]struct {
		url         string
		expectedErr error
	}{
		"valid url with https scheme": {
			url:         "https://github.com/Jh123x/go-validate",
			expectedErr: nil,
		},
		"valid url with http scheme": {
			url:         "http://google.com",
			expectedErr: nil,
		},
		"invalid url no scheme": {
			url:         "www.google.com",
			expectedErr: errs.InvalidURIError,
		},
		"invalid url no colon": {
			url:         "http//www.google.com",
			expectedErr: errs.InvalidURIError,
		},
		"invalid only scheme": {
			url:         "http",
			expectedErr: errs.InvalidURIError,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedErr, IsValidURI(testCase.url)())
		})
	}
}
