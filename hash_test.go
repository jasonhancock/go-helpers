package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMD5(t *testing.T) {
	require.Equal(t, "e99a18c428cb38d5f260853678922e03", MD5([]byte("abc123")))
}
