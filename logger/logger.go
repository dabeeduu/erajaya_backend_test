package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() *logrus.Logger {
	log = logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   false,
		ForceColors:     true,
	})
	log.SetLevel(logrus.DebugLevel)

	return log
}
