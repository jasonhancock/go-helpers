package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLCFirst(t *testing.T) {
	require.Equal(t, "hELLO", LCFirst("HELLO"))
	require.Equal(t, "hello", LCFirst("Hello"))
	require.Equal(t, "hello", LCFirst("hello"))
}

func TestUCFirst(t *testing.T) {
	require.Equal(t, "HELLO", UCFirst("HELLO"))
	require.Equal(t, "Hello", UCFirst("hello"))
	require.Equal(t, "Hello", UCFirst("Hello"))
}
