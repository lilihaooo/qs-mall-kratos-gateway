package api

import (
	"gateway/internal/service"
	"gateway/pb"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}
func (p *ProductHandler) Add(c *gin.Context) {
	var req pb.AddSpuInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, nil)
	}
	res, _ := p.service.AddProduct(c, &req)
	c.JSON(200, res)
}
