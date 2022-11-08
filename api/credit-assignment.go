package api

import (
	"encoding/json"
	"net/http"

	"github.com/BreCkver/Go-Investment/data"
	"github.com/BreCkver/Go-Investment/models"
	"github.com/BreCkver/Go-Investment/services"
)

type API struct {
}

func (api *API) creditAssignment(w http.ResponseWriter, r *http.Request) {

	var request models.CreditAssignmentParams
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Input received is wrong", 400)
		return
	}

	var data data.DBDataAccess = data.NewCreditAssignerData()
	service := services.NewCreditAssignmentService(data)
	typeCount300, typeCount500, typeCount700, err := service.Assign(request.Investment)
	result := models.NewCreditAssignment(typeCount300, typeCount500, typeCount700)

	w.Header().Set("context-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(result)

	service.SaveStatistics(request.Investment, err == nil)
}

func (api *API) statistics(w http.ResponseWriter, r *http.Request) {
	var data data.DBDataAccess = data.NewCreditAssignerData()
	service := services.NewCreditAssignmentService(data)
	summary, err := service.GetStatistics()
	w.Header().Set("context-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(summary)
}
