package home

import (
	"time"

	"github.com/ppcamp/e2e-testing/support"
)

func verifyThePageTitle(reporter support.Reporter) error {
	// expectedTitle := " - Google Search"
	// if err := common.VerifyTitle(reporter, expectedTitle); err != nil {
	// 	return err
	// }
	time.Sleep(1 * time.Second) // simulating some check
	return nil
}
