package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/internal/user"
	"github.com/ljb6/wavy-backend/middleware"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// public router
	public := router.Group("/user")
	public.POST("/register", userHandler.RegisterHandler)
	public.POST("/login", userHandler.LoginHandler)

	// private router
	private := router.Group("/private")
	private.Use(middleware.AuthMiddleware())
	//private.GET("/profile")
}