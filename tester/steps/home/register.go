package home

import (
	"github.com/cucumber/godog"
	s "github.com/ppcamp/e2e-testing/support"
)

func Register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I launch the home page$`, s.WithReporter(iLaunchTheHomePage))
	ctx.Step(`^I enter with text ([\w\s]+)$`, s.WithReporterT(iEnterWithText))
	ctx.Step(`^I hit the search button$`, s.WithReporter(iHitTheSearchButton))
	ctx.Step(`^Verify the page title`, s.WithReporter(verifyThePageTitle))
}
