package api

import (
	"build-version/model/response"
	"build-version/service"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func GetAllActiveSessionsApiHandler(w http.ResponseWriter, r *http.Request) {
	transactionTs := time.Now()
	if results, err := service.GetActiveSessions(); err != nil {
		errorJsonResponse(w, http.StatusInternalServerError,
			&response.Error{
				ErrorMessage:  err.Error(),
				CorrelationId: uuid.New().String(),
				TransactionTs: transactionTs,
			},
		)
		return
	} else {
		json.NewEncoder(w).Encode(results)
	}
}