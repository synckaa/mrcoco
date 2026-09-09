package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/service/masterdata"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service service.ItemService
}

func NewItemHandler(service service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponseList(data), "pagination": pagination})
}

func (h *ItemHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponse(item)})
}

func (h *ItemHandler) GetByKonsumenID(c *gin.Context) {
	konsumenID, err := strconv.ParseUint(c.Param("konsumen_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Konsumen ID"})
		return
	}

	pagination := helper.GetPagination(c)
	data, total, err := h.service.GetByKonsumenID(uint(konsumenID), pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponseList(data), "pagination": pagination})
}

func (h *ItemHandler) Search(c *gin.Context) {
	pagination := helper.GetPagination(c)
	kode := c.Query("kode")
	nama := c.Query("nama")
	var kategoriID uint
	if kID := c.Query("kategori_id"); kID != "" {
		parsedID, err := strconv.ParseUint(kID, 10, 32)
		if err == nil {
			kategoriID = uint(parsedID)
		}
	}
	var konsumenID uint
	if kID := c.Query("konsumen_id"); kID != "" {
		parsedID, err := strconv.ParseUint(kID, 10, 32)
		if err == nil {
			konsumenID = uint(parsedID)
		}
	}
	var status *bool
	if s := c.Query("status"); s != "" {
		parsedStatus, err := strconv.ParseBool(s)
		if err == nil {
			status = &parsedStatus
		}
	}

	data, total, err := h.service.Search(kode, nama, kategoriID, konsumenID, status, pagination.Offset, pagination.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponseList(data), "pagination": pagination})
}

func (h *ItemHandler) Count(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *ItemHandler) Create(c *gin.Context) {
	var item masterdata.DataItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	created, err := h.service.GetByID(item.ID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"data": masterdata.ToItemResponse(&item)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": masterdata.ToItemResponse(created)})
}

func (h *ItemHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.GetByID(item.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponse(item)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": masterdata.ToItemResponse(updated)})
}

func (h *ItemHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
