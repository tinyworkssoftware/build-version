package main

import (
	"build-version/api"
	"build-version/model/toml"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

var (
	apiPrefix = "/api/v1"
	port = ":8080"
)

func main() {
	var conf model.TomlConfig
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	app := App{}
	app.Initialize()
	app.Run(port)
}

func (app *App) Initialize() {
	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("%s/healthcheck", apiPrefix), api.HealthCheckApiHandler).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("%s/session/start", apiPrefix), api.StartSessionApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/session/end", apiPrefix), api.EndSessionApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.Use(mux.CORSMethodMiddleware(router))
	app.Router = router
}

func (app *App) Run(port string) {
	if err := http.ListenAndServe(port, app.Router); err != nil {
		panic(err)
	}
}

