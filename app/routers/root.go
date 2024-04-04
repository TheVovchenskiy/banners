package routers

import (
	"fmt"
	"net/http"

	"github.com/TheVovchenskiy/banners/app"
	"github.com/TheVovchenskiy/banners/configs"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Run() (err error) {
	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	rootRouter := mux.NewRouter()

	fmt.Printf("\tstarting server at %d\n", configs.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.ServerPort), rootRouter)

	return
}
