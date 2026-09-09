package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type GrupService interface {
	GetAll(offset, limit int) ([]masterdata.DataGrup, int64, error)
	GetByID(id uint) (*masterdata.DataGrup, error)
	Search(nama, noHp, alamat string, status *bool, offset, limit int) ([]masterdata.DataGrup, int64, error)
	Count() (int64, error)
	Create(grup *masterdata.DataGrup) error
	Update(grup *masterdata.DataGrup) error
	Delete(id uint) error
}

type grupService struct {
	repo repository.GrupRepository
	db   *gorm.DB
}

func NewGrupService(repo repository.GrupRepository, db *gorm.DB) GrupService {
	return &grupService{repo: repo, db: db}
}

func (s *grupService) GetAll(offset, limit int) ([]masterdata.DataGrup, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *grupService) GetByID(id uint) (*masterdata.DataGrup, error) {
	return s.repo.FindByID(id)
}

func (s *grupService) Search(nama, noHp, alamat string, status *bool, offset, limit int) ([]masterdata.DataGrup, int64, error) {
	return s.repo.Search(nama, noHp, alamat, status, offset, limit)
}

func (s *grupService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *grupService) Create(grup *masterdata.DataGrup) error {
	return s.repo.Create(grup)
}

func (s *grupService) Update(grup *masterdata.DataGrup) error {
	return s.repo.Update(grup)
}

func (s *grupService) Delete(id uint) error {
	return s.repo.Delete(id)
}
