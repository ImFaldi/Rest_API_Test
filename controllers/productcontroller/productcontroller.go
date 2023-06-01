package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indrabpn12/Rest_API_Test.git/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Store(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error!"})
			return
		}
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Destroy(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error!"})
			return
		}
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted!"})
}
