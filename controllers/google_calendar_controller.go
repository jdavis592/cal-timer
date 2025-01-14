package controllers

import (
	"cal-timer/services"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateRandomState creates a secure random state string
func GenerateRandomState() (string, error) {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GetCalendarEvents(c *gin.Context) {
	// Fetch token from storage (mock example here)
	token := c.Query("token")

	events, err := services.FetchCalendarEvents(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}
