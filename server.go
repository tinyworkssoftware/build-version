package main

import (
	"build-version/api"
	"build-version/config"
	"build-version/middleware"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

var (
	apiPrefix = "/api"
	port = ":8080"
)

func main() {
	initializeLogger()
	config.SetConfigAsEnvironmentVariables()
	app := App{}
	app.InitializeRoutes()
	app.InitializeMiddleware()
	app.Run(port)
}

func initializeLogger() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	logLevel := config.GetAppConfig().Application.LogLevel
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if logLevel == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if logLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

}

func (app *App) InitializeRoutes() {
	apiPrefix = fmt.Sprintf("%s/%s", apiPrefix, os.Getenv("API_VERSION"))
	router := mux.NewRouter()
	//NOTE: Admin APIs
	router.HandleFunc(fmt.Sprintf("%s/healthcheck", apiPrefix), api.HealthCheckApiHandler).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("%s/admin/sessions/active", apiPrefix), api.GetAllActiveSessionsApiHandler).Methods(http.MethodGet)

	//NOTE: Session APIs
	router.HandleFunc(fmt.Sprintf("%s/session/start", apiPrefix), api.StartSessionApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/session/end", apiPrefix), api.EndSessionApiHandler).Methods(http.MethodPost, http.MethodPut)

	//NOTE: Organisation & Project APIs
	router.HandleFunc(fmt.Sprintf("%s/organisation", apiPrefix), api.CreateOrganisationApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/organisation/{orgId}", apiPrefix), api.GetOrganisationApiHandler).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("%s/organisation/{orgId}/project", apiPrefix), api.CreateProjectApiHandler).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc(fmt.Sprintf("%s/organisation/{orgId}/project/{projId}", apiPrefix), api.GetProjectApiHandler).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("%s/organisation/{orgId}/project/{projId}/token", apiPrefix), api.RegenerateProjectTokenApiHandler).Methods(http.MethodPut)

	//NOTE: Plan APIs
	router.HandleFunc(fmt.Sprintf("%s/plans", apiPrefix), api.GetAvailablePlansApiHandler).Methods(http.MethodGet)

	app.Router = router
}

func (app *App) InitializeMiddleware() {
	app.Router.Use(mux.CORSMethodMiddleware(app.Router))
	app.Router.Use(middleware.LoggingMiddleware)
	app.Router.Use(middleware.AddJsonContentTypeMiddleware)
	app.Router.Use(middleware.VerifyAuthorizationTokenMiddleware)
}

func (app *App) Run(port string) {
	if err := http.ListenAndServe(port, app.Router); err != nil {
		panic(err)
	}
}


