package main

import (
	"log"

	"mrcoco/internal/config/database"
	"mrcoco/internal/models/masterdata"
	"mrcoco/internal/models/pembelian"
	"mrcoco/internal/models/bayar_hutang"
	"mrcoco/internal/routes"

	"gorm.io/gorm"
)

func main() {
	database.ConnectDB()
	db := database.GetDB()

	runMigrations(db)

	r := routes.SetupRoutes()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&masterdata.DataSatuan{},
		&masterdata.DataKategori{},
		&masterdata.DataJenis{},
		&masterdata.DataSubKlasifikasi{},
		&masterdata.DataKlasifikasi{},
		&masterdata.DataItem{},
		&masterdata.DataGrup{},
		&masterdata.DataKonsumen{},
		&masterdata.DataRekening{},
		&masterdata.DataSupplier{},
		&masterdata.DataGudang{},
		&pembelian.DataPembelian{},
		&pembelian.DataPembelianItem{},
		&bayar_hutang.DataBayarHutang{},
		&bayar_hutang.DataBayarHutangItem{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
