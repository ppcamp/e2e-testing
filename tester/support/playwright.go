package support

import (
	"fmt"

	"github.com/ppcamp/e2e-testing/config"
	log "github.com/sirupsen/logrus"

	"github.com/playwright-community/playwright-go"
)

type instance struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	page    playwright.Page
}

// playwrightInstance creates and return the playwright page and browser.
// Usually this function is called before every scenario.
// This function will use the chromium base, so it'll only work for browser chromium based
func playwrightInstance() (*instance, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start playwright: %v", err)
	}

	option := playwright.BrowserTypeLaunchOptions{
		Channel:  &config.Browser,
		Headless: &config.Headless,
		Args:     config.ChromiumArgs,
	}

	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	page.SetViewportSize(1920, 1080)

	return &instance{pw, browser, page}, nil
}

func ScreenShot(page playwright.Page, path string) error {
	log.Debug("debug", path)
	_, err := page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String(path)})
	if err != nil {
		log.WithError(err).Error("fail to save image ", path)
	}
	return err
}
