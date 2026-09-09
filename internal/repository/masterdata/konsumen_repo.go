package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type KonsumenRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataKonsumen, int64, error)
	FindByID(id uint) (*masterdata.DataKonsumen, error)
	Search(nama, alamat, noHp string, grupID uint, status *bool, offset, limit int) ([]masterdata.DataKonsumen, int64, error)
	Count() (int64, error)
	Create(konsumen *masterdata.DataKonsumen) error
	Update(konsumen *masterdata.DataKonsumen) error
	Delete(id uint) error
}

type konsumenRepository struct {
	db *gorm.DB
}

func NewKonsumenRepository(db *gorm.DB) KonsumenRepository {
	return &konsumenRepository{db: db}
}

func (r *konsumenRepository) FindAll(offset, limit int) ([]masterdata.DataKonsumen, int64, error) {
	var konsumens []masterdata.DataKonsumen
	var total int64
	r.db.Model(&masterdata.DataKonsumen{}).Count(&total)
	err := r.db.Preload("Grup").Offset(offset).Limit(limit).Find(&konsumens).Error
	return konsumens, total, err
}

func (r *konsumenRepository) FindByID(id uint) (*masterdata.DataKonsumen, error) {
	var konsumen masterdata.DataKonsumen
	err := r.db.Preload("Grup").First(&konsumen, id).Error
	return &konsumen, err
}

func (r *konsumenRepository) Search(nama, alamat, noHp string, grupID uint, status *bool, offset, limit int) ([]masterdata.DataKonsumen, int64, error) {
	var konsumens []masterdata.DataKonsumen
	var total int64
	query := r.db.Preload("Grup")
	query = helper.ApplySmartSearch(query, "nama", nama)
	query = helper.ApplySmartSearch(query, "alamat", alamat)
	query = helper.ApplySmartSearch(query, "no_hp", noHp)
	if grupID > 0 {
		query = query.Where("grup_id = ?", grupID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query.Model(&masterdata.DataKonsumen{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&konsumens).Error
	return konsumens, total, err
}

func (r *konsumenRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataKonsumen{}).Count(&count).Error
	return count, err
}

func (r *konsumenRepository) Create(konsumen *masterdata.DataKonsumen) error {
	return r.db.Create(konsumen).Error
}

func (r *konsumenRepository) Update(konsumen *masterdata.DataKonsumen) error {
	return r.db.Save(konsumen).Error
}

func (r *konsumenRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataKonsumen{}, id).Error
}
