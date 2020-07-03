package onesignal

import (
	"encoding/json"
	"net/url"
)

// AppsServices handles communication with onesignal api.
type AppsServices struct {
	client *Client
}

// Browse function
//
// Browse notifications, you can see all the notification you have
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/view-apps-apps
func (c *AppsServices) Browse() ([]interface{}, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "apps")
	if err != nil {
		return nil, err
	}

	c.client.UseAuthKey = true
	resp, err := GET(u.String(), c.client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
