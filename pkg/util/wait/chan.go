// Package ready provides a one-shot readiness signal that carries the result
// of an upstream component's health check.
package wait

import (
	"context"
	"sync"
)

// Chan is a one-shot readiness signal that reports the outcome of an upstream
// component's health check to any number of waiters. It allows dependent
// components to be started only after their upstream dependencies are verified
// to be fully ready to serve requests, instead of repeatedly retrying requests
// against an upstream that is not yet, or never will be, ready.
//
// A Chan is created with New, completed exactly once by its producer using
// MarkReady or MarkFailed, and observed by waiters using Done, Err, or Wait.
// A nil *Chan behaves like a signal that is never completed, so that waiters
// may also wait on components that are not running.
type Chan struct {
	once sync.Once
	mu   sync.Mutex
	err  error
	done chan struct{}
}

// New returns a new, incomplete Chan.
func New() *Chan {
	return &Chan{done: make(chan struct{})}
}

// MarkReady completes the signal, indicating that the upstream component
// passed its health check and is ready to serve requests. Only the first call
// to MarkReady or MarkFailed has any effect.
func (c *Chan) MarkReady() {
	c.complete(nil)
}

// MarkFailed completes the signal with the error that prevented the upstream
// component from becoming ready. Only the first call to MarkReady or
// MarkFailed has any effect.
func (c *Chan) MarkFailed(err error) {
	c.complete(err)
}

func (c *Chan) complete(err error) {
	if c == nil {
		return
	}
	c.once.Do(func() {
		c.mu.Lock()
		c.err = err
		c.mu.Unlock()
		close(c.done)
	})
}

// Done returns a channel that is closed once the signal is completed, allowing
// waiters to select on the readiness result. It returns nil for a nil *Chan.
func (c *Chan) Done() <-chan struct{} {
	if c == nil {
		return nil
	}
	return c.done
}

// Err returns the result of the health check: nil if the upstream component is
// ready, or the error that prevented it from becoming ready. It returns nil if
// the signal has not been completed yet; waiters should call it after Done is
// closed, or use Wait to observe the result.
func (c *Chan) Err() error {
	if c == nil {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

// Wait blocks until the signal is completed or the context is cancelled,
// returning the result of the health check, or the context error. Waiters that
// abort on a non-nil error avoid making requests to an upstream that is not, or
// never will be, ready. A nil *Chan blocks until the context is cancelled.
func (c *Chan) Wait(ctx context.Context) error {
	select {
	case <-c.Done():
		return c.Err()
	case <-ctx.Done():
		return ctx.Err()
	}
}
