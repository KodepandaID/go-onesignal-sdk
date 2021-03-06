package onesignal

import (
	"encoding/json"
	"net/url"
)

// NotificationServices handles communication with onesignal api.
type NotificationServices struct {
	client *Client
}

// Browse function
//
// Browse notifications, you can see all the notification you have
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/view-notifications
func (c *NotificationServices) Browse(limit, offset string) (Success, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "notifications?app_id=" + c.client.AppID + "&limit=" + limit + "&offset=" + offset)
	if err != nil {
		return Success{}, err
	}

	resp, err := GET(u.String(), c.client)
	if err != nil {
		return Success{}, err
	}
	defer resp.Body.Close()

	var data Success
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Success{}, err
	}

	return data, nil
}

// Get function
//
// View the details of a single notification. This function returns a JSON object not a struct like others
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/view-notification
func (c *NotificationServices) Get(id string) (interface{}, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "notifications/" + id + "?app_id=" + c.client.AppID)
	if err != nil {
		return Success{}, err
	}

	resp, err := GET(u.String(), c.client)
	if err != nil {
		return Success{}, err
	}
	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Success{}, err
	}

	return data, nil
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

// Cancel function
//
// Cancel a notification
//
// Stop a scheduled or currently outgoing notification
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/cancel-notification
func (c *NotificationServices) Cancel(id string) (Success, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "notifications/" + id + "?app_id=" + c.client.AppID)
	if err != nil {
		return Success{}, err
	}

	resp, err := DELETE(u.String(), c.client)
	if err != nil {
		return Success{}, err
	}

	return resp, nil
}
