package biz

import (
	"context"
	"fmt"
	"gateway/pb"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductBiz struct {
	log        *log.Helper
	productRpc pb.ProductClient
}

func NewProductBiz(logger log.Logger, productRpc pb.ProductClient) *ProductBiz {
	return &ProductBiz{
		log:        log.NewHelper(logger),
		productRpc: productRpc,
	}
}

func (uc *ProductBiz) AddProduct(ctx context.Context, in *pb.AddSpuInfoRequest) (*pb.Reply, error) {
	// 构建rpc请求
	ret, err := uc.productRpc.AddProduct(ctx, in)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ret, nil
}
