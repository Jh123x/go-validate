package validator

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

// TestLazyValidator tests the LazyValidator.
func TestLazyValidator(t *testing.T) {
	tests := map[string]struct {
		options     []ttypes.Validate
		expectedErr error
	}{
		"default case with options should return nil": {
			options:     []ttypes.Validate{},
			expectedErr: nil,
		},
		"with options with no errors should not return an error": {
			options:     []ttypes.Validate{validateWNil},
			expectedErr: nil,
		},
		"with options with errors should return an error": {
			options:     []ttypes.Validate{validateWErr},
			expectedErr: errTest,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			validator := NewLazyValidator()
			assert.Equal(t, testCase.expectedErr, validator.WithOptions(testCase.options...).Validate())
		})
	}
}

// TestLazyValidator_cache ensure that LazyValidator can be cached.
func TestLazyValidator_Caching(t *testing.T) {
	validator := NewLazyValidator()
	validator2 := validator.WithOptions(validateWErr)
	assert.Nil(t, validator.Validate())
	assert.Equal(t, errTest, validator2.Validate())
}

func TestLazyValidator_ReadMeExample(t *testing.T) {
	validator := NewLazyValidator()
	err := validator.WithOptions(
		options.IsNotEmpty("").WithError(fmt.Errorf("empty string")),             // Fails and returns error.
		options.IsLength([]string{}, 0, 3).WithError(fmt.Errorf("empty string")), // Will not be evaluated.
	).Validate()
	assert.Equal(t, fmt.Errorf("empty string"), err)
}
