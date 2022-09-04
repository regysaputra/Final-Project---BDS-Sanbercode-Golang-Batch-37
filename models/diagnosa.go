package models

import (
	"time"
)

type Diagnosa struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	Tanggal	      time.Time  `json:"tanggal"`
	BeratBadan    float64    `json:"berat_badan"`
	TekananDarah  string     `json:"tekanan_darah"`
	Keluhan	      string     `json:"keluhan"`
	HasilDiagnosa string     `json:"hasil_diagnosa"`
	Tindakan	  string     `json:"tindakan"`
	DokterID	  uint       `json:"dokterID"`
	PasienID	  uint		 `json:"pasienID"`
	Dokter 		  Dokter	 `json:"-"`
	Pasien		  Pasien	 `json:"-"`
	Transaksi     *Transaksi `json:"-"`
	Resep 		  []Resep    `json:"-"`
}