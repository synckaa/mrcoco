package repository

import (
	pembelianmodel "mrcoco/internal/models/pembelian"
	bayarhutangmodel "mrcoco/internal/models/bayar_hutang"

	"gorm.io/gorm"
)

type PembelianRepository interface {
	FindAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]pembelianmodel.DataPembelian, int64, error)
	FindByID(id uint) (*pembelianmodel.DataPembelian, error)
	Count(startDate, endDate, supplierID, noTransaksi string) (int64, error)
	Create(pembelian *pembelianmodel.DataPembelian) error
	Update(pembelian *pembelianmodel.DataPembelian) error
	Delete(id uint) error
	DeleteItemsByPembelianID(pembelianID uint) error
}

type pembelianRepository struct {
	db *gorm.DB
}

func NewPembelianRepository(db *gorm.DB) PembelianRepository {
	return &pembelianRepository{db: db}
}

func (r *pembelianRepository) FindAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]pembelianmodel.DataPembelian, int64, error) {
	var pembelians []pembelianmodel.DataPembelian
	var total int64

	query := r.db.Model(&pembelianmodel.DataPembelian{})
	query = r.applyFilters(query, startDate, endDate, supplierID, noTransaksi)
	query.Count(&total)

	err := query.Preload("Supplier").Preload("Gudang").Preload("Rekening").Preload("Items").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&pembelians).Error
	return pembelians, total, err
}

func (r *pembelianRepository) FindByID(id uint) (*pembelianmodel.DataPembelian, error) {
	var pembelian pembelianmodel.DataPembelian
	err := r.db.Preload("Supplier").Preload("Gudang").Preload("Rekening").
		Preload("Items").Preload("Items.Item").Preload("Items.Item.Satuan").
		First(&pembelian, id).Error
	return &pembelian, err
}

func (r *pembelianRepository) Count(startDate, endDate, supplierID, noTransaksi string) (int64, error) {
	var count int64
	query := r.db.Model(&pembelianmodel.DataPembelian{})
	query = r.applyFilters(query, startDate, endDate, supplierID, noTransaksi)
	err := query.Count(&count).Error
	return count, err
}

func (r *pembelianRepository) Create(pembelian *pembelianmodel.DataPembelian) error {
	return r.db.Create(pembelian).Error
}

func (r *pembelianRepository) Update(pembelian *pembelianmodel.DataPembelian) error {
	return r.db.Save(pembelian).Error
}

func (r *pembelianRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("pembelian_id = ?", id).Delete(&bayarhutangmodel.DataBayarHutangItem{}).Error; err != nil {
			return err
		}
		if err := tx.Where("pembelian_id = ?", id).Delete(&pembelianmodel.DataPembelianItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&pembelianmodel.DataPembelian{}, id).Error
	})
}

func (r *pembelianRepository) DeleteItemsByPembelianID(pembelianID uint) error {
	return r.db.Where("pembelian_id = ?", pembelianID).Delete(&pembelianmodel.DataPembelianItem{}).Error
}

func (r *pembelianRepository) applyFilters(query *gorm.DB, startDate, endDate, supplierID, noTransaksi string) *gorm.DB {
	if startDate != "" {
		query = query.Where("tanggal >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("tanggal <= ?", endDate)
	}
	if supplierID != "" {
		query = query.Where("supplier_id = ?", supplierID)
	}
	if noTransaksi != "" {
		query = query.Where("no_transaksi ILIKE ?", "%"+noTransaksi+"%")
	}
	return query
}
