package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabpn12/Rest_API_Test.git/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("api/products", productController.Index)
	r.POST("api/products/:id", productController.Show)
	r.POST("api/products", productController.Store)
	r.PUT("api/products/:id", productController.Update)
	r.DELETE("api/products/:id", productController.Destroy)

}
