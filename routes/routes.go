package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/internal/email"
	"github.com/ljb6/wavy-backend/internal/middleware"
	"github.com/ljb6/wavy-backend/internal/subscribers"
	"github.com/ljb6/wavy-backend/internal/user"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	// user
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// subscribers
	subscribersRepo := subscribers.NewRepository(db)
	subscribersService := subscribers.NewService(subscribersRepo)
	subscribersHandler := subscribers.NewHandler(subscribersService)

	// mail
	mailService := email.NewEmailService(userRepo, subscribersRepo)
	mailHandler := email.NewEmailHandler(mailService)

	// check auth
	router.GET("auth/check", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		userID := ctx.GetString("userID")

		user, err := userRepo.GetUserDataByID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"userID":     user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"plan":       user.Plan,
			"created_at": user.CreatedAt,
		})
	})

	// public routes
	public := router.Group("/user")
	public.POST("/register", userHandler.RegisterHandler)
	public.POST("/login", userHandler.LoginHandler)

	// private routes
	private := router.Group("/private")
	private.Use(middleware.AuthMiddleware())

	// GET
	private.GET("/auth/check")
	private.GET("/database/getsubs", subscribersHandler.HandleGetSubscribers)
	private.GET("/database/getsettings", userHandler.GetUserSettingsHandler)
	private.GET("database/download", subscribersHandler.HandleDataDownload)

	// POST
	private.POST("/auth/logout", userHandler.LogoutHandler)
	private.POST("/auth/changepassword", userHandler.ChangePasswordHandler)
	private.POST("/database/addsub", subscribersHandler.HandleNewSubscriberManually)
	private.POST("/database/clearsubs", subscribersHandler.HandleClearSubscribersData)
	private.POST("/database/setsettings", userHandler.SetUserSettingsHandler)
	private.POST("/mail/sendmail", mailHandler.SendEmailHandler)
}
