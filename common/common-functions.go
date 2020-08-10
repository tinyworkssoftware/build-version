package common

import (
	"build-version/model/response"
	"encoding/json"
	"net/http"
)

func ErrorJsonResponse(w http.ResponseWriter, statusCode int, errorContent *response.Error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorContent)
	return
}
