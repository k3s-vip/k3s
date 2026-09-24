package wait

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

type (
	ConditionFunc            = wait.ConditionFunc
	ConditionWithContextFunc = wait.ConditionWithContextFunc
	Backoff                  = wait.Backoff
)

var ErrWaitTimeout = wait.ErrWaitTimeout

type interruptedError struct {
	cause error
}

func (e *interruptedError) Error() string {
	if e.cause == nil {
		return ErrWaitTimeout.Error()
	}
	return ErrWaitTimeout.Error() + ": " + e.cause.Error()
}

func (e *interruptedError) Is(target error) bool {
	return target == ErrWaitTimeout || errors.Is(e.cause, target)
}

func (e *interruptedError) Unwrap() error { return e.cause }

func ErrorInterrupted(cause error) error {
	if cause == nil {
		return ErrWaitTimeout
	}
	return &interruptedError{cause: cause}
}

func Interrupted(err error) bool {
	return err != nil && (errors.Is(err, ErrWaitTimeout) ||
		errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded))
}

type intervalGen func() (d time.Duration, ok bool)

type pollConfig struct {
	immediate bool
	timeout   time.Duration
	next      intervalGen
}

func run(ctx context.Context, cfg pollConfig, cond ConditionWithContextFunc) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if cond == nil || cfg.next == nil {
		return ErrWaitTimeout
	}

	var deadlineCh <-chan time.Time
	if cfg.timeout > 0 {
		t := time.NewTimer(cfg.timeout)
		defer t.Stop()
		deadlineCh = t.C
	}

	if cfg.immediate {
		done, err := cond(ctx)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}

	for {
		d, ok := cfg.next()
		if !ok {
			return ErrWaitTimeout
		}

		timer := time.NewTimer(d)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-deadlineCh:
			timer.Stop()
			return ErrWaitTimeout
		}

		done, err := cond(ctx)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}
}

func backoffInterval(b Backoff) intervalGen {
	if b.Duration <= 0 {
		return func() (time.Duration, bool) { return 0, false }
	}
	local := b
	return func() (time.Duration, bool) {
		return local.Step(), true
	}
}

func chanToCtx(parent context.Context, stopCh <-chan struct{}) (context.Context, context.CancelFunc) {
	if parent == nil {
		parent = context.Background()
	}
	if stopCh == nil {
		return context.WithCancel(parent)
	}
	ctx, cancel := context.WithCancel(parent)
	go func() {
		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()
	return ctx, cancel
}

func fixedInterval(interval time.Duration, jitter float64) intervalGen {
	return func() (time.Duration, bool) {
		if interval <= 0 {
			return 0, false
		}
		if jitter <= 0 {
			return interval, true
		}
		//nolint:gosec
		return interval + time.Duration(rand.Float64()*jitter*float64(interval)), true
	}
}

func lift(f ConditionFunc) ConditionWithContextFunc {
	if f == nil {
		return nil
	}
	return func(context.Context) (bool, error) { return f() }
}

func BackoffUntil(f func(), backoff Backoff, sliding bool, stopCh <-chan struct{}) {
	if f == nil {
		return
	}
	ctx, cancel := chanToCtx(context.Background(), stopCh)
	defer cancel()

	next := backoffInterval(backoff)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		var start time.Time
		if !sliding {
			start = time.Now()
		}

		f()

		d, ok := next()
		if !ok {
			return
		}

		if !sliding {
			if elapsed := time.Since(start); elapsed < d {
				d -= elapsed
			} else {
				d = 0
			}
		}

		if d > 0 {
			timer := time.NewTimer(d)
			select {
			case <-timer.C:
			case <-ctx.Done():
				timer.Stop()
				return
			}
		}
	}
}

func ExponentialBackoff(backoff Backoff, condition ConditionFunc) error {
	return ExponentialBackoffWithContext(context.Background(), backoff, lift(condition))
}

func ExponentialBackoffWithContext(ctx context.Context, backoff Backoff, condition ConditionWithContextFunc) error {
	return run(ctx, pollConfig{
		immediate: true,
		next:      backoffInterval(backoff),
	}, condition)
}

