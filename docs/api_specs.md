# API Specifications — MrCoCo

**Base URL:** `http://localhost:8080/api`

---

## Daftar Isi

1. [Informasi Umum](#info-umum)
2. [Master Data](#master-data)
3. [Item](#item)
4. [Pembelian](#pembelian)
5. [Bayar Hutang](#bayar-hutang)
6. [Error Handling](#error-handling)
7. [Quick Reference](#quick-reference)

---

## Informasi Umum {#info-umum}

### Response Format

Semua response menggunakan format:

```json
{
  "data": [...],
  "pagination": { ... }
}
```

### Pagination

Semua endpoint `GET /api/{resource}` mendukung pagination:

| Param | Default | Max | Keterangan |
|-------|---------|-----|------------|
| `page` | 1 | - | Nomor halaman |
| `limit` | 10 | 100 | Jumlah data per halaman |

**Contoh:**
```
GET /api/satuan?page=1&limit=20
GET /api/item?page=2&limit=50
```

### Smart Search

Setiap kata dalam search query dicari secara terpisah (AND logic):

| Query | Yang Dicari | Contoh Match |
|-------|-------------|--------------|
| `"laptop asus"` | "laptop" **DAN** "asus" | "Laptop ASUS", "ASUS Laptop" |
| `"budi jakarta"` | "budi" **DAN** "jakarta" | "Budi Jakarta", "Budi di Jakarta" |
| `"0812"` | "0812" | "081234567890" |

---

## Master Data {#master-data}

Semua master data memiliki endpoint yang sama:

| Resource | Endpoint |
|----------|----------|
| Satuan | `/api/satuan` |
| Kategori | `/api/kategori` |
| Jenis | `/api/jenis` |
| Sub Klasifikasi | `/api/sub-klasifikasi` |
| Klasifikasi | `/api/klasifikasi` |
| Grup | `/api/grup` |
| Konsumen | `/api/konsumen` |
| Supplier | `/api/supplier` |
| Gudang | `/api/gudang` |
| Rekening | `/api/rekening` |

### Ambil Semua Data

```
GET /api/{resource}
```

**Query:** `page`, `limit`

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "nama": "Box",
      "catatan": "Satuan dalam box",
      "created_at": "2026-09-05T10:00:00Z",
      "updated_at": "2026-09-05T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 25
  }
}
```

### Cari Data (Smart Search)

```
GET /api/{resource}/search
```

**Query:** `nama`, `page`, `limit`

**Contoh:**
```
GET /api/satuan/search?nama=box
GET /api/item/search?nama=laptop+asus&page=1&limit=5
```

### Ambil Total Data

```
GET /api/{resource}/total
```

**Response:**
```json
{ "total": 25 }
```

### Ambil Data by ID

```
GET /api/{resource}/:id
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "nama": "Box",
    "catatan": "Satuan dalam box"
  }
}
```

### Buat Data Baru

```
POST /api/{resource}
```

**Request Body:**
```json
{
  "nama": "Box",
  "catatan": "Satuan dalam box"
}
```

### Update Data

```
PUT /api/{resource}/:id
```

**Request Body:** (sama dengan create)

### Hapus Data

```
DELETE /api/{resource}/:id
```

---

## Item {#item}

Item memiliki relasi ke banyak master data (jenis, kategori, satuan, rekening, konsumen).

### Struktur Data

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
  "konsumen_id": 1,
  "konsumen": "Budi Santoso",
  "harga_beli": 5000000,
  "harga_jual": 7500000,
  "status": true,
  "created_at": "2026-09-05T10:00:00Z",
  "updated_at": "2026-09-05T10:00:00Z"
}
```

### Endpoint

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | `/api/item` | Ambil semua item (paginated) |
| GET | `/api/item/search` | Cari item (paginated) |
| GET | `/api/item/total` | Total item |
| GET | `/api/item/:id` | Detail item |
| GET | `/api/item/konsumen/:konsumen_id` | Item by konsumen (paginated) |
| POST | `/api/item` | Buat item baru |
| PUT | `/api/item/:id` | Update item |
| DELETE | `/api/item/:id` | Hapus item |

### Search Fields

| Param | Keterangan |
|-------|------------|
| `kode` | Search by kode item |
| `nama` | Search by nama item |
| `kategori_id` | Filter by kategori ID |
| `status` | Filter by status (true/false) |

---

## Pembelian {#pembelian}

Pembelian = catatan bahwa kamu beli barang dari supplier.

### Struktur Data

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
  "grand_total": 500000,
  "jumlah_uang_muka": 100000,
  "jumlah_uang_sisa": 400000,
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
  ]
}
```

### Endpoint

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | `/api/pembelian` | History pembelian (paginated) |
| GET | `/api/pembelian/total` | Total pembelian |
| GET | `/api/pembelian/next-no` | Preview no transaksi berikutnya |
| GET | `/api/pembelian/:id` | Detail pembelian |
| POST | `/api/pembelian` | Buat pembelian baru |
| PUT | `/api/pembelian/:id` | Update pembelian |
| DELETE | `/api/pembelian/:id` | Hapus pembelian |

### Filter History

| Param | Keterangan |
|-------|------------|
| `start_date` | Dari tanggal (YYYY-MM-DD) |
| `end_date` | Sampai tanggal (YYYY-MM-DD) |
| `supplier_id` | Filter per supplier |
| `no_transaksi` | Cari by no transaksi |

### Request Body (Create/Update)

```json
{
  "tanggal": "2026-09-05",
  "supplier_id": 1,
  "gudang_id": 1,
  "tgl_jatuh_tempo": "2026-10-05",
  "rekening_id": 1,
  "grand_total": 500000,
  "jumlah_uang_muka": 100000,
  "jumlah_uang_sisa": 400000,
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

**Catatan:**
- `no_transaksi` auto-generated oleh backend
- `grand_total`, `jumlah_uang_muka`, `jumlah_uang_sisa` dihitung oleh frontend
- Saat update, items lama dihapus dan diganti dengan items baru

---

## Bayar Hutang {#bayar-hutang}

Bayar Hutang = catatan bahwa kamu membayar sisa hutang ke supplier.

### Struktur Data

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
    }
  ]
}
```

### Endpoint

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | `/api/bayar-hutang` | History bayar hutang (paginated) |
| GET | `/api/bayar-hutang/total` | Total bayar hutang |
| GET | `/api/bayar-hutang/next-no` | Preview no transaksi berikutnya |
| GET | `/api/bayar-hutang/pembelian-sisa/:supplier_id` | Lihat pembelian dengan sisa hutang |
| GET | `/api/bayar-hutang/:id` | Detail bayar hutang |
| POST | `/api/bayar-hutang` | Buat bayar hutang baru |
| PUT | `/api/bayar-hutang/:id` | Update bayar hutang |
| DELETE | `/api/bayar-hutang/:id` | Hapus bayar hutang |

### Filter History

| Param | Keterangan |
|-------|------------|
| `start_date` | Dari tanggal (YYYY-MM-DD) |
| `end_date` | Sampai tanggal (YYYY-MM-DD) |
| `supplier_id` | Filter per supplier |
| `no_transaksi` | Cari by no transaksi |

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
  "grand_total": 800000,
  "jumlah_bayar": 700000,
  "sisa": 100000,
  "catatan": "Bayar cicilan",
  "items": [
    { "pembelian_id": 1, "jumlah_bayar": 400000 },
    { "pembelian_id": 2, "jumlah_bayar": 300000 }
  ]
}
```

