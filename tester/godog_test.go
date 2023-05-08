package main

import (
	"os"
	"testing"

	"github.com/ppcamp/e2e-testing/config"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
)

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:      "pretty",
		Paths:       []string{"Features"},
		Concurrency: config.Concurrency,
	}

	godog.BindCommandLineFlags("godog.", &opts)

	pflag.Parse()

	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
