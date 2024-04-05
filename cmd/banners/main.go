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
	err := routers.Run()
	if err != nil {
		fmt.Println(err)
	}
}
