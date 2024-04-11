package routers

import (
	"net/http"

	"github.com/TheVovchenskiy/banners/internal/delivery"
	"github.com/TheVovchenskiy/banners/internal/usecase"
	"github.com/gorilla/mux"
)

func MountAuthRouter(router *mux.Router, bannerUsecase *usecase.BannerUsecase) {
	handler := delivery.NewBannerHandler(bannerUsecase)

	router.Handle("/banners", http.HandlerFunc(handler.HandleGetBanners)).Methods("GET")
	// router.Handle("/banners/{id}", http.HandlerFunc(handler.HandleLogin)).Methods("POST")
}
