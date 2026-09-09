package repository

import (
	"mrcoco/internal/helper"
	"mrcoco/internal/models/masterdata"

	"gorm.io/gorm"
)

type GrupRepository interface {
	FindAll(offset, limit int) ([]masterdata.DataGrup, int64, error)
	FindByID(id uint) (*masterdata.DataGrup, error)
	Search(nama, noHp, alamat string, status *bool, offset, limit int) ([]masterdata.DataGrup, int64, error)
	Count() (int64, error)
	Create(grup *masterdata.DataGrup) error
	Update(grup *masterdata.DataGrup) error
	Delete(id uint) error
}

type grupRepository struct {
	db *gorm.DB
}

func NewGrupRepository(db *gorm.DB) GrupRepository {
	return &grupRepository{db: db}
}

func (r *grupRepository) FindAll(offset, limit int) ([]masterdata.DataGrup, int64, error) {
	var grups []masterdata.DataGrup
	var total int64
	r.db.Model(&masterdata.DataGrup{}).Count(&total)
	err := r.db.Offset(offset).Limit(limit).Find(&grups).Error
	return grups, total, err
}

func (r *grupRepository) FindByID(id uint) (*masterdata.DataGrup, error) {
	var grup masterdata.DataGrup
	err := r.db.First(&grup, id).Error
	return &grup, err
}

func (r *grupRepository) Search(nama, noHp, alamat string, status *bool, offset, limit int) ([]masterdata.DataGrup, int64, error) {
	var grups []masterdata.DataGrup
	var total int64
	query := r.db
	query = helper.ApplySmartSearch(query, "nama", nama)
	query = helper.ApplySmartSearch(query, "no_hp", noHp)
	query = helper.ApplySmartSearch(query, "alamat", alamat)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	query.Model(&masterdata.DataGrup{}).Count(&total)
	err := query.Offset(offset).Limit(limit).Find(&grups).Error
	return grups, total, err
}

func (r *grupRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&masterdata.DataGrup{}).Count(&count).Error
	return count, err
}

func (r *grupRepository) Create(grup *masterdata.DataGrup) error {
	return r.db.Create(grup).Error
}

func (r *grupRepository) Update(grup *masterdata.DataGrup) error {
	return r.db.Save(grup).Error
}

func (r *grupRepository) Delete(id uint) error {
	return r.db.Delete(&masterdata.DataGrup{}, id).Error
}
