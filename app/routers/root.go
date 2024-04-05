package routers

import (
	"fmt"
	"net/http"

	"github.com/TheVovchenskiy/banners/app"
	"github.com/TheVovchenskiy/banners/configs"
	_ "github.com/TheVovchenskiy/banners/docs"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() (err error) {
	swaggerRouter := mux.NewRouter()

	swaggerRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", configs.SwaggerPort)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	go http.ListenAndServe(fmt.Sprintf(":%d", configs.SwaggerPort), swaggerRouter)

	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	rootRouter := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	fmt.Printf("\tstarting server at %d\n", configs.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.ServerPort), rootRouter)

	return
}
