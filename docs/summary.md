# Summary — MrCoCo Development (5-9 September 2026)

## Yang Sudah Dikerjakan

### 1. Smart Search

**File:** `internal/helper/search.go`

- Buat helper `ApplySmartSearch()` yang split search query jadi kata-kata terpisah
- Setiap kata dijadikan `ILIKE` condition dengan AND logic
- Contoh: `"per woi"` → `WHERE nama ILIKE '%per%' AND nama ILIKE '%woi%'`
- Jadi "per pcs woi" bakal muncul karena ada "per" DAN "woi"
- Semua repository sudah pakai helper ini

**Repository yang sudah pakai smart search:**
- satuan_repo.go (field: `nama`)
- kategori_repo.go (field: `nama`)
- jenis_repo.go (field: `nama`)
- sub_klasifikasi_repo.go (field: `nama`)
- klasifikasi_repo.go (field: `nama`)
- rekening_repo.go (field: `nama`)
- item_repo.go (fields: `kode`, `nama`)
- grup_repo.go (fields: `nama`, `no_hp`, `alamat`)
- konsumen_repo.go (fields: `nama`, `alamat`, `no_hp`)
- supplier_repo.go (fields: `nama`, `alamat`)
- gudang_repo.go (field: `nama`)

### 2. Hard Delete (Permanent Delete)

**Models:** `internal/models/masterdata.go`

- Hapus field `DeletedAt gorm.DeletedAt` dari semua model
- Hapus import `gorm.io/gorm` dari models (ga kepake lagi)
- Tabel sekarang ga ada kolom `deleted_at` yang ga kepake

**Repositories:**

- Hapus `Unscoped()` dari semua method `Delete()`
- Tanpa field `DeletedAt`, GORM otomatis hard delete

### 3. Table Names (Tanpa Auto Pluralize)

**Models:** `internal/models/masterdata.go`

- Tambah `TableName()` method di semua model
- GORM ga lagi auto tambah "s" di belakang nama tabel

| Model | Tabel |
|-------|-------|
| DataSatuan | `data_satuan` |
| DataKategori | `data_kategori` |
| DataJenis | `data_jenis` |
| DataSubKlasifikasi | `data_sub_klasifikasi` |
| DataKlasifikasi | `data_klasifikasi` |
| DataItem | `data_item` |
| DataGrup | `data_grup` |
| DataKonsumen | `data_konsumen` |
| DataRekening | `data_rekening` |
| DataSupplier | `data_supplier` |
| DataGudang | `data_gudang` |
| DataPembelian | `data_pembelian` |
| DataPembelianItem | `data_pembelian_item` |

### 4. Pagination

**File:** `internal/helper/pagination.go`

- Helper `GetPagination()` parse query params `page` (default: 1) dan `limit` (default: 10, max: 100)
- Otomatis hitung `offset`
- Response format:
  ```json
  {
    "data": [...],
    "pagination": {
      "page": 1,
      "limit": 10,
      "total": 50
    }
  }
  ```
- Berlaku untuk semua endpoint `GET /api/{resource}` dan `GET /api/{resource}/search`

### 5. Fix Validation Error (Nested Struct)

**Models:** `internal/models/masterdata.go`

- Ubah field relasi dari struct ke pointer:
  - `DataItem.Kategori` → `*DataKategori`
  - `DataItem.Satuan` → `*DataSatuan`
  - `DataKonsumen.Grup` → `*DataGrup`
- Tanpa pointer, GORM tetap validate nested struct walaupun client ga kirim field-nya
- Dengan pointer, kalau field ga dikirim, nil dan ga di-validate

### 6. Helper Dipindah ke Lokasi Shared

**Lama:** `internal/repository/masterdata/helper.go`
**Baru:** `internal/helper/search.go`

- Fungsi di-export: `applySmartSearch` → `ApplySmartSearch`
- Semua repository import `"mrcoco/internal/helper"` dan panggil `helper.ApplySmartSearch()`
- Bisa dipakai di domain lain selain master data

### 7. Rekening — Foreign Key Klasifikasi & Sub Klasifikasi

**Sebelum:** `DataRekening` menyimpan `sub_klasifikasi` dan `klasifikasi` sebagai teks biasa (`varchar(100)`), tidak ada relasi ke tabel `data_klasifikasi` dan `data_sub_klasifikasi`.

