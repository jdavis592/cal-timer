package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

func getLocalCredentials() ([]byte, error) {
	// Load your credentials.json file
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}
	return b, nil
}

func setOAuth2Config(b []byte) (*oauth2.Config, error) {
	// Set up OAuth 2.0 config
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}
	return config, nil
}

func createAuthUrl(config *oauth2.Config) {
	// Generate auth URL and print it
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the following URL to authorize:\n%v\n", authURL)
}

func codeForTokenExchange(config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	// Exchange the code for a token
	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
		return nil, err
	}

	return token, nil
}

func RunAuth() error {
	var err error
	credentials, err := getLocalCredentials()

	config, err := setOAuth2Config(credentials)

	createAuthUrl(config)

	// Start a local HTTP server to capture the callback
	codeChan := make(chan string)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No code in request", http.StatusBadRequest)
			err = fmt.Errorf("No code in request")
			return
		}
		fmt.Fprintf(w, "Authorization successful. You can close this window.")
		codeChan <- code
	})
	go func() {
		log.Fatal(http.ListenAndServe("localhost:3000", nil))
	}()

	// Wait for the authorization code
	authCode := <-codeChan

	token, err := codeForTokenExchange(config, authCode)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	fmt.Printf("Token: %+v\n", token)
	return nil
}
