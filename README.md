# Golang-Kasir-Api-With-Checkout-and-Report-System

Sumber Belajar: CodeWithUmam

Backend API sistem kasir menggunakan Go, PostgreSQL (Supabase),
dengan fitur:
- Product CRUD
- Checkout transaksi
- Report penjualan harian
- Report berdasarkan range tanggal

## Tech Stack
- Go
- PostgreSQL
- pgx
- net/http

## Run
```bash
go run main.go


---

## Note

- [x] Repo public
- [x] `.env` tidak ke-push
- [x] Bisa `go run`
- [x] Endpoint checkout & report jalan

---

## 📝 API Endpoints

| Method | Endpoint           | Description                      |
|--------|--------------------|----------------------------------|
| GET    | /test              | Test server status               |
| GET    | /api/procuts       | Get all products                 |
| GET    | /api/products/:id  | Get a products by ID             |
| POST   | /api/products      | Create a new prod                |
| PUT    | /api/products/:id  | Update a prod by ID              |
| DELETE | /api/products/:id  | Delete a prod by ID              |
| POST   | /api/checkout      | Create a checkout transaction    |
| GET    | /api/report        | Get today's sales summary        |
