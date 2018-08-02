package transactionserver

import (
	"context"
	"net/http"
	"simple-microservices/transaction/internal/common"
	productPb "simple-microservices/transaction/rpc/product"
	trxPb "simple-microservices/transaction/rpc/transaction"
)

type Server struct{}

const TRX_COLLECTION = "transactions"

func (s *Server) Order(ctx context.Context, trx *trxPb.TransactionModel) (*trxPb.ResultInfo, error) {
	dbcontext := NewContext()
	defer dbcontext.Close()

	collection := dbcontext.DbCollection(TRX_COLLECTION)

	err := collection.Insert(&trx)
	if err != nil {
		return &trxPb.ResultInfo{Message: "Failed to save transaction!", Status: "NOK"}, err
	}

	// --- Update stock product, hit the Product Service
	productClient := productPb.NewProductProtobufClient(common.AppConfig.ProductService, &http.Client{})
	_, e := productClient.UpdateStock(context.Background(), &productPb.UpdateStockParam{ProductId: trx.ProductId, Amount: trx.OrderAmount})

	if e != nil {
		return &trxPb.ResultInfo{Message: "Transaction has been successfully saved but Product stock is failed to update!", Status: "OK"}, nil
	}

	return &trxPb.ResultInfo{Message: "Transaction has been successfully saved!", Status: "OK"}, nil
}
