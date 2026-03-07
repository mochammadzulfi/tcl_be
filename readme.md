# TCL Inventory Backend

Backend API untuk aplikasi **TCL Inventory Management System** yang menangani manajemen produk, transaksi **Stock In**, **Stock Out**, inventory, dan laporan.

---

# Tech Stack

- Golang
- Gin Web Framework
- GORM
- PostgreSQL
- REST API

---

# Cara Menjalankan Aplikasi

```bash
git clone https://github.com/mochammadzulfi/tcl_be.git
cd tcl_be
Setup Database

File database sudah disediakan pada folder:

database/inventory_db.sql
Buat database baru
CREATE DATABASE inventory_db;
Import database

Menggunakan terminal:

psql -U postgres -d inventory_db -f database/inventory_db.sql

Atau dapat diimport menggunakan pgAdmin / database tools lainnya.

Menjalankan Backend

Masuk ke folder backend (jika repository dipisah, langsung di root project):

go mod tidy

Sesuaikan konfigurasi database pada file:

config/database.go

Contoh konfigurasi:

host=localhost
user=postgres
password=postgres
dbname=inventory_db
port=5433
sslmode=disable

Jalankan server:

go run main.go

Server akan berjalan di:

http://localhost:8080
Endpoint API
Products
GET  /products
POST /products
Stock In
GET   /stock-in
POST  /stock-in
PATCH /stock-in/{id}/start
PATCH /stock-in/{id}/complete
Stock Out
GET   /stock-out
POST  /stock-out
PATCH /stock-out/{id}/start
PATCH /stock-out/{id}/complete
PATCH /stock-out/{id}/cancel
Inventory
GET /inventory
Reports
GET /reports/stock-in
GET /reports/stock-out