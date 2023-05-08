package support

import (
	"context"
	"fmt"
)

var ReporterKey = reporterType{}

type reporterType struct{}

// contextWithReporter creates a new reporter and returns a new context with this value embeded
func contextWithReporter(ctx context.Context) (context.Context, error) {
	instance, err := playwrightInstance()
	if err != nil {
		return ctx, err
	}

	reporter := &implReporter{instance.page, instance.browser, instance.pw}
	return context.WithValue(ctx, ReporterKey, reporter), nil
}

// fromContext will try to get the Reporter from the current context
func fromContext(ctx context.Context) (Reporter, error) {
	value := ctx.Value(ReporterKey)
	v, ok := value.(Reporter)
	if !ok {
		return nil, fmt.Errorf("fail to cast to reporter interface")
	}
	return v, nil
}

// WithReporter is a decorator used to retrieve Reporter from the context and pass it as argument
// to the step functions.
func WithReporter(fn StepFunc) godogStepFunc {
	return func(ctx context.Context) error {
		reporter, err := fromContext(ctx)
		if err != nil {
			return err
		}
		return fn(reporter)
	}
}

// WithoutParam is a decorator to ignore the ctx parameter
func WithoutParam(fn func() error) godogStepFunc {
	return func(_ context.Context) error {
		return fn()
	}
}
