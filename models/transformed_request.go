package models

type TransformedRequest struct {
	Event      string               `json:"event"`
	EventType  string               `json:"event_type"`
	AppID      string               `json:"app_id"`
	UserID     string               `json:"user_id"`
	MessageID  string               `json:"message_id"`
	PageTitle  string               `json:"page_title"`
	PageURL    string               `json:"page_url"`
	Language   string               `json:"browser_language"`
	Screen     string               `json:"screen_size"`
	Attributes map[string]Attribute `json:"attributes"`
	Traits     map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
