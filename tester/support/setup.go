package support

import (
	"context"
	"fmt"
	"time"

	"github.com/cucumber/godog"
)

// Setup the before and after functions for scenarios and steps
func Setup(ctx *godog.ScenarioContext) {
	step := ctx.StepContext()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		fmt.Println("BeforeScenario:", sc.Name)
		return contextWithReporter(ctx)
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		fmt.Println("AfterScenario", sc.Name)
		if err != nil {
			return ctx, err
		}
		// time.Sleep(2 * time.Second)
		fmt.Println("loaded")
		return ctx, nil
	})

	step.Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		fmt.Println("BeforeStep", st.Text)
		time.Sleep(1 * time.Second)
		return ctx, nil
	})

	step.After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		fmt.Println("AfterStep", st.Text)
		time.Sleep(1 * time.Second)
		return ctx, nil
	})
}
