package onesignal

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestCreateNotification(t *testing.T) {
	client := NewClient()
	client.APIKey = os.Getenv("API_KEY")
	client.AppID = os.Getenv("APP_ID")

	opt := &NotificationOpt{
		NotificationContent: NotificationContent{
			Headings: map[string]string{"en": "Judul Percobaan Nih"},
			Contents: map[string]string{"en": "Percobaan Pesan"},
		},
		Appearance: Appearance{
			ChromeWebIcon: "https://kodepanda.com/assets/images/kodepanda-white.png",
		},
		SpecificDevice: SpecificDevice{
			IncludedSegments: []string{"Active Users"},
		},
	}
	_, err := client.Notification.Create(opt)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}
