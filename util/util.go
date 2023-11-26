package util

import (
	"fmt"
	"strings"

	"github.com/bala-golang/GoWebhookTransformer/config"
	"github.com/bala-golang/GoWebhookTransformer/models"
)

// TransformRequest transforms the OriginalRequest into a TransformedRequest.
func TransformRequest(originalReq models.OriginalRequest) models.TransformedRequest {
	transformedReq := models.TransformedRequest{
		Event:      originalReq.Event,
		EventType:  originalReq.EventType,
		AppID:      originalReq.ID,
		UserID:     originalReq.UID,
		MessageID:  originalReq.MessageID,
		PageTitle:  originalReq.PageTitle,
		PageURL:    originalReq.PageURL,
		Language:   originalReq.BrowserLanguage,
		Screen:     originalReq.ScreenSize,
		Attributes: make(map[string]models.Attribute),
		Traits:     make(map[string]models.Trait),
	}

	for key, attr := range originalReq.ATR {
		if strings.HasPrefix(key, config.AttributeKey) {
			attrKeyNo := strings.Split(key, config.AttributeKey)
			if len(attrKeyNo) >= 2 {
				value, valueOk := originalReq.ATR[fmt.Sprint(config.AttributeValue, attrKeyNo[1])]
				valueType, valueTypeOk := originalReq.ATR[fmt.Sprint(config.AttributeValueType, attrKeyNo[1])]
				if valueOk && valueTypeOk {
					transformedReq.Attributes[attr] = models.Attribute{
						Value: value,
						Type:  valueType,
					}
				}
			}
		}
	}

	for key, attr := range originalReq.UATR {
		if strings.HasPrefix(key, config.UserTraitsAttributeKey) {
			attrKeyNo := strings.Split(key, config.UserTraitsAttributeKey)
			if len(attrKeyNo) >= 2 {
				value, valueOk := originalReq.UATR[fmt.Sprint(config.UserTraitsAttributeValue, attrKeyNo[1])]
				valueType, valueTypeOk := originalReq.UATR[fmt.Sprint(config.UserTraitsAttributeValueType, attrKeyNo[1])]
				if valueOk && valueTypeOk {
					transformedReq.Traits[attr] = models.Trait{
						Value: value,
						Type:  valueType,
					}
				}
			}
		}
	}

	return transformedReq
}
