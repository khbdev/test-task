package api_server

import (
	"github.com/gin-gonic/gin"

	"test-task/internal/handler"
)

func NewRouter(productHandler *handler.ProductHandler) *gin.Engine {

	r := gin.Default()

	r.POST("/products/bulk", productHandler.BulkProductUpsert)

	return r
}
