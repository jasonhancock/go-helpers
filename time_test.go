package helpers

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSleepCtx(t *testing.T) {
	t.Run("ctx cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		wg.Add(1)
		var ts time.Time
		start := time.Now()
		go func() {
			defer wg.Done()
			SleepCtx(ctx, 1*time.Minute)
			ts = time.Now()
		}()
		cancel()
		wg.Wait()
		require.WithinDuration(t, start, ts, 1*time.Millisecond)
	})

	t.Run("ctx cancelled", func(t *testing.T) {
		ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(1)
		var ts time.Time
		start := time.Now()
		go func() {
			defer wg.Done()
			SleepCtx(ctx, 10*time.Millisecond)
			ts = time.Now()
		}()
		wg.Wait()
		require.WithinDuration(t, start, ts, 12*time.Millisecond)
	})
}
