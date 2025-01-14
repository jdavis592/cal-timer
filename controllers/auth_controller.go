package controllers

import (
	"cal-timer/auth"
	"cal-timer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGoogle(c *gin.Context) {
	state := utils.GenerateRandomState()
	c.SetCookie("oauth_state", state, 300, "/", "", false, true)
	authURL := auth.GetAuthURL(state)
	c.Redirect(http.StatusFound, authURL)
}

func AuthCallback(c *gin.Context) {
	// Retrieve the state value from the query parameter
	queryState := c.Query("state")

	// Retrieve the state value from the cookie
	cookieState, err := c.Cookie("oauth_state")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State cookie not found"})
		return
	}

	// Verify that the state matches
	if queryState != cookieState {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid state value"})
		return
	}

	// Proceed with exchanging the authorization code for a token
	code := c.Query("code")
	token, err := auth.ExchangeToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	// Optionally store the token securely
	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken})
}
