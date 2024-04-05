package response_test

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
)

func TestServeJSONError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"valid error", serverErrors.ErrInternal, `{"error":"internal server error"}`},
		{"empty error", errors.New(""), `{"error":"internal server error"}`},
		{"nil error", nil, `{"error":"internal server error"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			response.ServeJsonError(context.Background(), rr, tt.err)

			if rr.Body.String() != tt.want {
				t.Errorf("ServeJSONError() = %v, want %v", rr.Body.String(), tt.want)
			}
		})
	}
}
