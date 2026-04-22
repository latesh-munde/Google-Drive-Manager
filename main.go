package main

import (
	"log"

	"gin-drive-manager/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	auth := r.Group("/auth")
	{
		auth.GET("/login", handlers.Login)
		auth.GET("/callback", handlers.Callback)
	}

	drive := r.Group("/drive")
	{
		drive.GET("/files", handlers.ListFiles)
	}

	log.Println("Server running at :8080")
	r.Run(":8080")
}
