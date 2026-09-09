package service

import (
	"fmt"
	"mrcoco/internal/helper"
	bayarhutangmodel "mrcoco/internal/models/bayar_hutang"
	pembelianmodel "mrcoco/internal/models/pembelian"
	"mrcoco/internal/repository/bayar_hutang"
	"time"

	"gorm.io/gorm"
)

type BayarHutangService interface {
	GetAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]bayarhutangmodel.DataBayarHutang, int64, error)
	GetByID(id uint) (*bayarhutangmodel.DataBayarHutang, error)
	Count(startDate, endDate, supplierID, noTransaksi string) (int64, error)
	Create(bayarHutang *bayarhutangmodel.DataBayarHutang) error
	Update(id uint, bayarHutang *bayarhutangmodel.DataBayarHutang) error
	Delete(id uint) error
	GetNextNoTransaksi(tanggal string) (string, error)
	GetPembelianSisaBySupplier(supplierID uint) ([]pembelianmodel.DataPembelian, error)
	SearchPembelianSisaBySupplier(supplierID uint, search string) ([]pembelianmodel.DataPembelian, error)
}

type bayarHutangService struct {
	repo repository.BayarHutangRepository
	db   *gorm.DB
}

func NewBayarHutangService(repo repository.BayarHutangRepository, db *gorm.DB) BayarHutangService {
	return &bayarHutangService{repo: repo, db: db}
}

func (s *bayarHutangService) GetAll(offset, limit int, startDate, endDate, supplierID, noTransaksi string) ([]bayarhutangmodel.DataBayarHutang, int64, error) {
	return s.repo.FindAll(offset, limit, startDate, endDate, supplierID, noTransaksi)
}

func (s *bayarHutangService) GetByID(id uint) (*bayarhutangmodel.DataBayarHutang, error) {
	return s.repo.FindByID(id)
}

func (s *bayarHutangService) Count(startDate, endDate, supplierID, noTransaksi string) (int64, error) {
	return s.repo.Count(startDate, endDate, supplierID, noTransaksi)
}

