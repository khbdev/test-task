package handler

import (
	"errors"
	"net/http"
	"strconv"
	"test-task/internal/models/repostory"

	"test-task/internal/core"
	service2 "test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type SlotHandler struct {
	service *service2.SlotService
}

func NewSlotHandler(
	service *service2.SlotService,
) *SlotHandler {

	return &SlotHandler{
		service: service,
	}
}

type UpdateSlotRequest struct {
	Slot      int    `json:"slot"`
	ProductID string `json:"product_id"`
}

func (h *SlotHandler) UpdateSlots(c *gin.Context) {

	companyID := c.GetHeader("X-Company-Id")

	if companyID == "" {

		core.Error(
			c,
			http.StatusBadRequest,
			errors.New("company id required"),
		)

		return
	}

	var body []UpdateSlotRequest

	if err := c.ShouldBindJSON(&body); err != nil {

		core.Error(
			c,
			http.StatusBadRequest,
			err,
		)

		return
	}

	slots := make([]repostory.ShelfSlot, 0, len(body))

	for _, item := range body {

		slots = append(
			slots,
			repostory.ShelfSlot{
				Slot:      item.Slot,
				ProductID: &item.ProductID,
			},
		)
	}
	err := h.service.UpdateSlots(
		c.Request.Context(),
		companyID,
		slots,
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
		"slots updated",
	)
}

func (h *SlotHandler) GetSlots(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	slots, total, err := h.service.GetSlots(
		c,
		c.GetHeader("X-Company-Id"),
		page,
		limit,
		c.Query("search"),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data":  slots,
		"total": total,
	})
}
