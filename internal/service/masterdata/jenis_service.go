package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type JenisService interface {
	GetAll(offset, limit int) ([]masterdata.DataJenis, int64, error)
	GetByID(id uint) (*masterdata.DataJenis, error)
	Search(nama string, offset, limit int) ([]masterdata.DataJenis, int64, error)
	Count() (int64, error)
	Create(jenis *masterdata.DataJenis) error
	Update(jenis *masterdata.DataJenis) error
	Delete(id uint) error
}

type jenisService struct {
	repo repository.JenisRepository
	db   *gorm.DB
}

func NewJenisService(repo repository.JenisRepository, db *gorm.DB) JenisService {
	return &jenisService{repo: repo, db: db}
}

func (s *jenisService) GetAll(offset, limit int) ([]masterdata.DataJenis, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *jenisService) GetByID(id uint) (*masterdata.DataJenis, error) {
	return s.repo.FindByID(id)
}

func (s *jenisService) Search(nama string, offset, limit int) ([]masterdata.DataJenis, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *jenisService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *jenisService) Create(jenis *masterdata.DataJenis) error {
	return s.repo.Create(jenis)
}

func (s *jenisService) Update(jenis *masterdata.DataJenis) error {
	return s.repo.Update(jenis)
}

func (s *jenisService) Delete(id uint) error {
	return s.repo.Delete(id)
}
