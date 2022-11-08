package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/BreCkver/Go-Investment/api"
	"github.com/gorilla/mux"
)

func Handler() {
	r := mux.NewRouter()
	credit := &api.API{}
	credit.RegisterRoutes(r)

	PORT := os.Getenv("PORTS")
	if PORT == "" {
		PORT = "8081"
	}

	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}

	log.Println("Listening...")
	log.Fatal(srv.ListenAndServe())
}
