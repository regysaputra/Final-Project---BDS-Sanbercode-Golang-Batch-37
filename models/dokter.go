package models

type Dokter struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	Nama         string     `gorm:"not null" json:"nama"`
	TanggalLahir string     `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin string     `gorm:"not null" json:"jenis_kelamin"`
	NoHP         string     `gorm:"unique" json:"no_hp"`
	Spesialis    string     `gorm:"not null" json:"spesialis"`
	Jadwal       string     `gorm:"not null" json:"jadwal"`
	Diagnosa     *Diagnosa  `json:"-"`
	RawatInap	 *RawatInap `json:"-"`
}