package config

import (
	"final-project/models"
	"final-project/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() *gorm.DB {
    environment := utils.Getenv("ENVIRONMENT", "development")

    if environment == "production" {
        username := os.Getenv("DATABASE_USERNAME")
        password := os.Getenv("DATABASE_PASSWORD")
        host := os.Getenv("DATABASE_HOST")
        port := os.Getenv("DATABASE_PORT")
        database := os.Getenv("DATABASE_NAME")
        // production
        dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            panic(err.Error())
        }

        db.AutoMigrate(&models.Appointment{}, &models.Diagnosa{}, &models.Dokter{}, &models.Kamar{}, &models.Obat{}, &models.Pasien{}, &models.Perawat{}, &models.RawatInap{}, &models.Resep{}, &models.Transaksi{}, &models.User{})

        return db
    } else {
        // development
        username := "root"
        password := ""
        host := "tcp(127.0.0.1:3306)"
        database := "database_movie"

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
}