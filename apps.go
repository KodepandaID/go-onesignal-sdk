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
// Browse apps, you can see all the apps you have
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

// Get function
//
// View the details of a single OneSignal app. This function returns a JSON object not a struct like others
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/view-an-app
func (c *AppsServices) Get(id string) (interface{}, error) {
	u, err := url.Parse(c.client.BaseURL.String() + "apps/" + id)
	if err != nil {
		return Success{}, err
	}

	c.client.UseAuthKey = true
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
