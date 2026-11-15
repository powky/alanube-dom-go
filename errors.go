package alanube

import (
	"encoding/json"
	"fmt"
)

// Error represents an error response from the Alanube API.
// The API returns different error formats depending on the status code:
//   - 400: {"message": ["..."], "code": 400}
//   - 404: {"message": "...", "code": 404}
//   - 500: {"message": "...", "code": "EPR500"}
type Error struct {
	StatusCode int
	Code       string   // may be numeric string ("400") or error code ("EPR500")
	Messages   []string // one or more error messages
}

func (e *Error) Error() string {
	if len(e.Messages) > 0 {
		return fmt.Sprintf("alanube: %d %s: %s", e.StatusCode, e.Code, e.Messages[0])
	}
	return fmt.Sprintf("alanube: %d %s", e.StatusCode, e.Code)
}

// DecodeError is returned when the API answered with a success status but the
// body could not be decoded into the expected type. The operation DID succeed
// on the server side (e.g. the document was created), so Body is kept for
// recovery — never retry blindly on this error.
type DecodeError struct {
	StatusCode int
	Body       []byte
	Err        error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("alanube: decode %d response: %v", e.StatusCode, e.Err)
}

func (e *DecodeError) Unwrap() error { return e.Err }

// UnmarshalJSON handles the different error response formats from the API.
func (e *Error) UnmarshalJSON(data []byte) error {
	// Try format: {"message": [...], "code": ...} or {"message": "...", "code": ...}
	var raw struct {
		Message json.RawMessage `json:"message"`
		Code    json.RawMessage `json:"code"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Parse code (can be int or string)
	if raw.Code != nil {
		var codeInt int
		if err := json.Unmarshal(raw.Code, &codeInt); err == nil {
			e.Code = fmt.Sprintf("%d", codeInt)
		} else {
			var codeStr string
			if err := json.Unmarshal(raw.Code, &codeStr); err == nil {
				e.Code = codeStr
			}
		}
	}

	// Parse message (can be string or []string)
	if raw.Message != nil {
		var msgStr string
		if err := json.Unmarshal(raw.Message, &msgStr); err == nil {
			e.Messages = []string{msgStr}
		} else {
			var msgArr []string
			if err := json.Unmarshal(raw.Message, &msgArr); err == nil {
				e.Messages = msgArr
			}
		}
	}

	return nil
}
