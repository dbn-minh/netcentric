package controller

import (
	"github.com/gin-gonic/gin"
	"mockdb/mockdb"
	"mockdb/model"
	"net/http"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, mockdb.Products)
}

func AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = len(mockdb.Products) + 1
	mockdb.Products = append(mockdb.Products, product)
	c.JSON(http.StatusCreated, product)
}
