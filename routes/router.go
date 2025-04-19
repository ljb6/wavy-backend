package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeServer(db *sql.DB) {
	router := gin.Default()

	InitializeRoutes(router, db)

	router.Run() // listen and serve on 0.0.0.0:8080
}