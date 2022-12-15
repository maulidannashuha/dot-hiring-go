Table of contents
=================
<!--ts-->
   * [Table of contents](#table-of-contents)
   * [DOT Hiring Go](#dot-hiring-go)
      * [Instalisasi](#instalisasi)
      * [Kode](#kode)
        * [Patern](#pattern)
        * [Framework](#framework)
        * [Relation](#relation)
        * [Postman](#postman)
      * [Video](#video)
        
<!--te-->

# DOT Hiring Go

## Instalisasi

Anda bisa install dengan cara :
```bash
go get tidy
```

## Kode
Kemudian jalankan dengan perintah : 
```bash
go run main.go
```

### Pattern
Pattern yang digunakan adalah Model Controller Router. Pattern ini dipilih untuk memudahkan pembuatan aplikasi kecil terutama dengan hanya terdapat beberapa module. Dengan asumsi tersebut maka diambillah Model Controller Router sebagai pattern yang digunakan.

Namun untuk aplikasi yang lebih besar DDD lebih disarankan untuk digunakan. Hal ini nantinya juga akan memudahkan ketika aplikasi dibuat menjadi microservices karena setiap layer sudah terisolasi.

### Framework
Framework yang digunakan adalah Gin

### Relation
Relation yang digunakan adalah user yang dapat memiliki banyak buku

## Video
Video dapat dilihat pada https://www.loom.com/share/e46e4dae0f604c2bbe47a8c8e7da2960
