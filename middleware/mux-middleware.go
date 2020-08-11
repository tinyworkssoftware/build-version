package middleware

import (
	"build-version/common"
	"build-version/config"
	"build-version/model/response"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func AssignCorrelationId(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if corrId := r.Header.Get("Correlation-Id"); len(corrId) > 0 {
			if _, err := uuid.Parse(corrId); err != nil {
				log.Debugf("Invalid Correlation Id. Generating a new ID for this request.")
			} else {
				log.Debugf("Request already has correlation id [%v]. This is probably a upstream system api call\n", corrId)
				next.ServeHTTP(w, r)
				return
			}
		}
		corrId := uuid.New().String()
		log.Debugf("Setting Correlation Id for request [%v]\n", corrId)
		r.Header.Set("Correlation-Id", corrId)
		next.ServeHTTP(w, r)
		return
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("API Request to [%v]\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
func VerifyAuthorizationTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Verifying token...\n")
		if isBypassApi(r.RequestURI) {
			if len(r.URL.Query().Get("access_token")) == 0 {
				common.ErrorJsonResponse(w, http.StatusUnauthorized, &response.Error{
					ErrorMessage:  "Required param missing [access_token]",
					CorrelationId: r.Header.Get("Correlation-Id"),
					TransactionTs: time.Now(),
				})
				return
			}
		} else {
			if authHeader := r.Header.Get("Authorization"); len(authHeader) > 0 {
				//TODO: implement token code check here when its up.
			} else {
				common.ErrorJsonResponse(w, http.StatusUnauthorized, &response.Error{
					ErrorMessage: "Required header missing [Authorization]",
					CorrelationId: r.Header.Get("Correlation-Id"),
					TransactionTs: time.Time{},
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
			log.Debugln("Detected API endpoint, adding global json content-type.")
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func isBypassApi(uri string) bool {
	uri = strings.Split(uri, "?")[0]
	for _, bypassApi := range config.GetAppConfig().Application.ApiBypass {
		if strings.Contains(uri, bypassApi) {
			log.Debugf("Bypass API [%v] detected. This doesn't mean that the api is publicly accessible, just accessible with access_token", uri)
			return true
		}
	}
	return false
}
