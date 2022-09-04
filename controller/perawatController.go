package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type perawatInput struct {
    Nama  string `gorm:"not null" json:"nama"`
}

// GetAllPerawat godoc
// @Summary Get all Perawat.
// @Description Get a list of Perawat.
// @Tags Perawat
// @Produce json
// @Success 200 {object} []models.Perawat
// @Router /perawat [get]
func GetAllPerawat(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var perawat []models.Perawat
    db.Find(&perawat)

    c.JSON(http.StatusOK, gin.H{"data": perawat})
}

// CreatePerawat godoc
// @Summary Create New Perawat.
// @Description Creating a new Perawat.
// @Tags Perawat
// @Param Body body perawatInput true "the body to create a new Perawat"
// @Produce json
// @Success 200 {object} models.Perawat
// @Router /perawat [post]
func CreatePerawat(c *gin.Context) {
    // Validate input
    var input perawatInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Perawat
    perawat := models.Perawat{Nama: input.Nama}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&perawat)

    c.JSON(http.StatusOK, gin.H{"data": perawat})
}

// GetPerawatById godoc
// @Summary Get Perawat.
// @Description Get an Perawat by id.
// @Tags Perawat
// @Produce json
// @Param id path string true "Perawat id"
// @Success 200 {object} models.Perawat
// @Router /perawat/{id} [get]
func GetPerawatById(c *gin.Context) { // Get model if exist
    var perawat models.Perawat

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&perawat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": perawat})
}

// UpdatePerawat godoc
// @Summary Update Perawat.
// @Description Update Perawat by id.
// @Tags Perawat
// @Produce json
// @Param id path string true "Perawat id"
// @Param Body body perawatInput true "the body to update age perawat category"
// @Success 200 {object} models.Perawat
// @Router /perawat/{id} [patch]
func UpdatePerawat(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var perawat models.Perawat
    if err := db.Where("id = ?", c.Param("id")).First(&perawat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input perawatInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Perawat
    updatedInput.Nama = input.Nama

    db.Model(&perawat).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": perawat})
}

// DeletePerawat godoc
// @Summary Delete one Perawat.
// @Description Delete a Perawat by id.
// @Tags Perawat
// @Produce json
// @Param id path string true "Perawat id"
// @Success 200 {object} map[string]boolean
// @Router /perawat/{id} [delete]
func DeletePerawat(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var perawat models.Perawat
    if err := db.Where("id = ?", c.Param("id")).First(&perawat).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&perawat)

    c.JSON(http.StatusOK, gin.H{"data": true})
}