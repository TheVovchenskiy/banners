package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TheVovchenskiy/banners/pkg/contextManager"
	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
	"github.com/sirupsen/logrus"
)

func ServerJsonData(ctx context.Context, w http.ResponseWriter, data any) {
	contextLogger := contextManager.GetContextLogger(ctx)

	js, err := json.Marshal(data)
	if err != nil {
		ServeJsonError(ctx, w, serverErrors.ErrInternal)
	}

	contextLogger.WithFields(logrus.Fields{
		"data": js,
	}).
		Info("serving json data")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Info("cannot send data")
	}

}

func ServeJsonError(ctx context.Context, w http.ResponseWriter, err error) {
	contextLogger := contextManager.GetContextLogger(ctx)
	apiErr := serverErrors.MapHTTPError(err)

	contextLogger.WithFields(logrus.Fields{
		"error":     err,
		"api_error": apiErr.Error,
		"status":    apiErr.Code,
	}).
		Error("got error, serving api_error")

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(apiErr.Code)
	data, _ := json.Marshal(map[string]string{"error": apiErr.Error.Error()})
	_, err = w.Write(data)
	if err != nil {
		contextLogger.WithFields(logrus.Fields{
			"error": err,
		}).
			Info("cannot send error")
	}
}
