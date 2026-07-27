package api_server

import (
	"github.com/gin-gonic/gin"

	"test-task/internal/handler"
)

func NewRouter(productHandler *handler.ProductHandler, slotHandler *handler.SlotHandler) *gin.Engine {

	r := gin.Default()

	r.POST("/products/bulk", productHandler.BulkProductUpsert)
	r.GET("/products/search",
		productHandler.SearchProducts,
	)
	r.DELETE(
		"/products",
		productHandler.DeleteProducts,
	)
	r.PUT("/slots", slotHandler.UpdateSlots)

	r.GET(
		"/slots",
		slotHandler.GetSlots,
	)
	r.GET(
		"/reports/stock-value",
		slotHandler.StockValue,
	)
	return r
}
