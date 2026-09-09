package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	pembelianmodel "mrcoco/internal/models/pembelian"
	"mrcoco/internal/service/pembelian"

	"github.com/gin-gonic/gin"
)

type PembelianHandler struct {
	service service.PembelianService
}

func NewPembelianHandler(service service.PembelianService) *PembelianHandler {
	return &PembelianHandler{service: service}
}

func (h *PembelianHandler) GetAll(c *gin.Context) {
	pagination := helper.GetPagination(c)
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	supplierID := c.Query("supplier_id")
	noTransaksi := c.Query("no_transaksi")

	data, total, err := h.service.GetAll(pagination.Offset, pagination.Limit, startDate, endDate, supplierID, noTransaksi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Total = total
	c.JSON(http.StatusOK, gin.H{"data": pembelianmodel.ToPembelianResponseList(data), "pagination": pagination})
}

func (h *PembelianHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	pembelian, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pembelian not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pembelianmodel.ToPembelianDetailResponse(pembelian)})
}

func (h *PembelianHandler) Count(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	supplierID := c.Query("supplier_id")
	noTransaksi := c.Query("no_transaksi")

	count, err := h.service.Count(startDate, endDate, supplierID, noTransaksi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count})
}

func (h *PembelianHandler) GetNextNoTransaksi(c *gin.Context) {
	tanggal := c.Query("tanggal")
	nextNo, err := h.service.GetNextNoTransaksi(tanggal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"no_transaksi": nextNo})
}

func (h *PembelianHandler) Create(c *gin.Context) {
	var pembelian pembelianmodel.DataPembelian
	if err := c.ShouldBindJSON(&pembelian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&pembelian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.GetByID(pembelian.ID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"data": pembelianmodel.ToPembelianCreateResponse(&pembelian)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": pembelianmodel.ToPembelianCreateResponse(created)})
}

func (h *PembelianHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pembelian not found"})
		return
	}

	var pembelian pembelianmodel.DataPembelian
	if err := c.ShouldBindJSON(&pembelian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(uint(id), &pembelian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": pembelianmodel.ToPembelianCreateResponse(&pembelian)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pembelianmodel.ToPembelianDetailResponse(updated)})
}

func (h *PembelianHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pembelian deleted successfully"})
}
