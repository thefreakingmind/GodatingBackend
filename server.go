package main

import (
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
)

func main(){
  router := mux.NewRouter()
  router.Use(api.JwtAuthentication)

  err := http.ListenAndServe(":8080", router)
  if err!= nil{
	fmt.Println(err)
  }
}
