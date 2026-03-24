package alanube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WebhookEvent wraps a parsed webhook payload with its detected type.
type WebhookEvent struct {
	// RawBody contains the original request body bytes.
	RawBody []byte

	// GovernmentStatusChanged is set when the event is a government status change.
	GovernmentStatusChanged *GovernmentStatusChangedEvent

	// EmissionFinished is set when the event is a document emission finished.
	EmissionFinished *EmissionFinishedEvent

	// DocumentReception is set when the event is a document reception.
	DocumentReception *DocumentReceptionEvent
}

// ParseWebhook reads and parses a webhook request body.
// It attempts to detect the event type from the JSON structure.
func ParseWebhook(r *http.Request) (*WebhookEvent, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("alanube: read webhook body: %w", err)
	}
	defer r.Body.Close()

	event := &WebhookEvent{RawBody: body}

	// Try to detect the event type by unmarshaling into a generic map first.
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("alanube: parse webhook: %w", err)
	}

	// If the payload has a "company" key, it's a government status changed event.
	if _, ok := raw["company"]; ok {
		var ev GovernmentStatusChangedEvent
		if err := json.Unmarshal(body, &ev); err != nil {
			return nil, fmt.Errorf("alanube: parse government status event: %w", err)
		}
		event.GovernmentStatusChanged = &ev
		return event, nil
	}

	// If the payload has a "type" key, check its value.
	if typeRaw, ok := raw["type"]; ok {
		var typeStr string
		if err := json.Unmarshal(typeRaw, &typeStr); err == nil {
			switch typeStr {
			case "receivedDocument", "receivedCommercialApproval", "sentCommercialApproval":
				var ev DocumentReceptionEvent
				if err := json.Unmarshal(body, &ev); err != nil {
					return nil, fmt.Errorf("alanube: parse reception event: %w", err)
				}
				event.DocumentReception = &ev
				return event, nil
			default:
				// Other document types (e.g., "fiscalInvoice") are emission finished events.
				var ev EmissionFinishedEvent
				if err := json.Unmarshal(body, &ev); err != nil {
					return nil, fmt.Errorf("alanube: parse emission event: %w", err)
				}
				event.EmissionFinished = &ev
				return event, nil
			}
		}
	}

	return nil, fmt.Errorf("alanube: unknown webhook event format")
}
