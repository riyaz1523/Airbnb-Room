package main

import (
	"Airbnb-room/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitDB();
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
