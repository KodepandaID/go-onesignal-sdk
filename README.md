# Go SDK client for the OneSignal API

This package provides an SDK OneSignal v7.0. This is not an official package from OneSignal for Go.

If you found an issue, you can report to Github Issue or send a direct message to my [twitter](https://twitter.com/lordaur).

## Installation
```bash
go get github.com/KodepandaID/go-onesignal-sdk
```

## Usage

This SDK needs an API key, Auth Key, and App ID optionally.

```go
import "github.com/KodepandaID/go-onesignal-sdk"

func main() {
    client := onesignal.NewClient()
    client.APIKey = "YOUR_API_KEY"
    client.AuthKey = "YOUR_AUTH_KEY"
    client.AppID = "YOUR_APP_ID"
}
```