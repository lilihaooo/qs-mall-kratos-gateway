package service

import (
	"context"
	"gateway/internal/biz"
	"gateway/pb"
)

type ProductService struct {
	pb.UnimplementedProductServer
	uc *biz.ProductBiz
}

func NewProductService(uc *biz.ProductBiz) *ProductService {
	return &ProductService{
		uc: uc,
	}
}

func (s *ProductService) AddProduct(ctx context.Context, in *pb.AddSpuInfoRequest) (*pb.Reply, error) {
	return s.uc.AddProduct(ctx, in)
}
