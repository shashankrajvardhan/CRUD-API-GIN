package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range students {
		if strconv.Itoa(user.ID) == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
