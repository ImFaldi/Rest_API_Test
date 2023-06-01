package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indrabpn12/Rest_API_Test.git/controllers/productcontroller"
	"github.com/indrabpn12/Rest_API_Test.git/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show)
	r.POST("api/products", productcontroller.Store)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products/:id", productcontroller.Destroy)

	r.Run()

}
