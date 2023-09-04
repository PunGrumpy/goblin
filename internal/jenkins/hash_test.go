package jenkins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input    string
		expected uint32
	}{
		{"hello", 0xc8fd181b},
		{"world", 0x9099abe4},
		{"jenkins", 0x90f20366},
		{"golang", 0xa577347b},
		{"testing", 0x61553780},
	}

	for _, test := range tests {
		result := Hash(test.input)
		assert.Equal(t, test.expected, result, "Hash(%s) did not return the expected result", test.input)
	}
}
