# 📚 Sistem Manajemen Jadwal Perkuliahan Akademik (JadwalKu)

## Deskripsi

JadwalKu merupakan aplikasi berbasis bahasa pemrograman Go yang digunakan untuk mengelola jadwal perkuliahan akademik. Sistem ini membantu staf administrasi maupun mahasiswa dalam mengatur dan melihat jadwal mata kuliah sehingga tidak terjadi bentrok waktu maupun penggunaan ruangan.

Data yang dikelola meliputi data mata kuliah, data dosen, waktu perkuliahan, dan lokasi ruangan.

## Fitur Utama

### 1. Manajemen Data Jadwal

* Menambahkan data jadwal mata kuliah
* Mengubah data jadwal mata kuliah
* Menghapus data jadwal mata kuliah
* Menampilkan seluruh jadwal perkuliahan

### 2. Pengelolaan Informasi Perkuliahan

* Mencatat nama mata kuliah
* Mencatat nama dosen
* Mencatat waktu mulai kuliah
* Mencatat waktu selesai kuliah
* Mencatat lokasi atau ruangan perkuliahan

### 3. Pencarian Data

Sistem menyediakan dua metode pencarian:

* Sequential Search
* Binary Search

Pencarian dapat dilakukan berdasarkan:

* Nama mata kuliah
* Nama dosen

### 4. Pengurutan Data

Sistem menyediakan dua metode pengurutan:

* Selection Sort
* Insertion Sort

Pengurutan dilakukan berdasarkan:

* Jam mulai perkuliahan

### 5. Statistik Jadwal

* Menampilkan total jam kuliah per minggu
* Menampilkan jumlah mata kuliah yang tersisa dalam satu hari

## Struktur Data

Setiap data jadwal terdiri dari:

| Field       | Keterangan                |
| ----------- | ------------------------- |
| Mata Kuliah | Nama mata kuliah          |
| Dosen       | Nama dosen pengampu       |
| Jam Mulai   | Waktu mulai perkuliahan   |
| Jam Selesai | Waktu selesai perkuliahan |
| Ruangan     | Lokasi perkuliahan        |

## Algoritma yang Digunakan

### Searching

* Sequential Search
* Binary Search

### Sorting

* Selection Sort
* Insertion Sort

## Tujuan Pengembangan

Aplikasi ini dibuat sebagai tugas besar Mata Kuliah Algoritma Pemrograman 2 dengan tujuan mengimplementasikan konsep:

* Array
* Record/Struct
* Fungsi dan Prosedur
* Sequential Search
* Binary Search
* Selection Sort
* Insertion Sort
* Pengolahan Data dan Statistik Sederhana

## Cara Menjalankan Program

1. Pastikan Golang telah terpasang pada perangkat.
2. Clone repository:

```bash
git clone https://github.com/AlinKarisa/Sistem-Manajemen-Jadwal-Perkuliahan-Akademik.git
```

3. Masuk ke folder proyek:

```bash
cd Sistem-Manajemen-Jadwal-Perkuliahan-Akademik
```

4. Jalankan program:

```bash
go run main.go
```

## Contoh Menu Program

```text
=== SISTEM MANAJEMEN JADWAL PERKULIAHAN ===

1. Tambah Jadwal
2. Ubah Jadwal
3. Hapus Jadwal
4. Tampilkan Jadwal
5. Cari Jadwal
6. Urutkan Jadwal
7. Statistik Perkuliahan
0. Keluar

Pilih Menu:
```

## 👥 Disusun Oleh

**Kelompok I**

1. **Shifa Andien Widyanto** (109082500003)
2. **Alin Karisa Hizannah** (109082500010)
3. **M Mahdan Argya Syarif** (109082500059)

---

*Tugas Besar Mata Kuliah Algoritma Pemrograman 2*  
*PROGRAM STUDI S1 INFORMATIKA*
*FAKULTAS INFORMATIKA*  
*TELKOM UNIVERSITY PURWOKERTO*
