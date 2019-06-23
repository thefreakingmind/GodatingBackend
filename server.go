package main

import (
	"github.com/gorilla/mux"
	"github.com/thefreakingmind12/godating/api"
	"os"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(api.JwtAuthentication) //attach JWT auth middleware

	err := http.ListenAndServe(":8000", router) //Launch the app, visit localhost:8000/api
}
