package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var u Students
	err := db.QueryRow("SELECT * FROM room WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Email, &u.Mobile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else {
		_, err := db.Exec("DELETE FROM room WHERE id = $1", id)
		if err != nil {
			//todo : fix error handling
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
