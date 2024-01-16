package main

import "github.com/gin-gonic/gin"

func router() {
	r := gin.Default()

	// Initialize the users variable
	students = []Students{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	// Define routes for user API
	r.GET("/users", GetUsers)
	r.GET("/getuser/:id", GetUserByID)
	r.POST("/createuser", CreateUser)
	r.PUT("/updateuser/:id", UpdateUser)
	r.DELETE("/deleteuser/:id", DeleteUser)

	r.Run(":8080")
}
