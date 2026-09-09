# Catatan Proyek MrCoCo

Untuk AI baru: Baca file ini dulu sebelum mulai kerja.

---

## Tech Stack

- **Go 1.27** + **Gin** (web) + **GORM** (ORM) + **PostgreSQL**
- Port: `8080`
- DB: `localhost:5432/devdb`
- Module name: `mrcoco`

## Arsitektur

```
Handler → Service → Repository → GORM → PostgreSQL
```

Semua ada di `internal/`:
```
internal/
├── config/database/db.go          # Koneksi DB
├── handler/
│   ├── masterdata/                # Handler master data
│   ├── pembelian/                 # Handler pembelian
│   └── bayar_hutang/              # Handler bayar hutang
├── helper/
│   ├── pagination.go              # GetPagination() — parse page/limit
│   ├── search.go                  # ApplySmartSearch() — split kata, ILIKE, AND
│   └── no_transaksi.go            # GenerateNoTransaksi() — auto no transaksi
├── models/
│   ├── masterdata.go              # Semua struct entity + TableName()
│   ├── date.go                    # Custom Date type
│   ├── pembelian/                 # Model pembelian
│   └── bayar_hutang/              # Model bayar hutang
├── repository/
│   ├── masterdata/                # Repository master data
│   ├── pembelian/                 # Repository pembelian
│   └── bayar_hutang/              # Repository bayar hutang
├── routes/routes.go               # Registrasi route
└── service/
    ├── masterdata/                # Service master data
    ├── pembelian/                 # Service pembelian
    └── bayar_hutang/              # Service bayar hutang
```

## Entity & Table

| Struct | Tabel | Catatan |
|--------|-------|---------|
| DataSatuan | `data_satuan` | |
| DataKategori | `data_kategori` | |
| DataJenis | `data_jenis` | |
| DataSubKlasifikasi | `data_sub_klasifikasi` | |
| DataKlasifikasi | `data_klasifikasi` | |
| DataItem | `data_item` | FK → Jenis, Kategori, Satuan, Rekening, Konsumen |
| DataGrup | `data_grup` | |
| DataKonsumen | `data_konsumen` | FK → Grup |
| DataRekening | `data_rekening` | FK → Klasifikasi, SubKlasifikasi |
| DataSupplier | `data_supplier` | |
| DataGudang | `data_gudang` | |
| DataPembelian | `data_pembelian` | FK → Supplier, Gudang, Rekening |
| DataPembelianItem | `data_pembelian_item` | FK → Pembelian, Item |
| DataBayarHutang | `data_bayar_hutang` | FK → Supplier, Gudang, Rekening |
| DataBayarHutangItem | `data_bayar_hutang_item` | FK → BayarHutang, Pembelian |

## Pattern: Foreign Key Relation

Ikuti pola ini saat tambah relasi baru:

**Model:**
```go
// FK field — pakai *uint supaya data lama yang NULL ga error
KategoriID uint          `gorm:"not null" json:"kategori_id" binding:"required"`
Kategori   *DataKategori `gorm:"foreignKey:KategoriID" json:"kategori,omitempty"`
```

**Response custom (hanya nama):**
Kalau mau response relasi cuma tampilkan nama (bukan full object), bikin:
1. Type `json:"-"` di field relasi di struct utama
2. Struct response terpisah
3. Method `MarshalJSON()` di struct utama

Contoh: `DataRekening` — `json:"-"` di field `Klasifikasi`/`SubKlasifikasi`, lalu custom MarshalJSON yang return `"klasifikasi": "Bank"` (string).

**Repository:**
```go
r.db.Preload("Klasifikasi").Preload("SubKlasifikasi").Find(&rekenings)
```

**Handler (setelah Create/Update):**
Fetch ulang by ID setelah create/update supaya data relasi ter-load:
```go
h.service.Create(&rekening)
created, _ := h.service.GetByID(rekening.ID)
c.JSON(201, gin.H{"data": created})
```

## Pattern: Smart Search

Helper di `internal/helper/search.go` — `ApplySmartSearch(query, field, keyword)`

- Split keyword jadi kata-kata (space)
- Setiap kata jadi `field ILIKE '%kata%'` dengan AND logic
- Contoh: `"per woi"` → `WHERE nama ILIKE '%per%' AND nama ILIKE '%woi%'`
- Semua repo sudah pakai ini

## Pattern: Pagination

Helper di `internal/helper/pagination.go` — `GetPagination(c)`

