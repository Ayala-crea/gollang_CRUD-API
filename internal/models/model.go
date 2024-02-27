package model

type User struct {
	IdUser   int    `gorm:"primaryKey;column:IdUser;autoIncrement" json:"-"`
	Nama     string `gorm:"column:Nama" json:"nama"`
	Npm      string `gorm:"column:Npm" json:"npm"`
	Kelas    string `gorm:"column:Kelas" json:"kelas"`
	AsalKota string `gorm:"column:asal_kota" json:"asal_kota"`
}

type Sell struct {
	IdSell   int    `gorm:"primaryKey;column:IdSell;autoIncrement" json:"-"`
	Nama     string `gorm:"column:Nama" json:"nama"`
	Npm      string `gorm:"column:Npm" json:"npm"`
	Kelas    string `gorm:"column:Kelas" json:"kelas"`
	AsalKota string `gorm:"column:asal_kota" json:"asal_kota"`
}
