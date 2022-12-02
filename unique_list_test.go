package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueList(t *testing.T) {
	l := NewUniqueList[string]()
	l.Add("a", "c", "b", "a")
	l.Add("a")
	require.Equal(t, []string{"a", "c", "b"}, l.Data())
}
