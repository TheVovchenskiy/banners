package routers

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/delivery"
	"github.com/TheVovchenskiy/banners/internal/repository/psql"
	"github.com/TheVovchenskiy/banners/pkg/middleware"
	"github.com/gorilla/mux"
)

func MountBannerRouter(router *mux.Router, bannerStorage psql.BannerPsqlRepo) {
	handler := delivery.NewBannerHandler(bannerStorage)

	router.Handle("/banners", middleware.AuthMiddleware([]string{"admin"}, http.HandlerFunc(handler.HandleGetBanners))).Methods("GET")
	// router.Handle("/banners/{id}", http.HandlerFunc(handler.HandleLogin)).Methods("POST")
}
