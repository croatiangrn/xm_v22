package http

import (
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

func InitRouter(cfg config.Config) {
	router := gin.Default()

	// Define the route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server
	router.Run(cfg.ServerPort)
}
