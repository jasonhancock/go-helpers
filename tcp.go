package helpers

import (
	"context"
	"fmt"
	"net"
	"time"
)

// Logger is an interface for logging messages.
type Logger interface {
	Warn(msg any, keyvals ...any)
	Info(msg any, keyvals ...any)
}

// WaitTCP will wait until a TCP connection has been established to addr. addr
// should be in "host:port" or "ip:port" format. Will continue to poll until the
// context is cancelled, in which case an error will be returned. If the returned
// error is nil, that means a connection was successfully established. interval
// is how often we should attempt to establish a connection. The logger can be nil.
func WaitTCP(ctx context.Context, l Logger, addr string, interval time.Duration) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return fmt.Errorf("resolving tcp addr failed: %w", err)
	}

	ticker := time.NewTimer(0)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				if l != nil {
					l.Warn("connection failed", "addr", addr, "error", err)
				}
			} else {
				if l != nil {
					l.Info("connection succeeded", "addr", addr)
				}
				_ = conn.Close()
				return nil
			}
			ticker.Reset(interval)
		}
	}
}
