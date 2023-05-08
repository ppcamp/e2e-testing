package main

import (
	"github.com/ppcamp/e2e-testing/steps/login"
	"github.com/ppcamp/e2e-testing/support"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	support.Setup(ctx)

	login.Register(ctx)
}
