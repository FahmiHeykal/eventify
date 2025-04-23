# Eventify - Backend API

Eventify adalah aplikasi backend untuk manajemen acara dan tiket menggunakan **Golang** dan **PostgreSQL**. Aplikasi ini menyediakan API untuk mengelola pengguna, acara, tiket, dan transaksi.

## Fitur Utama
- **User Authentication :** Register dan login menggunakan email dan password.
- **Event Management :** Buat, update, hapus acara.
- **Ticket Management :** Pengguna dapat membeli tiket untuk acara tertentu.
- **Transaction Management :** Kelola transaksi pembelian tiket.

## Teknologi yang Digunakan
- **Golang** (Go) : Bahasa pemrograman untuk pengembangan backend.
- **PostgreSQL** : Database untuk menyimpan data acara, tiket, dan transaksi.
- **GORM** : ORM untuk mempermudah interaksi dengan database.
- **JWT** : Untuk autentikasi dan pengelolaan sesi pengguna.
- **Echo** : Framework web yang digunakan untuk membangun API.
- **Bcrypt** : Untuk mengenkripsi dan memverifikasi password.

## Persyaratan Sistem
Sebelum menjalankan proyek, pastikan sudah memiliki :
- **Go 1.18+** terinstall di sistem.
- **PostgreSQL** terinstall dan berjalan.
- **Golang JWT** dan **Echo** sudah di-install dengan menggunakan `go mod`.
