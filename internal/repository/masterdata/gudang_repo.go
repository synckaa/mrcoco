package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type GudangRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataGudang, int64, error)
	FindByID(id uint) (*masterdata.DataGudang, error)
	Search(nama string, offset, limit int) ([]masterdata.DataGudang, int64, error)
	Count() (int64, error)
	Create(gudang *masterdata.DataGudang) error
	Update(gudang *masterdata.DataGudang) error
	Delete(id uint) error
}

type gudangRepository struct {
	db *gorm.DB
}

func NewGudangRepository(db *gorm.DB) GudangRepository {
	return &gudangRepository{db: db}
}

func (r *gudangRepository) FindAll(offset, limit int) ([]masterdata.DataGudang, int64, error) {
	var gudangs []masterdata.DataGudang
	var total int64
	r.db.Model(&masterdata.DataGudang{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&gudangs).Error
	return gudangs, total, err
}

func (r *gudangRepository) FindByID(id uint) (*masterdata.DataGudang, error) {
	var gudang masterdata.DataGudang
	err := r.db.First(&gudang, id).Error
	return &gudang, err
}

func (r *gudangRepository) Search(nama string, offset, limit int) ([]masterdata.DataGudang, int64, error) {
	var gudangs []masterdata.DataGudang
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataGudang{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&gudangs).Error
	return gudangs, total, err
}

func (r *gudangRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataGudang{}).Count(&count).Error
	return count, err
}

func (r *gudangRepository) Create(gudang *masterdata.DataGudang) error {
	return r.db.Create(gudang).Error
}

func (r *gudangRepository) Update(gudang *masterdata.DataGudang) error {
	return r.db.Save(gudang).Error
}

func (r *gudangRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataGudang{}, id).Error
}
