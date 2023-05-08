package main

import (
	"testing"

	"github.com/ppcamp/e2e-testing/config"
	log "github.com/sirupsen/logrus"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
)

var opts godog.Options

func TestMain(m *testing.M) {
	opts = godog.Options{
		Format:      "pretty",
		Paths:       []string{"features"},
		Concurrency: int(config.Concurrency),
	}

	godog.BindCommandLineFlags("godog.", &opts)

	pflag.Parse()

	opts.Paths = pflag.Args()

	_ = m.Run()

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

func TestFeatures(t *testing.T) {
	if config.IsTest {
		log.Info("is running a test")
		t.Skip()
	}

	log.WithFields(log.Fields{
		"Concurrency":   opts.Concurrency,
		"Paths":         opts.Paths,
		"Format":        opts.Format,
		"Tags":          opts.Tags,
		"StopOnFailure": opts.StopOnFailure,
	}).Info("Initializing tests")

	status := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Errorf("zero status code expected, %d received", status)
	}
}
