package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type SatuanRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataSatuan, int64, error)
	FindByID(id uint) (*masterdata.DataSatuan, error)
	Search(nama string, offset, limit int) ([]masterdata.DataSatuan, int64, error)
	Count() (int64, error)
	Create(satuan *masterdata.DataSatuan) error
	Update(satuan *masterdata.DataSatuan) error
	Delete(id uint) error
}

type satuanRepository struct {
	db *gorm.DB
}

func NewSatuanRepository(db *gorm.DB) SatuanRepository {
	return &satuanRepository{db: db}
}

func (r *satuanRepository) FindAll(offset, limit int) ([]masterdata.DataSatuan, int64, error) {
	var satus []masterdata.DataSatuan
	var total int64
	r.db.Model(&masterdata.DataSatuan{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&satus).Error
	return satus, total, err
}

func (r *satuanRepository) FindByID(id uint) (*masterdata.DataSatuan, error) {
	var satuan masterdata.DataSatuan
	err := r.db.First(&satuan, id).Error
	return &satuan, err
}

func (r *satuanRepository) Search(nama string, offset, limit int) ([]masterdata.DataSatuan, int64, error) {
	var satus []masterdata.DataSatuan
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataSatuan{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&satus).Error
	return satus, total, err
}

func (r *satuanRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataSatuan{}).Count(&count).Error
	return count, err
}

func (r *satuanRepository) Create(satuan *masterdata.DataSatuan) error {
	return r.db.Create(satuan).Error
}

func (r *satuanRepository) Update(satuan *masterdata.DataSatuan) error {
	return r.db.Save(satuan).Error
}

func (r *satuanRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataSatuan{}, id).Error
}
