package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/internal/user"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	
	

	userGroup := router.Group("/user")
	userGroup.POST("/register", userHandler.RegisterHandler)
	// userGroup.POST("/login", func(ctx *gin.Context) { handlers.LoginUser(ctx, db) })
}