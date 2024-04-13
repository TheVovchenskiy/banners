package routers

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/delivery"
	"github.com/TheVovchenskiy/banners/internal/usecase"
	"github.com/gorilla/mux"
)

func MountAuthRouter(router *mux.Router, userStorage usecase.UserStorage, roleStorage usecase.RoleStorage) {
	handler := delivery.NewAuthHandler(userStorage, roleStorage)

	router.Handle("/register", http.HandlerFunc(handler.HandleRegistration)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(handler.HandleLogin)).Methods("POST")
}
