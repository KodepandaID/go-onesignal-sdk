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

// SpecificDevice model for request who wants to send to the specific device
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#send-to-specific-devices
type SpecificDevice struct {
	IncludedSegments       []string `json:"included_segments,omitempty"`
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
	Headings         interface{} `json:"headings,omitempty"`
	Subtitle         interface{} `json:"subtitle,omitempty"`
	TemplateID       string      `json:"template_id,omitempty"`
	ContentAvailable bool        `json:"content_available,omitempty"`
	Contents         interface{} `json:"contents,omitempty"`
}

// EmailContent model for request who send an email
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#email-content
type EmailContent struct {
	EmailSubject     string `json:"email_subject,omitempty"`
	EmailBody        string `json:"email_body,omitempty"`
	EmailFromName    string `json:"email_from_name,omitempty"`
	EmailFromAddress string `json:"email_fromaddress,omitempty"`
}

// Attachments model for additional content attached to push notifications, primarily images
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#attachments
type Attachments struct {
	Data             interface{} `json:"data,omitempty"`
	URL              string      `json:"url,omitempty"`
	WebURL           string      `json:"web_url,omitempty"`
	AppURL           string      `json:"app_url,omitempty"`
	IosAttachments   interface{} `json:"ios_attachments,omitempty"`
	BigPicture       string      `json:"big_picture,omitempty"`
	ChromeWebImage   string      `json:"chrome_web_image,omitempty"`
	AdmBigPicture    string      `json:"adm_big_picture,omitempty"`
	ChromeBigPicture string      `json:"chrome_big_picture,omitempty"`
}

// ActionButtons model for added buttons to push notifications
//
// For more information, you can read this parameter reference from
// OneSignal API Reference https://documentation.onesignal.com/reference/create-notification#action-buttons
type ActionButtons struct {
	Buttons     []interface{} `json:"buttons,omitempty"`
	WebButtons  []interface{} `json:"web_buttons,omitempty"`
	IosCategory string        `json:"ios_category,omitempty"`
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
	AppID   string      `json:"app_id,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
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
