package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Concurrency      uint32 = 0 // 0 also evals to 1
	SiteURL          string = "https://www.google.com"
	Browser          string = "chrome"
	Headless         bool   = false
	ScreenShotFolder string = "../images"
	IsTest           bool   = false
)

func init() {
	readEnv() // allow to pass over env

	err := os.MkdirAll(ScreenShotFolder, os.ModePerm)
	if err != nil {
		logrus.Panic(err)
	}
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

	if v, ok := os.LookupEnv(`TEST_SCREENSHOT_FOLDER`); ok {
		ScreenShotFolder = v
	}

	if v, ok := os.LookupEnv(`TEST_HEADLESS`); ok {
		if vparsed, err := strconv.ParseBool(strings.Trim(v, " ")); err == nil {
			Headless = vparsed
		}
	}

	if v, ok := os.LookupEnv(`TEST_IS_TEST`); ok {
		if vparsed, err := strconv.ParseBool(strings.Trim(v, " ")); err == nil {
			IsTest = vparsed
		}
	}

	if v, ok := os.LookupEnv(`TEST_THREAD`); ok {
		if vparsed, err := strconv.ParseUint(strings.Trim(v, " "), 10, 32); err == nil {
			Concurrency = uint32(vparsed)
		}
	}
}
