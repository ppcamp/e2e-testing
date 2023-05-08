package config

import log "github.com/sirupsen/logrus"

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
