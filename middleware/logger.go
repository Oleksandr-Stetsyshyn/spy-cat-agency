package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			c.Request.ParseForm()
			if salary := c.Request.Form.Get("salary"); salary != "" {
				c.Request.Form.Set("salary", "****")
			}
			if notes := c.Request.Form.Get("notes"); notes != "" {
				c.Request.Form.Set("notes", "****")
			}
		}

		statusCode := c.Writer.Status()
		logLevel := "INFO"
		if statusCode >= 400 {
			logLevel = "ERROR"
		} else if statusCode >= 300 {
			logLevel = "WARN"
		}

		log.Printf("[%s] %s %s %s %d %s",
			logLevel,
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			statusCode,
			latency,
		)
	}
}
