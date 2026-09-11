package pre_order

import (
	"mrcoco/internal/models"
	"mrcoco/internal/models/masterdata"
	"time"
)

func (DataPreOrder) TableName() string {
	return "data_pre_order"
}

type DataPreOrder struct {
	ID         uint                     `gorm:"primaryKey" json:"id"`
	Tanggal    models.Date              `gorm:"type:date;not null" json:"tanggal" binding:"required"`
	KonsumenID uint                     `gorm:"not null" json:"konsumen_id" binding:"required"`
	Konsumen   *masterdata.DataKonsumen `gorm:"foreignKey:KonsumenID" json:"konsumen,omitempty"`
	GrandTotal int                      `gorm:"not null;default:0" json:"grand_total"`
	Catatan    string                   `gorm:"type:text" json:"catatan"`
	Items      []DataPreOrderItem       `gorm:"foreignKey:PreOrderID" json:"items,omitempty"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}

func (DataPreOrderItem) TableName() string {
	return "data_pre_order_item"
}

type DataPreOrderItem struct {
	ID         uint                 `gorm:"primaryKey" json:"id"`
	PreOrderID uint                 `gorm:"not null" json:"pre_order_id"`
	ItemID     uint                 `gorm:"not null" json:"item_id" binding:"required"`
	Item       *masterdata.DataItem `gorm:"foreignKey:ItemID" json:"item,omitempty"`
	Qty        int                  `gorm:"not null;default:0" json:"qty" binding:"required"`
	Harga      int                  `gorm:"not null;default:0" json:"harga" binding:"required"`
	Subtotal   int                  `gorm:"not null;default:0" json:"subtotal"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}
