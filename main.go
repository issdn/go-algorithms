package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Id string `json:"id"`
}

func main() {
	window := FixedWindow{
		Interval:     1,
		UserRequests: map[string]uint16{},
	}

	r := gin.Default()
	// IP Will be used.
	r.POST("/", func(c *gin.Context) {
		var request UserRequest
		if err := c.BindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
			return
		}
		_, ok := window.UserRequests[request.Id]
		fmt.Println(ok)
		userId, requests := window.Request(request.Id)
		c.JSON(http.StatusOK, gin.H{
			"userId":   userId,
			"requests": requests,
		})
	})
	r.Run()
}
