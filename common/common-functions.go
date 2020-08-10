package common

import (
	"build-version/model/response"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ErrorJsonResponse(w http.ResponseWriter, statusCode int, errorContent *response.Error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorContent)
	return
}

func ValidateEmptyStrings(values []string) error {
	invalids := make([]string, 0)

	for _ , target := range values {
		if len(target) == 0 {
			invalids = append(invalids, target)
		}
	}

	if len(invalids) > 0 {
		return errors.New(fmt.Sprintf("Empty Strings detected [%v]", invalids))
	}
	return nil
}
