package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type KlasifikasiRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataKlasifikasi, int64, error)
	FindByID(id uint) (*masterdata.DataKlasifikasi, error)
	Search(nama string, offset, limit int) ([]masterdata.DataKlasifikasi, int64, error)
	Count() (int64, error)
	Create(klasifikasi *masterdata.DataKlasifikasi) error
	Update(klasifikasi *masterdata.DataKlasifikasi) error
	Delete(id uint) error
}

type klasifikasiRepository struct {
	db *gorm.DB
}

func NewKlasifikasiRepository(db *gorm.DB) KlasifikasiRepository {
	return &klasifikasiRepository{db: db}
}

func (r *klasifikasiRepository) FindAll(offset, limit int) ([]masterdata.DataKlasifikasi, int64, error) {
	var klasifikasis []masterdata.DataKlasifikasi
	var total int64
	r.db.Model(&masterdata.DataKlasifikasi{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&klasifikasis).Error
	return klasifikasis, total, err
}

func (r *klasifikasiRepository) FindByID(id uint) (*masterdata.DataKlasifikasi, error) {
	var klasifikasi masterdata.DataKlasifikasi
	err := r.db.First(&klasifikasi, id).Error
	return &klasifikasi, err
}

func (r *klasifikasiRepository) Search(nama string, offset, limit int) ([]masterdata.DataKlasifikasi, int64, error) {
	var klasifikasis []masterdata.DataKlasifikasi
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataKlasifikasi{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&klasifikasis).Error
	return klasifikasis, total, err
}

func (r *klasifikasiRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataKlasifikasi{}).Count(&count).Error
	return count, err
}

func (r *klasifikasiRepository) Create(klasifikasi *masterdata.DataKlasifikasi) error {
	return r.db.Create(klasifikasi).Error
}

func (r *klasifikasiRepository) Update(klasifikasi *masterdata.DataKlasifikasi) error {
	return r.db.Save(klasifikasi).Error
}

func (r *klasifikasiRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataKlasifikasi{}, id).Error
}
