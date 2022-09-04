package models

type Resep struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Jumlah     int       `gorm:"not null" json:"jumlah"`
	Dosis      string    `gorm:"not null" json:"dosis"`
	DiagnosaID uint      `gorm:"not null" json:"diagnosaID"`
	ObatID     uint      `gorm:"not null" json:"obatID"`
	Diagnosa   Diagnosa `json:"-"`
	Obat	   Obat      `json:"-"`
} 