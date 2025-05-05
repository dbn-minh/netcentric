package routes

import (
	"github.com/gin-gonic/gin"
	"mockdb/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Orders
	r.GET("/orders", controller.GetAllOrders)
	r.POST("/orders", controller.CreateOrder)

	// Add more: /users, /products, /categories

	return r
}
