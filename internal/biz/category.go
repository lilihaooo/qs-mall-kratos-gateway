package biz

import (
	"context"
	"fmt"
	"gateway/pb/category/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type CategoryBiz struct {
	log         *log.Helper
	CategoryRpc v1.CategoryClient
}

func NewCategoryBiz(logger log.Logger, categoryRpc v1.CategoryClient) *CategoryBiz {
	return &CategoryBiz{
		log:         log.NewHelper(logger),
		CategoryRpc: categoryRpc,
	}
}

func (uc *CategoryBiz) CategoryTreeList(ctx context.Context, in *v1.CategoryListRequest) (*v1.CategoryListReply, error) {
	// 构建rpc请求
	ret, err := uc.CategoryRpc.CategoryTreeList(ctx, in)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ret, err
}
