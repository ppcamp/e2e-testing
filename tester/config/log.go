package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var LogLevel log.Level = log.DebugLevel

func init() {
	if v, ok := os.LookupEnv(`TEST_LOG_LEVEL`); ok {
		levelStr := strings.Trim(v, " ")
		lvl, err := log.ParseLevel(levelStr)
		if err != nil {
			log.Panic("fail to parse log level", err)
		}
		LogLevel = lvl
	}

	log.SetLevel(LogLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
