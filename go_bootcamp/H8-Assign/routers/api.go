package routers

import (
	"go_bootcamp/H8-Assign/controllers"

	"github.com/gin-gonic/gin"

	// "go_bootcamp/H8-swagger/controllers"
	// "github.com/gin-gonic/gin"
	_ "go_bootcamp/H8-Assign/docs"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// "go_bootcamp/controller"

func StartServer() *gin.Engine {
	router := gin.Default()

	// router.GET("/cars/:id", controllers.GetOneCars)

	router.POST("/create", controllers.CreateItems)

	router.GET("/items", controllers.GetAllItems)

	router.PUT("/api/orders/:id", controllers.UpdateItems)

	router.DELETE("/order/:id", controllers.DeleteOrder)

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
