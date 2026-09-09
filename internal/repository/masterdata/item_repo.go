package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataItem, int64, error)
	FindByID(id uint) (*masterdata.DataItem, error)
	FindByKonsumenID(konsumenID uint, offset, limit int) ([]masterdata.DataItem, int64, error)
	Search(kode, nama string, kategoriID uint, konsumenID uint, status *bool, offset, limit int) ([]masterdata.DataItem, int64, error)
	Count() (int64, error)
	Create(item *masterdata.DataItem) error
	Update(item *masterdata.DataItem) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) FindAll(offset, limit int) ([]masterdata.DataItem, int64, error) {
	var items []masterdata.DataItem
	var total int64
	r.db.Model(&masterdata.DataItem{}).Count(&total)
	err := r.db.Preload("Jenis").Preload("Kategori").Preload("Satuan").Preload("Rekening").Preload("Konsumen").Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *itemRepository) FindByID(id uint) (*masterdata.DataItem, error) {
	var item masterdata.DataItem
	err := r.db.Preload("Jenis").Preload("Kategori").Preload("Satuan").Preload("Rekening").Preload("Konsumen").First(&item, id).Error
	return &item, err
}

func (r *itemRepository) FindByKonsumenID(konsumenID uint, offset, limit int) ([]masterdata.DataItem, int64, error) {
	var items []masterdata.DataItem
	var total int64
	query := r.db.Model(&masterdata.DataItem{}).Where("konsumen_id IS NULL OR konsumen_id = ?", konsumenID)
	query.Count(&total)
	err := r.db.Preload("Jenis").Preload("Kategori").Preload("Satuan").Preload("Rekening").Preload("Konsumen").Where("konsumen_id IS NULL OR konsumen_id = ?", konsumenID).Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *itemRepository) Search(kode, nama string, kategoriID uint, konsumenID uint, status *bool, offset, limit int) ([]masterdata.DataItem, int64, error) {
	var items []masterdata.DataItem
	var total int64
	query := r.db.Preload("Jenis").Preload("Kategori").Preload("Satuan").Preload("Rekening").Preload("Konsumen")
	query = helper.ApplySmartSearch(query, "kode", kode)
	query = helper.ApplySmartSearch(query, "nama", nama)
	if kategoriID > 0 {
		query = query.Where("kategori_id = ?", kategoriID)
	}
	if konsumenID > 0 {
		query = query.Where("konsumen_id IS NULL OR konsumen_id = ?", konsumenID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query.Model(&masterdata.DataItem{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *itemRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataItem{}).Count(&count).Error
	return count, err
}

func (r *itemRepository) Create(item *masterdata.DataItem) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *masterdata.DataItem) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataItem{}, id).Error
}
