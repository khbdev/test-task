package handler

import (
	"errors"
	"net/http"
	"test-task/internal/core"
	"test-task/internal/models/service"
	service2 "test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service2.ProductService
}

func NewProductHandler(service *service2.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}
func (h *ProductHandler) BulkProductUpsert(c *gin.Context) {

	companyID := c.GetHeader("X-Company-Id")

	if companyID == "" {
		core.Error(
			c,
			http.StatusBadRequest,
			errors.New("X-Company-Id required"),
		)
		return
	}

	var products []service.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		core.Error(
			c,
			http.StatusBadRequest,
			err,
		)
		return
	}

	err := h.service.BulkProductUpsert(
		c.Request.Context(),
		companyID,
		products,
	)

	if err != nil {
		core.Error(
			c,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	core.Success(
		c,
		http.StatusOK,
		"products upserted",
	)
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {

	companyID := c.GetHeader("X-Company-Id")

	if companyID == "" {
		core.Error(
			c,
			http.StatusBadRequest,
			errors.New("company id required"),
		)
		return
	}

	q := c.Query("q")

	products, err := h.service.SearchProducts(
		c.Request.Context(),
		companyID,
		q,
	)

	if err != nil {
		core.Error(
			c,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	core.Success(
		c,
		http.StatusOK,
		products,
	)
}

func (h *ProductHandler) DeleteProducts(c *gin.Context) {

	companyID := c.GetHeader("X-Company-Id")

	if companyID == "" {
		core.Error(
			c,
			http.StatusBadRequest,
			errors.New("company id required"),
		)
		return
	}

	var ids []string

	if err := c.ShouldBindJSON(&ids); err != nil {
		core.Error(
			c,
			http.StatusBadRequest,
			err,
		)
		return
	}

	err := h.service.DeleteProducts(
		c.Request.Context(),
		companyID,
		ids,
	)

	if err != nil {
		core.Error(
			c,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	core.Success(
		c,
		http.StatusOK,
		"products deleted",
	)
}
