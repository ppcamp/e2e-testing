package home

import (
	"time"

	"github.com/ppcamp/e2e-testing/support"
)

// verifyTitle
// TODO: check for the regexes to allow passing a variable from the feature file
// SEE https://github.com/cucumber/godog
func verifyTitle(reporter support.Reporter) error {
	// expectedTitle := " - Google Search"
	// if err := common.VerifyTitle(reporter, expectedTitle); err != nil {
	// 	return err
	// }
	time.Sleep(1 * time.Second) // simulating some check
	return nil
}
