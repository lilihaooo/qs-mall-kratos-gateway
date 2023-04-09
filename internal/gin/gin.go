package gin

import (
	"gateway/internal/api"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/google/wire"
)

type Gin struct {
	Router *gin.Engine
}

func NewGin(ProductHandler *api.ProductHandler, CategoryHandler *api.CategoryHandler) *Gin {
	router := gin.Default()
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery()))
	router.POST("/product", ProductHandler.Add)

	router.GET("/category", CategoryHandler.CategoryTreeList)
	return &Gin{
		Router: router,
	}
}

var ProviderSet = wire.NewSet(NewGin)
