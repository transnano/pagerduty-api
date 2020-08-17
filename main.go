package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v1.0.20200817",
		})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/v2/enqueue", func(c *gin.Context) {
		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		b := string(buf[0:n])
		fmt.Println(b)
		num := len(b)
		switch {
		case num == 0:
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "invalid event",
				"message": "Event object is invalid",
				"errors":  "Length of 'routing_key' is incorrect (should be 32 characters)",
			})
		case num > 500:
			c.AbortWithStatus(http.StatusTooManyRequests)
		default:
			c.JSON(http.StatusAccepted, gin.H{
				"status":    "success",
				"message":   "Event processed",
				"dedup_key": "samplekeyhere",
			})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
