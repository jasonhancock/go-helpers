package helpers

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDedup(t *testing.T) {
	data := []int{1, 1, 1, 2, 4, 3, 5, 5, 3, 2}
	result := Dedup(data)
	sort.Ints(result)
	require.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
