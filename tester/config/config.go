package config

import (
	"os"
	"strconv"
)

var (
	Concurrency uint32 = 0
	SiteURL     string = "https://www.google.com"
	Browser     string = "chrome"
	Headless    bool   = false
)

func init() {
	readEnv() // allow to pass over env
}

// readEnv will load values from args, if passed, othewerise, it'll keep the original values
func readEnv() {
	// try to load those values from environment
	if v, ok := os.LookupEnv(`TEST_SITE`); ok {
		SiteURL = v
	}

	if v, ok := os.LookupEnv(`TEST_BROWSER`); ok {
		Browser = v
	}

	if v, ok := os.LookupEnv(`TEST_HEADLESS`); ok {
		if vparsed, err := strconv.ParseBool(v); err == nil {
			Headless = vparsed
		}
	}

	if v, ok := os.LookupEnv(`TEST_THREAD`); ok {
		if vparsed, err := strconv.ParseUint(v, 10, 32); err == nil {
			Concurrency = uint32(vparsed)
		}
	}
}
