package user

import (
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

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.h{"message": "Logged with success", "user": user})
}