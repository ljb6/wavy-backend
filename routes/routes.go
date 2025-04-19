package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/handlers"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	userGroup := router.Group("/user")
	userGroup.POST("/register", func(ctx *gin.Context) { handlers.RegisterUser(ctx, db) })
}