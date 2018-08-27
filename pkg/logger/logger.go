package logger

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Setup(config *LoggerConfig, logger *logrus.Logger) error {
	//set log level
	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	//set log formatter
	SetFortmat(config.LogFormat, logger)
	return nil
}
