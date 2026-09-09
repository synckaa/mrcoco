package pembelian

import (
	"mrcoco/internal/models"
	"mrcoco/internal/models/masterdata"
	"time"
)

func (DataPembelian) TableName() string {
	return "data_pembelian"
}

type DataPembelian struct {
	ID             uint                  `gorm:"primaryKey" json:"id"`
	NoTransaksi    string                `gorm:"type:varchar(20);not null;uniqueIndex" json:"no_transaksi"`
	Tanggal        models.Date           `gorm:"type:date;not null" json:"tanggal" binding:"required"`
	SupplierID     uint                  `gorm:"not null" json:"supplier_id" binding:"required"`
	Supplier       *masterdata.DataSupplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	GudangID       uint                  `gorm:"not null" json:"gudang_id" binding:"required"`
	Gudang         *masterdata.DataGudang   `gorm:"foreignKey:GudangID" json:"gudang,omitempty"`
	TglJatuhTempo  models.Date           `gorm:"type:date" json:"tgl_jatuh_tempo"`
	RekeningID     *uint                 `json:"rekening_id"`
	Rekening       *masterdata.DataRekening `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
	GrandTotal     int                   `gorm:"not null;default:0" json:"grand_total"`
	JumlahUangMuka int                   `gorm:"not null;default:0" json:"jumlah_uang_muka"`
	JumlahUangSisa int                   `gorm:"not null;default:0" json:"jumlah_uang_sisa"`
	Catatan        string                `gorm:"type:text" json:"catatan"`
	Items          []DataPembelianItem   `gorm:"foreignKey:PembelianID" json:"items,omitempty"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
}

func (DataPembelianItem) TableName() string {
	return "data_pembelian_item"
}

type DataPembelianItem struct {
	ID          uint                    `gorm:"primaryKey" json:"id"`
	PembelianID uint                    `gorm:"not null" json:"pembelian_id"`
	ItemID      uint                    `gorm:"not null" json:"item_id" binding:"required"`
	Item        *masterdata.DataItem    `gorm:"foreignKey:ItemID" json:"item,omitempty"`
	Qty         int                     `gorm:"not null;default:0" json:"qty" binding:"required"`
	Harga       int                     `gorm:"not null;default:0" json:"harga" binding:"required"`
	Diskon      int                     `gorm:"not null;default:0" json:"diskon"`
	Subtotal    int                     `gorm:"not null;default:0" json:"subtotal"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
}