- Query params: `page` (default 1), `limit` (default 10, max 100)
- Return `Pagination{Offset, Limit, Page, Total}`
- Response format:
  ```json
  { "data": [...], "pagination": { "page": 1, "limit": 10, "total": 50 } }
  ```

## Pattern: No Transaksi Auto-Generated

Helper di `internal/helper/no_transaksi.go` — `GenerateNoTransaksi(db, tablePrefix, tanggal)`

- Format: `{PREFIX}-yyyymmNNNN`
- Prefix: `P` untuk pembelian, `H` untuk bayar hutang
- Contoh: `P-2026090001`, `H-2026090001`
- Increment per bulan

## API Endpoints

Semua di bawah `/api/`:

### Master Data

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/{resource}` | Get all (paginated) |
| GET | `/{resource}/search?nama=` | Smart search (paginated) |
| GET | `/{resource}/total` | Total count |
| GET | `/{resource}/:id` | Get by ID |
| POST | `/{resource}` | Create |
| PUT | `/{resource}/:id` | Update |
| DELETE | `/{resource}/:id` | Hard delete |

**Resources:** `satuan`, `kategori`, `jenis`, `sub-klasifikasi`, `klasifikasi`, `item`, `grup`, `konsumen`, `rekening`, `supplier`, `gudang`

### Item Spesifik

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/item/konsumen/:konsumen_id` | Get all items by konsumen ID |

### Pembelian

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/pembelian` | History dengan filter |
| GET | `/pembelian/total` | Total dengan filter |
| GET | `/pembelian/next-no` | Preview no transaksi |
| GET | `/pembelian/:id` | Detail dengan items |
| POST | `/pembelian` | Create |
| PUT | `/pembelian/:id` | Update (replace items) |
| DELETE | `/pembelian/:id` | Delete (cascade) |

### Bayar Hutang

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/bayar-hutang` | History dengan filter |
| GET | `/bayar-hutang/total` | Total dengan filter |
| GET | `/bayar-hutang/next-no` | Preview no transaksi |
| GET | `/bayar-hutang/pembelian-sisa/:supplier_id` | Pembelian dengan sisa hutang |
| GET | `/bayar-hutang/:id` | Detail dengan items |
| POST | `/bayar-hutang` | Create |
| PUT | `/bayar-hutang/:id` | Update |
| DELETE | `/bayar-hutang/:id` | Delete (restore sisa) |

## Route Registration

Semua route didaftarkan di `internal/routes/routes.go`. Kalau tambah handler baru, daftarkan di situ.

**Note:** Route spesifik (misal `/item/konsumen/:konsumen_id`) didaftarkan sebelum route parameter (`/:id`) supaya tidak conflict.

## Migration

Auto-migrate ada di `main.go` function `runMigrations()`. Kalau tambah model baru, tambahkan ke list `db.AutoMigrate(...)`.

## Yang Perlu Diingat

- **Hard delete** — ga ada soft delete, data benar-benar dihapus
- **Pointer relations** — selalu pakai `*DataType` untuk field relasi, supaya validasi ga bermasalah
- **Binding required** — pakai `binding:"required"` di request body fields
- **Nullable FK** — kalau FK optional (bisa null), pakai `*uint` tanpa `gorm:"not null"` dan tanpa `binding:"required"`
- **Timestamp** — `created_at` dan `updated_at` otomatis diisi GORM
- **Unique index** — field `nama` di semua entity pakai `uniqueIndex`
- **Search** — semua search pakai smart search (ILIKE + AND), bukan exact match
- **Route order** — route spesifik (misal `/konsumen/:id`) didaftarkan sebelum route parameter (`/:id`) supaya tidak conflict
- **No transaksi** — auto-generated berdasarkan tanggal (format: `P-yyyymmNNNN` atau `H-yyyymmNNNN`)
- **Bayar Hutang** — `grand_total`, `jumlah_bayar`, `sisa` dihitung oleh frontend
- **Bayar Hutang** — saat bayar, `jumlah_uang_sisa` di pembelian berkurang dan `jumlah_uang_muka` bertambah
- **Bayar Hutang** — update: kalau `items` tidak dikirim, cuma update header. Kalau `items` dikirim, old items di-restore dulu baru new items di-apply

## Dokumen

- `docs/api_specs.md` — Dokumentasi API lengkap (frontend-friendly)
- `docs/FOR_FRONTEND.md` — Panduan khusus untuk frontend developer
- `docs/summary.md` — Log perubahan / changelog
- `docs/note.md` — Catatan ini (untuk AI baru)
- `docs/flutter_client_prompt.md` — Prompt untuk Flutter client development
