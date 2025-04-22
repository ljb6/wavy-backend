package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/internal/middleware"
	"github.com/ljb6/wavy-backend/internal/user"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	// user
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// check auth
	router.GET("auth/check", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		userID := ctx.GetString("userID")

		user, err := userRepo.GetUserDataByID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"userID": user.ID,
			"name": user.Name,
			"email": user.Email,
			"plan": user.Plan,
			"created_at": user.CreatedAt,
		})
	})

	// public router
	public := router.Group("/user")
	public.POST("/register", userHandler.RegisterHandler)
	public.POST("/login", userHandler.LoginHandler)

	// private router
	private := router.Group("/private")
	private.Use(middleware.AuthMiddleware())
	private.GET("/auth/check")
}