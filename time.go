package helpers

import (
	"context"
	"time"
)

// SleepCtx sleeps for the d duration amount of time, but will return immediately
// if the context is cancelled.
func SleepCtx(ctx context.Context, d time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(d):
	}
}
