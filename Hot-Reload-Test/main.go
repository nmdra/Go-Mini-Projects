package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Serve static files (CSS, JS, images)
	r.Static("/static", "./static")

	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/*")

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "about.html", nil)
	})


	PORT := "8090"

	// Start the server
	r.Run(":" + PORT) // Listen and serve on http://localhost:8080
}
