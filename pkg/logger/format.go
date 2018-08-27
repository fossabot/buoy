package logger

import (
	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

const (
	JSON   = "json"
	Text   = "text"
	Fluent = "fluent"
)

func SetFortmat(format string, logger *logrus.Logger) {
	switch format {
	case Text:
		logger.Formatter = &logrus.TextFormatter{}
	case JSON:
		logger.Formatter = &logrus.JSONFormatter{}
	case Fluent:
		logger.Formatter = &joonix.FluentdFormatter{}
	}
}
