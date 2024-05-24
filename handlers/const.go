package handlers

import "github.com/ad/domru/pkg/auth"

const (
	API_USER_AGENT   = "myHomeErth/8 CFNetwork/1209 Darwin/20.2.0"
	CLIENT_USERAGENT = "iPhone11,6 | iOS 14.3 | erth | 6.4.6 (build 3) | %s | 2 | %s"

	API_HA_NETWORK = "http://supervisor/network/info"

	API_AUTH_LOGIN            = "https://myhome.novotelecom.ru/auth/v2/login/%s"
	API_AUTH_CONFIRMATION     = "https://myhome.novotelecom.ru/auth/v2/confirmation/%s"
	API_AUTH_CONFIRMATION_SMS = "https://myhome.novotelecom.ru/auth/v2/auth/%s/confirmation"

	API_AUTH = "https://api-auth.domru.ru/v1/person/auth"

	API_CAMERAS           = "https://myhome.novotelecom.ru/rest/v1/forpost/cameras"
	API_OPEN_DOOR         = "https://myhome.novotelecom.ru/rest/v1/places/%s/accesscontrols/%s/actions"
	API_FINANCES          = "https://myhome.novotelecom.ru/rest/v1/subscribers/profiles/finances"
	API_SUBSCRIBER_PLACES = "https://myhome.novotelecom.ru/rest/v1/subscriberplaces"
	API_VIDEO_SNAPSHOT    = "https://myhome.novotelecom.ru/rest/v1/places/%s/accesscontrols/%s/videosnapshots"
	API_CAMERA_GET_STREAM = "https://myhome.novotelecom.ru/rest/v1/forpost/cameras/%s/video"
	API_REFRESH_SESSION   = "https://myhome.novotelecom.ru/auth/v2/session/refresh"
	API_EVENTS            = "https://myhome.novotelecom.ru/rest/v1/places/%s/events?allowExtentedActions=true"
	API_OPERATORS         = "https://myhome.novotelecom.ru/public/v1/operators"
)

type AccountsPageData struct {
	Accounts      []auth.Account
	Phone         string
	HassioIngress string
	LoginError    string
}

type LoginPageData struct {
	LoginError    string
	Phone         string
	HassioIngress string
}

type SMSPageData struct {
	Phone         string
	Index         string
	HassioIngress string
	LoginError    string
}

type Places struct {
	Data []struct {
		ID    int `json:"id"`
		Place struct {
			ID      int `json:"id"`
			Address struct {
				VisibleAddress string `json:"visibleAddress"`
			} `json:"address"`
			AccessControls []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"accessControls"`
		} `json:"place"`
		Subscriber struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			AccountID string `json:"accountId"`
		} `json:"subscriber"`
		Blocked bool `json:"blocked"`
	} `json:"data"`
}

type Cameras struct {
	Data []struct {
		ID       int    `json:"ID"`
		Name     string `json:"Name"`
		IsActive int    `json:"IsActive"`
	} `json:"data"`
}

type HomePageData struct {
	HassioIngress string
	HostIP        string
	Port          string
	Host          string
	Scheme        string
	LoginError    string
	Phone         string
	Token         string
	RefreshToken  string
	Cameras       Cameras
	Places        Places
	Finances      Finances
}

type HAConfig struct {
	Result string `json:"result"`
	Data   struct {
		Interfaces []struct {
			Ipv4 struct {
				Address []string `json:"address"`
			} `json:"ipv4"`
		} `json:"interfaces"`
	} `json:"data"`
}
