package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type GudangService interface {
	GetAll(offset, limit int) ([]masterdata.DataGudang, int64, error)
	GetByID(id uint) (*masterdata.DataGudang, error)
	Search(nama string, offset, limit int) ([]masterdata.DataGudang, int64, error)
	Count() (int64, error)
	Create(gudang *masterdata.DataGudang) error
	Update(gudang *masterdata.DataGudang) error
	Delete(id uint) error
}

type gudangService struct {
	repo repository.GudangRepository
	db   *gorm.DB
}

func NewGudangService(repo repository.GudangRepository, db *gorm.DB) GudangService {
	return &gudangService{repo: repo, db: db}
}

func (s *gudangService) GetAll(offset, limit int) ([]masterdata.DataGudang, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *gudangService) GetByID(id uint) (*masterdata.DataGudang, error) {
	return s.repo.FindByID(id)
}

func (s *gudangService) Search(nama string, offset, limit int) ([]masterdata.DataGudang, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *gudangService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *gudangService) Create(gudang *masterdata.DataGudang) error {
	return s.repo.Create(gudang)
}

func (s *gudangService) Update(gudang *masterdata.DataGudang) error {
	return s.repo.Update(gudang)
}

func (s *gudangService) Delete(id uint) error {
	return s.repo.Delete(id)
}
