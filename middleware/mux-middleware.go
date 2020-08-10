package middleware

import (
	"build-version/common"
	"build-version/config"
	"build-version/model/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("API Request to [%v]", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
func VerifyAuthorizationTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Verifying token...\n")
		if isBypassApi(r.RequestURI) {
			vars := mux.Vars(r)
			if len(vars["access_token"]) == 0 {
				common.ErrorJsonResponse(w, http.StatusUnauthorized, &response.Error{
					ErrorMessage:  "Required param missing [access_token]",
					CorrelationId: uuid.New().String(),
					TransactionTs: time.Now(),
				})
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func AddJsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, "api") {
			log.Println("Detected API endpoint, adding global json content-type.")
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func isBypassApi(uri string) bool {
	for _, bypassApi := range config.GetAppConfig().Application.ApiBypass {
		if uri == bypassApi {
			log.Debugf("Bypass API [%v] detected. This doesn't mean that the api is publicly accessible, just accessible with access_token", uri)
			return true
		}
	}
	return false
}
