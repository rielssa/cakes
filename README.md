# cakes
technical test Ralali

# Cake Store RESTful API Service

Cake Store RESTful API Service adalah layanan API yang menyediakan fungsi-fungsi CRUD (Create, Read, Update, Delete) untuk mengelola informasi tentang kue dalam sebuah toko kue.

## Prasyarat

Sebelum menjalankan aplikasi ini, pastikan Anda telah memenuhi prasyarat berikut:

- Go Language: Pastikan Go telah terinstal di komputer Anda. Anda dapat mengunduh dan menginstal Go dari situs resmi Go: https://golang.org/dl/
- MySQL Database: Pastikan Anda telah mengatur dan mengonfigurasi MySQL database di lingkungan Anda. Anda juga harus memiliki informasi koneksi database yang diperlukan untuk menghubungkan aplikasi ke MySQL.

## Instalasi

1. Clone repositori ini ke direktori lokal Anda:

```shell
git clone https://github.com/rielssa/cakestore.git
```

2. Masuk ke direktori proyek:

```shell
cd cakestore
```

3. Install dependensi yang diperlukan menggunakan perintah go get:

```shell
go get
```

4. Konfigurasi Database:
   
   - Buka file `config/config.json` dan sesuaikan pengaturan koneksi basis data MySQL sesuai dengan pengaturan MySQL Anda.
   - Pastikan Anda telah membuat basis data yang sesuai dan memiliki tabel "cakes" dengan struktur yang diperlukan. Anda dapat menggunakan skema tabel yang diberikan pada dokumentasi projek.

5. Jalankan aplikasi:

```shell
go run main.go
```

Aplikasi akan dijalankan pada `localhost:8080`.

## Endpoint API

Berikut adalah daftar endpoint yang tersedia dalam API ini:

1. **List of Cakes**
   
   - URL: `/cakes`
   - Metode: GET
   - Deskripsi: Mendapatkan daftar semua kue yang tersedia di toko kue.
   - Contoh Response:
     ```json
     [
       {
         "id": 1,
         "title": "Lemon cheesecake",
         "description": "A cheesecake made of lemon",
         "rating": 7.0,
         "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
         "created_at": "2020-02-01 10:56:31",
         "updated_at": "2020-02-13 09:30:23"
       },
       ...
     ]
     ```

2. **Detail of Cake**
   
   - URL: `/cakes/{id}`
   - Metode: GET
   - Deskripsi: Mendapatkan detail kue berdasarkan ID.
   - Contoh Response:
     ```json
     {
       "id": 1,
       "title": "Lemon cheesecake",
       "description": "A cheesecake made of lemon",
       "rating": 7.0,
       "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
       "created_at": "2020-02-01 10:56:31",
       "updated_at": "2020-02-13 09:30:23"
     }
     ```

3. **Add Cake**
   
   - URL: `/cakes`
   - Metode: POST
   - Deskripsi: Menambahkan kue baru ke toko kue.
   - Contoh Request Body:
     ```json
     {
       "title": "Chocolate cake",
       "description": "A delicious chocolate cake",
       "rating": 8.5,
       "image": "https://example.com/chocolate-cake.jpg"
     }
     ```
   - Contoh Response:
     ```json
     {
       "id": 2,
       "title": "Chocolate cake",
       "description": "A delicious chocolate cake",
       "rating": 8.5,
       "image": "https://example.com/chocolate-cake.jpg",
       "created_at": "2023-07-15 10:00:00",
       "updated_at": "2023-07-15 10:00:00"
     }
     ```

4. **Update Cake**
   
   - URL: `/cakes/{id}`
   - Metode: PUT
   - Deskripsi: Mengupdate informasi kue berdasarkan ID.
   - Contoh Request Body:
     ```json
     {
       "title": "New Chocolate Cake",
       "description

": "A new version of the chocolate cake",
       "rating": 9.0,
       "image": "https://example.com/new-chocolate-cake.jpg"
     }
     ```
   - Contoh Response:
     ```json
     {
       "id": 2,
       "title": "New Chocolate Cake",
       "description": "A new version of the chocolate cake",
       "rating": 9.0,
       "image": "https://example.com/new-chocolate-cake.jpg",
       "created_at": "2023-07-15 10:00:00",
       "updated_at": "2023-07-15 12:00:00"
     }
     ```

5. **Delete Cake**
   
   - URL: `/cakes/{id}`
   - Metode: DELETE
   - Deskripsi: Menghapus kue berdasarkan ID.
   - Contoh Response:
     ```json
     {
       "message": "Cake with ID 2 has been deleted successfully"
     }
     ```

## Testing

Unit tests telah disediakan untuk beberapa fungsi dalam proyek ini. Untuk menjalankan unit tests, gunakan perintah berikut:

```shell
go test ./...
```

## Lisensi

Proyek ini menggunakan lisensi [MIT](LICENSE).
