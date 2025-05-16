package controller

import (
	"github.com/gin-gonic/gin"
	"mockdb/mockdb"
	"mockdb/model"
	"net/http"
)

func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, mockdb.Orders)
}

func CreateOrder(c *gin.Context) {
	var newOrder model.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newOrder.ID = len(mockdb.Orders) + 1

	total := 0.0
	for _, item := range newOrder.Items {
		for _, p := range mockdb.Products {
			if p.ID == item.ProductID {
				total += p.Price * float64(item.Quantity)
			}
		}
	}
	newOrder.Total = total
	mockdb.Orders = append(mockdb.Orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}
