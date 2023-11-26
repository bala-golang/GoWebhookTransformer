package models

type OriginalRequest struct {
	Event           string            `json:"ev"`
	EventType       string            `json:"et"`
	ID              string            `json:"id"`
	UID             string            `json:"uid"`
	MessageID       string            `json:"mid"`
	PageTitle       string            `json:"t"`
	PageURL         string            `json:"p"`
	BrowserLanguage string            `json:"l"`
	ScreenSize      string            `json:"sc"`
	ATR             map[string]string `json:"attributes"`
	UATR            map[string]string `json:"user_traits"`
}
