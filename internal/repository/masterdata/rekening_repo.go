package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type RekeningRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataRekening, int64, error)
	FindByID(id uint) (*masterdata.DataRekening, error)
	Search(nama string, offset, limit int) ([]masterdata.DataRekening, int64, error)
	Count() (int64, error)
	Create(rekening *masterdata.DataRekening) error
	Update(rekening *masterdata.DataRekening) error
	Delete(id uint) error
}

type rekeningRepository struct {
	db *gorm.DB
}

func NewRekeningRepository(db *gorm.DB) RekeningRepository {
	return &rekeningRepository{db: db}
}

func (r *rekeningRepository) FindAll(offset, limit int) ([]masterdata.DataRekening, int64, error) {
	var rekenings []masterdata.DataRekening
	var total int64
	r.db.Model(&masterdata.DataRekening{}).Count(&total)
	err := r.db.Preload("Klasifikasi").Preload("SubKlasifikasi").Offset(offset).Limit(limit).Find(&rekenings).Error
	return rekenings, total, err
}

func (r *rekeningRepository) FindByID(id uint) (*masterdata.DataRekening, error) {
	var rekening masterdata.DataRekening
	err := r.db.Preload("Klasifikasi").Preload("SubKlasifikasi").First(&rekening, id).Error
	return &rekening, err
}

func (r *rekeningRepository) Search(nama string, offset, limit int) ([]masterdata.DataRekening, int64, error) {
	var rekenings []masterdata.DataRekening
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query.Model(&masterdata.DataRekening{}).Count(&total)
	err := query.Preload("Klasifikasi").Preload("SubKlasifikasi").Offset(offset).Limit(limit).Find(&rekenings).Error
	return rekenings, total, err
}

func (r *rekeningRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataRekening{}).Count(&count).Error
	return count, err
}

func (r *rekeningRepository) Create(rekening *masterdata.DataRekening) error {
	return r.db.Create(rekening).Error
}

func (r *rekeningRepository) Update(rekening *masterdata.DataRekening) error {
	return r.db.Save(rekening).Error
}

func (r *rekeningRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataRekening{}, id).Error
}
