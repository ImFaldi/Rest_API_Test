package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indrabpn12/Rest_API_Test.git/models"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {

}

func Store(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Destroy(c *gin.Context) {

}
