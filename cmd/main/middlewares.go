package main

import "github.com/gin-gonic/gin"

func HtmxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isHtmxRequest := c.GetHeader("hx-request") == "true"
		c.Set("htmxRequest", isHtmxRequest)

		c.Next()
	}
}
