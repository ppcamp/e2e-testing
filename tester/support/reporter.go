package support

import (
	"github.com/playwright-community/playwright-go"
)

type implReporter struct {
	playwright.Page
	browser playwright.Browser
	pw      *playwright.Playwright
}

type Reporter interface {
	// Get is like a Goto with waituntil attached
	Get(url string) error

	// CloseAll connections and playwright support (clear scenario)
	CloseAll() error

	playwright.Page
}

func (s *implReporter) CloseAll() error {
	err := s.browser.Close()
	if err != nil {
		return err
	}

	err = s.pw.Stop()
	return err
}

func (s *implReporter) Get(url string) error {
	_, err := s.Goto(url, playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateNetworkidle})
	if err != nil {
		return err
	}
	return nil
}
