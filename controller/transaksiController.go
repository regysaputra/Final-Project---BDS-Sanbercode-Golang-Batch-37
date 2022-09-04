package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type transaksiInput struct {
    DiagnosaID  uint      `gorm:"not null" json:"diagnosaID"`
	BiayaObat   string    `json:"biaya_obat"`
	BiayaDokter string    `json:"biaya_dokter"`
	BiayaKamar  string    `json:"biaya_kamar"`
	Total 		string    `json:"total"`
}

// GetAllTransaksi godoc
// @Summary Get all Transaksi.
// @Description Get a list of Transaksi.
// @Tags Transaksi
// @Produce json
// @Success 200 {object} []models.Transaksi
// @Router /transaksi [get]
func GetAllTransaksi(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var transaksi []models.Transaksi
    db.Find(&transaksi)

    c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// CreateTransaksi godoc
// @Summary Create New Transaksi.
// @Description Creating a new Transaksi.
// @Tags Transaksi
// @Param Body body transaksiInput true "the body to create a new Transaksi"
// @Produce json
// @Success 200 {object} models.Transaksi
// @Router /transaksi [post]
func CreateTransaksi(c *gin.Context) {
    // Validate input
    var input transaksiInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Transaksi
    transaksi := models.Transaksi{DiagnosaID: input.DiagnosaID, BiayaObat: input.BiayaObat, BiayaDokter: input.BiayaDokter, BiayaKamar: input.BiayaKamar, Total: input.Total}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&transaksi)

    c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// GetTransaksiById godoc
// @Summary Get Transaksi.
// @Description Get an Transaksi by id.
// @Tags Transaksi
// @Produce json
// @Param id path string true "Transaksi id"
// @Success 200 {object} models.Transaksi
// @Router /transaksi/{id} [get]
func GetTransaksiById(c *gin.Context) { // Get model if exist
    var transaksi models.Transaksi

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&transaksi).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// UpdateTransaksi godoc
// @Summary Update Transaksi.
// @Description Update Transaksi by id.
// @Tags Transaksi
// @Produce json
// @Param id path string true "Transaksi id"
// @Param Body body transaksiInput true "the body to update age transaksi category"
// @Success 200 {object} models.Transaksi
// @Router /transaksi/{id} [patch]
func UpdateTransaksi(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var transaksi models.Transaksi
    if err := db.Where("id = ?", c.Param("id")).First(&transaksi).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input transaksiInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Transaksi
    updatedInput.DiagnosaID = input.DiagnosaID
    updatedInput.BiayaObat = input.BiayaObat
	updatedInput.BiayaDokter = input.BiayaDokter
	updatedInput.BiayaKamar = input.BiayaKamar
	updatedInput.Total = input.Total

    db.Model(&transaksi).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": transaksi})
}

// DeleteTransaksi godoc
// @Summary Delete one Transaksi.
// @Description Delete a Transaksi by id.
// @Tags Transaksi
// @Produce json
// @Param id path string true "Transaksi id"
// @Success 200 {object} map[string]boolean
// @Router /transaksi/{id} [delete]
func DeleteTransaksi(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var transaksi models.Transaksi
    if err := db.Where("id = ?", c.Param("id")).First(&transaksi).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&transaksi)

    c.JSON(http.StatusOK, gin.H{"data": true})
}