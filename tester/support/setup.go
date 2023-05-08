package support

import (
	"context"
	"fmt"
	"time"

	"github.com/cucumber/godog"
)

func BeforeScenario(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	fmt.Println("BeforeScenario:", sc.Name)

	return contextWithReporter(ctx)
}

func AfterScenario(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	fmt.Println("AfterScenario", sc.Name)
	if err != nil {
		return ctx, err
	}
	// time.Sleep(2 * time.Second)
	fmt.Println("loaded")
	return ctx, nil
}

func BeforeStep(ctx context.Context, st *godog.Step) (context.Context, error) {
	fmt.Println("BeforeStep", st.Text)
	time.Sleep(1 * time.Second)
	return ctx, nil
}
