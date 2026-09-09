package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type SubKlasifikasiService interface {
	GetAll(offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error)
	GetByID(id uint) (*masterdata.DataSubKlasifikasi, error)
	Search(nama string, offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error)
	Count() (int64, error)
	Create(sub *masterdata.DataSubKlasifikasi) error
	Update(sub *masterdata.DataSubKlasifikasi) error
	Delete(id uint) error
}

type subKlasifikasiService struct {
	repo repository.SubKlasifikasiRepository
	db   *gorm.DB
}

func NewSubKlasifikasiService(repo repository.SubKlasifikasiRepository, db *gorm.DB) SubKlasifikasiService {
	return &subKlasifikasiService{repo: repo, db: db}
}

func (s *subKlasifikasiService) GetAll(offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *subKlasifikasiService) GetByID(id uint) (*masterdata.DataSubKlasifikasi, error) {
	return s.repo.FindByID(id)
}

func (s *subKlasifikasiService) Search(nama string, offset, limit int) ([]masterdata.DataSubKlasifikasi, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *subKlasifikasiService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *subKlasifikasiService) Create(sub *masterdata.DataSubKlasifikasi) error {
	return s.repo.Create(sub)
}

func (s *subKlasifikasiService) Update(sub *masterdata.DataSubKlasifikasi) error {
	return s.repo.Update(sub)
}

func (s *subKlasifikasiService) Delete(id uint) error {
	return s.repo.Delete(id)
}
