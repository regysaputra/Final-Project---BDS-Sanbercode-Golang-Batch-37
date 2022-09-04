package models

type Kamar struct {
	ID			uint		`gorm:"primary_key" json:"id"`
	NoKamar		string 		`gorm:"not null" json:"no_kamar"`
	Lantai		string		`gorm:"not null" json:"lantai"`
	Type        string		`gorm:"not null" json:"type"`
	Harga		string		`gorm:"not null" json:"harga"`
	Status		bool		`gorm:"not null" json:"status"`
	RawatInap   *RawatInap  `json:"-"`
}