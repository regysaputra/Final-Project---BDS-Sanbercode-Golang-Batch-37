package models

type Transaksi struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	BiayaObat   string    `json:"biaya_obat"`
	BiayaDokter string    `json:"biaya_dokter"`
	BiayaKamar  string    `json:"biaya_kamar"`
	Total 		string    `json:"total"`
	DiagnosaID  uint      `json:"diagnosaID"`
	Diagnosa    *Diagnosa `json:"-"`
} 