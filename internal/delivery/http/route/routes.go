package route

import (
	"DRX_Test/internal/provider"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handlers *provider.Handlers) {
	r.NoRoute(handlers.Root.RouteNotFound)
	r.GET("/", handlers.Root.Index)
	r.GET("/v1/products", handlers.Product.GetAllProducts)
	r.POST("/v1/products", handlers.Product.CreateProduct)
}
