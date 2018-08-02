package main

import (
	"context"
	"fmt"
	"net/http"
	"simple-microservices/product/rpc/product"
)

func main() {
	client := product.NewProductProtobufClient("http://localhost:8080", &http.Client{})

	saveProduct, err := client.Save(context.Background(), &product.ProductModel{ProductId: "P-100", Stock: 100, Name: "Vibranium"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Save product response : ", saveProduct)

	updateStock, err := client.UpdateStock(context.Background(), &product.UpdateStockParam{ProductId: "P-100", Amount: 5})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Update product response : ", updateStock)

	// getProduct, err := client.Get(context.Background(), &product.ParamString{Id: "P-100"})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Get product response : ", getProduct)

	// deleteProduct, err := client.Delete(context.Background(), &product.ParamString{Id: "P-100"})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Delete product response : ", deleteProduct)

}