func Poll(interval, timeout time.Duration, condition ConditionFunc) error {
	return run(context.Background(), pollConfig{
		immediate: true,
		timeout:   timeout,
		next:      fixedInterval(interval, 0),
	}, lift(condition))
}

func PollImmediate(interval, timeout time.Duration, condition ConditionFunc) error {
	return Poll(interval, timeout, condition)
}

func PollInfinite(interval time.Duration, condition ConditionFunc) error {
	return run(context.Background(), pollConfig{
		immediate: true,
		next:      fixedInterval(interval, 0),
	}, lift(condition))
}

func PollImmediateInfinite(interval time.Duration, condition ConditionFunc) error {
	return PollInfinite(interval, condition)
}

func PollUntil(interval time.Duration, condition ConditionFunc, stopCh <-chan struct{}) error {
	ctx, cancel := chanToCtx(context.Background(), stopCh)
	defer cancel()
	return run(ctx, pollConfig{
		immediate: true,
		next:      fixedInterval(interval, 0),
	}, lift(condition))
}

func PollImmediateUntil(interval time.Duration, condition ConditionFunc, stopCh <-chan struct{}) error {
	return PollUntil(interval, condition, stopCh)
}

func PollUntilContextCancel(ctx context.Context, interval time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	return run(ctx, pollConfig{
		immediate: immediate,
		next:      fixedInterval(interval, 0),
	}, condition)
}

func PollUntilContextTimeout(ctx context.Context, interval, timeout time.Duration, immediate bool, condition ConditionWithContextFunc) error {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return run(ctx, pollConfig{
		immediate: immediate,
		next:      fixedInterval(interval, 0),
	}, condition)
}

func Forever(f func(), period time.Duration) {
	Until(f, period, nil)
}

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0 {
		return duration
	}
	//nolint:gosec
	return duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
}

func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
	JitterUntil(f, period, 0, true, stopCh)
}

func NonSlidingUntil(f func(), period time.Duration, stopCh <-chan struct{}) {
	JitterUntil(f, period, 0, false, stopCh)
}

func JitterUntil(f func(), period time.Duration, jitterFactor float64, sliding bool, stopCh <-chan struct{}) {
	if f == nil {
		return
	}
	ctx, cancel := chanToCtx(context.Background(), stopCh)
	defer cancel()
	JitterUntilWithContext(ctx, func(context.Context) { f() }, period, jitterFactor, sliding)
}

func UntilWithContext(ctx context.Context, f func(context.Context), period time.Duration) {
	JitterUntilWithContext(ctx, f, period, 0, true)
}

func NonSlidingUntilWithContext(ctx context.Context, f func(context.Context), period time.Duration) {
	JitterUntilWithContext(ctx, f, period, 0, false)
}

func JitterUntilWithContext(ctx context.Context, f func(context.Context), period time.Duration, jitterFactor float64, sliding bool) {
	if ctx == nil {
		ctx = context.Background()
	}
	if f == nil || period <= 0 {
		return
	}

	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()

	resetTimer := func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(Jitter(period, jitterFactor))
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if !sliding {
			resetTimer()
		}

		f(ctx)

		if sliding {
			resetTimer()
		}

		select {
		case <-ctx.Done():
			return
		case <-timer.C:
		}
	}
}

func PollJitter(ctx context.Context, interval time.Duration, jitter float64, condition ConditionFunc) error {
	return run(ctx, pollConfig{
		immediate: true,
		next:      fixedInterval(interval, jitter),
	}, lift(condition))
}

func PollJitterWithContext(ctx context.Context, interval time.Duration, jitter float64, immediate bool, condition ConditionWithContextFunc) error {
	return run(ctx, pollConfig{
		immediate: immediate,
		next:      fixedInterval(interval, jitter),
	}, condition)
}

func PollJitterUntil(interval time.Duration, jitter float64, condition ConditionFunc, stopCh <-chan struct{}) error {
	ctx, cancel := chanToCtx(context.Background(), stopCh)
	defer cancel()
	return run(ctx, pollConfig{
		immediate: true,
		next:      fixedInterval(interval, jitter),
	}, lift(condition))
}
