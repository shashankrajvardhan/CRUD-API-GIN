package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var u Students
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	id := c.Param("id")
	_, err := db.Exec("UPDATE room SET username = $1, email = $2, mobile = $3 WHERE id = $4 ", u.Username, u.Email, u.Mobile, id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, u)
}
