package middleware

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/pkg/serverErrors"
)

func PanicRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response.ServeJsonError(r.Context(), w, serverErrors.ErrInternal)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})

}
