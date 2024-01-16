package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range students {
		if strconv.Itoa(user.ID) == id {
			var updatedUser Students

			if err := c.BindJSON(&updatedUser); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = user.ID
			students[i] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
