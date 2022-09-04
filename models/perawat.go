package models

type Perawat struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Nama      string     `gorm:"not null" json:"nama"`
	RawatInap *RawatInap `json:"-"`
}