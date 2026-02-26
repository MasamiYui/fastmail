package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a middleware that checks for a valid Bearer token
func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Check if the header starts with "Bearer "
		const prefix = "Bearer "
		if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		// Extract the token
		providedToken := authHeader[len(prefix):]

		if providedToken != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
