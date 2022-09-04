package config

import (
	"final-project/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() *gorm.DB {
    username := "root"
    password := ""
    host := "tcp(127.0.0.1:3306)"
    database := "hospital"

    dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })

    if err != nil {
        panic(err.Error())
    }

    db.AutoMigrate(&models.Appointment{}, &models.Diagnosa{}, &models.Dokter{}, &models.Kamar{}, &models.Obat{}, &models.Pasien{}, &models.Perawat{}, &models.RawatInap{}, &models.Resep{}, &models.Transaksi{}, &models.User{})

    return db
}