package main

import (
	"github.com/gorilla/mux"
	"github.com/thefreakingmind12/godating/api"
	"net/http"
	"github.com/thefreakingmind12/godating/controller"
)

func main() {

	router := mux.NewRouter()
	router.Use(api.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/api/user/new", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.Authenticate).Methods("POST")
	http.ListenAndServe(":8000", router) //Launch the app, visit localhost:8000/api
}
