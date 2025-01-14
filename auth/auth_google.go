package auth

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConfig *oauth2.Config

func InitOAuthConfig() {
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
}

func GetAuthURL(state string) string {
	return oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func ExchangeToken(code string) (*oauth2.Token, error) {
	return oauthConfig.Exchange(context.Background(), code)
}

func GetClient(token *oauth2.Token) *http.Client {
	return oauthConfig.Client(context.Background(), token)
}
