package login

import (
	"fmt"
	"newtester/config"
	"newtester/support"
	"time"
)

const (
	xpathTextfield string = `//*[@type='search']`
	xpathButton    string = `//*[@value="Pesquisa Google"]`
	textInput      string = `search in google`
)

func LaunchHomePage(reporter support.Reporter) error {
	fmt.Println("home page")
	if err := reporter.Get(config.MauiLocal); err != nil {
		return fmt.Errorf("couldn't navigate to login url: %v", err)
	}

	time.Sleep(1 * time.Second)

	locator, err := reporter.Locator(xpathTextfield)
	if err != nil {
		return err
	}
	locator.TextContent()

	if err := locator.Type(textInput); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	if err := reporter.Click(xpathButton); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)
	return nil
}