**Sesudah:** Menggunakan foreign key (`*uint`) dengan GORM relation.

**Model `DataRekening`:**
```go
SubKlasifikasiID *uint               `json:"sub_klasifikasi_id" binding:"required"`
SubKlasifikasi   *DataSubKlasifikasi `gorm:"foreignKey:SubKlasifikasiID" json:"-"`
KlasifikasiID    *uint               `json:"klasifikasi_id" binding:"required"`
Klasifikasi      *DataKlasifikasi    `gorm:"foreignKey:KlasifikasiID" json:"-"`
```

**Custom MarshalJSON:**
- Response `klasifikasi` dan `sub_klasifikasi` hanya mengembalikan **nama** (string), bukan full object
- Contoh response:
  ```json
  {
    "sub_klasifikasi_id": 1,
    "sub_klasifikasi": "Bank Besar",
    "klasifikasi_id": 2,
    "klasifikasi": "Bank"
  }
  ```

### 8. DataItem — Foreign Key Rekening & Konsumen

**Model `DataItem`:**
```go
RekeningID *uint         `json:"rekening_id"`
Rekening   *DataRekening `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
KonsumenID *uint         `json:"konsumen_id"`
Konsumen   *DataKonsumen `gorm:"foreignKey:KonsumenID" json:"konsumen,omitempty"`
```

- `KonsumenID` tidak pakai `binding:"required"` supaya optional (bisa null)
- Semua query sudah `.Preload("Rekening").Preload("Konsumen")`

### 9. Endpoint Baru — Item by Konsumen

**Endpoint:** `GET /api/item/konsumen/:konsumen_id`

- Get semua item berdasarkan konsumen ID (paginated)
- Route didaftarkan sebelum `/:id` supaya tidak conflict

### 10. DataGudang — Master Data Baru

**Model:** `internal/models/masterdata.go`

```go
type DataGudang struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Nama      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"nama" binding:"required"`
    Catatan   string    `gorm:"type:text" json:"catatan"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

**File:**
- `internal/repository/masterdata/gudang_repo.go` — CRUD + search
- `internal/service/masterdata/gudang_service.go` — service layer
- `internal/handler/masterdata/gudang_handler.go` — handler layer

**Endpoint:** Sama seperti master data lainnya (GET, search, total, create, update, delete)

### 11. Custom Date Type

**File:** `internal/models/date.go`

Custom type `Date` yang wrap `time.Time` dengan fitur:

- **UnmarshalJSON:** Terima format `"YYYY-MM-DD"` dan `"YYYY-MM-DDTHH:MM:SSZ"`
- **MarshalJSON:** Selalu return format `"YYYY-MM-DD"`
- **Value (Valuer):** Simpan ke DB sebagai date string
- **Scan (Scanner):** Baca dari DB (handle `time.Time`, `[]byte`, `string`)
- **GORM tag:** `type:date` supaya column PostgreSQL bertipe `date`

Digunakan di model `DataPembelian` untuk field `Tanggal` dan `TglJatuhTempo`.

### 12. Pembelian Module (CRUD Lengkap)

**Model:** `internal/models/pembelian.go`

```go
type DataPembelian struct {
    ID              uint                  `gorm:"primaryKey" json:"id"`
    NoTransaksi     string                `gorm:"type:varchar(20);not null;uniqueIndex" json:"no_transaksi"`
    Tanggal         Date                  `gorm:"type:date;not null" json:"tanggal" binding:"required"`
    SupplierID      uint                  `gorm:"not null" json:"supplier_id" binding:"required"`
    Supplier        *DataSupplier         `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
    GudangID        uint                  `gorm:"not null" json:"gudang_id" binding:"required"`
    Gudang          *DataGudang           `gorm:"foreignKey:GudangID" json:"gudang,omitempty"`
    TglJatuhTempo   Date                  `gorm:"type:date" json:"tgl_jatuh_tempo"`
    RekeningID      *uint                 `json:"rekening_id"`
    Rekening        *DataRekening         `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
    GrandTotal      int                   `gorm:"not null;default:0" json:"grand_total"`
    JumlahUangMuka  int                   `gorm:"not null;default:0" json:"jumlah_uang_muka"`
    JumlahUangSisa  int                   `gorm:"not null;default:0" json:"jumlah_uang_sisa"`
    Items           []DataPembelianItem   `gorm:"foreignKey:PembelianID" json:"items,omitempty"`
    CreatedAt       time.Time             `json:"created_at"`
    UpdatedAt       time.Time             `json:"updated_at"`
}

type DataPembelianItem struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    PembelianID  uint      `gorm:"not null" json:"pembelian_id"`
    ItemID       uint      `gorm:"not null" json:"item_id" binding:"required"`
    Item         *DataItem `gorm:"foreignKey:ItemID" json:"item,omitempty"`
    Qty          int       `gorm:"not null;default:0" json:"qty" binding:"required"`
    Harga        int       `gorm:"not null;default:0" json:"harga" binding:"required"`
    Diskon       int       `gorm:"not null;default:0" json:"diskon"`
    Subtotal     int       `gorm:"not null;default:0" json:"subtotal"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
