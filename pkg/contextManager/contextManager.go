package contextManager

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ContextKey string

const (
	ContextLoggerKey = ContextKey("logger")
)

func GetContextLogger(ctx context.Context) *logrus.Entry {
	logger, ok := ctx.Value(ContextLoggerKey).(*logrus.Entry)
	if !ok {
		defaultLogger := logrus.New()
		defaultLogger.SetLevel(logrus.InfoLevel)
		return defaultLogger.WithField("default", true)
	}
	return logger
}
