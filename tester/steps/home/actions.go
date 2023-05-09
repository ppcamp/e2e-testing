package home

import (
	"fmt"

	"github.com/ppcamp/e2e-testing/config"
	"github.com/ppcamp/e2e-testing/support"
)

const (
	xpathTextfield string = `//*[@type='search']`
)

// iLaunchTheHomePage and await until network kiddle event
func iLaunchTheHomePage(reporter support.Reporter) error {
	if err := reporter.Get(config.SiteURL); err != nil {
		return fmt.Errorf("couldn't navigate to login url: %v", err)
	}

	return nil
}

func iEnterWithText(reporter support.Reporter, textInput string) error {
	locator, err := reporter.Locator(xpathTextfield)
	if err != nil {
		return err
	}

	return locator.Type(textInput)
}

func iHitTheSearchButton(reporter support.Reporter) error {
	xpathButton := `//*[@value="Pesquisa Google"]`

	return reporter.Click(xpathButton)
}
