## Description
Project mengenai simple CRUD absensi app 

## Stack 
- Golang(go1.19.4)
- GorillaMUX
- Database: MySQL
- MessageBroker: RabbitMQ

## Start Project
- go run ./app/main.go

## Endpoint
silahkan mengimport file postman yang ada di folder postman untuk melihat endpoint serta payload

## Penggunaan
- Buatlah akun terlebih dahulu pada endpoint Register
- Login untuk mendapatkan token sebagai authentikasi yang akan tersimpan di dalam cookie
- Setelah login user dapat melihat riwayat dari aktivitas yang telah di input ataupun absensinya
- User dapat melakukan checkin dan mendapatkan token checkin yang akan tersimpan di dalam cookie
- Setelah checkin, user dapat mengakses endpoint yang ada di dalam Activity untuk mengelola aktivitasnya
- Setelah mengelola aktifitas, maka user bisa melakukan checkout dan token checkin yang tersimpan di cookie akan terhapus
- User juga dapat melakukan Logout dan token yang tersimpan di cookie akan terhapus

## Testing
Terdapat unit testing di dalam masing - masing folder `absensi, activity, user`, silahkan masuk ke dalam salah satu folder melalui terminal lalu jalankan `go test`
