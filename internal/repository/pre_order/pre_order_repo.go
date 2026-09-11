package repository

import (
	preordermodel "mrcoco/internal/models/pre_order"

	"gorm.io/gorm"
)

type PreOrderRepository interface {
	FindAll(offset, limit int, tanggal, konsumenID string) ([]preordermodel.DataPreOrder, int64, error)
	FindByID(id uint) (*preordermodel.DataPreOrder, error)
	Count(tanggal, konsumenID string) (int64, error)
	Create(preOrder *preordermodel.DataPreOrder) error
	Update(preOrder *preordermodel.DataPreOrder) error
	Delete(id uint) error
	DeleteItemsByPreOrderID(preOrderID uint) error
}

type preOrderRepository struct {
	db *gorm.DB
}

func NewPreOrderRepository(db *gorm.DB) PreOrderRepository {
	return &preOrderRepository{db: db}
}

func (r *preOrderRepository) FindAll(offset, limit int, tanggal, konsumenID string) ([]preordermodel.DataPreOrder, int64, error) {
	var preOrders []preordermodel.DataPreOrder
	var total int64

	query := r.db.Model(&preordermodel.DataPreOrder{})
	query = r.applyFilters(query, tanggal, konsumenID)
	query.Count(&total)

	err := query.Preload("Konsumen").Preload("Items").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&preOrders).Error
	return preOrders, total, err
}

func (r *preOrderRepository) FindByID(id uint) (*preordermodel.DataPreOrder, error) {
	var preOrder preordermodel.DataPreOrder
	err := r.db.Preload("Konsumen").
		Preload("Items").Preload("Items.Item").Preload("Items.Item.Satuan").
		First(&preOrder, id).Error
	return &preOrder, err
}

func (r *preOrderRepository) Count(tanggal, konsumenID string) (int64, error) {
	var count int64
	query := r.db.Model(&preordermodel.DataPreOrder{})
	query = r.applyFilters(query, tanggal, konsumenID)
	err := query.Count(&count).Error
	return count, err
}

func (r *preOrderRepository) Create(preOrder *preordermodel.DataPreOrder) error {
	return r.db.Create(preOrder).Error
}

func (r *preOrderRepository) Update(preOrder *preordermodel.DataPreOrder) error {
	return r.db.Save(preOrder).Error
}

func (r *preOrderRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("pre_order_id = ?", id).Delete(&preordermodel.DataPreOrderItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&preordermodel.DataPreOrder{}, id).Error
	})
}

func (r *preOrderRepository) DeleteItemsByPreOrderID(preOrderID uint) error {
	return r.db.Where("pre_order_id = ?", preOrderID).Delete(&preordermodel.DataPreOrderItem{}).Error
}

func (r *preOrderRepository) applyFilters(query *gorm.DB, tanggal, konsumenID string) *gorm.DB {
	if tanggal != "" {
		query = query.Where("tanggal = ?", tanggal)
	}
	if konsumenID != "" {
		query = query.Where("konsumen_id = ?", konsumenID)
	}
	return query
}
