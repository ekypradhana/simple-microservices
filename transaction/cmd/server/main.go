package main

import (
	"fmt"
	"net/http"
	"simple-microservices/transaction/internal/common"
	"simple-microservices/transaction/internal/transactionserver"
	"simple-microservices/transaction/rpc/transaction"
)

func main() {
	common.StartUp("config/config.json")
	server := &transactionserver.Server{}
	twirpHandler := transaction.NewTransactionServer(server, nil)
	fmt.Println("transaction RPC Server is listening on port : ", common.AppConfig.ServerPort)
	http.ListenAndServe(":"+common.AppConfig.ServerPort, twirpHandler)
}
