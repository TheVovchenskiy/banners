package main

import (
	"fmt"
	// "net/http"

	"github.com/TheVovchenskiy/banners/app/routers"
	// "github.com/gorilla/mux"
)

// @title			Banners API
// @version		1.0
// @description	This is the first version of Banners API
//
// @contact.name	Vladimir Konoplyuk
// @contact.url	https://github.com/TheVovchenskiy
// @contact.email	konoplyuk.v@yandex.ru
//
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	// r := mux.NewRouter()

	// r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	// 	httpSwagger.DeepLinking(true),
	// 	httpSwagger.DocExpansion("none"),
	// 	httpSwagger.DomID("swagger-ui"),
	// )).Methods(http.MethodGet)

	// go http.ListenAndServe(":1323", r)

	err := routers.Run()
	if err != nil {
		fmt.Println(err)
	}
}
