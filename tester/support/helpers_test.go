package support_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ppcamp/e2e-testing/support"
	"github.com/stretchr/testify/assert"
)

func TestAwait(t *testing.T) {
	assert := assert.New(t)

	cbError := errors.New("some unhandled error")
	cbFactory := func(attempts int, sleep time.Duration) support.CallbackClosure {
		n := -1
		return func(_ context.Context) error {
			n++
			time.Sleep(sleep)
			if attempts == n {
				return nil
			}
			return cbError
		}
	}

	retriesAwaysFail := 100

	testCases := []struct {
		desc                  string
		expect                error
		retries, afterRetries int
		ctxTimer, afterSleep  time.Duration
	}{
		{"should run with default values", nil, 0, 0, 0, 0},
		{"should fail", cbError, 0, retriesAwaysFail, 0, 0},
		{"should work after 1st attempt", nil, 1, 1, 0, 0},
		{"should cancel after 200ms (retries pending)", context.DeadlineExceeded, 50, retriesAwaysFail, 123 * time.Millisecond, 0},
		{"should cancel after 200ms (taking too much time)", context.DeadlineExceeded, retriesAwaysFail, retriesAwaysFail, 200 * time.Millisecond, 500 * time.Millisecond},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			timer := 3 * time.Second
			if tC.ctxTimer > 0 {
				timer = tC.ctxTimer
			}
			ctx, cancel := context.WithTimeout(context.Background(), timer)
			defer cancel()

			err := support.Await(
				ctx,
				cbFactory(tC.afterRetries, tC.afterSleep),
				support.WithRetries(tC.retries),
				support.WithRunEveryTimer(100*time.Millisecond),
			)

			assert.ErrorIs(err, tC.expect)
		})
	}
}
