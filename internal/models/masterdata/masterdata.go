package masterdata

import (
	"encoding/json"
	"time"
)

func (DataSatuan) TableName() string {
	return "data_satuan"
}

type DataSatuan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataKategori) TableName() string {
	return "data_kategori"
}

type DataKategori struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataJenis) TableName() string {
	return "data_jenis"
}

type DataJenis struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataSubKlasifikasi) TableName() string {
	return "data_sub_klasifikasi"
}

type DataSubKlasifikasi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataKlasifikasi) TableName() string {
	return "data_klasifikasi"
}

type DataKlasifikasi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataItem) TableName() string {
	return "data_item"
}

type DataItem struct {
	ID         uint          `gorm:"primaryKey" json:"id"`
	Kode       string        `gorm:"type:varchar(50);not null;uniqueIndex" json:"kode" binding:"required"`
	Nama       string        `gorm:"type:varchar(200);not null" json:"nama" binding:"required"`
	JenisID    uint          `gorm:"not null" json:"jenis_id" binding:"required"`
	Jenis      *DataJenis    `gorm:"foreignKey:JenisID" json:"jenis,omitempty"`
	KategoriID uint          `gorm:"not null" json:"kategori_id" binding:"required"`
	Kategori   *DataKategori `gorm:"foreignKey:KategoriID" json:"kategori,omitempty"`
	SatuanID   uint          `gorm:"not null" json:"satuan_id" binding:"required"`
	Satuan     *DataSatuan   `gorm:"foreignKey:SatuanID" json:"satuan,omitempty"`
	RekeningID *uint         `json:"rekening_id" binding:"required"`
	Rekening   *DataRekening `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
	KonsumenID *uint         `json:"konsumen_id"`
	Konsumen   *DataKonsumen `gorm:"foreignKey:KonsumenID" json:"konsumen,omitempty"`
	HargaBeli  int           `gorm:"not null;default:0" json:"harga_beli" binding:"required"`
	HargaJual  int           `gorm:"not null;default:0" json:"harga_jual" binding:"required"`
	Status     bool          `gorm:"default:true" json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

func (DataGrup) TableName() string {
	return "data_grup"
}

type DataGrup struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	NoHp      string    `gorm:"type:varchar(20)" json:"no_hp"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Alamat    string    `gorm:"type:text" json:"alamat"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataKonsumen) TableName() string {
	return "data_konsumen"
}

type DataKonsumen struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null" json:"nama" binding:"required"`
	NoHp      string    `gorm:"type:varchar(20)" json:"no_hp"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Alamat    string    `gorm:"type:text" json:"alamat"`
	GrupID    uint      `gorm:"not null" json:"grup_id" binding:"required"`
	Grup      *DataGrup `gorm:"foreignKey:GrupID" json:"grup,omitempty"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	Status    bool      `gorm:"default:true" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataRekening) TableName() string {
	return "data_rekening"
}

type DataRekening struct {
	ID               uint                `gorm:"primaryKey" json:"id"`
	NoRek            string              `gorm:"type:varchar(50);not null;uniqueIndex" json:"no_rek" binding:"required"`
	Nama             string              `gorm:"type:varchar(100);not null" json:"nama" binding:"required"`
	SubKlasifikasiID *uint               `json:"sub_klasifikasi_id" binding:"required"`
	SubKlasifikasi   *DataSubKlasifikasi `gorm:"foreignKey:SubKlasifikasiID" json:"-"`
	KlasifikasiID    *uint               `json:"klasifikasi_id" binding:"required"`
	Klasifikasi      *DataKlasifikasi    `gorm:"foreignKey:KlasifikasiID" json:"-"`
	CreatedAt        time.Time           `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
}

type rekeningResponse struct {
	ID               uint      `json:"id"`
	NoRek            string    `json:"no_rek"`
	Nama             string    `json:"nama"`
	SubKlasifikasiID *uint     `json:"sub_klasifikasi_id"`
	SubKlasifikasi   string    `json:"sub_klasifikasi"`
	KlasifikasiID    *uint     `json:"klasifikasi_id"`
	Klasifikasi      string    `json:"klasifikasi"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (r DataRekening) MarshalJSON() ([]byte, error) {
	type Alias DataRekening
	a := (Alias)(r)

	resp := rekeningResponse{
		ID:               a.ID,
		NoRek:            a.NoRek,
		Nama:             a.Nama,
		SubKlasifikasiID: a.SubKlasifikasiID,
		KlasifikasiID:    a.KlasifikasiID,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
	}

	if a.SubKlasifikasi != nil {
		resp.SubKlasifikasi = a.SubKlasifikasi.Nama
	}
	if a.Klasifikasi != nil {
		resp.Klasifikasi = a.Klasifikasi.Nama
	}

	return json.Marshal(resp)
}

func (DataGudang) TableName() string {
	return "data_gudang"
}

type DataGudang struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DataSupplier) TableName() string {
	return "data_supplier"
}

type DataSupplier struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not null" json:"nama" binding:"required"`
	NoHp      string    `gorm:"type:varchar(20)" json:"no_hp"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Alamat    string    `gorm:"type:text" json:"alamat"`
	Bank      string    `gorm:"type:varchar(50)" json:"bank"`
	NoRek     string    `gorm:"type:varchar(50)" json:"no_rek"`
	AtasNama  string    `gorm:"type:varchar(100)" json:"atas_nama"`
	Catatan   string    `gorm:"type:text" json:"catatan"`
	Status    bool      `gorm:"default:true" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
