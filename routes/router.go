package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeServer() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "status": "working",
		})
	  })

	router.Run() // listen and serve on 0.0.0.0:8080
}