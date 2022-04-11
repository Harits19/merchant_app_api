# Merchant App Api

Ini adalah proyek untuk melengkapi kebutuhan rekrutmen pada PT Majoo Indonesia.

# Teknologi

- [**gin**](https://pkg.go.dev/github.com/gin-gonic/gin)
  Web framework yang membantu membuat aplikasi REST API

- [**jwt**](https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible)
  Untuk membantu membuat JSON Web Token di aplikasi golang

- [**validator**](https://pkg.go.dev/github.com/go-playground/validator/v10)
  Untuk membantu memvalidasi data yang dikirimkan dari end user agar sesuai dengan tag pada struct yang diberikan

- [**mysql**](https://pkg.go.dev/github.com/go-sql-driver/mysql)
  Driver untuk membantu menghubungkan aplikasi dengan database MySQL

# GIF

Berikut merupakan demo dari aplikasi yang telah dibuat
![Alt Text](https://github.com/Harits19/merchant_app_api/blob/main/demo/demo.gif?raw=true)

# Majoo Test Case Backend

## Nomor 1 :

### G

Untuk struktur ERD masih bisa ditingkatkan atau dibenahi dengan menghilangkan Foreign Key merchant_id di Table Transactions dikarenakan dari Table Merchant sudah memiliki hubungan one to many pada Tabel Outlet dan Tabel Outlet sendiri juga memiliki hubungan one to many pada Table Transaction. Oleh karena itu, merchant_id dapat dihilangkan dikarenakan Table Merchant dapat terhubung secara tidak langsung ke Table Transaction melalui Table Outlet.

### H

Berikut merupakan DML untuk mengambil data dari tabel transaksi yang dihubungkan oleh tabel outlet, merchant, dan user

```
select
    MIN(M.merchant_name) as merchant_name,
    MIN(O.outlet_name) as outlet_name ,
    SUM(T.bill_total) as total_omzet,
    DATE_FORMAT(T.updated_at, '%Y-%m-%d') as date
from
    users as U
    join merchants as M on
    M.user_id = U.id
    join outlets as O on
    M.id = O.merchant_id
    join transactions as T on
    T.outlet_id = O.id
where
    U.id = ?
    and DATE_FORMAT(T.updated_at, '%Y-%m') = ?
group by
    date
```

Berikut merupakan DML untuk mengambil data dari table user

```
select
	id,
	name,
	user_name,
	password,
	created_at,
	created_by,
	updated_at,
	updated_by
from
	users
where
	user_name =?
```

## Nomor 2

- Pembuatan struct salah. Variabel AreaValue diganti menjadi float menyamakan dengan nilai fungsi yang ada
- Tipe data dari param1 dan param2 disamakan agar tidak terjadi error
- Penamaan variabel type harus diganti karena type sendiri merupakan syntax dari bahasa Golang
- Tipe data type2 yang awalnya []string diganti menjadi string agar berjalan statement switch case
- Penulisan syntax Var diganti menjadi var
- Terdapat perulangan pembuatan variabel area
- Terdapat perulangan untuk menajalan fungsi create pada variabel Db
- Terdapat varibel inst yang tidakd digunakan

Code setelah diubah:

```
package main

import (
	"errors"
	"log"
)

type Area struct {
	ID        int64   `gorm:"column:id;primaryKey;"`
	AreaValue float64 `gorm:"column:area_value"`
	AreaType  string  `gorm:"column:type"`
}

func call() error {
	err := _u.repository.InsertArea(10, 10, "persegi")
	if err != nil {
		log.Error().Msg(err.Error())
		err = errors.New(en.ERROR_DATABASE)
		return err
	}
	return nil
}

func (_r *AreaRepository) InsertArea(param1 float64, param2 float64, type2 string, ar *Area) (err error) {
	inst := _r.DB.Model(ar)
	var area float64
	area = 0
	switch type2 {
	case "persegi panjang":
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"

	case "persegi":
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi"
	case "segitiga":

		area = 0.5 * (param1 * param2)
		ar.AreaValue = area
		ar.AreaType = "segitiga"
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
	}
	err = _r.DB.create(&ar).Error
	if err != nil {
		return err
	}
	return
}
```

## Nomor 3

```
variabel inputDeret0
variabel inputDeret1
variabel selisih = inputDeret1 - inputDeret0
variabel panjangDeret

variabel index = 2
variabel angka array
angka array dimasukkan variabel inputDeret0
angka array dimasukkan variabel inputDeret1

for selama index lebih kecil dari panjang deret{
	variabel nilai = angka array pada index - 1 * index + selisih
	array angka dimasukkan variabel nilai
}
menampilkan variabel angka

```

## Nomor4

```
variabel listAngka array
variabel isAscending = true

for index1 = 0; index1 < panjang listAngka; setiap perulangan selesai index1 ditambah 1{
	for index2 = index1 +1; index2 < panjang listAngka; setiap perulangan selesai index2 ditambah 1{
		variabel angka1 = listAngka pada index1
		variabel angka2 = listAngka pada index2
		if jika isAscending == true{
			if angka1 lebih besar dari angka2{
				listAngka pada index1 = angka2
				listAngka pada index2 = angka1
			}
		}else{
			if angka1 lebih kecil dari angka2{
				listAngka pada index1 = angka2
				listAngka pada index2 = angka1
			}
		}

	}

	menampilkan variabel listAngka
}


```
