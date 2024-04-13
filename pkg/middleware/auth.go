package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/TheVovchenskiy/banners/configs"
	"github.com/TheVovchenskiy/banners/pkg/contextManager"
	"github.com/TheVovchenskiy/banners/pkg/response"
	"github.com/TheVovchenskiy/banners/pkg/token"
	"github.com/TheVovchenskiy/banners/pkg/utils"

	"github.com/golang-jwt/jwt"
)

const (
	bearer_schema = "Bearer"
)

func AuthMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			response.ServeJsonError(r.Context(), w, token.ErrAuthorizationHeaderRequired)
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			response.ServeJsonError(r.Context(), w, token.ErrInvalidToken)
			return
		}

		if bearerToken[0] != bearer_schema {
			response.ServeJsonError(r.Context(), w, token.ErrInvalidToken)
			return
		}

		tokenString := bearerToken[1]
		claims := &token.Claims{}

		accessToken, err := jwt.ParseWithClaims(tokenString, claims, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w: %v", token.ErrUnexpectedSigningMethod, jwtToken.Header["alg"])
			}
			return configs.JwtKey, nil
		})

		if err != nil || !accessToken.Valid {
			response.ServeJsonError(r.Context(), w, token.ErrInvalidToken)
			return
		}

		if !utils.In(claims.Role, allowedRoles) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), contextManager.ContextUserId, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
