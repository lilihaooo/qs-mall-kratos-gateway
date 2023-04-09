package api

import (
	"fmt"
	"gateway/internal/service"
	v1 "gateway/pb/category/v1"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}
func (p *CategoryHandler) CategoryTreeList(c *gin.Context) {
	req := v1.CategoryListRequest{}
	res, err := p.service.CategoryTreeList(c, &req)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, res)
}
