package login

import (
	"github.com/cucumber/godog"
	"github.com/ppcamp/e2e-testing/support"
)

func Register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I launch the test site$`, support.WithReporter(launchHomePage))
}
