package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromMap(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		data := map[string]any{
			"foo": "bar",
		}

		result, err := FromMap[string](data, "foo")
		require.NoError(t, err)
		require.Equal(t, "bar", result)
	})

	t.Run("ints", func(t *testing.T) {
		data := map[int]any{
			12: "bar",
		}

		result, err := FromMap[string](data, 12)
		require.NoError(t, err)
		require.Equal(t, "bar", result)
	})

	t.Run("int key, float value", func(t *testing.T) {
		data := map[int]any{
			12: 3.14159,
		}
		result, err := FromMap[float64](data, 12)
		require.NoError(t, err)
		require.Equal(t, 3.14159, result)
	})

	t.Run("key not found", func(t *testing.T) {
		data := map[string]any{}
		_, err := FromMap[string](data, "foo")
		require.Error(t, err)
		require.Equal(t, ErrKeyNotFound, err)
	})

	t.Run("value incorrect type", func(t *testing.T) {
		data := map[string]any{
			"foo": 1234,
		}
		_, err := FromMap[string](data, "foo")
		require.Error(t, err)
		require.Equal(t, ErrIncorrectType, err)
	})
}
