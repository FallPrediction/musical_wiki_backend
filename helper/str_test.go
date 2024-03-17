package helper_test

import (
	. "musical_wiki/helper"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert := assert.New(t)
	characters := strings.Split("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

	tests := []struct {
		name         string
		length       int
		expectLength int
	}{
		{
			name:         "positive integer",
			length:       10,
			expectLength: 10,
		},
		{
			name:         "negative integer",
			length:       -1,
			expectLength: 0,
		},
		{
			name:         "zero",
			length:       0,
			expectLength: 0,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			randomStr := NewStr().Random(test.length)
			for _, ch := range randomStr {
				assert.True(slices.Contains(characters, string(ch)))
			}
			assert.Equal(len(randomStr), test.expectLength)
		})
	}
}
