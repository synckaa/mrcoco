package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type SubKlasifikasiHandler struct {
	service service.SubKlasifikasiService
}

func NewSubKlasifikasiHandler(service service.SubKlasifikasiService) *SubKlasifikasiHandler {
	return &SubKlasifikasiHandler{service: service}
}

func (h *SubKlasifikasiHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *SubKlasifikasiHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sub, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub Klasifikasi not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sub})
}

func (h *SubKlasifikasiHandler) Search(c *gin.Context) {
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

func (h *SubKlasifikasiHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *SubKlasifikasiHandler) Create(c *gin.Context) {
	var sub masterdata.DataSubKlasifikasi
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": sub})
}

func (h *SubKlasifikasiHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sub, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sub Klasifikasi not found"})
		return
	}

	if err := c.ShouldBindJSON(sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sub})
}

func (h *SubKlasifikasiHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sub Klasifikasi deleted successfully"})
}
