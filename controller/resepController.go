package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type resepInput struct {
    Jumlah     int       `gorm:"not null" json:"jumlah"`
	Dosis      string    `gorm:"not null" json:"dosis"`
	DiagnosaID uint      `gorm:"not null" json:"diagnosaID"`
	ObatID     uint      `gorm:"not null" json:"obatID"`
}

// GetAllResep godoc
// @Summary Get all Resep.
// @Description Get a list of Resep.
// @Tags Resep
// @Produce json
// @Success 200 {object} []models.Resep
// @Router /resep [get]
func GetAllResep(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var resep []models.Resep
    db.Find(&resep)

    c.JSON(http.StatusOK, gin.H{"data": resep})
}

// CreateResep godoc
// @Summary Create New Resep.
// @Description Creating a new Resep.
// @Tags Resep
// @Param Body body resepInput true "the body to create a new Resep"
// @Produce json
// @Success 200 {object} models.Resep
// @Router /resep [post]
func CreateResep(c *gin.Context) {
    // Validate input
    var input resepInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Resep
    resep := models.Resep{Jumlah: input.Jumlah, Dosis: input.Dosis, DiagnosaID: input.DiagnosaID, ObatID: input.ObatID}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&resep)

    c.JSON(http.StatusOK, gin.H{"data": resep})
}

// GetResepById godoc
// @Summary Get Resep.
// @Description Get an Resep by id.
// @Tags Resep
// @Produce json
// @Param id path string true "Resep id"
// @Success 200 {object} models.Resep
// @Router /resep/{id} [get]
func GetResepById(c *gin.Context) { // Get model if exist
    var resep models.Resep

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&resep).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": resep})
}

// UpdateResep godoc
// @Summary Update Resep.
// @Description Update Resep by id.
// @Tags Resep
// @Produce json
// @Param id path string true "Resep id"
// @Param Body body resepInput true "the body to update age resep category"
// @Success 200 {object} models.Resep
// @Router /resep/{id} [patch]
func UpdateResep(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var resep models.Resep
    if err := db.Where("id = ?", c.Param("id")).First(&resep).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input resepInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Resep
    updatedInput.Jumlah = input.Jumlah
    updatedInput.Dosis = input.Dosis
	updatedInput.DiagnosaID = input.DiagnosaID
	updatedInput.ObatID = input.ObatID

    db.Model(&resep).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": resep})
}

// DeleteResep godoc
// @Summary Delete one Resep.
// @Description Delete a Resep by id.
// @Tags Resep
// @Produce json
// @Param id path string true "Resep id"
// @Success 200 {object} map[string]boolean
// @Router /resep/{id} [delete]
func DeleteResep(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var resep models.Resep
    if err := db.Where("id = ?", c.Param("id")).First(&resep).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&resep)

    c.JSON(http.StatusOK, gin.H{"data": true})
}