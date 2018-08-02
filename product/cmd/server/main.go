package main

import (
	"fmt"
	"net/http"
	"simple-microservices/product/internal/common"
	"simple-microservices/product/internal/productserver"
	"simple-microservices/product/rpc/product"
)

func main() {
	common.StartUp("config/config.json")
	server := &productserver.Server{}
	twirpHandler := product.NewProductServer(server, nil)
	fmt.Println("Product RPC Server is listening on port : ", common.AppConfig.ServerPort)
	http.ListenAndServe(":"+common.AppConfig.ServerPort, twirpHandler)
}
