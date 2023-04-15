package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	limiter := NewSlidingWindow(5, 1)

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		requests, err := limiter.Request(c.Request.Header.Get("X-Real-IP"))
		if err != nil {
			if rateLimitErr, ok := err.(*RateLimitError); ok {
				c.AbortWithError(rateLimitErr.code, err)
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"requests": requests,
		})

	})
	r.Run()
}
