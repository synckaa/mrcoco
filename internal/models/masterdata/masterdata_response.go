package masterdata

import "time"

type DataItemResponse struct {
	ID         uint      `json:"id"`
	Kode       string    `json:"kode"`
	Nama       string    `json:"nama"`
	JenisID    uint      `json:"jenis_id"`
	Jenis      string    `json:"jenis"`
	KategoriID uint      `json:"kategori_id"`
	Kategori   string    `json:"kategori"`
	SatuanID   uint      `json:"satuan_id"`
	Satuan     string    `json:"satuan"`
	RekeningID *uint     `json:"rekening_id"`
	Rekening   string    `json:"rekening"`
	KonsumenID *uint     `json:"konsumen_id"`
	Konsumen   string    `json:"konsumen"`
	HargaBeli  int       `json:"harga_beli"`
	HargaJual  int       `json:"harga_jual"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DataKonsumenResponse struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	NoHp      string    `json:"no_hp"`
	Email     string    `json:"email"`
	Alamat    string    `json:"alamat"`
	GrupID    uint      `json:"grup_id"`
	Grup      string    `json:"grup"`
	Catatan   string    `json:"catatan"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToItemResponse(item *DataItem) DataItemResponse {
	resp := DataItemResponse{
		ID:         item.ID,
		Kode:       item.Kode,
		Nama:       item.Nama,
		JenisID:    item.JenisID,
		KategoriID: item.KategoriID,
		SatuanID:   item.SatuanID,
		RekeningID: item.RekeningID,
		KonsumenID: item.KonsumenID,
		HargaBeli:  item.HargaBeli,
		HargaJual:  item.HargaJual,
		Status:     item.Status,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
	if item.Jenis != nil {
		resp.Jenis = item.Jenis.Nama
	}
	if item.Kategori != nil {
		resp.Kategori = item.Kategori.Nama
	}
	if item.Satuan != nil {
		resp.Satuan = item.Satuan.Nama
	}
	if item.Rekening != nil {
		resp.Rekening = item.Rekening.Nama
	}
	if item.Konsumen != nil {
		resp.Konsumen = item.Konsumen.Nama
	}
	return resp
}

func ToItemResponseList(items []DataItem) []DataItemResponse {
	result := make([]DataItemResponse, len(items))
	for i, item := range items {
		result[i] = ToItemResponse(&item)
	}
	return result
}

func ToKonsumenResponse(konsumen *DataKonsumen) DataKonsumenResponse {
	resp := DataKonsumenResponse{
		ID:        konsumen.ID,
		Nama:      konsumen.Nama,
		NoHp:      konsumen.NoHp,
		Email:     konsumen.Email,
		Alamat:    konsumen.Alamat,
		GrupID:    konsumen.GrupID,
		Catatan:   konsumen.Catatan,
		Status:    konsumen.Status,
		CreatedAt: konsumen.CreatedAt,
		UpdatedAt: konsumen.UpdatedAt,
	}
	if konsumen.Grup != nil {
		resp.Grup = konsumen.Grup.Nama
	}
	return resp
}

func ToKonsumenResponseList(konsumens []DataKonsumen) []DataKonsumenResponse {
	result := make([]DataKonsumenResponse, len(konsumens))
	for i, k := range konsumens {
		result[i] = ToKonsumenResponse(&k)
	}
	return result
}
