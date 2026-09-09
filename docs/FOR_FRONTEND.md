# Panduan API untuk Frontend Developer

**Base URL:** `http://localhost:8080/api`

---

## Daftar Isi

1. [Cara Baca Dokumentasi Ini](#cara-baca)
2. [Master Data (CRUD Sederhana)](#master-data)
3. [Item](#item)
4. [Pembelian](#pembelian)
5. [Bayar Hutang](#bayar-hutang)
6. [Contoh Implementasi](#contoh)
7. [Error Handling](#error-handling)

---

## Cara Baca Dokumentasi Ini {#cara-baca}

Setiap endpoint punya pola yang sama:

```
METHOD /api/nama-endpoint
```

**Method:**
- `GET` = Ambil data
- `POST` = Buat data baru
- `PUT` = Update data
- `DELETE` = Hapus data

**Response selalu pakai pola:**
```json
{
  "data": [...],
  "pagination": { ... }
}
```

---

## Master Data {#master-data}

Semua master data (satuan, kategori, jenis, klasifikasi, sub-klasifikasi, grup, konsumen, supplier, gudang, rekening) punya pola yang sama:

### Endpoint Pattern

| Fungsi | Method | Endpoint |
|--------|--------|----------|
| Ambil semua | GET | `/api/{resource}` |
| Cari data | GET | `/api/{resource}/search?nama=...` |
| Ambil total | GET | `/api/{resource}/total` |
| Ambil by ID | GET | `/api/{resource}/:id` |
| Buat baru | POST | `/api/{resource}` |
| Update | PUT | `/api/{resource}/:id` |
| Hapus | DELETE | `/api/{resource}/:id` |

**Resources:** satuan, kategori, jenis, sub-klasifikasi, klasifikasi, grup, konsumen, supplier, gudang, rekening

### Pagination

| Param | Default | Max | Keterangan |
|-------|---------|-----|------------|
| `page` | 1 | - | Nomor halaman |
| `limit` | 10 | 100 | Jumlah data per halaman |

**Contoh:**
```
GET /api/satuan?page=1&limit=20
```

### Smart Search

Setiap kata dipisah dan dicari secara terpisah (AND logic):

```
GET /api/item/search?nama=laptop+asus
# Akan match: "Laptop ASUS", "ASUS Laptop", dll.

GET /api/konsumen/search?nama=budi+jakarta
# Akan match: "Budi Jakarta", "Budi di Jakarta", dll.
```

### Contoh Response

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

---

## Item {#item}

Item memiliki relasi ke banyak master data.

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
| GET | `/api/item` | Ambil semua item |
| GET | `/api/item/search` | Cari item |
| GET | `/api/item/total` | Total item |
| GET | `/api/item/:id` | Detail item |
| GET | `/api/item/konsumen/:konsumen_id` | Item by konsumen |
| POST | `/api/item` | Buat item baru |
| PUT | `/api/item/:id` | Update item |
| DELETE | `/api/item/:id` | Hapus item |

### Search Fields

| Param | Keterangan |
|-------|------------|
| `kode` | Search by kode |
| `nama` | Search by nama |
| `kategori_id` | Filter by kategori |
| `status` | Filter by status (true/false) |

### Request Body (Create/Update)

```json
{
  "kode": "ELG-001",
  "nama": "Laptop ASUS",
  "jenis_id": 1,
  "kategori_id": 1,
  "satuan_id": 1,
  "rekening_id": 1,
  "konsumen_id": 1,
  "harga_beli": 5000000,
  "harga_jual": 7500000,
  "status": true
}
```

**Catatan:**
- `konsumen_id` optional (bisa null)
- Semua `*_id` harus sesuai dengan ID yang ada di master data

---

## Pembelian {#pembelian}

Pembelian = catatan bahwa kamu beli barang dari supplier.

### Struktur Data

```json
{
  "id": 1,
  "no_transaksi": "P-2026090001",  // Auto-generated
  "tanggal": "2026-09-05",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",  // Nama supplier (otomatis)
  "gudang_id": 1,
  "gudang": "Gudang Utama",        // Nama gudang (otomatis)
  "tgl_jatuh_tempo": "2026-10-05",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",      // Nama rekening (otomatis)
  "grand_total": 500000,           // Total harga semua barang
  "jumlah_uang_muka": 100000,      // Uang yang sudah dibayar di awal
  "jumlah_uang_sisa": 400000,      // Sisa yang belum dibayar
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
| GET | `/api/pembelian` | History pembelian |
| GET | `/api/pembelian/total` | Total pembelian |
| GET | `/api/pembelian/next-no` | Preview no transaksi |
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
  "no_transaksi": "H-2026090001",  // Auto-generated
  "tanggal": "2026-09-15",
  "supplier_id": 1,
  "supplier": "PT Supplier Jaya",
  "gudang_id": 1,
  "gudang": "Gudang Utama",
  "rekening_id": 1,
  "rekening": "Bank Mandiri",
  "grand_total": 800000,   // Total sisa hutang dari semua pembelian yang dipilih
  "jumlah_bayar": 700000,  // Total yang dibayar
  "sisa": 100000,          // Sisa hutang
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
  ]
}
```

### Endpoint

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | `/api/bayar-hutang` | History bayar hutang |
| GET | `/api/bayar-hutang/total` | Total bayar hutang |
| GET | `/api/bayar-hutang/next-no` | Preview no transaksi |
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

## Contoh Implementasi {#contoh}

### Flow Pembelian

```
1. Ambil no transaksi
   → GET /api/pembelian/next-no?tanggal=2026-09-05
   → Response: { "no_transaksi": "P-2026090001" }

2. Ambil data untuk form
   → GET /api/supplier         → list supplier
   → GET /api/gudang           → list gudang
   → GET /api/rekening         → list rekening
   → GET /api/item             → list item

3. Submit pembelian
   → POST /api/pembelian
   → Body: { tanggal, supplier_id, gudang_id, items: [...] }
   → Response: { id, no_transaksi, grand_total, ... }
```

### Flow Bayar Hutang

```
1. Pilih supplier
   → GET /api/bayar-hutang/pembelian-sisa/:supplier_id
   → Response: list pembelian dengan sisa hutang

2. Input jumlah bayar, distribusi ke pembelian
   → Hitung: grand_total = total sisa, jumlah_bayar = input user, sisa = grand_total - jumlah_bayar

3. Submit bayar hutang
   → POST /api/bayar-hutang
   → Body: { tanggal, supplier_id, gudang_id, grand_total, jumlah_bayar, sisa, items: [...] }
   → Response: { id, no_transaksi, grand_total, jumlah_bayar, sisa }
```

### Contoh Distribusi Pembayaran (JavaScript)

```javascript
// Data dari API
const pembelians = [
  { id: 1, no_transaksi: "P-001", jumlah_uang_sisa: 400000 },
  { id: 2, no_transaksi: "P-002", jumlah_uang_sisa: 300000 }
];

const totalBayar = 700000;
let sisaBayar = totalBayar;
const items = [];

for (const p of pembelians) {
  if (sisaBayar <= 0) break;
  
  const bayar = Math.min(sisaBayar, p.jumlah_uang_sisa);
  items.push({
    pembelian_id: p.id,
    jumlah_bayar: bayar
  });
  sisaBayar -= bayar;
}

// items = [
//   { pembelian_id: 1, jumlah_bayar: 400000 },  // P-001 lunas
//   { pembelian_id: 2, jumlah_bayar: 300000 }   // P-002 lunas
// ]

// Kirim ke API
const payload = {
  tanggal: "2026-09-15",
  supplier_id: 1,
  gudang_id: 1,
  rekening_id: 1,
  grand_total: 700000,
  jumlah_bayar: 700000,
  sisa: 0,
  catatan: "Bayar lunas semua",
  items: items
};
```

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
  "error": "jumlah bayar melebihi sisa hutang untuk pembelian P-2026090001: sisa=100000, bayar=200000"
}
```

```json
{
  "error": "data konsistensi error: jumlah_uang_muka menjadi negatif untuk pembelian P-2026090001 (muka=50000, bayar=100000)"
}
```

```json
{
  "error": "Bayar hutang not found"
}
```

---

## Quick Reference

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
