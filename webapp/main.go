package main

import (
	"fmt"
	"net/http"
	"simple-microservices/webapp/common"
	"simple-microservices/webapp/controllers"
)

func main() {
	common.StartUp("config/config.json")
	api := controllers.Api{}
	mux := http.NewServeMux()
	mux.Handle("/PopulateProduct", http.HandlerFunc(api.PopulateProduct))
	mux.Handle("/Order", http.HandlerFunc(api.Order))
	mux.Handle("/Save", http.HandlerFunc(api.SaveProduct))
	fmt.Println("WebApp Server is listening on port : ", common.AppConfig.ListeningPort)
	http.ListenAndServe(":"+common.AppConfig.ListeningPort, mux)

	// --- protips : using mux instead of standard http.HandleFunc, so we can easily insert middleware if needed later on

	// http.HandleFunc("/PopulateProduct", api.PopulateProduct)
	// http.ListenAndServe(":"+common.AppConfig.ListeningPort, nil)
}
