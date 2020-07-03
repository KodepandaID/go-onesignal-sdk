package onesignal

import (
	"math/rand"
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

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func TestCreateApp(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	opt := &AppsOpt{
		Name:     "new" + randomString(5),
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
		Name:     "update" + randomString(5),
		SiteName: "Kodepanda Kreasi Media",
	}
	_, err := client.Apps.Update(os.Getenv("APP_ID"), opt)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestBrowseDevice(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Apps.BrowseDevice("10", "0")
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestGetDevice(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AuthKey = os.Getenv("AUTH_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Apps.GetDevice(os.Getenv("DEVICE_ID"))
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
