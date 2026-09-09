package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataSupplier, int64, error)
	FindByID(id uint) (*masterdata.DataSupplier, error)
	Search(nama, alamat string, status *bool, offset, limit int) ([]masterdata.DataSupplier, int64, error)
	Count() (int64, error)
	Create(supplier *masterdata.DataSupplier) error
	Update(supplier *masterdata.DataSupplier) error
	Delete(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) FindAll(offset, limit int) ([]masterdata.DataSupplier, int64, error) {
	var suppliers []masterdata.DataSupplier
	var total int64
	r.db.Model(&masterdata.DataSupplier{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&suppliers).Error
	return suppliers, total, err
}

func (r *supplierRepository) FindByID(id uint) (*masterdata.DataSupplier, error) {
	var supplier masterdata.DataSupplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *supplierRepository) Search(nama, alamat string, status *bool, offset, limit int) ([]masterdata.DataSupplier, int64, error) {
	var suppliers []masterdata.DataSupplier
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query = helper.ApplySmartSearch(query, "alamat", alamat)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query.Model(&masterdata.DataSupplier{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&suppliers).Error
	return suppliers, total, err
}

func (r *supplierRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataSupplier{}).Count(&count).Error
	return count, err
}

func (r *supplierRepository) Create(supplier *masterdata.DataSupplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) Update(supplier *masterdata.DataSupplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataSupplier{}, id).Error
}
