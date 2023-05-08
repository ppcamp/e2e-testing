package common

import (
	"fmt"

	"github.com/ppcamp/e2e-testing/support"
)

func VerifyTitle(reporter support.Reporter, title string) error {
	t, err := reporter.Title()
	if err != nil {
		return err
	}

	if t != title {
		return fmt.Errorf("title doesn't match")
	}

	return nil
}
