package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/wavy-backend/database"
	"github.com/ljb6/wavy-backend/models"
	"github.com/ljb6/wavy-backend/security"
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

func LoginUser(c *gin.Context, db *sql.DB) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in login"})
		return
	}

	var hashedPassword string
	var userID int

	query := "SELECT id, password FROM users WHERE email = $1"
	err = db.QueryRow(query, user.Email).Scan(&userID, &hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found in database"})
		return
	}

	if security.CheckPassword(hashedPassword, user.Password) {

		token, err := security.GenerateJWT(strconv.Itoa(userID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error in generating JWT token"})
			return
		}

		c.SetCookie("token", token, 3600, "/", "", false, true)

		c.Writer.Header().Set("Set-Cookie", fmt.Sprintf("token=%s; Path=/; HttpOnly; SameSite=None", token))

		c.JSON(http.StatusOK, gin.H{"message": "User logged with success", "token": token})
		return
	}
}