package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// GET /ping → {"message": "pong"}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST /hello → {"name": "..."} 받기
	router.POST("/hello", func(c *gin.Context) {
		var body struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + body.Name,
		})
	})

	// 서버 실행
	router.Run(":8080") // http://localhost:8080
}
