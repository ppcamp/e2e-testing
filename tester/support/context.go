package support

import (
	"context"
	"fmt"

	"golang.org/x/exp/constraints"
)

var reporterKey = reporterSymbol{}

type reporterSymbol struct{}

// contextWithReporter creates a new reporter and returns a new context with this value embeded
func contextWithReporter(ctx context.Context) (context.Context, error) {
	instance, err := playwrightInstance()
	if err != nil {
		return ctx, err
	}

	reporter := &implReporter{instance.page, instance.browser, instance.pw}
	return context.WithValue(ctx, reporterKey, reporter), nil
}

// fromContext will try to get the Reporter from the current context
func fromContext(ctx context.Context) (Reporter, error) {
	value := ctx.Value(reporterKey)
	v, ok := value.(Reporter)
	if !ok {
		return nil, fmt.Errorf("fail to cast to reporter interface")
	}
	return v, nil
}

// WithReporter is a decorator used to retrieve Reporter from the context and pass it as argument
// to the step functions.
func WithReporter(fn func(reporter Reporter) error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		reporter, err := fromContext(ctx)
		if err != nil {
			return err
		}
		return fn(reporter)
	}
}

// WithReporter is a decorator used to retrieve Reporter from the context and pass it as argument
// to the step functions.
func WithReporterT[T constraints.Ordered](fn func(reporter Reporter, extra T) error) func(ctx context.Context, extra T) error {
	return func(ctx context.Context, extra T) error {
		reporter, err := fromContext(ctx)
		if err != nil {
			return err
		}
		return fn(reporter, extra)
	}
}
