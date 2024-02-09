package routes

import (
	// "go_bootcamp/controller"

	"go_bootcamp/H8-Gin/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controller.CreateCar)

	router.PUT("/cars/:carId", controller.UpdateCar)

	router.GET("/cars/getCar", controller.GetCar)

	router.GET("/cars/:carID", controller.DeleteCar)

	return router
}
