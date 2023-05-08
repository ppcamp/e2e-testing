package support

import (
	"context"
	"time"

	"github.com/cucumber/godog"
	"github.com/sirupsen/logrus"
)

type StepFunc func(Reporter) error

type godogStepFunc func(ctx context.Context) error

// Setup the before and after functions for scenarios and steps
func Setup(ctx *godog.ScenarioContext) {
	step := ctx.StepContext()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		logrus.Debug("BeforeScenario:", sc.Name)
		return contextWithReporter(ctx)
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		logrus.Debug("AfterScenario", sc.Name)
		if err != nil {
			return ctx, err
		}
		// time.Sleep(2 * time.Second)
		logrus.Debug("loaded")
		return ctx, nil
	})

	step.Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		logrus.Debug("BeforeStep", st.Text)
		time.Sleep(1 * time.Second)
		return ctx, nil
	})

	step.After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		logrus.Debug("AfterStep", st.Text)
		time.Sleep(1 * time.Second)
		return ctx, nil
	})
}
