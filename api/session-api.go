package api

import (
	"build-version/common"
	"build-version/model/request"
	"build-version/model/response"
	"build-version/service"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func StartSessionApiHandler(w http.ResponseWriter, r *http.Request) {
	correlationId := uuid.New().String()
	params := r.URL.Query()
	if len(params.Get("access_token")) > 0 {
		var requestData request.CreateSession
		json.NewDecoder(r.Body).Decode(&requestData)
		requestData.AccessToken = params.Get("access_token")
		if res, err := service.StartSession(&requestData); err != nil {
			common.ErrorJsonResponse(w, http.StatusInternalServerError, &response.Error{
				ErrorMessage: err.Error(),
				CorrelationId: correlationId,
				TransactionTs: time.Time{},
			})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
		}
	} else {
		common.ErrorJsonResponse(w, http.StatusBadRequest, &response.Error{
			ErrorMessage: "Required params [access_token] not found.",
			CorrelationId: correlationId,
		})
	}
}

func EndSessionApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

