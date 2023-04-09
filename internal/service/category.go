package service

import (
	"context"
	"gateway/internal/biz"
	"gateway/pb/category/v1"
)

type CategoryService struct {
	v1.UnimplementedCategoryServer
	uc *biz.CategoryBiz
}

func NewCategoryService(uc *biz.CategoryBiz) *CategoryService {
	return &CategoryService{
		uc: uc,
	}
}

func (s *CategoryService) CategoryTreeList(ctx context.Context, in *v1.CategoryListRequest) (*v1.CategoryListReply, error) {
	return s.uc.CategoryTreeList(ctx, in)
}
