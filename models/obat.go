package models

type Obat struct {
	ID           uint         `gorm:"primary_key" json:"id"`
	Nama		 string       `gorm:"not null" json:"nama"`
	Satuan		 string       `gorm:"not null" json:"satuan"`
	Harga        string       `gorm:"not null" json:"harga"`
	Keterangan   string       `json:"keterangan"`
	Resep		 []Resep      `json:"-"`
}