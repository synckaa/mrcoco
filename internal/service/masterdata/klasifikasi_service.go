package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type KlasifikasiService interface {
	GetAll(offset, limit int) ([]masterdata.DataKlasifikasi, int64, error)
	GetByID(id uint) (*masterdata.DataKlasifikasi, error)
	Search(nama string, offset, limit int) ([]masterdata.DataKlasifikasi, int64, error)
	Count() (int64, error)
	Create(klasifikasi *masterdata.DataKlasifikasi) error
	Update(klasifikasi *masterdata.DataKlasifikasi) error
	Delete(id uint) error
}

type klasifikasiService struct {
	repo repository.KlasifikasiRepository
	db   *gorm.DB
}

func NewKlasifikasiService(repo repository.KlasifikasiRepository, db *gorm.DB) KlasifikasiService {
	return &klasifikasiService{repo: repo, db: db}
}

func (s *klasifikasiService) GetAll(offset, limit int) ([]masterdata.DataKlasifikasi, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *klasifikasiService) GetByID(id uint) (*masterdata.DataKlasifikasi, error) {
	return s.repo.FindByID(id)
}

func (s *klasifikasiService) Search(nama string, offset, limit int) ([]masterdata.DataKlasifikasi, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *klasifikasiService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *klasifikasiService) Create(klasifikasi *masterdata.DataKlasifikasi) error {
	return s.repo.Create(klasifikasi)
}

func (s *klasifikasiService) Update(klasifikasi *masterdata.DataKlasifikasi) error {
	return s.repo.Update(klasifikasi)
}

func (s *klasifikasiService) Delete(id uint) error {
	return s.repo.Delete(id)
}
