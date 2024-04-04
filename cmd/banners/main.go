package main

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/TheVovchenskiy/banners/app/routers"
	"github.com/TheVovchenskiy/banners/configs"
)

// @title Banners API
// @version 1.0
// @description This is the first version of Banners API

// @contact.name Vladimir Konoplyuk
// @contact.url https://github.com/TheVovchenskiy
// @contact.email konoplyuk.v@yandex.ru

// @host localhost:8080
// @BasePath /api/v1
func main() {
	http.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", configs.SwaggerPort)),
	))
	go http.ListenAndServe(fmt.Sprintf(":%d", configs.SwaggerPort), nil)

	err := routers.Run()
	if err != nil {
		fmt.Println(err)
	}
}
