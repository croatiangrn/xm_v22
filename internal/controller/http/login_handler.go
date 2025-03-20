package http

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// This is only for demonstration purposes, an actual login implementation is required
type LoginHandler struct {
}

// NewLoginHandler creates a new LoginHandler instance.
func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

// Login handles user login and JWT token generation.
func (h *LoginHandler) Login(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req LoginRequest

		// Bind the request body to LoginRequest struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// In a real-world application, we'd validate the username and password with a database and return a user record if successful.

		// For demonstration, assume successful login if credentials are correct
		if req.Username == "admin" && req.Password == "password" {
			jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":       123,
				"email":    "admin@example.com",
				"username": req.Username,
			}).SignedString([]byte(jwtSecret))

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating JWT"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": jwtToken})
			return
		}
	
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
