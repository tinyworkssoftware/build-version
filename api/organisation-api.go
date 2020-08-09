package api

import (
	request "build-version/model/request"
	"build-version/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateOrganisationApiHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody request.CreateOrg
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if response, err := service.CreateOrganisation(requestBody); err != nil {
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

func CreateProjectApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := uuid.Parse(vars["orgId"]); err != nil {
		http.Error(w, "Invalid organisation-id", http.StatusBadRequest)
		return
	}
	var requestBody *request.CreateProject
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestBody.Organisation = vars["orgId"]
	if response, err := service.CreateProject(requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func GetProjectApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := uuid.Parse(vars["orgId"]); err != nil {
		http.Error(w, "Invalid organisation-id", http.StatusBadRequest)
		return
	}
	if _, err := uuid.Parse(vars["projId"]); err != nil {
		http.Error(w, "Invalid organisation-id", http.StatusBadRequest)
		return
	}

	if response, err := service.FindProject(vars["projId"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func RegenerateProjectTokenApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := uuid.Parse(vars["orgId"]); err != nil {
		http.Error(w, "Invalid organisationId", http.StatusBadRequest)
		return
	}
	if _, err := uuid.Parse(vars["projId"]); err != nil {
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
		return
	}

	if response, err := service.RegenerateProjectToken(vars["projId"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func GetAvailablePlansApiHandler(w http.ResponseWriter, r *http.Request) {
	if data, err := service.GetAllPlans(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}