**Catatan:**
- `grand_total`, `jumlah_bayar`, `sisa` dihitung oleh frontend
- `no_transaksi` auto-generated oleh backend

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
→ Old items di-restore dulu, lalu new items di-apply.

---

## Error Handling {#error-handling}

### Status Codes

| Code | Arti |
|------|------|
| 200 | Sukses |
| 201 | Berhasil buat data baru |
| 400 | Request salah (data tidak lengkap) |
| 404 | Data tidak ditemukan |
| 500 | Error server |

### Contoh Error Response

```json
{
  "error": "Satuan not found"
}
```

```json
{
  "error": "jumlah bayar melebihi sisa hutang untuk pembelian P-2026090001: sisa=100000, bayar=200000"
}
```

---

## Quick Reference {#quick-reference}

### Master Data

| Fungsi | Endpoint |
|--------|----------|
| Ambil semua | `GET /api/{resource}` |
| Cari data | `GET /api/{resource}/search?nama=...` |
| Ambil total | `GET /api/{resource}/total` |
| Ambil by ID | `GET /api/{resource}/:id` |
| Buat baru | `POST /api/{resource}` |
| Update | `PUT /api/{resource}/:id` |
| Hapus | `DELETE /api/{resource}/:id` |

### Transaksi

| Fungsi | Endpoint |
|--------|----------|
| Next no pembelian | `GET /api/pembelian/next-no?tanggal=YYYY-MM-DD` |
| Next no bayar hutang | `GET /api/bayar-hutang/next-no?tanggal=YYYY-MM-DD` |
| Pembelian sisa hutang | `GET /api/bayar-hutang/pembelian-sisa/:supplier_id` |

### Resources

`satuan`, `kategori`, `jenis`, `sub-klasifikasi`, `klasifikasi`, `item`, `grup`, `konsumen`, `rekening`, `supplier`, `gudang`, `pembelian`, `bayar-hutang`
