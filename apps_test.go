package onesignal

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestBrowseApps(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Apps.Browse()
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestViewApp(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Apps.Get(os.Getenv("APP_ID"))
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestCreateApp(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	opt := &AppsOpt{
		Name:     "Test library",
		SiteName: "Kodepanda Kreasi Media",
	}
	_, err := client.Apps.Create(opt)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestUpdateApp(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	opt := &AppsOpt{
		Name:     "Test library 1",
		SiteName: "Kodepanda Kreasi Media",
	}
	_, err := client.Apps.Update(os.Getenv("APP_ID"), opt)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
