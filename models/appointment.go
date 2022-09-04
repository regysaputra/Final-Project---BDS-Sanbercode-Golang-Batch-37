package models

import (
	"time"
)

type Appointment struct {
	ID uint `gorm:"primary_key" json:"id"`
	Mode string `json:"mode"`
	Tanggal time.Time `json:"department"`
	PasienID uint `json:"pasienID"`
	Pasien *Pasien `json:"-"`
}