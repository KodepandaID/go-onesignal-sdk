package onesignal

import (
	"net/url"
)

// NotificationServices handles communication with onesignal api.
type NotificationServices struct {
	client *Client
}

// Create function
//
// Create a notification
//
// Sends notifications to your users. This function relates to https://onesignal.com/api/v1/notifications
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/create-notification
func (c *NotificationServices) Create(opt *NotificationOpt) (Success, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "notifications")
	if err != nil {
		return Success{}, err
	}

	opt.AppID = c.client.AppID
	b, err := EncodeBody(opt)
	if err != nil {
		return Success{}, err
	}

	resp, err := POST(u.String(), b, c.client)
	if err != nil {
		return Success{}, err
	}

	return resp, nil
}