```

**Response:** `internal/models/pembelian_response.go`
- `PembelianResponse` — flatten nama supplier, gudang, rekening
- `PembelianItemResponse` — flatten nama item, satuan
- `PembelianDetailResponse` — header + items
- `PembelianCreateResponse` — id, no_transaksi, tanggal, grand_total, uang_muka, sisa

**Helper:** `internal/helper/no_transaksi.go`
- `GenerateNoTransaksi(db, tanggal)` — generate no transaksi berdasarkan tanggal transaksi
- Pola: `YYYYMM` + 4 digit increment (contoh: `2026090001`, `2026100001`)
- Query max no transaksi di bulan tersebut → increment

**Repository:** `internal/repository/pembelian/pembelian_repo.go`
- `FindAll` — dengan filter: start_date, end_date, supplier_id, no_transaksi
- `FindByID` — dengan preload items + item + satuan
- `Count` — dengan filter yang sama
- `Create` — insert header + items dalam transaction
- `Update` — delete items lama, insert items baru dalam transaction
- `Delete` — delete items + header dalam transaction
- `DeleteItemsByPembelianID` — untuk update

**Service:** `internal/service/pembelian/pembelian_service.go`
- `Create` — auto-generate no transaksi berdasarkan tanggal sebelum insert
- `Update` — delete items lama, insert items baru
- `GetNextNoTransaksi(tanggal)` — preview no transaksi berikutnya

**Handler:** `internal/handler/pembelian/pembelian_handler.go`
- `GET /api/pembelian` — history dengan filter
- `GET /api/pembelian/total` — jumlah transaksi dengan filter
- `GET /api/pembelian/next-no?tanggal=YYYY-MM-DD` — preview no transaksi
- `GET /api/pembelian/:id` — detail dengan items
- `POST /api/pembelian` — create (auto no transaksi)
- `PUT /api/pembelian/:id` — update (replace items)
- `DELETE /api/pembelian/:id` — delete (cascade)

### 14. Bayar Hutang Module (CRUD Lengkap + Update Pembelian Sisa)

**Model:** `internal/models/bayar_hutang/bayar_hutang.go`

```go
type DataBayarHutang struct {
    ID          uint                       `gorm:"primaryKey" json:"id"`
    NoTransaksi string                     `gorm:"type:varchar(20);not null;uniqueIndex" json:"no_transaksi"`
    Tanggal     models.Date                `gorm:"type:date;not null" json:"tanggal" binding:"required"`
    SupplierID  uint                       `gorm:"not null" json:"supplier_id" binding:"required"`
    Supplier    *masterdata.DataSupplier   `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
    GudangID    uint                       `gorm:"not null" json:"gudang_id" binding:"required"`
    Gudang      *masterdata.DataGudang     `gorm:"foreignKey:GudangID" json:"gudang,omitempty"`
    RekeningID  *uint                      `json:"rekening_id"`
    Rekening    *masterdata.DataRekening   `gorm:"foreignKey:RekeningID" json:"rekening,omitempty"`
    GrandTotal  int                        `gorm:"not null;default:0" json:"grand_total"`
    JumlahBayar int                        `gorm:"not null;default:0" json:"jumlah_bayar"`
    Sisa        int                        `gorm:"not null;default:0" json:"sisa"`
    Catatan     string                     `gorm:"type:text" json:"catatan"`
    Items       []DataBayarHutangItem      `gorm:"foreignKey:BayarHutangID" json:"items,omitempty"`
    CreatedAt   time.Time                  `json:"created_at"`
    UpdatedAt   time.Time                  `json:"updated_at"`
}

type DataBayarHutangItem struct {
    ID            uint                      `gorm:"primaryKey" json:"id"`
    BayarHutangID uint                      `gorm:"not null" json:"bayar_hutang_id"`
    PembelianID   uint                      `gorm:"not null" json:"pembelian_id"`
    Pembelian     *pembelian.DataPembelian  `gorm:"foreignKey:PembelianID" json:"pembelian,omitempty"`
    SisaBayar     int                       `gorm:"not null;default:0" json:"sisa_bayar"`
    CreatedAt     time.Time                 `json:"created_at"`
    UpdatedAt     time.Time                 `json:"updated_at"`
}
```

**Helper:** `internal/helper/no_transaksi.go`
- `GenerateNoTransaksi(db, tablePrefix, tanggal)` — generate no transaksi dengan prefix
- Pola: `P-` untuk pembelian, `H-` untuk bayar hutang
- Contoh: `P-2026090001`, `H-2026090001`

**Repository:** `internal/repository/bayar_hutang/bayar_hutang_repo.go`
- `FindAll` — dengan filter: start_date, end_date, supplier_id, no_transaksi
- `FindByID` — dengan preload supplier, gudang, rekening, items, items.pembelian
- `Count` — dengan filter yang sama
- `Create/Update/Delete` — CRUD dalam transaction
- `FindPembelianSisaBySupplier` — ambil pembelian dengan jumlah_uang_sisa > 0
- `SearchPembelianSisaBySupplier` — search pembelian sisa by no transaksi

**Service:** `internal/service/bayar_hutang/bayar_hutang_service.go`
- `Create` — auto-generate no transaksi (H-), hitung grand_total & sisa otomatis, kurangi jumlah_uang_sisa pembelian
- `Update` — restore jumlah_uang_sisa lama, apply jumlah_uang_sisa baru
- `Delete` — restore jumlah_uang_sisa semua pembelian terkait
- `GetPembelianSisaBySupplier` — ambil pembelian yang masih ada sisa

**Handler:** `internal/handler/bayar_hutang/bayar_hutang_handler.go`
- `GET /api/bayar-hutang` — history dengan filter
- `GET /api/bayar-hutang/total` — jumlah transaksi dengan filter
- `GET /api/bayar-hutang/next-no?tanggal=YYYY-MM-DD` — preview no transaksi
- `GET /api/bayar-hutang/pembelian-sisa/:supplier_id` — ambil pembelian sisa by supplier
- `GET /api/bayar-hutang/:id` — detail dengan items
- `POST /api/bayar-hutang` — create (auto no transaksi, auto grand_total)
- `PUT /api/bayar-hutang/:id` — update (replace items, restore & apply sisa)
- `DELETE /api/bayar-hutang/:id` — delete (restore sisa pembelian)

**Flow Bayar Hutang:**
```
1. Pilih supplier → GET /api/bayar-hutang/pembelian-sisa/:supplier_id
   → Response: list pembelian dengan jumlah_uang_sisa > 0

2. Pilih pembelian, input jumlah_bayar
   → POST /api/bayar-hutang
   → grand_total = total sisa_bayar items (otomatis)
   → sisa = grand_total - jumlah_bayar (otomatis)
   → jumlah_uang_sisa di pembelian berkurang

3. Setelah bayar, pembelian yang sisa = 0 tidak muncul di list pembelian sisa
```

### 15. Auto-Migrate Update

**File:** `main.go`

```go
err := db.AutoMigrate(
    &models.DataSatuan{},
    &models.DataKategori{},
    &models.DataJenis{},
    &models.DataSubKlasifikasi{},
    &models.DataKlasifikasi{},
    &models.DataItem{},
    &models.DataGrup{},
    &models.DataKonsumen{},
    &models.DataRekening{},
    &models.DataSupplier{},
    &models.DataGudang{},
    &models.DataPembelian{},
    &models.DataPembelianItem{},
    &models.DataBayarHutang{},       // BARU
    &models.DataBayarHutangItem{},   // BARU
)
```

### 16. Routes Update

**File:** `internal/routes/routes.go`

- Import handler, repository, service pembelian dengan alias
- Import handler, repository, service bayar_hutang dengan alias
- Inisialisasi Gudang repo/service/handler
- Inisialisasi Pembelian repo/service/handler
- Inisialisasi Bayar Hutang repo/service/handler
- Daftarkan routes gudang + pembelian + bayar-hutang

### 17. API Docs Update

**File:** `docs/api_specs.md`

- Section 11: Gudang (CRUD endpoints)
- Section 12: Pembelian (CRUD endpoints + filter + next-no)
- Section 13: Bayar Hutang (CRUD endpoints + filter + next-no + pembelian sisa)
- Quick Reference: tambah gudang, pembelian, bayar-hutang

### 18. Documentation Improvements

**API Specs (`docs/api_specs.md`):**
- Restructure menjadi lebih rapi dan frontend-friendly
- Tambah Daftar Isi untuk navigasi mudah
- Group endpoint berdasarkan fungsi (Master Data, Item, Pembelian, Bayar Hutang)
- Tambah tabel Quick Reference
- Format konsisten untuk semua endpoint
- Contoh response yang lebih jelas

**Frontend Guide (`docs/FOR_FRONTEND.md`):**
- Update dengan struktur yang lebih baik
- Tambah contoh implementasi (flow pembelian & bayar hutang)
- Tambah contoh kode JavaScript untuk distribusi pembayaran
- Tambah tabel endpoint yang lebih lengkap
- Penjelasan cara kerja bayar hutang yang lebih jelas

### 19. Flutter Client Prompt

**File:** `docs/flutter_client_prompt.md`

- Responsive layout (desktop sidebar + mobile drawer)
- Master data landing page dengan grid buttons
- Transaksi landing page
- 11 master data CRUD pages (satu pola, field berbeda)
- Pembelian history dengan filter (date range, supplier, no transaksi)
- Pembelian create/edit page (dedicated full page)
- Bayar Hutang history dengan filter (date range, supplier, no transaksi)
- Bayar Hutang create/edit page (dedicated full page)
- Custom Date type handling (YYYY-MM-DD)
- Auto no transaksi (fetch dari next-no endpoint)
- Item search dan add multiple items
- Auto calculation (subtotal, grand total, uang sisa)
- State management: Provider
- Folder structure lengkap

---

## File yang Berubah/Ditambah

| # | File | Status | Perubahan |
|---|------|--------|-----------|
| 1 | `internal/helper/search.go` | Baru | ApplySmartSearch() |
| 2 | `internal/helper/pagination.go` | Baru | GetPagination() + Pagination struct |
| 3 | `internal/helper/no_transaksi.go` | Ubah | GenerateNoTransaksi() dengan prefix P-/H- |
| 4 | `internal/models/masterdata.go` | Ubah | Hapus DeletedAt, pointer relations, TableName(), DataGudang |
| 5 | `internal/models/pembelian/pembelian.go` | Baru | DataPembelian + DataPembelianItem |
| 6 | `internal/models/pembelian/pembelian_response.go` | Baru | Response structs + converters |
| 7 | `internal/models/bayar_hutang/bayar_hutang.go` | Baru | DataBayarHutang + DataBayarHutangItem |
| 8 | `internal/models/bayar_hutang/bayar_hutang_response.go` | Baru | Response structs + converters |
| 9 | `internal/models/date.go` | Baru | Custom Date type (Valuer/Scanner/Marshaler) |
| 10 | `internal/repository/masterdata/gudang_repo.go` | Baru | Gudang CRUD |
| 11 | `internal/service/masterdata/gudang_service.go` | Baru | Gudang service |
| 12 | `internal/handler/masterdata/gudang_handler.go` | Baru | Gudang handler |
| 13 | `internal/repository/pembelian/pembelian_repo.go` | Baru | Pembelian CRUD + filters |
| 14 | `internal/service/pembelian/pembelian_service.go` | Baru | Pembelian service |
| 15 | `internal/handler/pembelian/pembelian_handler.go` | Baru | Pembelian handler |
| 16 | `internal/repository/bayar_hutang/bayar_hutang_repo.go` | Baru | Bayar Hutang CRUD + pembelian sisa |
| 17 | `internal/service/bayar_hutang/bayar_hutang_service.go` | Baru | Bayar Hutang service + update pembelian sisa |
| 18 | `internal/handler/bayar_hutang/bayar_hutang_handler.go` | Baru | Bayar Hutang handler |
| 19 | `internal/routes/routes.go` | Ubah | Tambah gudang + pembelian + bayar-hutang routes |
| 20 | `main.go` | Ubah | Auto-migrate DataGudang, DataPembelian, DataPembelianItem, DataBayarHutang, DataBayarHutangItem |
| 21 | `docs/api_specs.md` | Ubah | Tambah Section 11 Gudang + Section 12 Pembelian + Section 13 Bayar Hutang |
| 22 | `docs/flutter_client_prompt.md` | Baru | Flutter client development prompt |
| 23 | `docs/summary.md` | Ubah | Update summary |
| 24 | `docs/api_specs.md` | Ubah | Restructure menjadi lebih rapi dan frontend-friendly |
| 25 | `docs/FOR_FRONTEND.md` | Ubah | Update dengan struktur lebih baik + contoh implementasi |

---

## Architecture Recap

```
Handler (parse request, validate, call service)
  → Service (business logic, auto-generate no transaksi)
    → Repository (query with filters, transaction for create/update/delete)
      → GORM → PostgreSQL
```

**Pembelian flow:**
```
Create:
  Frontend → POST /api/pembelian (tanggal, supplier_id, items[])
    → Service: GenerateNoTransaksi(P, tanggal) → set NoTransaksi
      → Repository: DB Transaction (insert header + items)
        → Response: { id, no_transaksi, tanggal, grand_total, uang_muka, sisa }

Update:
  Frontend → PUT /api/pembelian/:id (tanggal, supplier_id, items[])
    → Repository: DB Transaction (delete old items → update header → insert new items)
      → Response: { header + items with relations }

Delete:
  Frontend → DELETE /api/pembelian/:id
    → Repository: DB Transaction (delete items → delete header)
      → Response: { message: "deleted" }
```

**Bayar Hutang flow:**
```
Ambil Pembelian Sisa:
  Frontend → GET /api/bayar-hutang/pembelian-sisa/:supplier_id
    → Repository: query pembelian WHERE jumlah_uang_sisa > 0
      → Response: list pembelian dengan sisa

Create:
  Frontend → POST /api/bayar-hutang (tanggal, supplier_id, gudang_id, jumlah_bayar, items[])
    → Service: GenerateNoTransaksi(H, tanggal) → set NoTransaksi
    → Service: grand_total = sum(sisa_bayar), sisa = grand_total - jumlah_bayar
    → Service: DB Transaction:
        1. Insert bayar_hutang header + items
        2. Kurangi jumlah_uang_sisa di tiap pembelian
      → Response: { id, no_transaksi, grand_total, jumlah_bayar, sisa }

Update:
  Frontend → PUT /api/bayar-hutang/:id
    → Service: DB Transaction:
        1. Restore jumlah_uang_sisa pembelian lama
        2. Delete items lama
        3. Insert items baru + kurangi jumlah_uang_sisa
      → Response: { header + items with relations }

Delete:
  Frontend → DELETE /api/bayar-hutang/:id
    → Service: DB Transaction:
        1. Restore jumlah_uang_sisa semua pembelian terkait
        2. Delete items + header
      → Response: { message: "deleted" }
```

---

## Teknologi

- Go 1.27 + Gin + GORM + PostgreSQL
- Clean Architecture: Handler → Service → Repository
- Smart Search (split kata, AND logic, ILIKE)
- Pagination (page/limit, total count)
- Hard Delete (permanent)
- Pointer relations (fix validation)
- Foreign Key relations (Rekening → Klasifikasi/SubKlasifikasi, Item → Rekening/Konsumen, Pembelian → Supplier/Gudang/Rekening, BayarHutang → Supplier/Gudang/Rekening)
- Custom MarshalJSON (response hanya nama, bukan full object)
- Custom Date type (Valuer/Scanner/MarshalJSON — date-only format)
- Auto-generated no transaksi dengan prefix (P- untuk pembelian, H- untuk bayar hutang)
- Transaction untuk create/update/delete (pembelian & bayar hutang)
- Auto-update jumlah_uang_sisa di pembelian saat bayar hutang
- Endpoint baru: Get items by konsumen ID, Pembelian CRUD + next-no, Bayar Hutang CRUD + next-no + pembelian sisa
