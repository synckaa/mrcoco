package service

import (
	preordermodel "mrcoco/internal/models/pre_order"
	"mrcoco/internal/repository/pre_order"

	"gorm.io/gorm"
)

type PreOrderService interface {
	GetAll(offset, limit int, tanggal, konsumenID string) ([]preordermodel.DataPreOrder, int64, error)
	GetByID(id uint) (*preordermodel.DataPreOrder, error)
	Count(tanggal, konsumenID string) (int64, error)
	Create(preOrder *preordermodel.DataPreOrder) error
	Update(id uint, preOrder *preordermodel.DataPreOrder) error
	Delete(id uint) error
}

type preOrderService struct {
	repo repository.PreOrderRepository
	db   *gorm.DB
}

func NewPreOrderService(repo repository.PreOrderRepository, db *gorm.DB) PreOrderService {
	return &preOrderService{repo: repo, db: db}
}

func (s *preOrderService) GetAll(offset, limit int, tanggal, konsumenID string) ([]preordermodel.DataPreOrder, int64, error) {
	return s.repo.FindAll(offset, limit, tanggal, konsumenID)
}

func (s *preOrderService) GetByID(id uint) (*preordermodel.DataPreOrder, error) {
	return s.repo.FindByID(id)
}

func (s *preOrderService) Count(tanggal, konsumenID string) (int64, error) {
	return s.repo.Count(tanggal, konsumenID)
}

func (s *preOrderService) Create(preOrder *preordermodel.DataPreOrder) error {
	return s.repo.Create(preOrder)
}

func (s *preOrderService) Update(id uint, preOrder *preordermodel.DataPreOrder) error {
	if err := s.repo.DeleteItemsByPreOrderID(id); err != nil {
		return err
	}
	preOrder.ID = id
	return s.repo.Update(preOrder)
}

func (s *preOrderService) Delete(id uint) error {
	return s.repo.Delete(id)
}
