package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/TheVovchenskiy/banners/pkg/contextManager"
	"github.com/TheVovchenskiy/banners/pkg/logging"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestId := uuid.NewString()

		contextLogger := logging.Logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"URL":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"request_id":  requestId,
		})

		contextLogger.Info("got request")
		ctx := context.WithValue(r.Context(), contextManager.ContextLoggerKey, contextLogger)

		next.ServeHTTP(w, r.WithContext(ctx))

		contextLogger.WithFields(logrus.Fields{
			"working_time": time.Since(start).String(),
		}).Info("working time")
	})
}
