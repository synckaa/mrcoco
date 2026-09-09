package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type KonsumenHandler struct {
	service service.KonsumenService
}

func NewKonsumenHandler(service service.KonsumenService) *KonsumenHandler {
	return &KonsumenHandler{service: service}
}

func (h *KonsumenHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToKonsumenResponseList(data), "pagination": pagination})
}

func (h *KonsumenHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	konsumen, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Konsumen not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToKonsumenResponse(konsumen)})
}

func (h *KonsumenHandler) Search(c *gin.Context) {
	pagination := helper.GetPagination(c)
	nama := c.Query("nama")
	alamat := c.Query("alamat")
	noHp := c.Query("no_hp")
	var grupID uint
	if gID := c.Query("grup_id"); gID != "" {
		parsedID, err := strconv.ParseUint(gID, 10, 32)
		if err == nil {
			grupID = uint(parsedID)
		}
	}
	var status *bool
	if s := c.Query("status"); s != "" {
		parsedStatus, err := strconv.ParseBool(s)
		if err == nil {
			status = &parsedStatus
		}
	}

	data, total, err := h.service.Search(nama, alamat, noHp, grupID, status, pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToKonsumenResponseList(data), "pagination": pagination})
}

func (h *KonsumenHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *KonsumenHandler) Create(c *gin.Context) {
	var konsumen masterdata.DataKonsumen
	if err := c.ShouldBindJSON(&konsumen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&konsumen); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.GetByID(konsumen.ID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"data": masterdata.ToKonsumenResponse(&konsumen)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": masterdata.ToKonsumenResponse(created)})
}

func (h *KonsumenHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	konsumen, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Konsumen not found"})
		return
	}

	if err := c.ShouldBindJSON(konsumen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(konsumen); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.GetByID(konsumen.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": masterdata.ToKonsumenResponse(konsumen)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToKonsumenResponse(updated)})
}

func (h *KonsumenHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Konsumen deleted successfully"})
}
