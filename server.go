package main

import (
	"build-version/api"
	"fmt"
	"github.com/gorilla/mux"
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
	app := App{}
	app.Initialize()
	app.Run(port)
}

func (app *App) Initialize() {
	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("%s/healthcheck", apiPrefix), api.HealthCheckApiHandler).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("%s/session/start", apiPrefix), api.StartSessionApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/session/end", apiPrefix), api.EndSessionApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/plan", apiPrefix), api.GetAvailablePlansApiHandler).Methods(http.MethodGet)
	router.Use(mux.CORSMethodMiddleware(router))
	app.Router = router
}

func (app *App) Run(port string) {
	if err := http.ListenAndServe(port, app.Router); err != nil {
		panic(err)
	}
}

