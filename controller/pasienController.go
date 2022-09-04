package controller

import (
    "net/http"
    "time"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type pasienInput struct {
    Nama	     string	`json:"nama"`
    JenisKelamin string `gorm:"not null" json:"jenis_kelamin"`
    NoHP         string `gorm:"unique" json:"no_hp"`
	Alamat       string `json:"alamat"`
	TanggalLahir time.Time `gorm:"not null" json:"tanggal_lahir"`
}

// GetAllPasien godoc
// @Summary Get all Pasien.
// @Description Get a list of Pasien.
// @Tags Pasien
// @Produce json
// @Success 200 {object} []models.Pasien
// @Router /pasien [get]
func GetAllPasien(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var pasien []models.Pasien
    db.Find(&pasien)

    c.JSON(http.StatusOK, gin.H{"data": pasien})
}

// CreatePasien godoc
// @Summary Create New Pasien.
// @Description Creating a new Pasien.
// @Tags Pasien
// @Param Body body pasienInput true "the body to create a new Pasien"
// @Produce json
// @Success 200 {object} models.Pasien
// @Router /pasien [post]
func CreatePasien(c *gin.Context) {
    // Validate input
    var input pasienInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Pasien
    pasien := models.Pasien{Nama: input.Nama, JenisKelamin: input.JenisKelamin, NoHP: input.NoHP, Alamat: input.Alamat, TanggalLahir: input.TanggalLahir}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&pasien)

    c.JSON(http.StatusOK, gin.H{"data": pasien})
}

// GetPasienById godoc
// @Summary Get Pasien.
// @Description Get an Pasien by id.
// @Tags Pasien
// @Produce json
// @Param id path string true "Pasien id"
// @Success 200 {object} models.Pasien
// @Router /pasien/{id} [get]
func GetPasienById(c *gin.Context) { // Get model if exist
    var pasien models.Pasien

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&pasien).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": pasien})
}

// UpdatePasien godoc
// @Summary Update Pasien.
// @Description Update Pasien by id.
// @Tags Pasien
// @Produce json
// @Param id path string true "Pasien id"
// @Param Body body pasienInput true "the body to update age pasien category"
// @Success 200 {object} models.Pasien
// @Router /pasien/{id} [patch]
func UpdatePasien(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var pasien models.Pasien
    if err := db.Where("id = ?", c.Param("id")).First(&pasien).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input pasienInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Pasien
    updatedInput.Nama = input.Nama
    updatedInput.JenisKelamin = input.JenisKelamin
	updatedInput.NoHP = input.NoHP
	updatedInput.Alamat = input.Alamat
	updatedInput.TanggalLahir = input.TanggalLahir
    updatedInput.UpdatedAt = time.Now()

    db.Model(&pasien).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": pasien})
}

// DeletePasien godoc
// @Summary Delete one Pasien.
// @Description Delete a Pasien by id.
// @Tags Pasien
// @Produce json
// @Param id path string true "Pasien id"
// @Success 200 {object} map[string]boolean
// @Router /pasien/{id} [delete]
func DeletePasien(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var pasien models.Pasien
    if err := db.Where("id = ?", c.Param("id")).First(&pasien).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&pasien)

    c.JSON(http.StatusOK, gin.H{"data": true})
}