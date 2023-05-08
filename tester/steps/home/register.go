package home

import (
	"github.com/cucumber/godog"
	"github.com/ppcamp/e2e-testing/support"
)

func Register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I launch the home page$`, support.WithReporter(launchHomePage))
	ctx.Step(`^Verify the page title$`, support.WithReporter(verifyTitle))
}
