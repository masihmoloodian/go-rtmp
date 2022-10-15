package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masihmoloodian/go-rtmp/controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	RegisterRoutes(router)

	log.Println(fmt.Sprintf("Server listen on port: %v", 8000))
	http.ListenAndServe(fmt.Sprintf(":%v", 8000), router)
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateOrUpdate).Methods("POST")
}