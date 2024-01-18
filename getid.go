package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var u Students
	err := db.QueryRow("SELECT * FROM room WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Email, &u.Mobile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, u)

}
