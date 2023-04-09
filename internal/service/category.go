package service

import (
	"context"
	"fmt"
	"gateway/pb/category/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type CategoryService struct {
	v1.UnimplementedCategoryServer

	log         *log.Helper
	categoryRpc v1.CategoryClient
}

func NewCategoryService(logger log.Logger, categoryRpc v1.CategoryClient) *CategoryService {
	return &CategoryService{
		log:         log.NewHelper(logger),
		categoryRpc: categoryRpc,
	}
}

func (s *CategoryService) CategoryTreeList(ctx context.Context, in *v1.CategoryListRequest) (*v1.CategoryListReply, error) {
	// 构建rpc请求
	ret, err := s.categoryRpc.CategoryTreeList(ctx, in)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ret, err
}
