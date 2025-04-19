package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/database"
	"github.com/ljb6/wavy-backend/models"
)

func RegisterUser(c *gin.Context, db *sql.DB) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	user.Plan = "free"

	userID, err := database.CreateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user in database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created with success", "id": userID})
}
