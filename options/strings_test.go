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

func TestIsValidJson(t *testing.T) {
	tests := map[string]struct {
		url         string
		expectedErr error
	}{
		"valid json": {
			url:         `{"name":"jh123x"}`,
			expectedErr: nil,
		},
		"valid empty arr": {
			url:         `[]`,
			expectedErr: nil,
		},
		"valid empty obj": {
			url:         `{}`,
			expectedErr: nil,
		},
		"invalid json": {
			url:         `{"name":"jh123x"`,
			expectedErr: errs.InvalidJsonError,
		},
		"empty json string": {
			url:         ``,
			expectedErr: errs.InvalidJsonError,
		},
		"invalid arr": {
			url:         `[1,2,3`,
			expectedErr: errs.InvalidJsonError,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedErr, IsValidJson(testCase.url)())
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := map[string]struct {
		email       string
		expectedErr error
	}{
		"valid email": {
			email:       "email@email.com",
			expectedErr: nil,
		},
		"invalid email": {
			email:       "email",
			expectedErr: errs.InvalidEmailError,
		},
		"empty email": {
			email:       "",
			expectedErr: errs.InvalidEmailError,
		},
		"email with space": {
			email:       "email @email.com",
			expectedErr: errs.InvalidEmailError,
		},
		"email with multiple @": {
			email:       "email@@email.com",
			expectedErr: errs.InvalidEmailError,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedErr, IsValidEmail(testCase.email)())
		})
	}
}
