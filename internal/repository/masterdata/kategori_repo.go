package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataKategori, int64, error)
	FindByID(id uint) (*masterdata.DataKategori, error)
	Search(nama string, offset, limit int) ([]masterdata.DataKategori, int64, error)
	Count() (int64, error)
	Create(kategori *masterdata.DataKategori) error
	Update(kategori *masterdata.DataKategori) error
	Delete(id uint) error
}

type kategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) KategoriRepository {
	return &kategoriRepository{db: db}
}

func (r *kategoriRepository) FindAll(offset, limit int) ([]masterdata.DataKategori, int64, error) {
	var kategoris []masterdata.DataKategori
	var total int64
	r.db.Model(&masterdata.DataKategori{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&kategoris).Error
	return kategoris, total, err
}

func (r *kategoriRepository) FindByID(id uint) (*masterdata.DataKategori, error) {
	var kategori masterdata.DataKategori
	err := r.db.First(&kategori, id).Error
	return &kategori, err
}

func (r *kategoriRepository) Search(nama string, offset, limit int) ([]masterdata.DataKategori, int64, error) {
	var kategoris []masterdata.DataKategori
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataKategori{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&kategoris).Error
	return kategoris, total, err
}

func (r *kategoriRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataKategori{}).Count(&count).Error
	return count, err
}

func (r *kategoriRepository) Create(kategori *masterdata.DataKategori) error {
	return r.db.Create(kategori).Error
}

func (r *kategoriRepository) Update(kategori *masterdata.DataKategori) error {
	return r.db.Save(kategori).Error
}

func (r *kategoriRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataKategori{}, id).Error
}
