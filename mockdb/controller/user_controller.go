package controller

import (
	"github.com/gin-gonic/gin"
	"mockdb/mockdb"
	"mockdb/model"
	"net/http"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, mockdb.Users)
}

func AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = len(mockdb.Users) + 1
	mockdb.Users = append(mockdb.Users, user)
	c.JSON(http.StatusCreated, user)
}