func (s *bayarHutangService) Create(bayarHutang *bayarhutangmodel.DataBayarHutang) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		noTransaksi, err := helper.GenerateNoTransaksi(tx, "H", bayarHutang.Tanggal.Time)
		if err != nil {
			return err
		}
		bayarHutang.NoTransaksi = noTransaksi

		if err := tx.Create(bayarHutang).Error; err != nil {
			return err
		}

		for _, item := range bayarHutang.Items {
			var pembelian pembelianmodel.DataPembelian
			if err := tx.First(&pembelian, item.PembelianID).Error; err != nil {
				return fmt.Errorf("pembelian %d not found", item.PembelianID)
			}

			newSisa := pembelian.JumlahUangSisa - item.JumlahBayar
			if newSisa < 0 {
				return fmt.Errorf("jumlah bayar melebihi sisa hutang untuk pembelian %s: sisa=%d, bayar=%d", pembelian.NoTransaksi, pembelian.JumlahUangSisa, item.JumlahBayar)
			}
			newMuka := pembelian.JumlahUangMuka + item.JumlahBayar

			if err := tx.Model(&pembelianmodel.DataPembelian{}).Where("id = ?", item.PembelianID).Updates(map[string]interface{}{
				"jumlah_uang_sisa": newSisa,
				"jumlah_uang_muka": newMuka,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *bayarHutangService) Update(id uint, bayarHutang *bayarhutangmodel.DataBayarHutang) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var oldBayarHutang bayarhutangmodel.DataBayarHutang
		if err := tx.Preload("Items").First(&oldBayarHutang, id).Error; err != nil {
			return err
		}

		hasNewItems := len(bayarHutang.Items) > 0

		if hasNewItems {
			for _, oldItem := range oldBayarHutang.Items {
				var pembelian pembelianmodel.DataPembelian
				if err := tx.First(&pembelian, oldItem.PembelianID).Error; err != nil {
					return err
				}
				restoreSisa := pembelian.JumlahUangSisa + oldItem.JumlahBayar
				restoreMuka := pembelian.JumlahUangMuka - oldItem.JumlahBayar
				if restoreMuka < 0 {
					return fmt.Errorf("data konsistensi error: jumlah_uang_muka menjadi negatif untuk pembelian %s (muka=%d, bayar=%d)", pembelian.NoTransaksi, pembelian.JumlahUangMuka, oldItem.JumlahBayar)
				}
				if err := tx.Model(&pembelianmodel.DataPembelian{}).Where("id = ?", oldItem.PembelianID).Updates(map[string]interface{}{
					"jumlah_uang_sisa": restoreSisa,
					"jumlah_uang_muka": restoreMuka,
				}).Error; err != nil {
					return err
				}
			}

			if err := tx.Where("bayar_hutang_id = ?", id).Delete(&bayarhutangmodel.DataBayarHutangItem{}).Error; err != nil {
				return err
			}
		}

		bayarHutang.NoTransaksi = oldBayarHutang.NoTransaksi
		if err := tx.Model(&bayarhutangmodel.DataBayarHutang{}).Where("id = ?", id).Updates(bayarHutang).Error; err != nil {
			return err
		}

		if hasNewItems {
			for _, item := range bayarHutang.Items {
				var pembelian pembelianmodel.DataPembelian
				if err := tx.First(&pembelian, item.PembelianID).Error; err != nil {
					return fmt.Errorf("pembelian %d not found", item.PembelianID)
				}

				newSisa := pembelian.JumlahUangSisa - item.JumlahBayar
				if newSisa < 0 {
					return fmt.Errorf("jumlah bayar melebihi sisa hutang untuk pembelian %s: sisa=%d, bayar=%d", pembelian.NoTransaksi, pembelian.JumlahUangSisa, item.JumlahBayar)
				}
				newMuka := pembelian.JumlahUangMuka + item.JumlahBayar

				if err := tx.Model(&pembelianmodel.DataPembelian{}).Where("id = ?", item.PembelianID).Updates(map[string]interface{}{
					"jumlah_uang_sisa": newSisa,
					"jumlah_uang_muka": newMuka,
				}).Error; err != nil {
					return err
				}

				item.BayarHutangID = id
				if err := tx.Create(&item).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (s *bayarHutangService) Delete(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var bayarHutang bayarhutangmodel.DataBayarHutang
		if err := tx.Preload("Items").First(&bayarHutang, id).Error; err != nil {
			return err
		}

		for _, item := range bayarHutang.Items {
			var pembelian pembelianmodel.DataPembelian
			if err := tx.First(&pembelian, item.PembelianID).Error; err != nil {
				return err
			}
			restoreSisa := pembelian.JumlahUangSisa + item.JumlahBayar
			restoreMuka := pembelian.JumlahUangMuka - item.JumlahBayar
			if restoreMuka < 0 {
				return fmt.Errorf("data konsistensi error: jumlah_uang_muka menjadi negatif untuk pembelian %s (muka=%d, bayar=%d)", pembelian.NoTransaksi, pembelian.JumlahUangMuka, item.JumlahBayar)
			}
			if err := tx.Model(&pembelianmodel.DataPembelian{}).Where("id = ?", item.PembelianID).Updates(map[string]interface{}{
				"jumlah_uang_sisa": restoreSisa,
				"jumlah_uang_muka": restoreMuka,
			}).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("bayar_hutang_id = ?", id).Delete(&bayarhutangmodel.DataBayarHutangItem{}).Error; err != nil {
			return err
		}

		return tx.Delete(&bayarhutangmodel.DataBayarHutang{}, id).Error
	})
}

func (s *bayarHutangService) GetNextNoTransaksi(tanggal string) (string, error) {
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
	return helper.GenerateNoTransaksi(s.db, "H", t)
}

func (s *bayarHutangService) GetPembelianSisaBySupplier(supplierID uint) ([]pembelianmodel.DataPembelian, error) {
	return s.repo.FindPembelianSisaBySupplier(supplierID)
}

func (s *bayarHutangService) SearchPembelianSisaBySupplier(supplierID uint, search string) ([]pembelianmodel.DataPembelian, error) {
	return s.repo.SearchPembelianSisaBySupplier(supplierID, search)
}
