package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Example1 is a my experiment
func Example1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example1", "lovely middlewares")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
