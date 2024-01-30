package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func partial(c *gin.Context) {
	username := c.Param("username")

	var u Students
	err := db.QueryRow("SELECT * FROM room WHERE username LIKE  $1", "%"+username+"%").Scan(&u.ID, &u.Username, &u.Email, &u.Mobile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}
