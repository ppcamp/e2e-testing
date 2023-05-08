package main

import (
	"os"
	"testing"

	"github.com/ppcamp/e2e-testing/config"
	"github.com/sirupsen/logrus"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
)

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:      "pretty",
		Paths:       []string{"features"},
		Concurrency: int(config.Concurrency),
	}

	godog.BindCommandLineFlags("godog.", &opts)

	pflag.Parse()

	opts.Paths = pflag.Args()

	logrus.WithFields(logrus.Fields{
		"Concurrency":   opts.Concurrency,
		"Paths":         opts.Paths,
		"Format":        opts.Format,
		"Tags":          opts.Tags,
		"StopOnFailure": opts.StopOnFailure,
	}).Info("Initializing tests")

	_ = godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	_ = m.Run()

	os.Exit(0)

	// To use the tests status ()
	// status := godog.TestSuite{
	// 	ScenarioInitializer: InitializeScenario,
	// 	Options:             &opts,
	// }.Run()
	// Optional: Run `testing` package's logic besides godog.
	// if st := m.Run(); st > status {
	// 	status = st
	// }
	// os.Exit(status)
}
