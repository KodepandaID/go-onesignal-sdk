package onesignal

import (
	"log"
	"net/url"
)

const (
	defaultBaseURL = "https://onesignal.com/api/v1/"
)

// A Client manages communication with the OneSignal API.
type Client struct {
	BaseURL *url.URL
	APIKey  string
	AppID   string

	Notification *NotificationServices
}

// NewClient returns a new OneSignal API client
func NewClient() *Client {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		BaseURL: baseURL,
	}
	c.Notification = &NotificationServices{client: c}

	return c
}
