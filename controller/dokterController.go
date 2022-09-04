package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type dokterInput struct {
    Nama         string     `gorm:"not null" json:"nama"`
	TanggalLahir string     `gorm:"not null" json:"tanggal_lahir"`
	JenisKelamin string     `gorm:"not null" json:"jenis_kelamin"`
	NoHP         string     `gorm:"unique" json:"no_hp"`
	Spesialis    string     `gorm:"not null" json:"spesialis"`
	Jadwal       string     `gorm:"not null" json:"jadwal"`
}

// GetAllDokter godoc
// @Summary Get all Dokter.
// @Description Get a list of Dokter.
// @Tags Dokter
// @Produce json
// @Success 200 {object} []models.Dokter
// @Router /dokter [get]
func GetAllDokter(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var dokter []models.Dokter
    db.Find(&dokter)

    c.JSON(http.StatusOK, gin.H{"data": dokter})
}

// CreateDokter godoc
// @Summary Create New Dokter.
// @Description Creating a new Dokter.
// @Tags Dokter
// @Param Body body dokterInput true "the body to create a new Dokter"
// @Produce json
// @Success 200 {object} models.Dokter
// @Router /dokter [post]
func CreateDokter(c *gin.Context) {
    // Validate input
    var input dokterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Dokter
    dokter := models.Dokter{Nama: input.Nama, TanggalLahir: input.TanggalLahir, JenisKelamin: input.JenisKelamin, NoHP: input.NoHP, Spesialis: input.Spesialis, Jadwal: input.Jadwal}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&dokter)

    c.JSON(http.StatusOK, gin.H{"data": dokter})
}

// GetDokterById godoc
// @Summary Get Dokter.
// @Description Get an Dokter by id.
// @Tags Dokter
// @Produce json
// @Param id path string true "Dokter id"
// @Success 200 {object} models.Dokter
// @Router /dokter/{id} [get]
func GetDokterById(c *gin.Context) { // Get model if exist
    var dokter models.Dokter

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&dokter).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": dokter})
}

// UpdateDokter godoc
// @Summary Update Dokter.
// @Description Update Dokter by id.
// @Tags Dokter
// @Produce json
// @Param id path string true "Dokter id"
// @Param Body body dokterInput true "the body to update age dokter category"
// @Success 200 {object} models.Dokter
// @Router /dokter/{id} [patch]
func UpdateDokter(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var dokter models.Dokter
    if err := db.Where("id = ?", c.Param("id")).First(&dokter).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input dokterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Dokter
    updatedInput.Nama = input.Nama
    updatedInput.TanggalLahir = input.TanggalLahir
	updatedInput.JenisKelamin = input.JenisKelamin
	updatedInput.NoHP = input.NoHP
	updatedInput.Spesialis = input.Spesialis
	updatedInput.Jadwal = input.Jadwal

    db.Model(&dokter).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": dokter})
}

// DeleteDokter godoc
// @Summary Delete one Dokter.
// @Description Delete a Dokter by id.
// @Tags Dokter
// @Produce json
// @Param id path string true "Dokter id"
// @Success 200 {object} map[string]boolean
// @Router /dokter/{id} [delete]
func DeleteDokter(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var dokter models.Dokter
    if err := db.Where("id = ?", c.Param("id")).First(&dokter).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&dokter)

    c.JSON(http.StatusOK, gin.H{"data": true})
}