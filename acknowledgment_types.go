package alanube

// Acknowledgment represents an external reception acknowledgment.
type Acknowledgment struct {
	ID                     string  `json:"id"`
	RNC                    string  `json:"rnc"`
	DocumentNumber         string  `json:"documentNumber"`
	Status                 string  `json:"status"` // RECEIVED, NOT_RECEIVED
	ErrorCode              *int    `json:"errorCode,omitempty"`
	ErrorMsg               string  `json:"errorMsg,omitempty"`
	SupplierIdentification string  `json:"supplierIdentification,omitempty"`
	AcknowledgedAt         string  `json:"acknowledgedAt,omitempty"`
	XML                    string  `json:"xml,omitempty"`
}

// AcknowledgmentsQueryOptions represents query parameters for listing acknowledgments.
type AcknowledgmentsQueryOptions struct {
	DocumentNumber string // max 13, requires RNC
	RNC            string // 9-11 chars
	Status         string // RECEIVED, NOT_RECEIVED
	Start          string // YYYY-MM-DD or YYYY-MM-DD hh:mm:ss
	End            string // YYYY-MM-DD or YYYY-MM-DD hh:mm:ss
	Limit          int    // 1-100, default 25
	Page           int    // >= 1, default 1
}
