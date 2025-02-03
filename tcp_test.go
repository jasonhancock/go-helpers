package helpers

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	h "github.com/jasonhancock/go-testhelpers/generic"
	"github.com/stretchr/testify/require"
)

func TestWaitTCP(t *testing.T) {
	t.Run("socket listening", func(t *testing.T) {
		port := h.NewRandomPort(t)
		list, err := net.Listen("tcp", port)
		require.NoError(t, err)
		defer list.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		require.NoError(t, WaitTCP(ctx, nil, port, 1*time.Second))
	})

	t.Run("socket not listening", func(t *testing.T) {
		port := h.NewRandomPort(t)

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		err := WaitTCP(ctx, nil, port, 100*time.Millisecond)
		require.Error(t, err)
		require.Equal(t, context.DeadlineExceeded, err)
	})

	t.Run("socket eventually listens", func(t *testing.T) {
		port := h.NewRandomPort(t)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			require.NoError(t, WaitTCP(ctx, nil, port, 10*time.Millisecond))
		}()

		time.Sleep(500 * time.Millisecond)

		list, err := net.Listen("tcp", port)
		require.NoError(t, err)
		defer list.Close()

		wg.Wait()
	})
}
