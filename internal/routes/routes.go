package routes

import (
	"mrcoco/internal/config/database"
	bayarHutangHandler "mrcoco/internal/handler/bayar_hutang"
	"mrcoco/internal/handler/masterdata"
	pembelianHandler "mrcoco/internal/handler/pembelian"
	bayarHutangRepo "mrcoco/internal/repository/bayar_hutang"
	"mrcoco/internal/repository/masterdata"
	pembelianRepo "mrcoco/internal/repository/pembelian"
	bayarHutangService "mrcoco/internal/service/bayar_hutang"
	"mrcoco/internal/service/masterdata"
	pembelianService "mrcoco/internal/service/pembelian"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	db := database.GetDB()

	satuanRepo := repository.NewSatuanRepository(db)
	kategoriRepo := repository.NewKategoriRepository(db)
	jenisRepo := repository.NewJenisRepository(db)
	subKlasifikasiRepo := repository.NewSubKlasifikasiRepository(db)
	klasifikasiRepo := repository.NewKlasifikasiRepository(db)
	itemRepo := repository.NewItemRepository(db)
	grupRepo := repository.NewGrupRepository(db)
	konsumenRepo := repository.NewKonsumenRepository(db)
	rekeningRepo := repository.NewRekeningRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	gudangRepo := repository.NewGudangRepository(db)
	pembelianRepoInstance := pembelianRepo.NewPembelianRepository(db)
	bayarHutangRepoInstance := bayarHutangRepo.NewBayarHutangRepository(db)

	satuanService := service.NewSatuanService(satuanRepo, db)
	kategoriService := service.NewKategoriService(kategoriRepo, db)
	jenisService := service.NewJenisService(jenisRepo, db)
	subKlasifikasiService := service.NewSubKlasifikasiService(subKlasifikasiRepo, db)
	klasifikasiService := service.NewKlasifikasiService(klasifikasiRepo, db)
	itemService := service.NewItemService(itemRepo, db)
	grupService := service.NewGrupService(grupRepo, db)
	konsumenService := service.NewKonsumenService(konsumenRepo, db)
	rekeningService := service.NewRekeningService(rekeningRepo, db)
	supplierService := service.NewSupplierService(supplierRepo, db)
	gudangService := service.NewGudangService(gudangRepo, db)
	pembelianServiceInstance := pembelianService.NewPembelianService(pembelianRepoInstance, db)
	bayarHutangServiceInstance := bayarHutangService.NewBayarHutangService(bayarHutangRepoInstance, db)

	satuanHandler := handler.NewSatuanHandler(satuanService)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)
	jenisHandler := handler.NewJenisHandler(jenisService)
	subKlasifikasiHandler := handler.NewSubKlasifikasiHandler(subKlasifikasiService)
	klasifikasiHandler := handler.NewKlasifikasiHandler(klasifikasiService)
	itemHandler := handler.NewItemHandler(itemService)
	grupHandler := handler.NewGrupHandler(grupService)
	konsumenHandler := handler.NewKonsumenHandler(konsumenService)
	rekeningHandler := handler.NewRekeningHandler(rekeningService)
	supplierHandler := handler.NewSupplierHandler(supplierService)
	gudangHandler := handler.NewGudangHandler(gudangService)
	pembelianHandlerInstance := pembelianHandler.NewPembelianHandler(pembelianServiceInstance)
	bayarHutangHandlerInstance := bayarHutangHandler.NewBayarHutangHandler(bayarHutangServiceInstance)

	api := r.Group("/api")
	{
		satuan := api.Group("/satuan")
		{
			satuan.GET("", satuanHandler.GetAll)
			satuan.GET("/search", satuanHandler.Search)
			satuan.GET("/total", satuanHandler.Count)
			satuan.GET("/:id", satuanHandler.GetByID)
			satuan.POST("", satuanHandler.Create)
			satuan.PUT("/:id", satuanHandler.Update)
			satuan.DELETE("/:id", satuanHandler.Delete)
		}

		kategori := api.Group("/kategori")
		{
			kategori.GET("", kategoriHandler.GetAll)
			kategori.GET("/search", kategoriHandler.Search)
			kategori.GET("/total", kategoriHandler.Count)
			kategori.GET("/:id", kategoriHandler.GetByID)
			kategori.POST("", kategoriHandler.Create)
			kategori.PUT("/:id", kategoriHandler.Update)
			kategori.DELETE("/:id", kategoriHandler.Delete)
		}

		jenis := api.Group("/jenis")
		{
			jenis.GET("", jenisHandler.GetAll)
			jenis.GET("/search", jenisHandler.Search)
			jenis.GET("/total", jenisHandler.Count)
			jenis.GET("/:id", jenisHandler.GetByID)
			jenis.POST("", jenisHandler.Create)
			jenis.PUT("/:id", jenisHandler.Update)
			jenis.DELETE("/:id", jenisHandler.Delete)
		}

		subKlasifikasi := api.Group("/sub-klasifikasi")
		{
			subKlasifikasi.GET("", subKlasifikasiHandler.GetAll)
			subKlasifikasi.GET("/search", subKlasifikasiHandler.Search)
			subKlasifikasi.GET("/total", subKlasifikasiHandler.Count)
			subKlasifikasi.GET("/:id", subKlasifikasiHandler.GetByID)
			subKlasifikasi.POST("", subKlasifikasiHandler.Create)
			subKlasifikasi.PUT("/:id", subKlasifikasiHandler.Update)
			subKlasifikasi.DELETE("/:id", subKlasifikasiHandler.Delete)
		}

		klasifikasi := api.Group("/klasifikasi")
		{
			klasifikasi.GET("", klasifikasiHandler.GetAll)
			klasifikasi.GET("/search", klasifikasiHandler.Search)
			klasifikasi.GET("/total", klasifikasiHandler.Count)
			klasifikasi.GET("/:id", klasifikasiHandler.GetByID)
			klasifikasi.POST("", klasifikasiHandler.Create)
			klasifikasi.PUT("/:id", klasifikasiHandler.Update)
			klasifikasi.DELETE("/:id", klasifikasiHandler.Delete)
		}

		item := api.Group("/item")
		{
			item.GET("", itemHandler.GetAll)
			item.GET("/search", itemHandler.Search)
			item.GET("/total", itemHandler.Count)
			item.GET("/konsumen/:konsumen_id", itemHandler.GetByKonsumenID)
			item.GET("/:id", itemHandler.GetByID)
			item.POST("", itemHandler.Create)
			item.PUT("/:id", itemHandler.Update)
			item.DELETE("/:id", itemHandler.Delete)
		}

		grup := api.Group("/grup")
		{
			grup.GET("", grupHandler.GetAll)
			grup.GET("/search", grupHandler.Search)
			grup.GET("/total", grupHandler.Count)
			grup.GET("/:id", grupHandler.GetByID)
			grup.POST("", grupHandler.Create)
			grup.PUT("/:id", grupHandler.Update)
			grup.DELETE("/:id", grupHandler.Delete)
		}

		konsumen := api.Group("/konsumen")
		{
			konsumen.GET("", konsumenHandler.GetAll)
			konsumen.GET("/search", konsumenHandler.Search)
			konsumen.GET("/total", konsumenHandler.Count)
			konsumen.GET("/:id", konsumenHandler.GetByID)
			konsumen.POST("", konsumenHandler.Create)
			konsumen.PUT("/:id", konsumenHandler.Update)
			konsumen.DELETE("/:id", konsumenHandler.Delete)
		}

		rekening := api.Group("/rekening")
		{
			rekening.GET("", rekeningHandler.GetAll)
			rekening.GET("/search", rekeningHandler.Search)
			rekening.GET("/total", rekeningHandler.Count)
			rekening.GET("/:id", rekeningHandler.GetByID)
			rekening.POST("", rekeningHandler.Create)
			rekening.PUT("/:id", rekeningHandler.Update)
			rekening.DELETE("/:id", rekeningHandler.Delete)
		}

		supplier := api.Group("/supplier")
		{
			supplier.GET("", supplierHandler.GetAll)
			supplier.GET("/search", supplierHandler.Search)
			supplier.GET("/total", supplierHandler.Count)
			supplier.GET("/:id", supplierHandler.GetByID)
			supplier.POST("", supplierHandler.Create)
			supplier.PUT("/:id", supplierHandler.Update)
			supplier.DELETE("/:id", supplierHandler.Delete)
		}

		gudang := api.Group("/gudang")
		{
			gudang.GET("", gudangHandler.GetAll)
			gudang.GET("/search", gudangHandler.Search)
			gudang.GET("/total", gudangHandler.Count)
			gudang.GET("/:id", gudangHandler.GetByID)
			gudang.POST("", gudangHandler.Create)
			gudang.PUT("/:id", gudangHandler.Update)
			gudang.DELETE("/:id", gudangHandler.Delete)
		}

		pembelian := api.Group("/pembelian")
		{
			pembelian.GET("", pembelianHandlerInstance.GetAll)
			pembelian.GET("/total", pembelianHandlerInstance.Count)
			pembelian.GET("/next-no", pembelianHandlerInstance.GetNextNoTransaksi)
			pembelian.GET("/:id", pembelianHandlerInstance.GetByID)
			pembelian.POST("", pembelianHandlerInstance.Create)
			pembelian.PUT("/:id", pembelianHandlerInstance.Update)
			pembelian.DELETE("/:id", pembelianHandlerInstance.Delete)
		}

		bayarHutang := api.Group("/bayar-hutang")
		{
			bayarHutang.GET("", bayarHutangHandlerInstance.GetAll)
			bayarHutang.GET("/total", bayarHutangHandlerInstance.Count)
			bayarHutang.GET("/next-no", bayarHutangHandlerInstance.GetNextNoTransaksi)
			bayarHutang.GET("/pembelian-sisa/:supplier_id", bayarHutangHandlerInstance.GetPembelianSisa)
			bayarHutang.GET("/:id", bayarHutangHandlerInstance.GetByID)
			bayarHutang.POST("", bayarHutangHandlerInstance.Create)
			bayarHutang.PUT("/:id", bayarHutangHandlerInstance.Update)
			bayarHutang.DELETE("/:id", bayarHutangHandlerInstance.Delete)
		}
	}

	return r
}
