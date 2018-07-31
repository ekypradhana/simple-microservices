package main

import (
	"context"
	"fmt"
	"net/http"
	"simple-microservices/product/rpc/product"
)

func main() {
	client := product.NewProductProtobufClient("http://localhost:8080", &http.Client{})

	saveProduct, err := client.Save(context.Background(), &product.ProductModel{ProductId: "P-04", Stock: 100})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(saveProduct)

	updateStock, err := client.UpdateStock(context.Background(), &product.UpdateStockParam{ProductId: "P-04", Amount: 5})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(updateStock)
}
