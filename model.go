package onesignal

// Success model for success responses from OneSignal API
type Success struct {
	ID            string        `json:"id,omitempty"`
	Recipients    int           `json:"recipients,omitempty"`
	ExternalID    string        `json:"external_id,omitempty"`
	Success       bool          `json:"success,omitempty"`
	Errors        interface{}   `json:"errors,omitempty"`
	TotalCount    int           `json:"total_count,omitempty"`   // Show on BrowseNotification
	Offset        int           `json:"offset,omitempty"`        // Show on BrowseNotification
	Limit         int           `json:"limit,omitempty"`         // Show on BrowseNotification
	Notifications []interface{} `json:"notifications,omitempty"` // Show on BrowseNotification
}

// Errors model for all the error request
type Errors struct {
	Errors []string `json:"errors,omitempty"`
}

// DeviceResponse model for response browse devices
type DeviceResponse struct {
	TotalCount int           `json:"total_count,omitempty"`
	Offset     int           `json:"offset,omitempty"`
	Limit      int           `json:"limit,omitempty"`
	Players    []interface{} `json:"players,omitempty"`
}

// SpecificDevice model for request who wants to send to the specific device
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#send-to-specific-devices
type SpecificDevice struct {
	// IncludedSegments, you can filled with ["Active Users", "Inactive Users", "Subscribed Users", "Unsubscribed Users"].
	// Included segment required to send what you want to your target. Users in these segments will receive a notification
	IncludedSegments []string `json:"included_segments,omitempty"`
	// IncludePlayerIDS, you can filled with users device ID
	// With this settings, you can send your notification to specific device
	IncludePlayerIDS       []string `json:"include_player_ids,omitempty"`
	IncludeExternalUserIDS []string `json:"include_external_user_ids,omitempty"`
	IncludeEmailTokens     []string `json:"include_email_tokens,omitempty"`
	IncludeIosTokens       []string `json:"include_ios_tokens,omitempty"`
	IncludeWpWnsUris       []string `json:"include_wp_wns_uris,omitempty"`
	IncludeAmazonRegIDS    []string `json:"include_amazon_reg_ids,omitempty"`
	IncludeChromeRegIDS    []string `json:"include_chrome_reg_ids,omitempty"`
	IncludeAndroidRegIDS   []string `json:"include_android_reg_ids,omitempty"`
}

// NotificationContent model for request who send a push notifications
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#content--language
type NotificationContent struct {
	// Set your notification title with these, then you can set a different title for the specific language
	Headings interface{} `json:"headings,omitempty"`
	// Set your notification subtitle with these, then you can set a different subtitle for the specific language
	Subtitle interface{} `json:"subtitle,omitempty"`
	// If you already set up a template with your account, and want to use that template, filled with your template ID
	TemplateID string `json:"template_id,omitempty"`
	// These configs just for iOS to displaying a notification content
	ContentAvailable bool `json:"content_available,omitempty"`
	// Set your notification content with these, then you can set a different content for the specific language
	Contents interface{} `json:"contents,omitempty"`
}

// EmailContent model for request who send an email
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#email-content
type EmailContent struct {
	// These config required for the subject of the email
	EmailSubject string `json:"email_subject,omitempty"`
	// These config required for the content of the email
	EmailBody string `json:"email_body,omitempty"`
	// The name the email is from. If this is not specified, this will use your default from name set in email set up
	EmailFromName string `json:"email_from_name,omitempty"`
	// The email address the email is from. If this is not specified, this will use your default from email in email set up
	EmailFromAddress string `json:"email_fromaddress,omitempty"`
}

