package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-microservices/webapp/common"
	ProductPB "simple-microservices/webapp/rpc/product"
	TrxPB "simple-microservices/webapp/rpc/transaction"
)

type Api struct{}

func (c *Api) PopulateProduct(w http.ResponseWriter, r *http.Request) {

	client := ProductPB.NewProductProtobufClient(common.AppConfig.ProductService, &http.Client{})
	ret, err := client.Get(context.Background(), &ProductPB.ParamString{Id: ""})
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
	}
	jsonResponse(w, http.StatusOK, ret)
}

func (c *Api) SaveProduct(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		ProductId string
		Stock     int32
		Name      string
	}{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Unmarshal
	err = json.Unmarshal(b, &payload)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	client := ProductPB.NewProductProtobufClient(common.AppConfig.ProductService, &http.Client{})
	ret, err := client.Save(context.Background(), &ProductPB.ProductModel{ProductId: payload.ProductId, Name: payload.Name, Stock: payload.Stock})
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
	}
	jsonResponse(w, http.StatusOK, ret)
}

func (c *Api) Order(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		ProductId   string
		OrderAmount int32
	}{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Unmarshal
	err = json.Unmarshal(b, &payload)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	client := TrxPB.NewTransactionProtobufClient(common.AppConfig.TransactionService, &http.Client{})
	ret, err := client.Order(context.Background(), &TrxPB.TransactionModel{ProductId: payload.ProductId, OrderAmount: payload.OrderAmount})
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
	}
	jsonResponse(w, http.StatusOK, ret)
}

func jsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res, _ := json.Marshal(data)
	fmt.Fprint(w, string(res))
}
