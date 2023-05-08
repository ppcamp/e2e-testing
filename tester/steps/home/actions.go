package home

import (
	"fmt"

	"github.com/ppcamp/e2e-testing/config"
	"github.com/ppcamp/e2e-testing/support"
)

const (
	xpathTextfield string = `//*[@type='search']`
)

func launchHomePage(reporter support.Reporter) error {
	if err := reporter.Get(config.SiteURL); err != nil {
		return fmt.Errorf("couldn't navigate to login url: %v", err)
	}

	return nil
}

func enterWithText(reporter support.Reporter) error {
	textInput := `google`

	locator, err := reporter.Locator(xpathTextfield)
	if err != nil {
		return err
	}

	return locator.Type(textInput)
}

func enterWithTextGlobo(reporter support.Reporter) error {
	textInput := `globo`

	locator, err := reporter.Locator(xpathTextfield)
	if err != nil {
		return err
	}

	return locator.Type(textInput)
}

func clickButton(reporter support.Reporter) error {
	xpathButton := `//*[@value="Pesquisa Google"]`

	return reporter.Click(xpathButton)
}
