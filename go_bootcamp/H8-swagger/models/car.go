package models

type Car struct {
	Pemilik  string `gorm:"varchar(100)"`
	Merk     string `gorm:"varchar(100)`
	Harga    int    `gorm:"integer(11)"`
	Typecars string `gorm:"varchar(100)"`
	id       uint   `gorm:"primaryKey"`
}