// Attachments model for additional content attached to push notifications, primarily images
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#attachments
type Attachments struct {
	// A custom map of data that is passed back to your app.
	Data interface{} `json:"data,omitempty"`
	// The URL to open in the browser when a user clicks on the notification.
	URL string `json:"url,omitempty"`
	// Same as URL but only sent to web push platforms.
	// Including Chrome, Firefox, Safari, Opera, etc.
	WebURL string `json:"web_url,omitempty"`
	// Same as URL but only sent to app platforms.
	// Including iOS, Android, macOS, Windows, ChromeApps, etc.
	AppURL string `json:"app_url,omitempty"`
	// Adds media attachments to notifications. Set as JSON object, key as a media id of your choice and the value as a valid local filename or URL.
	// User must press and hold on the notification to view.
	IosAttachments interface{} `json:"ios_attachments,omitempty"`
	// Picture to display in the expanded view. Can be a drawable resource name or a URL.
	BigPicture string `json:"big_picture,omitempty"`
	// Sets the web push notification's large image to be shown below the notification's title and text. Please see Web Push Notification Icons.
	ChromeWebImage string `json:"chrome_web_image,omitempty"`
	// Picture to display in the expanded view. Can be a drawable resource name or a URL.
	// Only used for Amazon platform
	AdmBigPicture string `json:"adm_big_picture,omitempty"`
	// Picture to display in the expanded view. Can be a drawable resource name or a URL.
	// Only used for ChromeApp platform
	ChromeBigPicture string `json:"chrome_big_picture,omitempty"`
}

// ActionButtons model for added buttons to push notifications
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#action-buttons
type ActionButtons struct {
	// Buttons to add to the notification. Icon only works for Android.
	Buttons []interface{} `json:"buttons,omitempty"`
	// Add action buttons to the notification. The id field is required.
	WebButtons []interface{} `json:"web_buttons,omitempty"`
	// Category APS payload, use with registerUserNotificationSettings:categories in your Objective-C / Swift code.
	IosCategory string `json:"ios_category,omitempty"`
}

// Appearance model for adjust icons, badges, and other appearance changes to your push notifications
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#appearance
type Appearance struct {
	AndroidChannelID         string      `json:"android_channel_id,omitempty"`
	HuaweiChannelID          string      `json:"huawei_channel_id,omitempty"`
	ExistingAndroidChannelID string      `json:"existing_android_channel_id,omitempty"`
	HuaweiExistingChannelID  string      `json:"huawei_existing_channel_id,omitempty"`
	AndroidBackgroundLayout  interface{} `json:"android_background_layout,omitempty"`
	SmallIcon                string      `json:"small_icon,omitempty"`
	HuaweiSmallIcon          string      `json:"huawei_small_icon,omitempty"`
	LargeIcon                string      `json:"large_icon,omitempty"`
	HuaweiLargeIcon          string      `json:"huawei_large_icon,omitempty"`
	AdmSmallIcon             string      `json:"adm_small_icon,omitempty"`
	AdmLargeIcon             string      `json:"adm_large_icon,omitempty"`
	ChromeWebIcon            string      `json:"chrome_web_icon,omitempty"`
	ChromeWebBadge           string      `json:"chrome_web_badge,omitempty"`
	FirefoxIcon              string      `json:"firefox_icon,omitempty"`
	ChromeIcon               string      `json:"chrome_icon,omitempty"`
	IosSound                 string      `json:"ios_sound,omitempty"`
	AndroidSound             string      `json:"android_sound,omitempty"`
	HuaweiSound              string      `json:"huawei_sound,omitempty"`
	AdmSound                 string      `json:"adm_sound,omitempty"`
	WpWnsSound               string      `json:"WpWnsSound,omitempty"`
	AndroidLedColor          string      `json:"android_led_color,omitempty"`
	HuaweiLedColor           string      `json:"huawei_led_color,omitempty"`
	AndroidAccentColor       string      `json:"android_accent_color,omitempty"`
	HuaweiAccentColor        string      `json:"huawei_accent_color,omitempty"`
	AndroidVisibility        int         `json:"android_visibility,omitempty"`
	HuaweiVisibility         int         `json:"huawei_visibility,omitempty"`
	IosBadgeType             string      `json:"ios_badgeType,omitempty"`
	IosBadgeCount            int         `json:"ios_badgeCount,omitempty"`
	CollapseID               string      `json:"collapse_id,omitempty"`
	ApnsAlert                interface{} `json:"apns_alert,omitempty"`
}

