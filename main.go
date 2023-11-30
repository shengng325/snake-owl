package main

import (
	"log"
	"net/http"

	"snake/controller"
	"snake/validator"

	"github.com/gorilla/mux"
)

func main() {
	// init controllers
	v := validator.NewValidator()
	s := controller.NewSnakeController(v)

	// init routers
	r := mux.NewRouter()
	r.HandleFunc("/new", s.NewGame).Methods("GET")
	r.HandleFunc("/validate", s.ValidateGame).Methods("POST")

	log.Println("Listening on port 8808")
	http.ListenAndServe(":8808", r)
}
