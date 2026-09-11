package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	preordermodel "mrcoco/internal/models/pre_order"
	"mrcoco/internal/service/pre_order"

	"github.com/gin-gonic/gin"
)

type PreOrderHandler struct {
	service service.PreOrderService
}

func NewPreOrderHandler(service service.PreOrderService) *PreOrderHandler {
	return &PreOrderHandler{service: service}
}

func (h *PreOrderHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	tanggal := c.Query("tanggal")
	konsumenID := c.Query("konsumen_id")

	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit, tanggal, konsumenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": preordermodel.ToPreOrderResponseList(data), "pagination": pagination})
}

func (h *PreOrderHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	preOrder, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pre Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": preordermodel.ToPreOrderDetailResponse(preOrder)})
}

func (h *PreOrderHandler) Count(c *gin.Context) {
	tanggal := c.Query("tanggal")
	konsumenID := c.Query("konsumen_id")

	count, err := h.service.Count(tanggal, konsumenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *PreOrderHandler) Create(c *gin.Context) {
	var preOrder preordermodel.DataPreOrder
	if err := c.ShouldBindJSON(&preOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&preOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.GetByID(preOrder.ID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"data": preordermodel.ToPreOrderCreateResponse(&preOrder)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": preordermodel.ToPreOrderCreateResponse(created)})
}

func (h *PreOrderHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pre Order not found"})
		return
	}

	var preOrder preordermodel.DataPreOrder
	if err := c.ShouldBindJSON(&preOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(uint(id), &preOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": preordermodel.ToPreOrderCreateResponse(&preOrder)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": preordermodel.ToPreOrderDetailResponse(updated)})
}

func (h *PreOrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pre Order deleted successfully"})
}
