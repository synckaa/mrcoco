package bayar_hutang

import (
	"mrcoco/internal/models"
	"mrcoco/internal/models/pembelian"
	"time"
)

type BayarHutangResponse struct {
	ID          uint        `json:"id"`
	NoTransaksi string      `json:"no_transaksi"`
	Tanggal     models.Date `json:"tanggal"`
	SupplierID  uint        `json:"supplier_id"`
	Supplier    string      `json:"supplier"`
	GudangID    uint        `json:"gudang_id"`
	Gudang      string      `json:"gudang"`
	RekeningID  *uint       `json:"rekening_id"`
	Rekening    string      `json:"rekening"`
	GrandTotal  int         `json:"grand_total"`
	JumlahBayar int         `json:"jumlah_bayar"`
	Sisa        int         `json:"sisa"`
	Catatan     string      `json:"catatan"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type BayarHutangItemResponse struct {
	ID          uint   `json:"id"`
	PembelianID uint   `json:"pembelian_id"`
	NoTransaksi string `json:"no_transaksi"`
	JumlahBayar int    `json:"jumlah_bayar"`
}

type BayarHutangDetailResponse struct {
	BayarHutangResponse
	Items []BayarHutangItemResponse `json:"items"`
}

type BayarHutangCreateResponse struct {
	ID          uint   `json:"id"`
	NoTransaksi string `json:"no_transaksi"`
	Tanggal     string `json:"tanggal"`
	SupplierID  uint   `json:"supplier_id"`
	Supplier    string `json:"supplier"`
	GrandTotal  int    `json:"grand_total"`
	JumlahBayar int    `json:"jumlah_bayar"`
	Sisa        int    `json:"sisa"`
	Catatan     string `json:"catatan"`
}

type PembelianSisaResponse struct {
	ID             uint        `json:"id"`
	NoTransaksi    string      `json:"no_transaksi"`
	Tanggal        models.Date `json:"tanggal"`
	GrandTotal     int         `json:"grand_total"`
	JumlahUangMuka int         `json:"jumlah_uang_muka"`
	JumlahUangSisa int         `json:"jumlah_uang_sisa"`
}

func ToBayarHutangResponse(b *DataBayarHutang) BayarHutangResponse {
	resp := BayarHutangResponse{
		ID:          b.ID,
		NoTransaksi: b.NoTransaksi,
		Tanggal:     b.Tanggal,
		SupplierID:  b.SupplierID,
		GudangID:    b.GudangID,
		RekeningID:  b.RekeningID,
		GrandTotal:  b.GrandTotal,
		JumlahBayar: b.JumlahBayar,
		Sisa:        b.Sisa,
		Catatan:     b.Catatan,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
	if b.Supplier != nil {
		resp.Supplier = b.Supplier.Nama
	}
	if b.Gudang != nil {
		resp.Gudang = b.Gudang.Nama
	}
	if b.Rekening != nil {
		resp.Rekening = b.Rekening.Nama
	}
	return resp
}

func ToBayarHutangResponseList(bayarHutangs []DataBayarHutang) []BayarHutangResponse {
	result := make([]BayarHutangResponse, len(bayarHutangs))
	for i, b := range bayarHutangs {
		result[i] = ToBayarHutangResponse(&b)
	}
	return result
}

func ToBayarHutangItemResponse(item *DataBayarHutangItem) BayarHutangItemResponse {
	resp := BayarHutangItemResponse{
		ID:          item.ID,
		PembelianID: item.PembelianID,
		JumlahBayar: item.JumlahBayar,
	}
	if item.Pembelian != nil {
		resp.NoTransaksi = item.Pembelian.NoTransaksi
	}
	return resp
}

func ToBayarHutangDetailResponse(b *DataBayarHutang) BayarHutangDetailResponse {
	resp := BayarHutangDetailResponse{
		BayarHutangResponse: ToBayarHutangResponse(b),
		Items:               make([]BayarHutangItemResponse, len(b.Items)),
	}
	for i, item := range b.Items {
		resp.Items[i] = ToBayarHutangItemResponse(&item)
	}
	return resp
}

func ToBayarHutangCreateResponse(b *DataBayarHutang) BayarHutangCreateResponse {
	resp := BayarHutangCreateResponse{
		ID:          b.ID,
		NoTransaksi: b.NoTransaksi,
		Tanggal:     b.Tanggal.Format("2006-01-02"),
		SupplierID:  b.SupplierID,
		GrandTotal:  b.GrandTotal,
		JumlahBayar: b.JumlahBayar,
		Sisa:        b.Sisa,
		Catatan:     b.Catatan,
	}
	if b.Supplier != nil {
		resp.Supplier = b.Supplier.Nama
	}
	return resp
}

func ToPembelianSisaResponse(p *pembelian.DataPembelian) PembelianSisaResponse {
	return PembelianSisaResponse{
		ID:             p.ID,
		NoTransaksi:    p.NoTransaksi,
		Tanggal:        p.Tanggal,
		GrandTotal:     p.GrandTotal,
		JumlahUangMuka: p.JumlahUangMuka,
		JumlahUangSisa: p.JumlahUangSisa,
	}
}

func ToPembelianSisaResponseList(pembelians []pembelian.DataPembelian) []PembelianSisaResponse {
	result := make([]PembelianSisaResponse, len(pembelians))
	for i, p := range pembelians {
		result[i] = ToPembelianSisaResponse(&p)
	}
	return result
}
