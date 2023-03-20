package main

import (
	"exampleProject/functions"
	"exampleProject/pkg"
	"exampleProject/types"
	"github.com/gin-gonic/gin"
)

func main() {

	types.CreateSlices()
	types.CopyData()
	types.GetSliceData()
	types.MapOperations()
	types.IterateWithRanges()
	types.CreateArray()

	functions.CallMultipleReturns()
	functions.CallVariadic()
	functions.CallPointerValue()

	router := gin.Default()
	router.GET("/locations/init", pkg.InitLocations)
	router.GET("/locations/get", pkg.GetLocations)
	router.POST("/locations/append", pkg.AddLocation)
	_ = router.Run("localhost:8000")

}
