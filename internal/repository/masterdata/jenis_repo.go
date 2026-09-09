package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type JenisRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataJenis, int64, error)
	FindByID(id uint) (*masterdata.DataJenis, error)
	Search(nama string, offset, limit int) ([]masterdata.DataJenis, int64, error)
	Count() (int64, error)
	Create(jenis *masterdata.DataJenis) error
	Update(jenis *masterdata.DataJenis) error
	Delete(id uint) error
}

type jenisRepository struct {
	db *gorm.DB
}

func NewJenisRepository(db *gorm.DB) JenisRepository {
	return &jenisRepository{db: db}
}

func (r *jenisRepository) FindAll(offset, limit int) ([]masterdata.DataJenis, int64, error) {
	var jenisList []masterdata.DataJenis
	var total int64
	r.db.Model(&masterdata.DataJenis{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&jenisList).Error
	return jenisList, total, err
}

func (r *jenisRepository) FindByID(id uint) (*masterdata.DataJenis, error) {
	var jenis masterdata.DataJenis
	err := r.db.First(&jenis, id).Error
	return &jenis, err
}

func (r *jenisRepository) Search(nama string, offset, limit int) ([]masterdata.DataJenis, int64, error) {
	var jenisList []masterdata.DataJenis
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataJenis{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&jenisList).Error
	return jenisList, total, err
}

func (r *jenisRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataJenis{}).Count(&count).Error
	return count, err
}

func (r *jenisRepository) Create(jenis *masterdata.DataJenis) error {
	return r.db.Create(jenis).Error
}

func (r *jenisRepository) Update(jenis *masterdata.DataJenis) error {
	return r.db.Save(jenis).Error
}

func (r *jenisRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataJenis{}, id).Error
}
