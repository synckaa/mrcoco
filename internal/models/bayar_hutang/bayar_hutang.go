package bayar_hutang

import (
	"mrcoco/internal/models"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/models/pembelian"
	"time"
)

func (DataBayarHutang) TableName() string {
	return "data_bayar_hutang"
}

type DataBayarHutang struct {
	ID          uint                           `gorm:"primaryKey" json:"id"`
	NoTransaksi string                         `gorm:"type:varchar(20);not null;uniqueIndex" json:"no_transaksi"`
	Tanggal     models.Date                    `gorm:"type:date;not null" json:"tanggal" binding:"required"`
	SupplierID  uint                           `gorm:"not null" json:"supplier_id" binding:"required"`
	Supplier    *masterdata.DataSupplier       `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	GudangID    uint                           `gorm:"not null" json:"gudang_id" binding:"required"`
	Gudang      *masterdata.DataGudang         `gorm:"foreignKey:GudangID" json:"gudang,omitempty"`
	RekeningID  *uint                          `json:"rekening_id"`
	Rekening    *masterdata.DataRekening       `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
	GrandTotal  int                            `gorm:"not null;default:0" json:"grand_total"`
	JumlahBayar int                            `gorm:"not null;default:0" json:"jumlah_bayar"`
	Sisa        int                            `gorm:"not null;default:0" json:"sisa"`
	Catatan     string                         `gorm:"type:text" json:"catatan"`
	Items       []DataBayarHutangItem          `gorm:"foreignKey:BayarHutangID" json:"items,omitempty"`
	CreatedAt   time.Time                      `json:"created_at"`
	UpdatedAt   time.Time                      `json:"updated_at"`
}

func (DataBayarHutangItem) TableName() string {
	return "data_bayar_hutang_item"
}

type DataBayarHutangItem struct {
	ID            uint                      `gorm:"primaryKey" json:"id"`
	BayarHutangID uint                      `gorm:"not null" json:"bayar_hutang_id"`
	PembelianID   uint                      `gorm:"not null" json:"pembelian_id"`
	Pembelian     *pembelian.DataPembelian  `gorm:"foreignKey:PembelianID" json:"pembelian,omitempty"`
	JumlahBayar   int                       `gorm:"not null;default:0" json:"jumlah_bayar"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
}
