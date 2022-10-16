package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanupCommaDelimited(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a,b,c", []string{"a", "b", "c"}},
		{"a, b ,c", []string{"a", "b", "c"}},
		{"a,,c", []string{"a", "c"}},
		{"", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			results := CleanupCommaDelimited(tt.input)
			require.Equal(t, tt.expected, results)
		})
	}

}
