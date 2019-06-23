package main

import (
	"github.com/gorilla/mux"
	"github.com/thefreakingmind12/godating/api"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(api.JwtAuthentication) //attach JWT auth middleware
	http.ListenAndServe(":8000", router) //Launch the app, visit localhost:8000/api
}
