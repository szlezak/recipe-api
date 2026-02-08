package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 1. Logic BEFORE the request
		fmt.Printf("--- START: %s %s ---\n", c.Request.Method, c.Request.URL.Path)

		c.Next() // 2. This tells Gin to go run the actual Handler

		// 3. Logic AFTER the request
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("--- END: Status %d | Time: %v ---\n", status, latency)
	}
}