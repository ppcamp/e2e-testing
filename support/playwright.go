// go install github.com/playwright-community/playwright-go/cmd/playwright
// playwright install --with-deps

package support

import (
	"fmt"
	"newtester/config"

	"github.com/playwright-community/playwright-go"
)

// playwrightInstance creates and return the playwright page and browser.
// Usually this function is called before every scenario.
// This function will use the chromium base, so it'll only work for browser chromium based
func playwrightInstance() (playwright.Browser, playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("could not start playwright: %v", err)
	}

	option := playwright.BrowserTypeLaunchOptions{
		Channel:  playwright.String(config.MauiBrowser),
		Headless: playwright.Bool(config.MauiHeadless),
		//SlowMo:   playwright.Float(100),
	}

	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		return nil, nil, fmt.Errorf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create page: %v", err)
	}

	return browser, page, nil
}
