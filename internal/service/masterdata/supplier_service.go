package service

import (
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/repository/masterdata"

	"gorm.io/gorm"
)

type SupplierService interface {
	GetAll(offset, limit int) ([]masterdata.DataSupplier, int64, error)
	GetByID(id uint) (*masterdata.DataSupplier, error)
	Search(nama, alamat string, status *bool, offset, limit int) ([]masterdata.DataSupplier, int64, error)
	Count() (int64, error)
	Create(supplier *masterdata.DataSupplier) error
	Update(supplier *masterdata.DataSupplier) error
	Delete(id uint) error
}

type supplierService struct {
	repo repository.SupplierRepository
	db   *gorm.DB
}

func NewSupplierService(repo repository.SupplierRepository, db *gorm.DB) SupplierService {
	return &supplierService{repo: repo, db: db}
}

func (s *supplierService) GetAll(offset, limit int) ([]masterdata.DataSupplier, int64, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *supplierService) GetByID(id uint) (*masterdata.DataSupplier, error) {
	return s.repo.FindByID(id)
}

func (s *supplierService) Search(nama, alamat string, status *bool, offset, limit int) ([]masterdata.DataSupplier, int64, error) {
	return s.repo.Search(nama, alamat, status, offset, limit)
}

func (s *supplierService) Count() (int64, error) {
	return s.repo.Count()
}

func (s *supplierService) Create(supplier *masterdata.DataSupplier) error {
	return s.repo.Create(supplier)
}

func (s *supplierService) Update(supplier *masterdata.DataSupplier) error {
	return s.repo.Update(supplier)
}

func (s *supplierService) Delete(id uint) error {
	return s.repo.Delete(id)
}
