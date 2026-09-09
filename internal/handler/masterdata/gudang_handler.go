package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type GudangHandler struct {
	service service.GudangService
}

func NewGudangHandler(service service.GudangService) *GudangHandler {
	return &GudangHandler{service: service}
}

func (h *GudangHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *GudangHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	gudang, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gudang not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gudang})
}

func (h *GudangHandler) Search(c *gin.Context) {
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

func (h *GudangHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *GudangHandler) Create(c *gin.Context) {
	var gudang masterdata.DataGudang
	if err := c.ShouldBindJSON(&gudang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&gudang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": gudang})
}

func (h *GudangHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	gudang, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gudang not found"})
		return
	}

	if err := c.ShouldBindJSON(gudang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(gudang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gudang})
}

func (h *GudangHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gudang deleted successfully"})
}
