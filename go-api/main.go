package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tayalone/example-gin-101/go-api/example"
	"github.com/tayalone/example-gin-101/go-api/middlewares"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// fmt.Println("rdb", rdb)

	// var dsnBuilder strings.Builder
	// fmt.Fprintf(&dsnBuilder, "host=db user=%q password=%q dbname=%q port=5432 sslmode=disable TimeZone=Asia/Bangkok", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	// dsn := dsnBuilder.String()

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// fmt.Println("db", db)
	// fmt.Println("err", err)

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

	r.GET("/env", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messge":      "OK",
			"TEST_GO_ENV": os.Getenv("TEST_GO_ENV"),
		})
	})

	r.GET("/handling-panic", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"messge": "error because go lang panic",
				})

			}
		}()
		panic("test panic")
	})

	r.GET("/test-middleware", middlewares.Example1(), func(c *gin.Context) {
		example1, existing := c.Get("example1")
		if !existing {
			c.JSON(http.StatusInternalServerError, gin.H{
				"messge": "example1 does not existing",
			})
			return
		}
		c.JSON(200, gin.H{
			"messge":   "OK",
			"example1": example1,
		})
	})

	r.GET("/redis-test-set", func(c *gin.Context) {
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"messge": "can not set value @ redis",
			})
			return
		}
		c.JSON(200, gin.H{
			"messge": "OK",
		})
	})

	r.GET("/redis-test-get", func(c *gin.Context) {
		res, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			fmt.Println("err", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"messge": "can not get value @ redis",
			})
			return
		}
		c.JSON(200, gin.H{
			"messge": "OK",
			"key":    res,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
