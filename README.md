# ReviewPiLem
Dokumentasi API dapat diakses pada [Postman Documentation](https://documenter.getpostman.com/view/40297962/2sB2iwEDrf)

## Prasyarat
Pastikan memiliki  
- Go Version >= go 1.20
- PostgreSQL >= version 15.0

### Instalasi
1. Clone repositori ini:
    ```bash
    git clone https://github.com/username/reviewpilem.git
    ```
2. Masuk ke direktori proyek:
    ```bash
    cd reviewpilem
    ```
3. Copy env:
    ```bash
    cp .env.example .env
    ```
    Dan isi varibale yang kosong pada .env

### Menjalankan Server
1. Migrasi
    ```bash
    go run main.go --migration
    ```
2. Seeding
    ```bash
    go run main.go --seed
    ```
3. Run
    ```bash
    go run main.go --run
    ```

Atau bisa lagsung menjalankan ketiganya dengan 

```bash
go run main.go --migration --seed --run
```

