package support

import (
	"context"
	"time"
)

type CallbackClosure func(context.Context) error

type awaitImplOpt struct {
	duration time.Duration
	retry    int
}

type awaitOption func(*awaitImplOpt)

func WithRunEveryTimer(t time.Duration) awaitOption {
	return func(o *awaitImplOpt) {
		o.duration = t
	}
}
func WithRetries(attempts int) awaitOption {
	return func(o *awaitImplOpt) {
		o.retry = attempts
	}
}

// Await is a simple wrapper function that will execute a callback method every amount of time.
//
// The function is canceled if the passed context is canceled, or if it reaches the number of retries
// As default, this method will execute the callback every 0.2s, with NO retries, and using an empty
// context.
//
// In the example bellow, the Await will run 4 times, every 100ms, if it find the element first, the
// Await will close and return nil. If we increase the timer or the number of retries and the ctx
// is closed first, we'll return the ctx DeadlineExceed error.
//
// Note that we're gonna await a full function cycle. This means that, even if you already reached
// the execution max timer, you're function will still run and, only after it finishes, it'll check
// again.
//
// It forwards the context to the callback closure.
//
// Example
//
//	maxAwaitTime := 3 * time.Second
//	ctx, cancel := context.WithTimeout(context.TODO(), maxAwaitTime)
//	defer cancel()
//	var button playwright.Locator
//	err := Await(
//		ctx,
//		func(_ context.Context) error {
//	    	button, err := support.PlaywrightPage.Locator(locator)
//	    	if button != nil || err != nil {
//	    	    return fmt.Errorf("failed to get the App Button under Product Catalog: %v", err)
//	    	}
//		},
//		WithRetries(3),
//		WithTime(200),
//	)
//	if err != nil {
//	    return err
//	}
func Await(ctx context.Context, cb CallbackClosure, options ...awaitOption) error {
	o := &awaitImplOpt{100 * time.Millisecond, 0}
	for _, v := range options {
		v(o)
	}

	ticker := time.NewTicker(o.duration)
	retries := 0

	for {
		select {
		case <-ticker.C:
			if err := cb(ctx); err == nil {
				return nil
			} else if retries == o.retry {
				return err
			}
			retries++
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}
