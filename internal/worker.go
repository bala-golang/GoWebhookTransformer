package internal

import (
	"github.com/bala-golang/GoWebhookTransformer/models"
	"github.com/bala-golang/GoWebhookTransformer/util"
)

// Worker processes incoming OriginalRequests from a channel.
func Worker(channel <-chan models.OriginalRequest) {
	for {
		select {
		case req := <-channel:
			// Transform the OriginalRequest into a TransformedRequest
			transformedReq := util.TransformRequest(req)

			// Send the TransformedRequest to the webhook
			sendToWebhook(transformedReq)
		}
	}
}
