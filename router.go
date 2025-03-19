package main

import "github.com/gin-gonic/gin"

func InitRouter() {
	// Initialize the router
	router := gin.Default()

	// Define the route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server
	router.Run(":8080")
}
