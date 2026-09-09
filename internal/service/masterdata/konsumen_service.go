package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type KonsumenService interface {
	GetAll(offset, limit int) ([]masterdata.DataKonsumen, int64, error)
	GetByID(id uint) (*masterdata.DataKonsumen, error)
	Search(nama, alamat, noHp string, grupID uint, status *bool, offset, limit int) ([]masterdata.DataKonsumen, int64, error)
	Count() (int64, error)
	Create(konsumen *masterdata.DataKonsumen) error
	Update(konsumen *masterdata.DataKonsumen) error
	Delete(id uint) error
}

type konsumenService struct {
	repo repository.KonsumenRepository
	db   *gorm.DB
}

func NewKonsumenService(repo repository.KonsumenRepository, db *gorm.DB) KonsumenService {
	return &konsumenService{repo: repo, db: db}
}

func (s *konsumenService) GetAll(offset, limit int) ([]masterdata.DataKonsumen, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *konsumenService) GetByID(id uint) (*masterdata.DataKonsumen, error) {
	return s.repo.FindByID(id)
}

func (s *konsumenService) Search(nama, alamat, noHp string, grupID uint, status *bool, offset, limit int) ([]masterdata.DataKonsumen, int64, error) {
	return s.repo.Search(nama, alamat, noHp, grupID, status, offset, limit)
}

func (s *konsumenService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *konsumenService) Create(konsumen *masterdata.DataKonsumen) error {
	return s.repo.Create(konsumen)
}

func (s *konsumenService) Update(konsumen *masterdata.DataKonsumen) error {
	return s.repo.Update(konsumen)
}

func (s *konsumenService) Delete(id uint) error {
	return s.repo.Delete(id)
}
