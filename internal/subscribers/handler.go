package subscribers

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

func (h *Handler) HandleNewSubscriberManually(c *gin.Context) {

	userID := c.GetString("userID")

	var req SubRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscriber data"})
	}

	err = h.service.AddSubscriber(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in adding subscriber"})
	}
}

func (h *Handler) HandleGetSubscribers(c *gin.Context) {

	userID := c.GetString("userID")

	jsonSubs, err := h.service.GetSubscribers(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get subscribers"})
	}

	c.JSON(http.StatusOK, jsonSubs)
}