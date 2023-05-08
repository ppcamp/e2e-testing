package main

import (
	"github.com/ppcamp/e2e-testing/steps/login"
	"github.com/ppcamp/e2e-testing/support"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(support.BeforeScenario)
	ctx.After(support.AfterScenario)

	step := ctx.StepContext()
	step.Before(support.BeforeStep)

	ctx.Step(`^I launch the test site$`, support.WithReporter(login.LaunchHomePage))
}
