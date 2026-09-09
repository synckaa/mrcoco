package repository

import (
	bayarhutangmodel "mrcoco/internal/models/bayar_hutang"
	pembelianmodel "mrcoco/internal/models/pembelian"

	"gorm.io/gorm"
)

type BayarHutangRepository interface {
	FindAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]bayarhutangmodel.DataBayarHutang, int64, error)
	FindByID(id uint) (*bayarhutangmodel.DataBayarHutang, error)
	Count(startDate, endDate, supplierID, noTransaksi string) (int64, error)
	Create(bayarHutang *bayarhutangmodel.DataBayarHutang) error
	Update(bayarHutang *bayarhutangmodel.DataBayarHutang) error
	Delete(id uint) error
	DeleteItemsByBayarHutangID(bayarHutangID uint) error
	FindPembelianSisaBySupplier(supplierID uint) ([]pembelianmodel.DataPembelian, error)
	SearchPembelianSisaBySupplier(supplierID uint, search string) ([]pembelianmodel.DataPembelian, error)
	UpdatePembelianSisa(pembelianID uint, sisa int) error
}

type bayarHutangRepository struct {
	db *gorm.DB
}

func NewBayarHutangRepository(db *gorm.DB) BayarHutangRepository {
	return &bayarHutangRepository{db: db}
}

func (r *bayarHutangRepository) FindAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]bayarhutangmodel.DataBayarHutang, int64, error) {
	var bayarHutangs []bayarhutangmodel.DataBayarHutang
	var total int64

	query := r.db.Model(&bayarhutangmodel.DataBayarHutang{})
	query = r.applyFilters(query, startDate, endDate, supplierID, noTransaksi)
	query.Count(&total)

	err := query.Preload("Supplier").Preload("Gudang").Preload("Rekening").Preload("Items").Preload("Items.Pembelian").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&bayarHutangs).Error
	return bayarHutangs, total, err
}

func (r *bayarHutangRepository) FindByID(id uint) (*bayarhutangmodel.DataBayarHutang, error) {
	var bayarHutang bayarhutangmodel.DataBayarHutang
	err := r.db.Preload("Supplier").Preload("Gudang").Preload("Rekening").
		Preload("Items").Preload("Items.Pembelian").
		First(&bayarHutang, id).Error
	return &bayarHutang, err
}

func (r *bayarHutangRepository) Count(startDate, endDate, supplierID, noTransaksi string) (int64, error) {
	var count int64
	query := r.db.Model(&bayarhutangmodel.DataBayarHutang{})
	query = r.applyFilters(query, startDate, endDate, supplierID, noTransaksi)
	err := query.Count(&count).Error
	return count, err
}

func (r *bayarHutangRepository) Create(bayarHutang *bayarhutangmodel.DataBayarHutang) error {
	return r.db.Create(bayarHutang).Error
}

func (r *bayarHutangRepository) Update(bayarHutang *bayarhutangmodel.DataBayarHutang) error {
	return r.db.Save(bayarHutang).Error
}

func (r *bayarHutangRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("bayar_hutang_id = ?", id).Delete(&bayarhutangmodel.DataBayarHutangItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&bayarhutangmodel.DataBayarHutang{}, id).Error
	})
}

func (r *bayarHutangRepository) DeleteItemsByBayarHutangID(bayarHutangID uint) error {
	return r.db.Where("bayar_hutang_id = ?", bayarHutangID).Delete(&bayarhutangmodel.DataBayarHutangItem{}).Error
}

func (r *bayarHutangRepository) FindPembelianSisaBySupplier(supplierID uint) ([]pembelianmodel.DataPembelian, error) {
	var pembelians []pembelianmodel.DataPembelian
	err := r.db.Where("supplier_id = ? AND jumlah_uang_sisa > ?", supplierID, 0).
		Order("created_at DESC").Find(&pembelians).Error
	return pembelians, err
}

func (r *bayarHutangRepository) SearchPembelianSisaBySupplier(supplierID uint, search string) ([]pembelianmodel.DataPembelian, error) {
	var pembelians []pembelianmodel.DataPembelian
	query := r.db.Where("supplier_id = ? AND jumlah_uang_sisa > ?", supplierID, 0)
	if search != "" {
		query = query.Where("no_transaksi ILIKE ?", "%"+search+"%")
	}
	err := query.Order("created_at DESC").Find(&pembelians).Error
	return pembelians, err
}

func (r *bayarHutangRepository) UpdatePembelianSisa(pembelianID uint, sisa int) error {
	return r.db.Model(&pembelianmodel.DataPembelian{}).Where("id = ?", pembelianID).Update("jumlah_uang_sisa", sisa).Error
}

func (r *bayarHutangRepository) applyFilters(query *gorm.DB, startDate, endDate, supplierID, noTransaksi string) *gorm.DB {
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
