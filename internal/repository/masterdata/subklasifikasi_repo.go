package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type SubKlasifikasiRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error)
	FindByID(id uint) (*masterdata.DataSubKlasifikasi, error)
	Search(nama string, offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error)
	Count() (int64, error)
	Create(sub *masterdata.DataSubKlasifikasi) error
	Update(sub *masterdata.DataSubKlasifikasi) error
	Delete(id uint) error
}

type subKlasifikasiRepository struct {
	db *gorm.DB
}

func NewSubKlasifikasiRepository(db *gorm.DB) SubKlasifikasiRepository {
	return &subKlasifikasiRepository{db: db}
}

func (r *subKlasifikasiRepository) FindAll(offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error) {
	var subs []masterdata.DataSubKlasifikasi
	var total int64
	r.db.Model(&masterdata.DataSubKlasifikasi{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&subs).Error
	return subs, total, err
}

func (r *subKlasifikasiRepository) FindByID(id uint) (*masterdata.DataSubKlasifikasi, error) {
	var sub masterdata.DataSubKlasifikasi
	err := r.db.First(&sub, id).Error
	return &sub, err
}

func (r *subKlasifikasiRepository) Search(nama string, offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error) {
	var subs []masterdata.DataSubKlasifikasi
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataSubKlasifikasi{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&subs).Error
	return subs, total, err
}

func (r *subKlasifikasiRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataSubKlasifikasi{}).Count(&count).Error
	return count, err
}

func (r *subKlasifikasiRepository) Create(sub *masterdata.DataSubKlasifikasi) error {
	return r.db.Create(sub).Error
}

func (r *subKlasifikasiRepository) Update(sub *masterdata.DataSubKlasifikasi) error {
	return r.db.Save(sub).Error
}

func (r *subKlasifikasiRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataSubKlasifikasi{}, id).Error
}
