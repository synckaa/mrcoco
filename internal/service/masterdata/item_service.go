package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type ItemService interface {
	GetAll(offset, limit int) ([]masterdata.DataItem, int64, error)
	GetByID(id uint) (*masterdata.DataItem, error)
	GetByKonsumenID(konsumenID uint, offset, limit int) ([]masterdata.DataItem, int64, error)
	Search(kode, nama string, kategoriID uint, konsumenID uint, status *bool, offset, limit int) ([]masterdata.DataItem, int64, error)
	Count() (int64, error)
	Create(item *masterdata.DataItem) error
	Update(item *masterdata.DataItem) error
	Delete(id uint) error
}

type itemService struct {
	repo repository.ItemRepository
	db   *gorm.DB
}

func NewItemService(repo repository.ItemRepository, db *gorm.DB) ItemService {
	return &itemService{repo: repo, db: db}
}

func (s *itemService) GetAll(offset, limit int) ([]masterdata.DataItem, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *itemService) GetByID(id uint) (*masterdata.DataItem, error) {
	return s.repo.FindByID(id)
}

func (s *itemService) GetByKonsumenID(konsumenID uint, offset, limit int) ([]masterdata.DataItem, int64, error) {
	return s.repo.FindByKonsumenID(konsumenID, offset, limit)
}

func (s *itemService) Search(kode, nama string, kategoriID uint, konsumenID uint, status *bool, offset, limit int) ([]masterdata.DataItem, int64, error) {
	return s.repo.Search(kode, nama, kategoriID, konsumenID, status, offset, limit)
}

func (s *itemService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *itemService) Create(item *masterdata.DataItem) error {
	return s.repo.Create(item)
}

func (s *itemService) Update(item *masterdata.DataItem) error {
	return s.repo.Update(item)
}

func (s *itemService) Delete(id uint) error {
	return s.repo.Delete(id)
}
