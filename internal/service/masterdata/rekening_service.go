package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type RekeningService interface {
	GetAll(offset, limit int) ([]masterdata.DataRekening, int64, error)
	GetByID(id uint) (*masterdata.DataRekening, error)
	Search(nama string, offset, limit int) ([]masterdata.DataRekening, int64, error)
	Count() (int64, error)
	Create(rekening *masterdata.DataRekening) error
	Update(rekening *masterdata.DataRekening) error
	Delete(id uint) error
}

type rekeningService struct {
	repo repository.RekeningRepository
	db   *gorm.DB
}

func NewRekeningService(repo repository.RekeningRepository, db *gorm.DB) RekeningService {
	return &rekeningService{repo: repo, db: db}
}

func (s *rekeningService) GetAll(offset, limit int) ([]masterdata.DataRekening, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *rekeningService) GetByID(id uint) (*masterdata.DataRekening, error) {
	return s.repo.FindByID(id)
}

func (s *rekeningService) Search(nama string, offset, limit int) ([]masterdata.DataRekening, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *rekeningService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *rekeningService) Create(rekening *masterdata.DataRekening) error {
	return s.repo.Create(rekening)
}

func (s *rekeningService) Update(rekening *masterdata.DataRekening) error {
	return s.repo.Update(rekening)
}

func (s *rekeningService) Delete(id uint) error {
	return s.repo.Delete(id)
}
