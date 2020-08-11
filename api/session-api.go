package api

import (
	"build-version/common"
	"build-version/model/dto"
	"build-version/model/request"
	"build-version/model/response"
	"build-version/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func StartSessionApiHandler(w http.ResponseWriter, r *http.Request) {
	correlationId := r.Header.Get("Correlation-Id")
	params := r.URL.Query()
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

}

func EndSessionApiHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("access_token")
	correlationId := r.Header.Get("Correlation-Id")
	sessionId := mux.Vars(r)["sessionId"]
	if err := common.ValidateEmptyStrings([]string{accessToken, sessionId}); err != nil {
		common.ErrorJsonResponse(w, http.StatusBadRequest, &response.Error{
			ErrorMessage:  err.Error(),
			CorrelationId: correlationId,
			TransactionTs: time.Now(),
		})
		return
	}
	requestData := &dto.UpdateSessionDTO{
		CorrelationId: correlationId,
		SessionId:     sessionId,
		AccessToken:   accessToken,
	}
	if err := service.EndSession(requestData); err != nil {
		common.ErrorJsonResponse(w, http.StatusBadRequest, &response.Error{
			ErrorMessage:  err.Error(),
			CorrelationId: correlationId,
			TransactionTs: time.Now(),
		})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

