package home

import (
	"github.com/cucumber/godog"
	"github.com/ppcamp/e2e-testing/support"
)

func Register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I launch the home page$`, support.WithReporter(launchHomePage))
	ctx.Step(`^I enter with text globo$`, support.WithReporter(enterWithTextGlobo))
	ctx.Step(`^I enter with text msn$`, support.WithReporter(enterWithText))
	ctx.Step(`^I hit the search button$`, support.WithReporter(clickButton))
	ctx.Step(`^Verify the page title`, support.WithReporter(verifyTitle))
}
