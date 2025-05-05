package controller

import (
	"github.com/gin-gonic/gin"
	"mockdb/mockdb"
	"net/http"
)

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, mockdb.Categories)
}
