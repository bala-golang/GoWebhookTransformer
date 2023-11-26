package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bala-golang/GoWebhookTransformer/config"
	"github.com/bala-golang/GoWebhookTransformer/models"

	"github.com/gin-gonic/gin"
)

var (
	// RequestChannel is a channel to pass the original request for further processing.
	RequestChannel = make(chan models.OriginalRequest)

	invalidRequestError    = gin.H{"error": "Invalid JSON format"}
	successResponseMessage = gin.H{"message": "Request received and processing..."}
)

// ProcessEventHandler handles incoming webhook events.
func ProcessEventHandler(c *gin.Context) {
	e := map[string]interface{}{}
	var event models.OriginalRequest

	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, invalidRequestError)
		return
	}

	data, err := json.Marshal(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, invalidRequestError)
		return
	}
	if err := json.Unmarshal(data, &event); err != nil {
		c.JSON(http.StatusBadRequest, invalidRequestError)
		return
	}

	event.ATR = map[string]string{}
	event.UATR = map[string]string{}

	var ok bool
	for key, val := range e {
		// Check if the key has a prefix matching uatr.
		if strings.HasPrefix(key, config.UserTraits) {
			event.UATR[key], ok = val.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, invalidRequestError)
				return
			}
		}

		// Check if the key has a prefix matching atr.
		if strings.HasPrefix(key, config.Attribute) {
			event.ATR[key], ok = val.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, invalidRequestError)
				return
			}
		}
	}

	go func(req models.OriginalRequest) {
		RequestChannel <- req
	}(event)

	c.JSON(http.StatusOK, successResponseMessage)
}
