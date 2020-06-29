package onesignal

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

var notificationID string

func TestBrowseNotification(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Notification.Browse("50", "0")
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestCreateNotification(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AppID = os.Getenv("APP_ID")

	opt := &NotificationOpt{
		NotificationContent: NotificationContent{
			Headings: map[string]string{"en": "Judul Percobaan Nih Ahay"},
			Contents: map[string]string{"en": "Percobaan Pesan Nih"},
		},
		Appearance: Appearance{
			ChromeWebIcon: "https://kodepanda.com/assets/images/kodepanda-white.png",
		},
		SpecificDevice: SpecificDevice{
			IncludedSegments: []string{"Active Users"},
		},
	}
	resp, err := client.Notification.Create(opt)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	} else {
		notificationID = resp.ID
	}
}

func TestGetNotification(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Notification.Get(notificationID)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func TestCancelNotification(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AppID = os.Getenv("APP_ID")

	_, err := client.Notification.Cancel(notificationID)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
