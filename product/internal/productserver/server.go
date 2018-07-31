package productserver

import (
	"context"
	"fmt"
	pb "simple-microservices/product/rpc/product"

	"gopkg.in/mgo.v2/bson"
)

type Server struct{}

const PRODUCT_COLLECTION = "products"

func (s *Server) Save(ctx context.Context, product *pb.ProductModel) (*pb.ResultInfo, error) {
	dbcontext := NewContext()
	defer dbcontext.Close()

	collection := dbcontext.DbCollection(PRODUCT_COLLECTION)

	// --- we can add validation if the productid is exist here, but for now just skip it

	err := collection.Insert(&product)
	if err != nil {
		return &pb.ResultInfo{Message: "Failed to save product!", Status: "NOK"}, err
	}

	return &pb.ResultInfo{Message: "Product has been successfully saved!", Status: "OK"}, nil
}

func (s *Server) UpdateStock(ctx context.Context, param *pb.UpdateStockParam) (*pb.ResultInfo, error) {

	// --- the process is to select the respective document first, decrease the stock with the param.Amount, then save it again

	dbcontext := NewContext()
	defer dbcontext.Close()

	collection := dbcontext.DbCollection(PRODUCT_COLLECTION)
	ProductId := param.ProductId
	var ProductModel pb.ProductModel
	err := collection.Find(bson.M{"productid": ProductId}).One(&ProductModel)

	if err != nil {
		return &pb.ResultInfo{Message: "Can't find specified product id", Status: "NOK"}, err
	}

	// --- update stock
	NewStock := ProductModel.Stock - param.Amount
	ProductModel.Stock = NewStock
	selector := bson.M{"productid": param.ProductId}
	err = collection.Update(selector, ProductModel)
	if err != nil {
		return &pb.ResultInfo{Message: "Failed to update product stock!", Status: "NOK"}, err
	}

	return &pb.ResultInfo{Message: "Product stock updated successfully!", Status: "OK"}, nil
}

func (s *Server) Get(ctx context.Context, param *pb.ParamString) (*pb.ListProductModel, error) {
	dbcontext := NewContext()
	defer dbcontext.Close()

	collection := dbcontext.DbCollection(PRODUCT_COLLECTION)
	ProductId := param.Id
	productModel := []*pb.ProductModel{}
	var err error
	if ProductId == "" {
		// populate all datas
		err = collection.Find(bson.M{}).All(&productModel)
	} else {
		// get data based on specific productid
		err = collection.Find(bson.M{"productid": ProductId}).All(&productModel)
	}
	return &pb.ListProductModel{Result: productModel}, err
}

func (s *Server) Delete(ctx context.Context, param *pb.ParamString) (*pb.ResultInfo, error) {
	if param.Id == "" {
		return &pb.ResultInfo{Message: "Failed to delete data!", Status: "NOK"}, fmt.Errorf("ID is required")
	}
	dbcontext := NewContext()
	defer dbcontext.Close()

	c := dbcontext.DbCollection(PRODUCT_COLLECTION)
	ProductId := param.Id

	err := c.Remove(bson.M{"productid": ProductId})

	if err != nil {
		return &pb.ResultInfo{Message: "Failed to delete data!", Status: "NOK"}, err
	}

	return &pb.ResultInfo{Message: "Data successfully deleted!", Status: "OK"}, nil
}
