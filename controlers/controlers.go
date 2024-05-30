package controlers

import (
	"net/http"
	"pincodeapp/models"

	"github.com/gin-gonic/gin"
)

// Get Function
func GetPincodes(c *gin.Context) {
	var pincodes []models.Pincode
	query := models.DB

	if pincode := c.Query("pincode"); pincode != "" {
		query = query.Where("pincode = ?", pincode)
	}
	if district := c.Query("district"); district != "" {
		query = query.Where("districtname = ?", district)
	}
	if state := c.Query("state"); state != "" {
		query = query.Where("statename = ?", state)
	}
	if locality := c.Query("locality"); locality != "" {
		query = query.Where("officename = ?", locality)
	}

	query.Find(&pincodes)
	c.JSON(http.StatusOK, pincodes)
}

// Create a new pincode
func CreatePincode(c *gin.Context) {
	var newPincode models.Pincode
	if err := c.ShouldBindJSON(&newPincode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&newPincode)
	c.JSON(http.StatusCreated, newPincode)
}

// Updating pincode
func UpdatePincode(c *gin.Context) {
	var pincode models.Pincode
	id := c.Param("id")
	if err := models.DB.First(&pincode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pincode not found"})
		return
	}

	var input models.Pincode
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Updating fields
	pincode.Officename = input.Officename
	pincode.OfficeType = input.OfficeType
	pincode.Deliverystatus = input.Deliverystatus

	if err := models.DB.Save(&pincode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pincode"})
		return
	}

	c.JSON(http.StatusOK, pincode)
}

// Deleting a pincode
func DeletePincode(c *gin.Context) {
	var pincode models.Pincode
	id := c.Param("id")
	if err := models.DB.First(&pincode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pincode not found"})
		return
	}
	models.DB.Delete(&pincode)
	c.JSON(http.StatusOK, gin.H{"message": "Pincode deleted successfully"})
}
