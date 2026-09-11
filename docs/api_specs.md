# API Specifications — MrCoCo

**Base URL:** `http://localhost:8080/api`
**Framework:** Go (Gin) + GORM + PostgreSQL
**Authentication:** Tidak ada (semua endpoint public)

---

## Daftar Isi

1. [Informasi Umum](#info-umum)
2. [Satuan](#satuan)
3. [Kategori](#kategori)
4. [Jenis](#jenis)
5. [Sub-Klasifikasi](#sub-klasifikasi)
6. [Klasifikasi](#klasifikasi)
7. [Grup](#grup)
8. [Konsumen](#konsumen)
9. [Rekening](#rekening)
10. [Supplier](#supplier)
11. [Gudang](#gudang)
12. [Item](#item)
13. [Pembelian](#pembelian)
14. [Bayar Hutang](#bayar-hutang)
15. [Pre Order](#pre-order)
16. [Error Handling](#error-handling)
17. [Quick Reference](#quick-reference)

---

## Informasi Umum {#info-umum}

### Response Format

Semua response menggunakan envelope format:

```json
{
  "data": [...],
  "pagination": { "page": 1, "limit": 10, "total": 25 }
}
```

### Pagination

Semua endpoint `GET /api/{resource}` (list) mendukung pagination:

| Param | Default | Max | Keterangan |
|-------|---------|-----|------------|
| `page` | 1 | - | Nomor halaman (1-indexed) |
| `limit` | 10 | 100 | Jumlah data per halaman |

**Contoh:**
```
GET /api/satuan?page=1&limit=20
GET /api/item?page=2&limit=50
```

### Smart Search

Setiap kata dalam search query dicari secara terpisah dengan AND logic menggunakan PostgreSQL `ILIKE`:

| Query | Yang Dicari | Contoh Match |
|-------|-------------|--------------|
| `"laptop asus"` | "laptop" **DAN** "asus" | "Laptop ASUS", "ASUS Laptop" |
| `"budi jakarta"` | "budi" **DAN** "jakarta" | "Budi Jakarta", "Budi di Jakarta" |

### HTTP Status Codes

| Code | Arti |
|------|------|
| 200 | Sukses (GET, PUT, DELETE) |
| 201 | Berhasil buat data baru (POST) |
| 400 | Request salah / invalid ID / validation error |
| 404 | Data tidak ditemukan |
| 500 | Error server |

### Response Format Detail

| Endpoint Type | Response |
|---|---|
| List (GET all) | `{ "data": [...], "pagination": { "page", "limit", "total" } }` |
| Detail (GET by ID) | `{ "data": { ... } }` |
| Count (GET total) | `{ "total": N }` |
| Next No (GET next-no) | `{ "no_transaksi": "P-YYYYMM0001" }` |
| Search | `{ "data": [...], "pagination": { "page", "limit", "total" } }` |
| Delete success | `{ "message": "... deleted successfully" }` |
| Error | `{ "error": "..." }` |

---

## Satuan {#satuan}

**Endpoint:** `/api/satuan`

Data satuan unit pengukuran (Box, Pcs, Unit, dll).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama satuan |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/satuan` | Ambil semua satuan (paginated) |
| 2 | **GET** | `/api/satuan/search` | Cari satuan (smart search by `nama`) |
| 3 | **GET** | `/api/satuan/total` | Total jumlah satuan |
| 4 | **GET** | `/api/satuan/:id` | Detail satuan by ID |
| 5 | **POST** | `/api/satuan` | Buat satuan baru |
| 6 | **PUT** | `/api/satuan/:id` | Update satuan |
| 7 | **DELETE** | `/api/satuan/:id` | Hapus satuan |

### Search

**Query params:**

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/satuan/search?nama=box
GET /api/satuan/search?nama=box+pcs&page=1&limit=5
```

### Request Body (Create/Update)

```json
{
  "nama": "Box",
  "catatan": "Satuan dalam box"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Box",
    "catatan": "Satuan dalam box",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Kategori {#kategori}

**Endpoint:** `/api/kategori`

Data kategori item (Elektronik, Furniture, dll).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama kategori |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/kategori` | Ambil semua kategori (paginated) |
| 2 | **GET** | `/api/kategori/search` | Cari kategori (smart search by `nama`) |
| 3 | **GET** | `/api/kategori/total` | Total jumlah kategori |
| 4 | **GET** | `/api/kategori/:id` | Detail kategori by ID |
| 5 | **POST** | `/api/kategori` | Buat kategori baru |
| 6 | **PUT** | `/api/kategori/:id` | Update kategori |
| 7 | **DELETE** | `/api/kategori/:id` | Hapus kategori |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

### Request Body (Create/Update)

```json
{
  "nama": "Elektronik",
  "catatan": "Semua jenis elektronik"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Elektronik",
    "catatan": "Semua jenis elektronik",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Jenis {#jenis}

**Endpoint:** `/api/jenis`

Data jenis item (Laptop, HP, Printer, dll).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama jenis |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/jenis` | Ambil semua jenis (paginated) |
| 2 | **GET** | `/api/jenis/search` | Cari jenis (smart search by `nama`) |
| 3 | **GET** | `/api/jenis/total` | Total jumlah jenis |
| 4 | **GET** | `/api/jenis/:id` | Detail jenis by ID |
| 5 | **POST** | `/api/jenis` | Buat jenis baru |
| 6 | **PUT** | `/api/jenis/:id` | Update jenis |
| 7 | **DELETE** | `/api/jenis/:id` | Hapus jenis |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

### Request Body (Create/Update)

```json
{
  "nama": "Laptop",
  "catatan": "Semua jenis laptop"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Laptop",
    "catatan": "Semua jenis laptop",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Sub-Klasifikasi {#sub-klasifikasi}

**Endpoint:** `/api/sub-klasifikasi`

Data sub-klasifikasi rekening.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama sub-klasifikasi |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/sub-klasifikasi` | Ambil semua sub-klasifikasi (paginated) |
| 2 | **GET** | `/api/sub-klasifikasi/search` | Cari sub-klasifikasi (smart search by `nama`) |
| 3 | **GET** | `/api/sub-klasifikasi/total` | Total jumlah sub-klasifikasi |
| 4 | **GET** | `/api/sub-klasifikasi/:id` | Detail sub-klasifikasi by ID |
| 5 | **POST** | `/api/sub-klasifikasi` | Buat sub-klasifikasi baru |
| 6 | **PUT** | `/api/sub-klasifikasi/:id` | Update sub-klasifikasi |
| 7 | **DELETE** | `/api/sub-klasifikasi/:id` | Hapus sub-klasifikasi |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

### Request Body (Create/Update)

```json
{
  "nama": "Kas",
  "catatan": "Sub-klasifikasi kas"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Kas",
    "catatan": "Sub-klasifikasi kas",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Klasifikasi {#klasifikasi}

**Endpoint:** `/api/klasifikasi`

Data klasifikasi rekening.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama klasifikasi |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/klasifikasi` | Ambil semua klasifikasi (paginated) |
| 2 | **GET** | `/api/klasifikasi/search` | Cari klasifikasi (smart search by `nama`) |
| 3 | **GET** | `/api/klasifikasi/total` | Total jumlah klasifikasi |
| 4 | **GET** | `/api/klasifikasi/:id` | Detail klasifikasi by ID |
| 5 | **POST** | `/api/klasifikasi` | Buat klasifikasi baru |
| 6 | **PUT** | `/api/klasifikasi/:id` | Update klasifikasi |
| 7 | **DELETE** | `/api/klasifikasi/:id` | Hapus klasifikasi |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

### Request Body (Create/Update)

```json
{
  "nama": "Aset Lancar",
  "catatan": "Klasifikasi aset lancar"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Aset Lancar",
    "catatan": "Klasifikasi aset lancar",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Grup {#grup}

**Endpoint:** `/api/grup`

Data grup konsumen (perorangan, perusahaan, dll).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama grup |
| `no_hp` | string | Tidak | - | 20 | Nomor HP |
| `email` | string | Tidak | - | 100 | Alamat email |
| `alamat` | string | Tidak | - | text | Alamat lengkap |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/grup` | Ambil semua grup (paginated) |
| 2 | **GET** | `/api/grup/search` | Cari grup (multi-field search) |
| 3 | **GET** | `/api/grup/total` | Total jumlah grup |
| 4 | **GET** | `/api/grup/:id` | Detail grup by ID |
| 5 | **POST** | `/api/grup` | Buat grup baru |
| 6 | **PUT** | `/api/grup/:id` | Update grup |
| 7 | **DELETE** | `/api/grup/:id` | Hapus grup |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `no_hp` | Filter by no_hp (ILIKE) |
| `alamat` | Filter by alamat (ILIKE) |
| `status` | Filter by status (true/false) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/grup/search?nama=toko
GET /api/grup/search?nama=toko&alamat=jakarta&page=1&limit=5
```

### Request Body (Create/Update)

```json
{
  "nama": "Toko Berkah",
  "no_hp": "081234567890",
  "email": "berkah@email.com",
  "alamat": "Jl. Merdeka No. 10, Jakarta",
  "catatan": "Grup toko"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `no_hp`: optional, string, max 20 char
- `email`: optional, string, max 100 char
- `alamat`: optional, text
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Toko Berkah",
    "no_hp": "081234567890",
    "email": "berkah@email.com",
    "alamat": "Jl. Merdeka No. 10, Jakarta",
    "catatan": "Grup toko",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Konsumen {#konsumen}

**Endpoint:** `/api/konsumen`

Data konsumen/pelanggan dengan relasi ke grup.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | Tidak | 100 | Nama konsumen |
| `no_hp` | string | Tidak | - | 20 | Nomor HP |
| `email` | string | Tidak | - | 100 | Alamat email |
| `alamat` | string | Tidak | - | text | Alamat lengkap |
| `grup_id` | uint | **Ya** | - | - | FK ke data_grup |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `status` | bool | Tidak | - | - | Default: true (aktif) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Response Fields (flattened)

```json
{
  "id": 1,
  "nama": "Budi Santoso",
  "no_hp": "081234567890",
  "email": "budi@email.com",
  "alamat": "Jl. Sudirman No. 5, Bandung",
  "grup_id": 1,
  "grup": "Toko Berkah",
  "catatan": "Pelanggan tetap",
  "status": true,
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

> Field `grup` adalah nama grup (string) yang di-flatten dari relasi `grup_id`.

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/konsumen` | Ambil semua konsumen (paginated) |
| 2 | **GET** | `/api/konsumen/search` | Cari konsumen (multi-field search) |
| 3 | **GET** | `/api/konsumen/total` | Total jumlah konsumen |
| 4 | **GET** | `/api/konsumen/:id` | Detail konsumen by ID |
| 5 | **POST** | `/api/konsumen` | Buat konsumen baru |
| 6 | **PUT** | `/api/konsumen/:id` | Update konsumen |
| 7 | **DELETE** | `/api/konsumen/:id` | Hapus konsumen |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `alamat` | Filter by alamat (ILIKE) |
| `no_hp` | Filter by no_hp (ILIKE) |
| `grup_id` | Filter by grup ID (exact match) |
| `status` | Filter by status (true/false) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/konsumen/search?nama=budi
GET /api/konsumen/search?nama=budi&grup_id=1&status=true
GET /api/konsumen/search?alamat=jakarta&page=1&limit=20
```

### Request Body (Create/Update)

```json
{
  "nama": "Budi Santoso",
  "no_hp": "081234567890",
  "email": "budi@email.com",
  "alamat": "Jl. Sudirman No. 5, Bandung",
  "grup_id": 1,
  "catatan": "Pelanggan tetap",
  "status": true
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char
- `grup_id`: **required**, uint (harus ada di data_grup)
- `no_hp`: optional, string, max 20 char
- `email`: optional, string, max 100 char
- `alamat`: optional, text
- `catatan`: optional, text
- `status`: optional, bool (default: true)

---

## Rekening {#rekening}

**Endpoint:** `/api/rekening`

Data rekening bank dengan relasi ke sub-klasifikasi dan klasifikasi.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `no_rek` | string | **Ya** | **Ya** | 50 | Nomor rekening |
| `nama` | string | **Ya** | Tidak | 100 | Nama rekening |
| `sub_klasifikasi_id` | *uint | **Ya** | - | - | FK ke data_sub_klasifikasi |
| `klasifikasi_id` | *uint | **Ya** | - | - | FK ke data_klasifikasi |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Response Fields (flattened via custom MarshalJSON)

```json
{
  "id": 1,
  "no_rek": "1234567890",
  "nama": "Bank Mandiri",
  "sub_klasifikasi_id": 1,
  "sub_klasifikasi": "Kas",
  "klasifikasi_id": 1,
  "klasifikasi": "Aset Lancar",
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

> Field `sub_klasifikasi` dan `klasifikasi` adalah nama (string) yang di-flatten dari relasi. Menggunakan custom `MarshalJSON`.

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/rekening` | Ambil semua rekening (paginated) |
| 2 | **GET** | `/api/rekening/search` | Cari rekening (smart search by `nama`) |
| 3 | **GET** | `/api/rekening/total` | Total jumlah rekening |
| 4 | **GET** | `/api/rekening/:id` | Detail rekening by ID |
| 5 | **POST** | `/api/rekening` | Buat rekening baru |
| 6 | **PUT** | `/api/rekening/:id` | Update rekening |
| 7 | **DELETE** | `/api/rekening/:id` | Hapus rekening |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/rekening/search?nama=mandiri
```

### Request Body (Create/Update)

```json
{
  "no_rek": "1234567890",
  "nama": "Bank Mandiri",
  "sub_klasifikasi_id": 1,
  "klasifikasi_id": 1
}
```

**Validasi:**
- `no_rek`: **required**, string, max 50 char, harus unik
- `nama`: **required**, string, max 100 char
- `sub_klasifikasi_id`: **required**, *uint (harus ada di data_sub_klasifikasi)
- `klasifikasi_id`: **required**, *uint (harus ada di data_klasifikasi)

---

## Supplier {#supplier}

**Endpoint:** `/api/supplier`

Data supplier/vendor.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | Tidak | 100 | Nama supplier |
| `no_hp` | string | Tidak | - | 20 | Nomor HP |
| `email` | string | Tidak | - | 100 | Alamat email |
| `alamat` | string | Tidak | - | text | Alamat lengkap |
| `bank` | string | Tidak | - | 50 | Nama bank |
| `no_rek` | string | Tidak | - | 50 | Nomor rekening |
| `atas_nama` | string | Tidak | - | 100 | Nama pemilik rekening |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `status` | bool | Tidak | - | - | Default: true (aktif) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/supplier` | Ambil semua supplier (paginated) |
| 2 | **GET** | `/api/supplier/search` | Cari supplier (multi-field search) |
| 3 | **GET** | `/api/supplier/total` | Total jumlah supplier |
| 4 | **GET** | `/api/supplier/:id` | Detail supplier by ID |
| 5 | **POST** | `/api/supplier` | Buat supplier baru |
| 6 | **PUT** | `/api/supplier/:id` | Update supplier |
| 7 | **DELETE** | `/api/supplier/:id` | Hapus supplier |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `alamat` | Filter by alamat (ILIKE) |
| `status` | Filter by status (true/false) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/supplier/search?nama=pt+jaya
GET /api/supplier/search?nama=jaya&alamat=jakarta&status=true
```

### Request Body (Create/Update)

```json
{
  "nama": "PT Supplier Jaya",
  "no_hp": "081298765432",
  "email": "jaya@supplier.com",
  "alamat": "Jl. Industri No. 20, Surabaya",
  "bank": "BCA",
  "no_rek": "9876543210",
  "atas_nama": "PT Supplier Jaya",
  "catatan": "Supplier utama",
  "status": true
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char
- `no_hp`: optional, string, max 20 char
- `email`: optional, string, max 100 char
- `alamat`: optional, text
- `bank`: optional, string, max 50 char
- `no_rek`: optional, string, max 50 char
- `atas_nama`: optional, string, max 100 char
- `catatan`: optional, text
- `status`: optional, bool (default: true)

---

## Gudang {#gudang}

**Endpoint:** `/api/gudang`

Data gudang penyimpanan.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `nama` | string | **Ya** | **Ya** | 100 | Nama gudang |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/gudang` | Ambil semua gudang (paginated) |
| 2 | **GET** | `/api/gudang/search` | Cari gudang (smart search by `nama`) |
| 3 | **GET** | `/api/gudang/total` | Total jumlah gudang |
| 4 | **GET** | `/api/gudang/:id` | Detail gudang by ID |
| 5 | **POST** | `/api/gudang` | Buat gudang baru |
| 6 | **PUT** | `/api/gudang/:id` | Update gudang |
| 7 | **DELETE** | `/api/gudang/:id` | Hapus gudang |

### Search

| Param | Keterangan |
|-------|------------|
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

### Request Body (Create/Update)

```json
{
  "nama": "Gudang Utama",
  "catatan": "Gudang pusat"
}
```

**Validasi:**
- `nama`: **required**, string, max 100 char, harus unik
- `catatan`: optional, text

### Response (Create/Update)

```json
{
  "data": {
    "id": 1,
    "nama": "Gudang Utama",
    "catatan": "Gudang pusat",
    "created_at": "2026-09-05T10:00:00Z",
    "updated_at": "2026-09-05T10:00:00Z"
  }
}
```

---

## Item {#item}

**Endpoint:** `/api/item`

Data item/barang dengan relasi ke jenis, kategori, satuan, rekening, dan konsumen.

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `kode` | string | **Ya** | **Ya** | 50 | Kode item (unik) |
| `nama` | string | **Ya** | Tidak | 200 | Nama item |
| `jenis_id` | uint | **Ya** | - | - | FK ke data_jenis |
| `kategori_id` | uint | **Ya** | - | - | FK ke data_kategori |
| `satuan_id` | uint | **Ya** | - | - | FK ke data_satuan |
| `rekening_id` | *uint | **Ya** | - | - | FK ke data_rekening |
| `konsumen_id` | *uint | Tidak | - | - | FK ke data_konsumen (nullable) |
| `harga_beli` | int | **Ya** | - | - | Harga beli item |
| `harga_jual` | int | **Ya** | - | - | Harga jual item |
| `status` | bool | Tidak | - | - | Default: true (aktif) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Response Fields (flattened)

```json
{
  "id": 1,
  "kode": "ELG-001",
  "nama": "Laptop ASUS",
  "jenis_id": 1,
  "jenis": "Laptop",
  "kategori_id": 1,
  "kategori": "Elektronik",
  "satuan_id": 1,
  "satuan": "Unit",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "konsumen_id": null,
  "konsumen": "",
  "harga_beli": 5000000,
  "harga_jual": 7500000,
  "status": true,
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

> Field `jenis`, `kategori`, `satuan`, `rekening`, `konsumen` adalah nama (string) yang di-flatten dari relasi.

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/item` | Ambil semua item (paginated) |
| 2 | **GET** | `/api/item/search` | Cari item (multi-field search + filter) |
| 3 | **GET** | `/api/item/total` | Total jumlah item |
| 4 | **GET** | `/api/item/konsumen/:konsumen_id` | Item by konsumen (paginated) |
| 5 | **GET** | `/api/item/:id` | Detail item by ID |
| 6 | **POST** | `/api/item` | Buat item baru |
| 7 | **PUT** | `/api/item/:id` | Update item |
| 8 | **DELETE** | `/api/item/:id` | Hapus item |

### Search

| Param | Keterangan |
|-------|------------|
| `kode` | Filter by kode item (ILIKE) |
| `nama` | Smart search on nama (ILIKE, AND logic per word) |
| `kategori_id` | Filter by kategori ID (exact match) |
| `konsumen_id` | Filter by konsumen ID (exact match) |
| `status` | Filter by status (true/false) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/item/search?nama=laptop
GET /api/item/search?nama=laptop+asus&kategori_id=1&status=true
GET /api/item/search?kode=ELG&page=1&limit=20
```

### Get Item by Konsumen

```
GET /api/item/konsumen/:konsumen_id?page=1&limit=10
```

Response: list item (paginated) yang dimiliki oleh konsumen tertentu.

### Request Body (Create/Update)

```json
{
  "kode": "ELG-001",
  "nama": "Laptop ASUS",
  "jenis_id": 1,
  "kategori_id": 1,
  "satuan_id": 1,
  "rekening_id": 1,
  "konsumen_id": null,
  "harga_beli": 5000000,
  "harga_jual": 7500000,
  "status": true
}
```

**Validasi:**
- `kode`: **required**, string, max 50 char, harus unik
- `nama`: **required**, string, max 200 char
- `jenis_id`: **required**, uint (harus ada di data_jenis)
- `kategori_id`: **required**, uint (harus ada di data_kategori)
- `satuan_id`: **required**, uint (harus ada di data_satuan)
- `rekening_id`: **required**, *uint (harus ada di data_rekening)
- `konsumen_id`: optional, *uint (nullable, bisa null)
- `harga_beli`: **required**, int (min: 0)
- `harga_jual`: **required**, int (min: 0)
- `status`: optional, bool (default: true)

---

## Pembelian {#pembelian}

**Endpoint:** `/api/pembelian`

Data pembelian barang dari supplier. Setiap pembelian memiliki items (detail barang yang dibeli).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `no_transaksi` | string | Auto | **Ya** | 20 | Auto-generated: `P-YYYYMMNNNN` |
| `tanggal` | date | **Ya** | - | - | Format: `YYYY-MM-DD` |
| `supplier_id` | uint | **Ya** | - | - | FK ke data_supplier |
| `gudang_id` | uint | **Ya** | - | - | FK ke data_gudang |
| `tgl_jatuh_tempo` | date | Tidak | - | - | Format: `YYYY-MM-DD` |
| `rekening_id` | *uint | Tidak | - | - | FK ke data_rekening |
| `grand_total` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `jumlah_uang_muka` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `jumlah_uang_sisa` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `items` | array | **Ya** | - | - | Array of items (detail pembelian) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Item Fields (DataPembelianItem)

| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| `item_id` | uint | **Ya** | FK ke data_item |
| `qty` | int | **Ya** | Jumlah qty |
| `harga` | int | **Ya** | Harga per unit |
| `diskon` | int | Tidak | Default: 0 |
| `subtotal` | int | Tidak | Default: 0 |

### Response List (PembelianResponse)

```json
{
  "id": 1,
  "no_transaksi": "P-2026090001",
  "tanggal": "2026-09-05",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "gudang_id": 1,
  "gudang": "Gudang Utama",
  "tgl_jatuh_tempo": "2026-10-05",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "total_qty": 10,
  "grand_total": 50000000,
  "jumlah_uang_muka": 10000000,
  "jumlah_uang_sisa": 40000000,
  "catatan": "",
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

> Field `total_qty` adalah total qty dari semua items (dihitung backend).

### Response Detail (PembelianDetailResponse)

```json
{
  "id": 1,
  "no_transaksi": "P-2026090001",
  "tanggal": "2026-09-05",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "gudang_id": 1,
  "gudang": "Gudang Utama",
  "tgl_jatuh_tempo": "2026-10-05",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "total_qty": 10,
  "grand_total": 50000000,
  "jumlah_uang_muka": 10000000,
  "jumlah_uang_sisa": 40000000,
  "catatan": "",
  "items": [
    {
      "id": 1,
      "item_id": 1,
      "kode": "ELG-001",
      "nama": "Laptop ASUS",
      "satuan": "Unit",
      "qty": 10,
      "harga": 5000000,
      "diskon": 0,
      "subtotal": 50000000
    }
  ],
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

### Response Create (PembelianCreateResponse)

```json
{
  "id": 1,
  "no_transaksi": "P-2026090001",
  "tanggal": "2026-09-05",
  "grand_total": 50000000,
  "jumlah_uang_muka": 10000000,
  "jumlah_uang_sisa": 40000000,
  "catatan": ""
}
```

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/pembelian` | History pembelian (paginated + filter) |
| 2 | **GET** | `/api/pembelian/total` | Total pembelian (dengan filter) |
| 3 | **GET** | `/api/pembelian/next-no` | Preview no transaksi berikutnya |
| 4 | **GET** | `/api/pembelian/:id` | Detail pembelian (dengan items) |
| 5 | **POST** | `/api/pembelian` | Buat pembelian baru |
| 6 | **PUT** | `/api/pembelian/:id` | Update pembelian |
| 7 | **DELETE** | `/api/pembelian/:id` | Hapus pembelian (cascade delete items) |

### Filter History

| Param | Keterangan |
|-------|------------|
| `tanggal_awal` | Filter tanggal transaksi. Jika hanya ini diisi → exact match. Jika ada `tanggal_akhir` → range (>=) |
| `tanggal_akhir` | Filter tanggal jatuh tempo. Jika hanya ini diisi → exact match. Jika ada `tanggal_awal` → range (<=) |
| `jatuh_tempo_awal` | Filter tanggal jatuh tempo. Jika hanya ini diisi → exact match. Jika ada `jatuh_tempo_akhir` → range (>=) |
| `jatuh_tempo_akhir` | Filter tanggal jatuh tempo. Jika hanya ini diisi → exact match. Jika ada `jatuh_tempo_awal` → range (<=) |
| `supplier_id` | Filter per supplier ID |
| `gudang_id` | Filter per gudang ID |
| `rekening_id` | Filter per rekening ID |
| `no_transaksi` | Cari by no transaksi (ILIKE) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/pembelian?tanggal_awal=2026-09-01&tanggal_akhir=2026-09-30
GET /api/pembelian?tanggal_awal=2026-09-05
GET /api/pembelian?jatuh_tempo_awal=2026-10-01&jatuh_tempo_akhir=2026-10-31
GET /api/pembelian?supplier_id=1&page=1&limit=20
```

### Get Next No Transaksi

```
GET /api/pembelian/next-no?tanggal=2026-09-15
```

**Response:**
```json
{ "no_transaksi": "P-2026090001" }
```

Format auto-generated: `P-YYYYMMNNNN` (contoh: `P-2026090001`)

### Request Body (Create)

```json
{
  "tanggal": "2026-09-05",
  "supplier_id": 1,
  "gudang_id": 1,
  "tgl_jatuh_tempo": "2026-10-05",
  "rekening_id": 1,
  "grand_total": 50000000,
  "jumlah_uang_muka": 10000000,
  "jumlah_uang_sisa": 40000000,
  "catatan": "Pembelian awal bulan",
  "items": [
    {
      "item_id": 1,
      "qty": 10,
      "harga": 5000000,
      "diskon": 0,
      "subtotal": 50000000
    }
  ]
}
```

**Validasi:**
- `tanggal`: **required**, date (YYYY-MM-DD)
- `supplier_id`: **required**, uint (harus ada di data_supplier)
- `gudang_id`: **required**, uint (harus ada di data_gudang)
- `tgl_jatuh_tempo`: optional, date (YYYY-MM-DD)
- `rekening_id`: optional, *uint
- `grand_total`: optional, int (dihitung frontend)
- `jumlah_uang_muka`: optional, int (dihitung frontend)
- `jumlah_uang_sisa`: optional, int (dihitung frontend)
- `catatan`: optional, text
- `items`: **required**, array, minimal 1 item
- `items[].item_id`: **required**, uint
- `items[].qty`: **required**, int (min: 1)
- `items[].harga`: **required**, int (min: 0)
- `items[].diskon`: optional, int (default: 0)
- `items[].subtotal`: optional, int (default: 0)

**Backend logic:**
- `no_transaksi` auto-generated oleh backend
- `grand_total`, `jumlah_uang_muka`, `jumlah_uang_sisa` dihitung oleh frontend
- Saat update, items lama dihapus dan diganti dengan items baru

### Request Body (Update)

Sama dengan create. Semua items lama akan diganti dengan items baru.

---

## Bayar Hutang {#bayar-hutang}

**Endpoint:** `/api/bayar-hutang`

Data pembayaran sisa hutang ke supplier. Setiap bayar hutang memiliki items (detail pembayaran per pembelian).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `no_transaksi` | string | Auto | **Ya** | 20 | Auto-generated: `H-YYYYMMNNNN` |
| `tanggal` | date | **Ya** | - | - | Format: `YYYY-MM-DD` |
| `supplier_id` | uint | **Ya** | - | - | FK ke data_supplier |
| `gudang_id` | uint | **Ya** | - | - | FK ke data_gudang |
| `rekening_id` | *uint | Tidak | - | - | FK ke data_rekening |
| `grand_total` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `jumlah_bayar` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `sisa` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `items` | array | **Ya** | - | - | Array of items (detail pembayaran) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Item Fields (DataBayarHutangItem)

| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| `pembelian_id` | uint | **Ya** | FK ke data_pembelian |
| `jumlah_bayar` | int | **Ya** | Jumlah bayar untuk pembelian ini |

### Response List (BayarHutangResponse)

```json
{
  "id": 1,
  "no_transaksi": "H-2026090001",
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "gudang_id": 1,
  "gudang": "Gudang Utama",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "grand_total": 800000,
  "jumlah_bayar": 700000,
  "sisa": 100000,
  "catatan": "Bayar cicilan",
  "created_at": "2026-09-15T10:00:00Z",
  "updated_at": "2026-09-15T10:00:00Z"
}
```

### Response Detail (BayarHutangDetailResponse)

```json
{
  "id": 1,
  "no_transaksi": "H-2026090001",
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "gudang_id": 1,
  "gudang": "Gudang Utama",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "grand_total": 800000,
  "jumlah_bayar": 700000,
  "sisa": 100000,
  "catatan": "Bayar cicilan",
  "items": [
    {
      "id": 1,
      "pembelian_id": 1,
      "no_transaksi": "P-2026090001",
      "jumlah_bayar": 400000
    },
    {
      "id": 2,
      "pembelian_id": 2,
      "no_transaksi": "P-2026090002",
      "jumlah_bayar": 300000
    }
  ],
  "created_at": "2026-09-15T10:00:00Z",
  "updated_at": "2026-09-15T10:00:00Z"
}
```

### Response Create (BayarHutangCreateResponse)

```json
{
  "id": 1,
  "no_transaksi": "H-2026090001",
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "grand_total": 800000,
  "jumlah_bayar": 700000,
  "sisa": 100000,
  "catatan": "Bayar cicilan"
}
```

### Response Pembelian Sisa

```json
{
  "id": 1,
  "no_transaksi": "P-2026090001",
  "tanggal": "2026-09-05",
  "grand_total": 500000,
  "jumlah_uang_muka": 100000,
  "jumlah_uang_sisa": 400000
}
```

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/bayar-hutang` | History bayar hutang (paginated + filter) |
| 2 | **GET** | `/api/bayar-hutang/total` | Total bayar hutang (dengan filter) |
| 3 | **GET** | `/api/bayar-hutang/next-no` | Preview no transaksi berikutnya |
| 4 | **GET** | `/api/bayar-hutang/pembelian-sisa/:supplier_id` | Lihat pembelian dengan sisa hutang |
| 5 | **GET** | `/api/bayar-hutang/:id` | Detail bayar hutang (dengan items) |
| 6 | **POST** | `/api/bayar-hutang` | Buat bayar hutang baru |
| 7 | **PUT** | `/api/bayar-hutang/:id` | Update bayar hutang |
| 8 | **DELETE** | `/api/bayar-hutang/:id` | Hapus bayar hutang (restore sisa) |

### Filter History

| Param | Keterangan |
|-------|------------|
| `tanggal_awal` | Filter tanggal. Jika hanya ini diisi → exact match. Jika ada `tanggal_akhir` → range (>=) |
| `tanggal_akhir` | Filter tanggal. Jika hanya ini diisi → exact match. Jika ada `tanggal_awal` → range (<=) |
| `supplier_id` | Filter per supplier ID |
| `gudang_id` | Filter per gudang ID |
| `rekening_id` | Filter per rekening ID |
| `no_transaksi` | Cari by no transaksi (ILIKE) |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/bayar-hutang?tanggal_awal=2026-09-01&tanggal_akhir=2026-09-30
GET /api/bayar-hutang?tanggal_awal=2026-09-15
GET /api/bayar-hutang?supplier_id=1&page=1&limit=20
```

### Get Next No Transaksi

```
GET /api/bayar-hutang/next-no?tanggal=2026-09-15
```

**Response:**
```json
{ "no_transaksi": "H-2026090001" }
```

Format auto-generated: `H-YYYYMMNNNN` (contoh: `H-2026090001`)

### Get Pembelian Sisa (Belum Lunas)

```
GET /api/bayar-hutang/pembelian-sisa/:supplier_id
GET /api/bayar-hutang/pembelian-sisa/:supplier_id?search=P-202609
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "no_transaksi": "P-2026090001",
      "tanggal": "2026-09-05",
      "grand_total": 500000,
      "jumlah_uang_muka": 100000,
      "jumlah_uang_sisa": 400000
    }
  ]
}
```

> Hanya menampilkan pembelian yang masih memiliki `jumlah_uang_sisa > 0`. Query `search` opsional untuk filter by `no_transaksi`.

### Cara Kerja Bayar Hutang

Saat bayar hutang, backend akan:

1. **Kurangi** `jumlah_uang_sisa` di pembelian sesuai `jumlah_bayar`
2. **Tambah** `jumlah_uang_muka` di pembelian sesuai `jumlah_bayar`

**Contoh:**
```
Sebelum bayar:
  P-001: grand_total=500000, uang_muka=100000, sisa=400000
  P-002: grand_total=300000, uang_muka=0, sisa=300000

Bayar hutang: bayar 700000 (P-001=400000, P-002=300000)

Sesudah bayar:
  P-001: uang_muka=500000, sisa=0 (LUNAS)
  P-002: uang_muka=300000, sisa=0 (LUNAS)
```

### Request Body (Create)

```json
{
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "gudang_id": 1,
  "rekening_id": 1,
  "grand_total": 700000,
  "jumlah_bayar": 700000,
  "sisa": 0,
  "catatan": "Bayar lunas",
  "items": [
    { "pembelian_id": 1, "jumlah_bayar": 400000 },
    { "pembelian_id": 2, "jumlah_bayar": 300000 }
  ]
}
```

**Validasi:**
- `tanggal`: **required**, date (YYYY-MM-DD)
- `supplier_id`: **required**, uint (harus ada di data_supplier)
- `gudang_id`: **required**, uint (harus ada di data_gudang)
- `rekening_id`: optional, *uint
- `grand_total`: optional, int (dihitung frontend)
- `jumlah_bayar`: optional, int (dihitung frontend)
- `sisa`: optional, int (dihitung frontend)
- `catatan`: optional, text
- `items`: **required**, array, minimal 1 item
- `items[].pembelian_id`: **required**, uint (harus ada di data_pembelian)
- `items[].jumlah_bayar`: **required**, int (min: 1)

**Business rules:**
- `jumlah_bayar` per pembelian tidak boleh melebihi `jumlah_uang_sisa` pembelian tersebut
- `jumlah_uang_muka` tidak boleh negatif setelah pembayaran

### Request Body (Update)

**Mode 1: Update header saja (tanpa items)**
```json
{
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "gudang_id": 1,
  "rekening_id": 1,
  "grand_total": 800000,
  "jumlah_bayar": 700000,
  "sisa": 100000,
  "catatan": "Catatan updated"
}
```
→ Items tidak berubah, uang_muka/uang_sisa tidak berubah.

**Mode 2: Update items + header**
```json
{
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "gudang_id": 1,
  "rekening_id": 1,
  "grand_total": 400000,
  "jumlah_bayar": 400000,
  "sisa": 0,
  "catatan": "Bayar lunas",
  "items": [
    { "pembelian_id": 1, "jumlah_bayar": 400000 }
  ]
}
```
→ Old items di-restore dulu (jumlah_uang_sisa dan jumlah_uang_muka dikembalikan ke kondisi sebelum bayar hutang ini), lalu new items di-apply.

---

## Pre Order {#pre-order}

**Endpoint:** `/api/pre-order`

Data pre order dari konsumen. Setiap pre order memiliki items (detail barang yang dipesan).

### Fields

| Field | Type | Required | Unique | Max Length | Keterangan |
|-------|------|----------|--------|------------|------------|
| `id` | uint | - | - | - | Auto-generated primary key |
| `tanggal` | date | **Ya** | - | - | Format: `YYYY-MM-DD` |
| `konsumen_id` | uint | **Ya** | - | - | FK ke data_konsumen |
| `grand_total` | int | Tidak | - | - | Default: 0 (dihitung frontend) |
| `catatan` | string | Tidak | - | text | Keterangan tambahan |
| `items` | array | **Ya** | - | - | Array of items (detail pre order) |
| `created_at` | datetime | - | - | - | Auto-generated |
| `updated_at` | datetime | - | - | - | Auto-generated |

### Item Fields (DataPreOrderItem)

| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| `item_id` | uint | **Ya** | FK ke data_item |
| `qty` | int | **Ya** | Jumlah qty |
| `harga` | int | **Ya** | Harga per unit |
| `subtotal` | int | Tidak | Default: 0 |

### Response List (PreOrderResponse)

```json
{
  "id": 1,
  "tanggal": "2026-09-05",
  "konsumen_id": 1,
  "konsumen": "Budi Santoso",
  "total_qty": 5,
  "grand_total": 25000000,
  "catatan": "",
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

> Field `total_qty` adalah total qty dari semua items (dihitung backend).

### Response Detail (PreOrderDetailResponse)

```json
{
  "id": 1,
  "tanggal": "2026-09-05",
  "konsumen_id": 1,
  "konsumen": "Budi Santoso",
  "total_qty": 5,
  "grand_total": 25000000,
  "catatan": "",
  "items": [
    {
      "id": 1,
      "item_id": 1,
      "kode": "ELG-001",
      "nama": "Laptop ASUS",
      "satuan": "Unit",
      "qty": 3,
      "harga": 5000000,
      "subtotal": 15000000
    },
    {
      "id": 2,
      "item_id": 2,
      "kode": "ELG-002",
      "nama": "Mouse Logitech",
      "satuan": "Pcs",
      "qty": 2,
      "harga": 5000000,
      "subtotal": 10000000
    }
  ],
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

### Response Create (PreOrderCreateResponse)

```json
{
  "id": 1,
  "tanggal": "2026-09-05",
  "grand_total": 25000000,
  "catatan": ""
}
```

### Endpoints

| # | Method | Path | Keterangan |
|---|--------|------|------------|
| 1 | **GET** | `/api/pre-order` | History pre order (paginated + filter) |
| 2 | **GET** | `/api/pre-order/total` | Total pre order (dengan filter) |
| 3 | **GET** | `/api/pre-order/:id` | Detail pre order (dengan items) |
| 4 | **POST** | `/api/pre-order` | Buat pre order baru |
| 5 | **PUT** | `/api/pre-order/:id` | Update pre order |
| 6 | **DELETE** | `/api/pre-order/:id` | Hapus pre order (cascade delete items) |

### Filter History

| Param | Keterangan |
|-------|------------|
| `tanggal` | Filter by tanggal (YYYY-MM-DD) |
| `konsumen_id` | Filter per konsumen ID |
| `page` | Halaman (default: 1) |
| `limit` | Jumlah data (default: 10, max: 100) |

```
GET /api/pre-order?tanggal=2026-09-05
GET /api/pre-order?konsumen_id=1&page=1&limit=20
GET /api/pre-order?tanggal=2026-09-05&konsumen_id=1
```

> Untuk search item saat membuat pre-order, gunakan endpoint `GET /api/item/search?nama=xxx`.

### Request Body (Create)

```json
{
  "tanggal": "2026-09-05",
  "konsumen_id": 1,
  "grand_total": 25000000,
  "catatan": "Pre order bulan depan",
  "items": [
    {
      "item_id": 1,
      "qty": 3,
      "harga": 5000000,
      "subtotal": 15000000
    },
    {
      "item_id": 2,
      "qty": 2,
      "harga": 5000000,
      "subtotal": 10000000
    }
  ]
}
```

**Validasi:**
- `tanggal`: **required**, date (YYYY-MM-DD)
- `konsumen_id`: **required**, uint (harus ada di data_konsumen)
- `grand_total`: optional, int (dihitung frontend)
- `catatan`: optional, text
- `items`: **required**, array, minimal 1 item
- `items[].item_id`: **required**, uint (harus ada di data_item)
- `items[].qty`: **required**, int (min: 1)
- `items[].harga`: **required**, int (min: 0)
- `items[].subtotal`: optional, int (default: 0)

**Backend logic:**
- `grand_total` dihitung oleh frontend

### Request Body (Update)

Sama dengan create. Semua items lama akan diganti dengan items baru.

---

## Error Handling {#error-handling}

### Status Codes

| Code | Arti |
|------|------|
| 200 | Sukses (GET, PUT, DELETE) |
| 201 | Berhasil buat data baru (POST) |
| 400 | Request salah / invalid ID / validation error |
| 404 | Data tidak ditemukan |
| 500 | Error server |

### Contoh Error Response

```json
{ "error": "Invalid ID" }
```

```json
{ "error": "Satuan not found" }
```

```json
{ "error": "Konsumen not found" }
```

```json
{ "error": "jumlah bayar melebihi sisa hutang untuk pembelian P-2026090001: sisa=100000, bayar=200000" }
```

```json
{ "error": "Key: 'DataSatuan.Nama' Error:Field validation for 'Nama' failed on the 'required' tag" }
```

---

## Quick Reference {#quick-reference}

### Total Endpoints: 99

| Module | GET | POST | PUT | DELETE | Total |
|--------|-----|------|-----|--------|-------|
| Satuan | 4 | 1 | 1 | 1 | **7** |
| Kategori | 4 | 1 | 1 | 1 | **7** |
| Jenis | 4 | 1 | 1 | 1 | **7** |
| Sub-Klasifikasi | 4 | 1 | 1 | 1 | **7** |
| Klasifikasi | 4 | 1 | 1 | 1 | **7** |
| Grup | 4 | 1 | 1 | 1 | **7** |
| Konsumen | 4 | 1 | 1 | 1 | **7** |
| Rekening | 4 | 1 | 1 | 1 | **7** |
| Supplier | 4 | 1 | 1 | 1 | **7** |
| Gudang | 4 | 1 | 1 | 1 | **7** |
| Item | 5 | 1 | 1 | 1 | **8** |
| Pembelian | 4 | 1 | 1 | 1 | **7** |
| Bayar Hutang | 5 | 1 | 1 | 1 | **8** |
| Pre Order | 3 | 1 | 1 | 1 | **6** |
| **TOTAL** | **57** | **14** | **14** | **14** | **99** |

### Resources

`satuan`, `kategori`, `jenis`, `sub-klasifikasi`, `klasifikasi`, `grup`, `konsumen`, `rekening`, `supplier`, `gudang`, `item`, `pembelian`, `bayar-hutang`, `pre-order`

### Master Data CRUD Pattern

| Fungsi | Endpoint |
|--------|----------|
| Ambil semua | `GET /api/{resource}` |
| Cari data | `GET /api/{resource}/search?nama=...` |
| Ambil total | `GET /api/{resource}/total` |
| Ambil by ID | `GET /api/{resource}/:id` |
| Buat baru | `POST /api/{resource}` |
| Update | `PUT /api/{resource}/:id` |
| Hapus | `DELETE /api/{resource}/:id` |

### Transaksi Endpoints

| Fungsi | Endpoint |
|--------|----------|
| Next no pembelian | `GET /api/pembelian/next-no?tanggal=YYYY-MM-DD` |
| Next no bayar hutang | `GET /api/bayar-hutang/next-no?tanggal=YYYY-MM-DD` |
| Pembelian sisa hutang | `GET /api/bayar-hutang/pembelian-sisa/:supplier_id` |
| Item by konsumen | `GET /api/item/konsumen/:konsumen_id` |

### Field Mapping: Request vs Response

| Resource | Request Body | Response (flattened) |
|----------|-------------|---------------------|
| Konsumen | `grup_id` (uint) | `grup` (string nama grup) |
| Rekening | `sub_klasifikasi_id`, `klasifikasi_id` | `sub_klasifikasi` (string), `klasifikasi` (string) |
| Item | `jenis_id`, `kategori_id`, `satuan_id`, `rekening_id`, `konsumen_id` | `jenis`, `kategori`, `satuan`, `rekening`, `konsumen` (string) |
| Pembelian | `supplier_id`, `gudang_id`, `rekening_id` | `supplier`, `gudang`, `rekening` (string) |
| Bayar Hutang | `supplier_id`, `gudang_id`, `rekening_id` | `supplier`, `gudang`, `rekening` (string) |
| Pre Order | `konsumen_id` | `konsumen` (string nama konsumen) |
