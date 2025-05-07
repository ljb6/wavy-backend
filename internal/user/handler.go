package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterHandler(c *gin.Context) {

	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	user.Plan = "free"

	if err := h.service.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created with success"})
}

func (h *Handler) LoginHandler(c *gin.Context) {

	var req LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)

	c.Writer.Header().Set("Set-Cookie", fmt.Sprintf("token=%s; Path=/; HttpOnly; SameSite=None", token))

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged with success",
		"token":   token,
	})
}

func (h *Handler) LogoutHandler(c *gin.Context) {

	c.SetCookie("token", "", 3600, "/", "", false, true)

	c.Writer.Header().Set("Set-Cookie", fmt.Sprint("token=; Path=/; HttpOnly; SameSite=None", ""))

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out with success",
		"token":   "",
	})
}

func (h *Handler) ChangePasswordHandler(c *gin.Context) {

	var req ChangePasswordRequest

	userID := c.GetString("userID")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
	}

	err = h.service.ChangePassword(userID, req.Password, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "password changed with sucess"})
}

func (h *Handler) SetUserSettingsHandler(c *gin.Context) {

	var req UserSettings

	userID := c.GetString("userID")

	req.User_ID = userID

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	}

	err = h.service.CreateUserSettings(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "settings saved with success"})
}

func (h *Handler) GetUserSettingsHandler(c *gin.Context) {

	userID := c.GetString("userID")

	settings, err := h.service.GetUserSettings(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data is null"})
	}

	c.JSON(http.StatusOK, gin.H{
		"host": "host",
		"port": "port",
		"username": "username",
	})
}