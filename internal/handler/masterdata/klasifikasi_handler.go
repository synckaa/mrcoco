package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type KlasifikasiHandler struct {
	service service.KlasifikasiService
}

func NewKlasifikasiHandler(service service.KlasifikasiService) *KlasifikasiHandler {
	return &KlasifikasiHandler{service: service}
}

func (h *KlasifikasiHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *KlasifikasiHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	klasifikasi, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Klasifikasi not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": klasifikasi})
}

func (h *KlasifikasiHandler) Search(c *gin.Context) {
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

func (h *KlasifikasiHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *KlasifikasiHandler) Create(c *gin.Context) {
	var klasifikasi masterdata.DataKlasifikasi
	if err := c.ShouldBindJSON(&klasifikasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&klasifikasi); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": klasifikasi})
}

func (h *KlasifikasiHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	klasifikasi, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Klasifikasi not found"})
		return
	}

	if err := c.ShouldBindJSON(klasifikasi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(klasifikasi); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": klasifikasi})
}

func (h *KlasifikasiHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Klasifikasi deleted successfully"})
}
