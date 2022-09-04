package controller

import (
	"time"
    "net/http"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type diagnosaInput struct {
    Tanggal	      time.Time `json:"tanggal"`
	BeratBadan    float64   `json:"berat_badan"`
	TekananDarah  string    `json:"tekanan_darah"`
	Keluhan	      string    `json:"keluhan"`
	HasilDiagnosa string    `json:"hasil_diagnosa"`
	Tindakan	  string    `json:"tindakan"`
	DokterID	  uint      `json:"dokterID"`
	PasienID	  uint		`json:"pasienID"`
}

// GetAllDiagnosa godoc
// @Summary Get all Diagnosa.
// @Description Get a list of Diagnosa.
// @Tags Diagnosa
// @Produce json
// @Success 200 {object} []models.Diagnosa
// @Router /diagnosa [get]
func GetAllDiagnosa(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var diagnosa []models.Diagnosa
    db.Find(&diagnosa)

    c.JSON(http.StatusOK, gin.H{"data": diagnosa})
}

// CreateDiagnosa godoc
// @Summary Create New Diagnosa.
// @Description Creating a new Diagnosa.
// @Tags Diagnosa
// @Param Body body diagnosaInput true "the body to create a new Diagnosa"
// @Produce json
// @Success 200 {object} models.Diagnosa
// @Router /diagnosa [post]
func CreateDiagnosa(c *gin.Context) {
    // Validate input
    var input diagnosaInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Diagnosa
    diagnosa := models.Diagnosa{Tanggal: input.Tanggal, BeratBadan: input.BeratBadan, TekananDarah: input.TekananDarah, Keluhan: input.Keluhan, HasilDiagnosa: input.HasilDiagnosa, Tindakan: input.Tindakan, DokterID: input.DokterID, PasienID: input.PasienID}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&diagnosa)

    c.JSON(http.StatusOK, gin.H{"data": diagnosa})
}

// GetDiagnosaById godoc
// @Summary Get Diagnosa.
// @Description Get an Diagnosa by id.
// @Tags Diagnosa
// @Produce json
// @Param id path string true "Diagnosa id"
// @Success 200 {object} models.Diagnosa
// @Router /diagnosa/{id} [get]
func GetDiagnosaById(c *gin.Context) { // Get model if exist
    var diagnosa models.Diagnosa

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&diagnosa).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": diagnosa})
}

// UpdateDiagnosa godoc
// @Summary Update Diagnosa.
// @Description Update Diagnosa by id.
// @Tags Diagnosa
// @Produce json
// @Param id path string true "Diagnosa id"
// @Param Body body diagnosaInput true "the body to update age diagnosa category"
// @Success 200 {object} models.Diagnosa
// @Router /diagnosa/{id} [patch]
func UpdateDiagnosa(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var diagnosa models.Diagnosa
    if err := db.Where("id = ?", c.Param("id")).First(&diagnosa).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input diagnosaInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Diagnosa
    updatedInput.Tanggal = input.Tanggal
	updatedInput.BeratBadan = input.BeratBadan
	updatedInput.TekananDarah = input.TekananDarah
	updatedInput.Keluhan = input.Keluhan
	updatedInput.HasilDiagnosa = input.HasilDiagnosa
	updatedInput.Tindakan = input.Tindakan
	updatedInput.DokterID = input.DokterID
	updatedInput.PasienID = input.PasienID

    db.Model(&diagnosa).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": diagnosa})
}

// DeleteDiagnosa godoc
// @Summary Delete one Diagnosa.
// @Description Delete a Diagnosa by id.
// @Tags Diagnosa
// @Produce json
// @Param id path string true "Diagnosa id"
// @Success 200 {object} map[string]boolean
// @Router /diagnosa/{id} [delete]
func DeleteDiagnosa(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var diagnosa models.Diagnosa
    if err := db.Where("id = ?", c.Param("id")).First(&diagnosa).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&diagnosa)

    c.JSON(http.StatusOK, gin.H{"data": true})
}