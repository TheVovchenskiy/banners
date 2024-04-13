package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger(logFile *os.File, logLevel logrus.Level) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: "Mon, 02 Jan 2006 15:04:05 MST",
	}

	Logger.SetOutput(logFile)
	Logger.SetFormatter(jsonFormatter)
	Logger.SetLevel(logLevel)
}

func LogError(logger *logrus.Entry, err error, while string) {
	logger.WithFields(logrus.Fields{
		"error": err,
	}).
		Error(while)
}
