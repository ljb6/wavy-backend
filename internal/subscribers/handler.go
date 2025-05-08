package subscribers

import (
	"encoding/csv"
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
	
	fmt.Println(userID)

	jsonSubs, err := h.service.GetSubscribers(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "impossible to get subscribers"})
	}

	c.Data(http.StatusOK, "application/json", jsonSubs)
}

func (h *Handler) HandleClearSubscribersData(c *gin.Context) {

	userID := c.GetString("userID")

	err := h.service.ClearSubscribers(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to clear subscribers"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscribers cleared with success"})
}

func (h *Handler) HandleDataDownload(c *gin.Context) {

	userID := c.GetString("userID")

    c.Header("Content-Disposition", "attachment; filename=data.csv")
    c.Header("Content-Type", "text/csv")
    c.Header("Cache-Control", "no-cache")

    writer := csv.NewWriter(c.Writer)
    defer writer.Flush()

	data, err := h.service.DownloadData(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while downloading data"})
	}

    for _, row := range data {
        if err := writer.Write(row); err != nil {
            c.JSON(http.StatusOK, gin.H{"message": "error while generating csv"})
        }
    }
}