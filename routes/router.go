package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeServer(db *sql.DB) {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "working",
		})
	})

	InitializeRoutes(router, db)

	router.Run() // listen and serve on 0.0.0.0:8080
}