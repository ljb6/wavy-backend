package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	emailService *EmailService
}

func NewEmailHandler(service *EmailService) *EmailHandler {
	return &EmailHandler{emailService: service}
}

func (e *EmailHandler) SendEmailHandler(c *gin.Context) {
	
	userID := c.GetString("userID")

	var req EmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
	}

	err = e.emailService.SendEmail(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "email sent"})
}