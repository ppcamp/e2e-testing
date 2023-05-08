package support

import (
	"github.com/playwright-community/playwright-go"
)

type implReporter struct {
	playwright.Page
	browser playwright.Browser
}

// type playwrightMinimum interface {
// 	Type(text string, options ...playwright.PageTypeOptions) error
// 	Locator(selector string, options ...playwright.PageLocatorOptions) (playwright.Locator, error)
// 	Click(selector string, options ...playwright.PageClickOptions) error
// }

type Reporter interface {
	Get(url string) error
	playwright.Page
}

func (s *implReporter) Get(url string) error {
	_, err := s.Goto(url, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
	if err != nil {
		return err
	}
	return nil
}
