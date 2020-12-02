package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tayalone/example-gin-101/go-api/example"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/simple-pi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": example.SimplePi,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
