package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (credit *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/credit-assignment", credit.creditAssignment).Methods(http.MethodPost)
	r.HandleFunc("/statistics", credit.statistics).Methods(http.MethodPost)
}
