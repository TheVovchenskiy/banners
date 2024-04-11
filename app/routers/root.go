package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TheVovchenskiy/banners/app"
	"github.com/TheVovchenskiy/banners/configs"
	_ "github.com/TheVovchenskiy/banners/docs"
	"github.com/TheVovchenskiy/banners/pkg/logging"
	"github.com/sirupsen/logrus"

	"github.com/TheVovchenskiy/banners/pkg/middleware"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() (err error) {
	logFile, err := os.OpenFile(configs.LogsDir+configs.LogfileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logging.InitLogger(logFile, configs.LogLevel)

	swaggerRouter := mux.NewRouter()

	swaggerRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", configs.SwaggerPort)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	go func () {
		err := http.ListenAndServe(fmt.Sprintf(":%d", configs.SwaggerPort), swaggerRouter)
		if err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"error": err,
			}).
				Error("error while running swagger router")
		}
	}()

	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	rootRouter := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()

	// MountAuthRouter()

	rootRouter.Use(middleware.LoggerMiddleware)
	rootRouter.Use(middleware.PanicRecoverMiddleware)

	fmt.Printf("\tstarting server at %d\n", configs.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.ServerPort), rootRouter)

	return
}
