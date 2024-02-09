package routers

import (
	"go_bootcamp/H8-swagger/controllers"

	"github.com/gin-gonic/gin"

	_ "go_bootcamp/H8-swagger/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// "go_bootcamp/controller"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars/:id", controllers.GetOneCars)

	router.POST("/cars", controllers.CreateCars)

	router.GET("/cars", controllers.GetAllCar)

	router.PATCH("/cars/:id", controllers.UpdateCars)

	router.DELETE("/cars/:id", controllers.DeleteCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
