package services

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func CreateEvent(token *oauth2.Token) (*calendar.Event, error) {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	service, err := calendar.New(client)
	if err != nil {
		return nil, err
	}

	event := &calendar.Event{
		Summary:     "Sample Event",
		Location:    "Virtual",
		Description: "A test event",
		Start: &calendar.EventDateTime{
			DateTime: "2025-01-15T10:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2025-01-15T11:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
	}

	createdEvent, err := service.Events.Insert("primary", event).Do()
	if err != nil {
		return nil, err
	}
	return createdEvent, nil
}

func FetchCalendarEvents(accessToken string) ([]*calendar.Event, error) {
	token := &oauth2.Token{AccessToken: accessToken}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	service, err := calendar.New(client)
	if err != nil {
		return nil, err
	}

	events, err := service.Events.List("primary").Do()
	if err != nil {
		return nil, err
	}

	return events.Items, nil
}
