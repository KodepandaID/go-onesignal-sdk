package onesignal

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// POST function
// to handle http POST method
func POST(url string, payload io.Reader, c *Client) (Success, error) {
	var client = &http.Client{}

	req, err := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+c.APIKey)
	if err != nil {
		return Success{}, err
	}

	resp, err := client.Do(req)
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

// EncodeBody function
//
// Converting a struct to io.ReadWriter for http.request
func EncodeBody(body interface{}) (io.ReadWriter, error) {
	var buf io.ReadWriter
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return nil, err
	}
	buf = b

	return buf, nil
}
