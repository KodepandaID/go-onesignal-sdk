package onesignal

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestBrowseApps(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	data, err := client.Apps.Browse()
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Print(data)
}
