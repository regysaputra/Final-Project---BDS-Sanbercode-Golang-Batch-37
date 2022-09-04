package controller

import (
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
	"time"
)

type rawatInapInput struct {
    TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`
	PasienID	  uint 		`json:"pasienID"`
	DokterID	  uint		`json:"dokterInapID"`
	PerawatID     uint 		`json:"perawatID"`
	KamarID		  string	`json:"kamarID"`
}

// GetAllRawatInap godoc
// @Summary Get all RawatInap.
// @Description Get a list of RawatInap.
// @Tags RawatInap
// @Produce json
// @Success 200 {object} []models.RawatInap
// @Router /rawatInap [get]
func GetAllRawatInap(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var rawatInap []models.RawatInap
    db.Find(&rawatInap)

    c.JSON(http.StatusOK, gin.H{"data": rawatInap})
}

// CreateRawatInap godoc
// @Summary Create New RawatInap.
// @Description Creating a new RawatInap.
// @Tags RawatInap
// @Param Body body rawatInapInput true "the body to create a new RawatInap"
// @Produce json
// @Success 200 {object} models.RawatInap
// @Router /rawatInap [post]
func CreateRawatInap(c *gin.Context) {
    // Validate input
    var input rawatInapInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create RawatInap
    rawatInap := models.RawatInap{TanggalMasuk: input.TanggalMasuk, TanggalKeluar: input.TanggalKeluar, PasienID: input.PasienID, DokterID: input.DokterID, PerawatID: input.PerawatID, KamarID: input.KamarID}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&rawatInap)

    c.JSON(http.StatusOK, gin.H{"data": rawatInap})
}

// GetRawatInapById godoc
// @Summary Get RawatInap.
// @Description Get an RawatInap by id.
// @Tags RawatInap
// @Produce json
// @Param id path string true "RawatInap id"
// @Success 200 {object} models.RawatInap
// @Router /rawatInap/{id} [get]
func GetRawatInapById(c *gin.Context) { // Get model if exist
    var rawatInap models.RawatInap

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&rawatInap).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": rawatInap})
}

// UpdateRawatInap godoc
// @Summary Update RawatInap.
// @Description Update RawatInap by id.
// @Tags RawatInap
// @Produce json
// @Param id path string true "RawatInap id"
// @Param Body body rawatInapInput true "the body to update age rawatInap category"
// @Success 200 {object} models.RawatInap
// @Router /rawatInap/{id} [patch]
func UpdateRawatInap(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var rawatInap models.RawatInap
    if err := db.Where("id = ?", c.Param("id")).First(&rawatInap).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input rawatInapInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.RawatInap
    updatedInput.TanggalMasuk = input.TanggalMasuk
    updatedInput.TanggalKeluar = input.TanggalKeluar
	updatedInput.PasienID = input.PasienID
	updatedInput.DokterID = input.DokterID
	updatedInput.PerawatID = input.PerawatID
	updatedInput.KamarID = input.KamarID

    db.Model(&rawatInap).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": rawatInap})
}

// DeleteRawatInap godoc
// @Summary Delete one RawatInap.
// @Description Delete a RawatInap by id.
// @Tags RawatInap
// @Produce json
// @Param id path string true "RawatInap id"
// @Success 200 {object} map[string]boolean
// @Router /rawatInap/{id} [delete]
func DeleteRawatInap(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var rawatInap models.RawatInap
    if err := db.Where("id = ?", c.Param("id")).First(&rawatInap).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&rawatInap)

    c.JSON(http.StatusOK, gin.H{"data": true})
}