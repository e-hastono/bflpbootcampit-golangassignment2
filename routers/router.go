package routers

import (
	"github.com/e-hastono/bflpbootcamptit-golangassignment2/controllers"
	"github.com/e-hastono/bflpbootcamptit-golangassignment2/models"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	models.StartDB()
	r := gin.Default()

	groupedR := r.Group("/api/v1/orders")
	{
		groupedR.POST("/", controllers.CreateOrder)
		groupedR.GET("/", controllers.GetAllData)
		groupedR.PATCH("/:orderid", controllers.UpdateOrderPatch)
		groupedR.PUT("/:orderid", controllers.UpdateOrderPut)
		groupedR.DELETE("/:orderid", controllers.DeleteOrder)
	}

	return r
}