// Delivery model for set schedule notification for future delivery.
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#delivery
type Delivery struct {
	SendAfter            string `json:"send_after,omitempty"`
	DelayedOption        string `json:"delayed_option,omitempty"`
	DeliveryTimeOfDay    string `json:"delivery_time_of_day,omitempty"`
	TTL                  int    `json:"ttl,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	ApnsPushTypeOverride string `json:"apns_push_type_override,omitempty"`
}

// Grouping model for combine multiple notifications into a single notification to improve the user experience.
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#grouping--collapsing
type Grouping struct {
	AndroidGroup        string      `json:"android_group,omitempty"`
	AndroidGroupMessage interface{} `json:"android_group_message,omitempty"`
	AdmGroup            string      `json:"adm_group,omitempty"`
	AdmGroupMessage     interface{} `json:"adm_group_message,omitempty"`
	ThreadID            string      `json:"thread_id,omitempty"`
	SummaryArg          string      `json:"summary_arg,omitempty"`
	SummaryArgCunt      int         `json:"summaryArgCount,omitempty"`
}

// Platform model for only send to specific platforms.
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#platform-to-deliver-to
type Platform struct {
	IsIos                     bool   `json:"isIos,omitempty"`
	IsAndroid                 bool   `json:"isAndroid,omitempty"`
	IsHuawei                  bool   `json:"isHuawei,omitempty"`
	IsAnyWeb                  bool   `json:"isAnyWeb,omitempty"`
	IsEmail                   bool   `json:"isEmail,omitempty"`
	IsChromeWeb               bool   `json:"isChromeWeb,omitempty"`
	IsFirefox                 bool   `json:"isFirefox,omitempty"`
	IsSafari                  bool   `json:"isSafari,omitempty"`
	IsWpWns                   bool   `json:"isWP_WNS,omitempty"`
	IsAdm                     bool   `json:"isAdm,omitempty"`
	IsChrome                  bool   `json:"isChrome,omitempty"`
	ChannelForExternalUserIds string `json:"channel_for_external_user_ids,omitempty"`
}

// NotificationOpt model for request create a notification.
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/create-notification
type NotificationOpt struct {
	AppID   string        `json:"app_id,omitempty"`
	Filters []interface{} `json:"filters,omitempty"`
	SpecificDevice
	NotificationContent
	EmailContent
	Attachments
	ActionButtons
	Appearance
	Delivery
	Grouping
	Platform
}

// AppsOpt model for request create an OneSignal application.
//
// For your reference, you can read this API reference on https://documentation.onesignal.com/reference/create-an-app
type AppsOpt struct {
	Name                             string `json:"name,omitempty"`
	ApnsEnv                          string `json:"apns_env,omitempty"`
	ApnsP12                          string `json:"apns_p12,omitempty"`
	ApnsP12Password                  string `json:"apns_p12_password,omitempty"`
	GcmKey                           string `json:"gcm_key,omitempty"`
	AndroidGcmSenderID               string `json:"android_gcm_sender_id,omitempty"`
	ChromeWebOrigin                  string `json:"chrome_web_origin,omitempty"`
	ChromeWebDefaultNotificationIcon string `json:"chrome_web_notification_icon,omitempty"`
	ChromeWebSubDomain               string `json:"chrome_web_sub_domain,omitempty"`
	SafariApnsP12                    string `json:"safari_apns_p12,omitempty"`
	SafariApnsP12Password            string `json:"safari_apns_p12_password,omitempty"`
	SiteName                         string `json:"site_name,omitempty"`
	SafariSiteOrigin                 string `json:"safari_site_origin,omitempty"`
	SafariIcon256                    string `json:"safari_icon_256_256,omitempty"`
	ChromeKey                        string `json:"chrome_key,omitempty"`
	AdditionalDataIsRootPayload      bool   `json:"additional_data_is_root_payload,omitempty"`
	OrganizationID                   string `json:"organization_id,omitempty"`
}
