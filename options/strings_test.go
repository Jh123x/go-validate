package options

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	tests := map[string]struct {
		url         string
		expectedErr error
	}{
		"valid url with https scheme": {
			url:         "https://www.google.com",
			expectedErr: nil,
		},
		"valid url with http scheme": {
			url:         "http://google.com",
			expectedErr: nil,
		},
		"invalid url no scheme": {
			url:         "www.google.com",
			expectedErr: errs.InvalidURLError,
		},
		"invalid url no colon": {
			url:         "http//www.google.com",
			expectedErr: errs.InvalidURLError,
		},
		"invalid only scheme": {
			url:         "http",
			expectedErr: errs.InvalidURLError,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedErr, IsValidURL(testCase.url)())
		})
	}
}
