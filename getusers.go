package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, username, email, mobile FROM room")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var students []Students
	for rows.Next() {
		var student Students
		err := rows.Scan(&student.ID, &student.Username, &student.Email, &student.Mobile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}
	c.JSON(http.StatusOK, students)
}
