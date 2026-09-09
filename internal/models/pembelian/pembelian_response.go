package pembelian

import (
	"mrcoco/internal/models"
	"time"
)

type PembelianResponse struct {
	ID             uint      `json:"id"`
	NoTransaksi    string    `json:"no_transaksi"`
	Tanggal        models.Date      `json:"tanggal"`
	SupplierID     uint      `json:"supplier_id"`
	Supplier       string    `json:"supplier"`
	GudangID       uint      `json:"gudang_id"`
	Gudang         string    `json:"gudang"`
	TglJatuhTempo  models.Date      `json:"tgl_jatuh_tempo"`
	RekeningID     *uint     `json:"rekening_id"`
	Rekening       string    `json:"rekening"`
	TotalQty       int       `json:"total_qty"`
	GrandTotal     int       `json:"grand_total"`
	JumlahUangMuka int       `json:"jumlah_uang_muka"`
	JumlahUangSisa int       `json:"jumlah_uang_sisa"`
	Catatan        string    `json:"catatan"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PembelianItemResponse struct {
	ID       uint   `json:"id"`
	ItemID   uint   `json:"item_id"`
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Satuan   string `json:"satuan"`
	Qty      int    `json:"qty"`
	Harga    int    `json:"harga"`
	Diskon   int    `json:"diskon"`
	Subtotal int    `json:"subtotal"`
}

type PembelianDetailResponse struct {
	PembelianResponse
	Items []PembelianItemResponse `json:"items"`
}

type PembelianCreateResponse struct {
	ID             uint   `json:"id"`
	NoTransaksi    string `json:"no_transaksi"`
	Tanggal        string `json:"tanggal"`
	GrandTotal     int    `json:"grand_total"`
	JumlahUangMuka int    `json:"jumlah_uang_muka"`
	JumlahUangSisa int    `json:"jumlah_uang_sisa"`
	Catatan        string `json:"catatan"`
}

func ToPembelianResponse(p *DataPembelian) PembelianResponse {
	resp := PembelianResponse{
		ID:             p.ID,
		NoTransaksi:    p.NoTransaksi,
		Tanggal:        p.Tanggal,
		SupplierID:     p.SupplierID,
		GudangID:       p.GudangID,
		TglJatuhTempo:  p.TglJatuhTempo,
		RekeningID:     p.RekeningID,
		GrandTotal:     p.GrandTotal,
		JumlahUangMuka: p.JumlahUangMuka,
		JumlahUangSisa: p.JumlahUangSisa,
		Catatan:        p.Catatan,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
	if p.Supplier != nil {
		resp.Supplier = p.Supplier.Nama
	}
	if p.Gudang != nil {
		resp.Gudang = p.Gudang.Nama
	}
	if p.Rekening != nil {
		resp.Rekening = p.Rekening.Nama
	}
	for _, item := range p.Items {
		resp.TotalQty += item.Qty
	}
	return resp
}

func ToPembelianResponseList(pembelians []DataPembelian) []PembelianResponse {
	result := make([]PembelianResponse, len(pembelians))
	for i, p := range pembelians {
		result[i] = ToPembelianResponse(&p)
	}
	return result
}

func ToPembelianItemResponse(item *DataPembelianItem) PembelianItemResponse {
	resp := PembelianItemResponse{
		ID:       item.ID,
		ItemID:   item.ItemID,
		Qty:      item.Qty,
		Harga:    item.Harga,
		Diskon:   item.Diskon,
		Subtotal: item.Subtotal,
	}
	if item.Item != nil {
		resp.Kode = item.Item.Kode
		resp.Nama = item.Item.Nama
		if item.Item.Satuan != nil {
			resp.Satuan = item.Item.Satuan.Nama
		}
	}
	return resp
}

func ToPembelianDetailResponse(p *DataPembelian) PembelianDetailResponse {
	resp := PembelianDetailResponse{
		PembelianResponse: ToPembelianResponse(p),
		Items:             make([]PembelianItemResponse, len(p.Items)),
	}
	for i, item := range p.Items {
		resp.Items[i] = ToPembelianItemResponse(&item)
	}
	return resp
}

func ToPembelianCreateResponse(p *DataPembelian) PembelianCreateResponse {
	return PembelianCreateResponse{
		ID:             p.ID,
		NoTransaksi:    p.NoTransaksi,
		Tanggal:        p.Tanggal.Format("2006-01-02"),
		GrandTotal:     p.GrandTotal,
		JumlahUangMuka: p.JumlahUangMuka,
		JumlahUangSisa: p.JumlahUangSisa,
		Catatan:        p.Catatan,
	}
}
