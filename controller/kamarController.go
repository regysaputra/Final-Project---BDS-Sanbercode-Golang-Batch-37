package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type kamarInput struct {
    NoKamar     string      `gorm:"not null" json:"no_kamar"`
    Lantai		string		`gorm:"not null" json:"lantai"`
	Type        string		`gorm:"not null" json:"type"`
	Harga		string		`gorm:"not null" json:"harga"`
	Status		bool		`gorm:"not null" json:"status"`
}

// GetAllKamar godoc
// @Summary Get all Kamar.
// @Description Get a list of Kamar.
// @Tags Kamar
// @Produce json
// @Success 200 {object} []models.Kamar
// @Router /kamar [get]
func GetAllKamar(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var kamar []models.Kamar
    db.Find(&kamar)

    c.JSON(http.StatusOK, gin.H{"data": kamar})
}

// CreateKamar godoc
// @Summary Create New Kamar.
// @Description Creating a new Kamar.
// @Tags Kamar
// @Param Body body kamarInput true "the body to create a new Kamar"
// @Produce json
// @Success 200 {object} models.Kamar
// @Router /kamar [post]
func CreateKamar(c *gin.Context) {
    // Validate input
    var input kamarInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Kamar
    kamar := models.Kamar{NoKamar: input.NoKamar, Lantai: input.Lantai, Type: input.Type, Harga: input.Harga, Status: input.Status}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&kamar)

    c.JSON(http.StatusOK, gin.H{"data": kamar})
}

// GetKamarById godoc
// @Summary Get Kamar.
// @Description Get an Kamar by id.
// @Tags Kamar
// @Produce json
// @Param id path string true "Kamar id"
// @Success 200 {object} models.Kamar
// @Router /kamar/{id} [get]
func GetKamarById(c *gin.Context) { // Get model if exist
    var kamar models.Kamar

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&kamar).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": kamar})
}

// UpdateKamar godoc
// @Summary Update Kamar.
// @Description Update Kamar by id.
// @Tags Kamar
// @Produce json
// @Param id path string true "Kamar id"
// @Param Body body kamarInput true "the body to update age kamar category"
// @Success 200 {object} models.Kamar
// @Router /kamar/{id} [patch]
func UpdateKamar(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var kamar models.Kamar
    if err := db.Where("id = ?", c.Param("id")).First(&kamar).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input kamarInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Kamar
    updatedInput.NoKamar = input.NoKamar
    updatedInput.Lantai = input.Lantai
    updatedInput.Type = input.Type
	updatedInput.Harga = input.Harga
	updatedInput.Status = input.Status

    db.Model(&kamar).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": kamar})
}

// DeleteKamar godoc
// @Summary Delete one Kamar.
// @Description Delete a Kamar by id.
// @Tags Kamar
// @Produce json
// @Param id path string true "Kamar id"
// @Success 200 {object} map[string]boolean
// @Router /kamar/{id} [delete]
func DeleteKamar(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var kamar models.Kamar
    if err := db.Where("id = ?", c.Param("id")).First(&kamar).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&kamar)

    c.JSON(http.StatusOK, gin.H{"data": true})
}