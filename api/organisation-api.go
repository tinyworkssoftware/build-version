package api

import (
	request2 "build-version/model/request"
	"build-version/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateOrganisationApiHandler(w http.ResponseWriter, r *http.Request) {
	var request request2.CreateOrg
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if response, err := service.CreateOrganisation(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func GetOrganisationApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, err := uuid.Parse(vars["orgId"]); err != nil {
		http.Error(w, "Invalid organisation-id", http.StatusBadRequest)
		return
	}

	if record, err := service.FindOrganisation(vars["orgId"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(record)
		return
	}
}
