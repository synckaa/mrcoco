package handler

import (
	"net/http"
	"strconv"

	"mrcoco/internal/helper"
	bayarhutangmodel "mrcoco/internal/models/bayar_hutang"
	"mrcoco/internal/service/bayar_hutang"

	"github.com/gin-gonic/gin"
)

type BayarHutangHandler struct {
	service service.BayarHutangService
}

func NewBayarHutangHandler(service service.BayarHutangService) *BayarHutangHandler {
	return &BayarHutangHandler{service: service}
}

func (h *BayarHutangHandler) GetAll(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToBayarHutangResponseList(data), "pagination": pagination})
}

func (h *BayarHutangHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	bayarHutang, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bayar hutang not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToBayarHutangDetailResponse(bayarHutang)})
}

func (h *BayarHutangHandler) Count(c *gin.Context) {
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

func (h *BayarHutangHandler) GetNextNoTransaksi(c *gin.Context) {
	tanggal := c.Query("tanggal")
	nextNo, err := h.service.GetNextNoTransaksi(tanggal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"no_transaksi": nextNo})
}

func (h *BayarHutangHandler) Create(c *gin.Context) {
	var bayarHutang bayarhutangmodel.DataBayarHutang
	if err := c.ShouldBindJSON(&bayarHutang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&bayarHutang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.GetByID(bayarHutang.ID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"data": bayarhutangmodel.ToBayarHutangCreateResponse(&bayarHutang)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": bayarhutangmodel.ToBayarHutangCreateResponse(created)})
}

func (h *BayarHutangHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bayar hutang not found"})
		return
	}

	var bayarHutang bayarhutangmodel.DataBayarHutang
	if err := c.ShouldBindJSON(&bayarHutang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(uint(id), &bayarHutang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToBayarHutangCreateResponse(&bayarHutang)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToBayarHutangDetailResponse(updated)})
}

func (h *BayarHutangHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bayar hutang deleted successfully"})
}

func (h *BayarHutangHandler) GetPembelianSisa(c *gin.Context) {
	supplierID, err := strconv.ParseUint(c.Param("supplier_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supplier ID"})
		return
	}

	search := c.Query("search")
	if search != "" {
		data, err := h.service.SearchPembelianSisaBySupplier(uint(supplierID), search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToPembelianSisaResponseList(data)})
	} else {
		data, err := h.service.GetPembelianSisaBySupplier(uint(supplierID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": bayarhutangmodel.ToPembelianSisaResponseList(data)})
	}
}
