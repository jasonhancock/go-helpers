package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		haystack := []string{"a", "b", "c"}
		require.True(t, Contains(haystack, "c"))
		require.False(t, Contains(haystack, "d"))
	})

	t.Run("ints", func(t *testing.T) {
		haystack := []int{1, 2, 3}
		require.True(t, Contains(haystack, 2))
		require.False(t, Contains(haystack, 4))
	})
}
