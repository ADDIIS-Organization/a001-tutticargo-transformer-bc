package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware es un middleware que registra cada solicitud entrante.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()

		log.Printf(
			"%s %s %s %d %s",
			clientIP,
			method,
			path,
			statusCode,
			latency,
		)
	}
}

func main() {
	initDB()

	r := gin.Default()
	r.Use(LoggingMiddleware())

	r.POST("/orders/excel/upload", uploadFile)
	r.GET("/test", test)
	r.GET("/test2", test)

	r.Run(":8081")
}
