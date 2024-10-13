package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// LoginRequest represents the request body for the login API
type LoginRequest struct {
	APIKey string `json:"api_key" example:"my_secure_api_key"` // Example added for Swagger
}

// Login handles the login request and returns a JWT token if the API key is valid
// @Summary Logs in with an API key and returns a JWT token
// @Description Validates the API key and returns a JWT token for authentication
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param   loginRequest body LoginRequest true "API key for login"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Invalid API key"
// @Failure 500 {object} map[string]string "Failed to generate token"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the API key
	if loginRequest.APIKey != os.Getenv("API_KEY") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
