package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type obatInput struct {
    Nama		 string       `gorm:"not null" json:"nama"`
	Satuan		 string       `gorm:"not null" json:"satuan"`
	Harga        string       `gorm:"not null" json:"harga"`
	Keterangan   string       `json:"keterangan"`
}

// GetAllObat godoc
// @Summary Get all Obat.
// @Description Get a list of Obat.
// @Tags Obat
// @Produce json
// @Success 200 {object} []models.Obat
// @Router /obat [get]
func GetAllObat(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var obat []models.Obat
    db.Find(&obat)

    c.JSON(http.StatusOK, gin.H{"data": obat})
}

// CreateObat godoc
// @Summary Create New Obat.
// @Description Creating a new Obat.
// @Tags Obat
// @Param Body body obatInput true "the body to create a new Obat"
// @Produce json
// @Success 200 {object} models.Obat
// @Router /obat [post]
func CreateObat(c *gin.Context) {
    // Validate input
    var input obatInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Obat
    obat := models.Obat{Nama: input.Nama, Satuan: input.Satuan, Harga: input.Harga, Keterangan: input.Keterangan}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&obat)

    c.JSON(http.StatusOK, gin.H{"data": obat})
}

// GetObatById godoc
// @Summary Get Obat.
// @Description Get an Obat by id.
// @Tags Obat
// @Produce json
// @Param id path string true "Obat id"
// @Success 200 {object} models.Obat
// @Router /obat/{id} [get]
func GetObatById(c *gin.Context) { // Get model if exist
    var obat models.Obat

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&obat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": obat})
}

// UpdateObat godoc
// @Summary Update Obat.
// @Description Update Obat by id.
// @Tags Obat
// @Produce json
// @Param id path string true "Obat id"
// @Param Body body obatInput true "the body to update age obat category"
// @Success 200 {object} models.Obat
// @Router /obat/{id} [patch]
func UpdateObat(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var obat models.Obat
    if err := db.Where("id = ?", c.Param("id")).First(&obat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input obatInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Obat
    updatedInput.Nama = input.Nama
    updatedInput.Satuan = input.Satuan
	updatedInput.Harga = input.Harga
	updatedInput.Keterangan = input.Keterangan

    db.Model(&obat).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": obat})
}

// DeleteObat godoc
// @Summary Delete one Obat.
// @Description Delete a Obat by id.
// @Tags Obat
// @Produce json
// @Param id path string true "Obat id"
// @Success 200 {object} map[string]boolean
// @Router /obat/{id} [delete]
func DeleteObat(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var obat models.Obat
    if err := db.Where("id = ?", c.Param("id")).First(&obat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&obat)

    c.JSON(http.StatusOK, gin.H{"data": true})
}