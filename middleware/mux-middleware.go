package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("API Request to [%v]", r.RequestURI)
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
