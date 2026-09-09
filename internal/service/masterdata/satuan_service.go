package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type SatuanService interface {
	GetAll(offset, limit int) ([]masterdata.DataSatuan, int64, error)
	GetByID(id uint) (*masterdata.DataSatuan, error)
	Search(nama string, offset, limit int) ([]masterdata.DataSatuan, int64, error)
	Count() (int64, error)
	Create(satuan *masterdata.DataSatuan) error
	Update(satuan *masterdata.DataSatuan) error
	Delete(id uint) error
}

type satuanService struct {
	repo repository.SatuanRepository
	db   *gorm.DB
}

func NewSatuanService(repo repository.SatuanRepository, db *gorm.DB) SatuanService {
	return &satuanService{repo: repo, db: db}
}

func (s *satuanService) GetAll(offset, limit int) ([]masterdata.DataSatuan, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *satuanService) GetByID(id uint) (*masterdata.DataSatuan, error) {
	return s.repo.FindByID(id)
}

func (s *satuanService) Search(nama string, offset, limit int) ([]masterdata.DataSatuan, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *satuanService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *satuanService) Create(satuan *masterdata.DataSatuan) error {
	return s.repo.Create(satuan)
}

func (s *satuanService) Update(satuan *masterdata.DataSatuan) error {
	return s.repo.Update(satuan)
}

func (s *satuanService) Delete(id uint) error {
	return s.repo.Delete(id)
}
