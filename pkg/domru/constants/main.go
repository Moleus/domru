package constants

import "fmt"

const (
	BaseUrl            = "https://myhome.novotelecom.ru"
	USERAGENT_TEMPLATE = "Google sdkgphone64x8664 | Android 14 | erth | 8.9.2 (8090200)"

	API_HA_NETWORK = "http://supervisor/network/info"

	API_AUTH = "https://api-auth.domru.ru/v1/person/auth"

	API_AUTH_LOGIN            = "%s/auth/v2/login/%s"
	API_AUTH_CONFIRMATION     = "%s/auth/v2/confirmation/%s"
	API_AUTH_CONFIRMATION_SMS = "%s/auth/v2/auth/%s/confirmation"

	API_CAMERAS           = "%s/rest/v1/forpost/cameras"
	API_OPEN_DOOR         = "%s/rest/v1/places/%d/accesscontrols/%d/actions"
	API_FINANCES          = "%s/rest/v1/subscribers/profiles/finances"
	API_SUBSCRIBER_PLACES = "%s/rest/v1/subscriberplaces"
	API_VIDEO_SNAPSHOT    = "%s/rest/v1/places/%d/accesscontrols/%d/videosnapshots"
	API_CAMERA_GET_STREAM = "%s/rest/v1/forpost/cameras/%d/video"
	API_REFRESH_SESSION   = "%s/auth/v2/session/refresh"
	API_EVENTS            = "%s/rest/v1/places/%s/events?allowExtentedActions=true"
	API_OPERATORS         = "%s/public/v1/operators"

	CUSTOM_STREAM_URL = "%s/stream/%d"
)

func GetUserAgent(login string) string {
	return USERAGENT_TEMPLATE
	//return fmt.Sprintf(USERAGENT_TEMPLATE, login, "f3fbe9cc-e969-4282-9835-a9cb5c94eb79")
}

func GetSnapshotUrl(baseUrl string, placeId, accessControlId int) string {
	return fmt.Sprintf(API_VIDEO_SNAPSHOT, baseUrl, placeId, accessControlId)
}

func GetOpenDoorUrl(baseUrl string, placeId, accessControlId int) string {
	return fmt.Sprintf(API_OPEN_DOOR, baseUrl, placeId, accessControlId)
}

func GetCameraStreamUrl(baseUrl string, cameraId int) string {
	return fmt.Sprintf(CUSTOM_STREAM_URL, baseUrl, cameraId)
}

func GetEventsUrl(baseUrl, placeId string) string {
	return fmt.Sprintf(API_EVENTS, baseUrl, placeId)
}

func GetAuthConfirmationUrl(baseUrl, phone string) string {
	return fmt.Sprintf(API_AUTH_CONFIRMATION, baseUrl, phone)
}

func GetAuthConfirmationSmsUrl(baseUrl, phone string) string {
	return fmt.Sprintf(API_AUTH_CONFIRMATION_SMS, baseUrl, phone)
}
