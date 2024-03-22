package main

import (
	"fmt"
	"testing"

	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

var (
	errTest = fmt.Errorf("test error")
)

/* Test Scenario for Benchmarking */
type Response struct {
	Code        int            // Must be non-zero.
	Message     string         // Must be non-empty.
	Extras      map[string]any // Must be non-nil.
	Optional    string         // Optional.
	SetIfOptSet string         // Set if Optional is set, empty otherwise.
}

// benchmarkValidator benchmarks the validateFn.
func benchmarkValidator(b *testing.B, response *Response, validateFn ttypes.ValTest[*Response], hasErr bool) {
	for i := 0; i < b.N; i++ {
		err := validateFn(response)
		assert.Equal(b, hasErr, err != nil, fmt.Sprintf("expected error: %v, got: %v", hasErr, err))
	}
}
