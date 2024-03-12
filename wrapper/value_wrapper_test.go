package wrapper

import (
	"testing"

	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestValueWrapper(t *testing.T) {
	tests := map[string]struct {
		value       int
		options     ttypes.ValTest[int]
		expectedErr error
	}{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value)
			err := valueWrapper.WithOptions(tc.options).Validate()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestOptionCompatibility(t *testing.T) {
	tests := map[string]struct {
		value               int
		options             ttypes.ValTest[int]
		expectedValidateErr error
	}{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			valueWrapper := NewValueWrapper(tc.value)
			option := valueWrapper.WithOptions(tc.options).ToOption()
			assert.Equal(t, tc.expectedValidateErr, option())
		})
	}
}
