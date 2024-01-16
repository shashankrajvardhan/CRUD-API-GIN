package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range students {
		if strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
