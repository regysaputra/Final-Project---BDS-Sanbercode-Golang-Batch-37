package controller

import (
    "net/http"
    "time"
    "final-project/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type appointmentInput struct {
    Mode         string     `gorm:"not null" json:"mode"`
    Tanggal	     time.Time	`json:"tanggal"`
	PasienID     uint     `gorm:"not null" json:"pasienID"`
}

// GetAllAppointment godoc
// @Summary Get all Appointment.
// @Description Get a list of Appointment.
// @Tags Appointment
// @Produce json
// @Success 200 {object} []models.Appointment
// @Router /appointment [get]
func GetAllAppointment(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var appointment []models.Appointment
    db.Find(&appointment)

    c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// CreateAppointment godoc
// @Summary Create New Appointment.
// @Description Creating a new Appointment.
// @Tags Appointment
// @Param Body body appointmentInput true "the body to create a new Appointment"
// @Produce json
// @Success 200 {object} models.Appointment
// @Router /appointment [post]
func CreateAppointment(c *gin.Context) {
    // Validate input
    var input appointmentInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Appointment
    appointment := models.Appointment{Mode: input.Mode, Tanggal: input.Tanggal, PasienID: input.PasienID}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&appointment)

    c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GetAppointmentById godoc
// @Summary Get Appointment.
// @Description Get an Appointment by id.
// @Tags Appointment
// @Produce json
// @Param id path string true "Appointment id"
// @Success 200 {object} models.Appointment
// @Router /appointment/{id} [get]
func GetAppointmentById(c *gin.Context) { // Get model if exist
    var appointment models.Appointment

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// UpdateAppointment godoc
// @Summary Update Appointment.
// @Description Update Appointment by id.
// @Tags Appointment
// @Produce json
// @Param id path string true "Appointment id"
// @Param Body body appointmentInput true "the body to update age appointment category"
// @Success 200 {object} models.Appointment
// @Router /appointment/{id} [patch]
func UpdateAppointment(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var appointment models.Appointment
    if err := db.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input appointmentInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Appointment
    updatedInput.Mode = input.Mode
    updatedInput.Tanggal = input.Tanggal
	updatedInput.PasienID = input.PasienID

    db.Model(&appointment).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// DeleteAppointment godoc
// @Summary Delete one Appointment.
// @Description Delete a Appointment by id.
// @Tags Appointment
// @Produce json
// @Param id path string true "Appointment id"
// @Success 200 {object} map[string]boolean
// @Router /appointment/{id} [delete]
func DeleteAppointment(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var appointment models.Appointment
    if err := db.Where("id = ?", c.Param("id")).First(&appointment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&appointment)

    c.JSON(http.StatusOK, gin.H{"data": true})
}