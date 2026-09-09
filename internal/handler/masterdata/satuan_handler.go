package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type SatuanHandler struct {
	service service.SatuanService
}

func NewSatuanHandler(service service.SatuanService) *SatuanHandler {
	return &SatuanHandler{service: service}
}

func (h *SatuanHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *SatuanHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	satuan, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satuan not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": satuan})
}

func (h *SatuanHandler) Search(c *gin.Context) {
	pagination := helper.GetPagination(c)
	nama := c.Query("nama")
	data, total, err := h.service.Search(nama, pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *SatuanHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *SatuanHandler) Create(c *gin.Context) {
	var satuan masterdata.DataSatuan
	if err := c.ShouldBindJSON(&satuan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&satuan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": satuan})
}

func (h *SatuanHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	satuan, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Satuan not found"})
		return
	}

	if err := c.ShouldBindJSON(satuan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(satuan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": satuan})
}

func (h *SatuanHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Satuan deleted successfully"})
}
