package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type KategoriService interface {
	GetAll(offset, limit int) ([]masterdata.DataKategori, int64, error)
	GetByID(id uint) (*masterdata.DataKategori, error)
	Search(nama string, offset, limit int) ([]masterdata.DataKategori, int64, error)
	Count() (int64, error)
	Create(kategori *masterdata.DataKategori) error
	Update(kategori *masterdata.DataKategori) error
	Delete(id uint) error
}

type kategoriService struct {
	repo repository.KategoriRepository
	db   *gorm.DB
}

func NewKategoriService(repo repository.KategoriRepository, db *gorm.DB) KategoriService {
	return &kategoriService{repo: repo, db: db}
}

func (s *kategoriService) GetAll(offset, limit int) ([]masterdata.DataKategori, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *kategoriService) GetByID(id uint) (*masterdata.DataKategori, error) {
	return s.repo.FindByID(id)
}

func (s *kategoriService) Search(nama string, offset, limit int) ([]masterdata.DataKategori, int64, error) {
	return s.repo.Search(nama, offset, limit)
}

func (s *kategoriService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *kategoriService) Create(kategori *masterdata.DataKategori) error {
	return s.repo.Create(kategori)
}

func (s *kategoriService) Update(kategori *masterdata.DataKategori) error {
	return s.repo.Update(kategori)
}

func (s *kategoriService) Delete(id uint) error {
	return s.repo.Delete(id)
}
