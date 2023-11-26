package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bala-golang/GoWebhookTransformer/config"
	"github.com/bala-golang/GoWebhookTransformer/models"
)

// sendToWebhook sends a TransformedRequest to the configured webhook URL.
func sendToWebhook(req models.TransformedRequest) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error while marshaling request:", err)
		return
	}

	resp, err := http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("Error sending to webhook:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response from webhook:", resp.Status)
}
