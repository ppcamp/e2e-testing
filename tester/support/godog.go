package support

import (
	"context"

	"github.com/cucumber/godog"
	log "github.com/sirupsen/logrus"
)

type StepFunc func(Reporter) error

type godogStepFunc func(ctx context.Context) error

// Setup the before and after functions for scenarios and steps
func Setup(ctx *godog.ScenarioContext) {
	step := ctx.StepContext()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		log.Debug("BeforeScenario:", sc.Name)
		return contextWithReporter(ctx)
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		log.Debug("AfterScenario", sc.Name)
		if err != nil {
			log.WithError(err).WithField("Scenario", sc.Name)
		}
		return ctx, err
	})

	step.Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		log.Debug("BeforeStep", st.Text)
		return ctx, nil
	})

	step.After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		log.Debug("AfterStep", st.Text)
		if err != nil {
			log.WithError(err).WithField("Step", st.Text)
		}
		return ctx, err
	})
}
