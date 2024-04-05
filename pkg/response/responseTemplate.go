package response

import (
	"encoding/json"
	"net/http"

	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
)

type ErrToSend struct {
	Message string `json:"message"`
}

func MarshalAndSend(w http.ResponseWriter, data any) {
	js, err := json.Marshal(data)
	if err != nil {
		ServeJsonError(w, serverErrors.ErrInternal)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func ServeJsonError(w http.ResponseWriter, err error) {
	apiErr := serverErrors.MapHTTPError(err)

	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(apiErr.Code)
	data, _ := json.Marshal(apiErr)
	w.Write(data)
}
