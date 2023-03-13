package main

import (
	"exampleProject/pkg"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/locations/init", pkg.InitLocations)
	router.GET("/locations/get", pkg.GetLocations)
	router.POST("/locations/append", pkg.AddLocation)
	router.Run("localhost:8000")

}
