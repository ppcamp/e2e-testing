package support

import (
	"context"
	"fmt"
)

var ReporterKey = reporterType{}

type reporterType struct{}

// contextWithReporter creates a new reporter and returns a new context with this value embeded
func contextWithReporter(ctx context.Context) (context.Context, error) {
	browser, page, err := playwrightInstance()
	if err != nil {
		return ctx, err
	}

	reporter := &implReporter{page, browser}
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

type StepFunc func(Reporter) error
type GodogStepFunc func(ctx context.Context) error

// WithReporter is a decorator used to retrieve Reporter from the context and pass it as argument
// to the step functions.
func WithReporter(fn StepFunc) GodogStepFunc {
	return func(ctx context.Context) error {
		reporter, err := fromContext(ctx)
		if err != nil {
			return err
		}
		return fn(reporter)
	}
}
