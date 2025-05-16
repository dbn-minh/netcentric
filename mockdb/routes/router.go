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

	// User
	r.GET("/users", controller.GetUsers)

	// Products
	r.GET("/products", controller.GetProducts)

	return r
}
