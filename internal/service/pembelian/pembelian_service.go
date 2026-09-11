package service

import (
	"mrcoco/internal/helper"
	pembelianmodel "mrcoco/internal/models/pembelian"
	"mrcoco/internal/repository/pembelian"
	"time"

	"gorm.io/gorm"
)

type PembelianService interface {
	GetAll(offset, limit int, tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi string) ([]pembelianmodel.DataPembelian, int64, error)
	GetByID(id uint) (*pembelianmodel.DataPembelian, error)
	Count(tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi string) (int64, error)
	Create(pembelian *pembelianmodel.DataPembelian) error
	Update(id uint, pembelian *pembelianmodel.DataPembelian) error
	Delete(id uint) error
	GetNextNoTransaksi(tanggal string) (string, error)
}

type pembelianService struct {
	repo repository.PembelianRepository
	db   *gorm.DB
}

func NewPembelianService(repo repository.PembelianRepository, db *gorm.DB) PembelianService {
	return &pembelianService{repo: repo, db: db}
}

func (s *pembelianService) GetAll(offset, limit int, tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi string) ([]pembelianmodel.DataPembelian, int64, error) {
	return s.repo.FindAll(offset, limit, tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi)
}

func (s *pembelianService) GetByID(id uint) (*pembelianmodel.DataPembelian, error) {
	return s.repo.FindByID(id)
}

func (s *pembelianService) Count(tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi string) (int64, error) {
	return s.repo.Count(tanggalAwal, tanggalAkhir, jatuhTempoAwal, jatuhTempoAkhir, supplierID, gudangID, rekeningID, noTransaksi)
}

func (s *pembelianService) Create(pembelian *pembelianmodel.DataPembelian) error {
	noTransaksi, err := helper.GenerateNoTransaksi(s.db, "P", pembelian.Tanggal.Time)
	if err != nil {
		return err
	}
	pembelian.NoTransaksi = noTransaksi
	return s.repo.Create(pembelian)
}

func (s *pembelianService) Update(id uint, pembelian *pembelianmodel.DataPembelian) error {
	if err := s.repo.DeleteItemsByPembelianID(id); err != nil {
		return err
	}
	pembelian.ID = id
	return s.repo.Update(pembelian)
}

func (s *pembelianService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *pembelianService) GetNextNoTransaksi(tanggal string) (string, error) {
	var t time.Time
	var err error
	if tanggal != "" {
		t, err = time.Parse("2006-01-02", tanggal)
		if err != nil {
			t = time.Now()
		}
	} else {
		t = time.Now()
	}
	return helper.GenerateNoTransaksi(s.db, "P", t)
}
