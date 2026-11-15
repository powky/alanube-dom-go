package alanube

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// FlexString is a string that also accepts a JSON number when decoding.
//
// The API documents some DGII response codes (e.g. governmentResponse.value[].codigo)
// as strings but returns them as numbers in practice, so both forms must decode.
type FlexString string

// UnmarshalJSON accepts a JSON string, number or null.
func (f *FlexString) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*f = ""
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*f = FlexString(s)
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return fmt.Errorf("alanube: FlexString: expected string or number, got %s", data)
	}
	*f = FlexString(n.String())
	return nil
}

// String returns the underlying string value.
func (f FlexString) String() string { return string(f) }
