package main

import (
	"pincodeapp/controlers"
	"pincodeapp/initializers"
	"pincodeapp/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	models.InitDatabase()
	r := gin.Default()
	r.StaticFile("/", "./index.html")

	// CRUD API Endpoints
	r.GET("/pincodes", controlers.GetPincodes)
	r.POST("/pincodes", controlers.CreatePincode)
	r.PUT("/pincodes/:id", controlers.UpdatePincode)
	r.DELETE("/pincodes/:id", controlers.DeletePincode)
	r.Run()
}
