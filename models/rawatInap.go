package models

import (
	"time"
)

type RawatInap struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`
	PasienID	  uint 		`json:"pasienID"`
	DokterID	  uint		`json:"dokterID"`
	PerawatID     uint 		`json:"perawatID"`
	KamarID		  string	`json:"kamarID"`
	Pasien 		  Pasien	`json:"-"`
	Dokter		  Dokter    `json:"-"`
	Perawat       Perawat   `json:"-"`
	Kamar         Kamar     `json:"-"`
}