package routers

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/rest"
	"github.com/gorilla/mux"
)

func MountAuthRouter(router *mux.Router) {
	handler := rest.NewBannerHandler()

	router.Handle("/banners", http.HandlerFunc(handler.HandleGetBanners)).Methods("GET")
	// router.Handle("/banners/{id}", http.HandlerFunc(handler.HandleLogin)).Methods("POST")
}
