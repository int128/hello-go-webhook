package main

// WebhookEvent represents a JIRA event via a webhook.
type WebhookEvent struct {
	ID   int `json:"id"`
	User struct {
		Name string `json:"name"`
	} `json:"user"`
}
