package handlers

import (
	"log"
	"net/http"

	"github.com/BreCkver/Go-Investment/api"
	"github.com/gorilla/mux"
)

func Handler() {
	r := mux.NewRouter()
	credit := &api.API{}
	credit.RegisterRoutes(r)

	PORT := "8081"

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}

	log.Printf("Listening... %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
