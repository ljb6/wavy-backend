package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/internal/user"
	"github.com/ljb6/wavy-backend/internal/middleware"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	// check auth
	router.GET("auth/check", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		userID := ctx.GetString("userID")
		ctx.JSON(200, gin.H{"userID": userID})
	})

	// user
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
	private.GET("/auth/check")
}