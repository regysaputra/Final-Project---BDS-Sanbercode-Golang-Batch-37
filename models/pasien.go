package models

import (
	"time"
)

type Pasien struct {
	ID           uint         `gorm:"primary_key" json:"id"`
	Nama	     string	      `json:"nama"`
    JenisKelamin string       `gorm:"not null" json:"jenis_kelamin"`
    NoHP         string       `gorm:"unique" json:"no_hp"`
	Alamat       string       `json:"alamat"`
	TanggalLahir time.Time    `gorm:"not null" json:"tanggal_lahir"`
    CreatedAt    time.Time    `json:"created_at"`
    UpdatedAt    time.Time    `json:"updated_at"`
	Appointment  *Appointment `json:"-"`
	RawatInap    *RawatInap   `json:"-"` 
}