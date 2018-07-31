package productserver

import (
	"testing"

	"context"
	"simple-microservices/product/internal/common"
	pb "simple-microservices/product/rpc/product"
)

func TestStartUp(t *testing.T) {
	common.StartUp("../../cmd/server/config.json")
}

func TestSave(t *testing.T) {
	s := Server{}
	_, err := s.Save(context.Background(), &pb.ProductModel{ProductId: "P-99", Stock: 100})
	if err != nil {
		t.Errorf("[error]: %s\n", err)
	}
}

func TestGet(t *testing.T) {
	s := Server{}
	_, err := s.Save(context.Background(), &pb.ProductModel{ProductId: "P-99", Stock: 100})
	if err != nil {
		t.Errorf("[error]: %s\n", err)
	}
}
