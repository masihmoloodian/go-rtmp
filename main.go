package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()

	router := mux.NewRouter().StrictSlash(true)

	log.Println(fmt.Sprintf("Server listen on port: %v", AppConfig.ServicePort))
	http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.ServicePort), router)
}

func RegisterRoutes(router *mux.Router) {
	
}