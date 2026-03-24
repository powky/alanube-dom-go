package alanube

// RangeCancelledENCF represents a range of cancelled eNCF sequences.
type RangeCancelledENCF struct {
	ENCFFrom string `json:"encfFrom"` // required, max 13
	ENCFUntil string `json:"encfUntil"` // required, max 13
}

// CancellationLine represents a single cancellation line within the request.
type CancellationLine struct {
	LineNumber            int                    `json:"lineNumber"`            // required, 1-10
	ECFType               int                    `json:"ecfType"`               // required, 31-47
	RangeCancelledENCF    []RangeCancelledENCF   `json:"rangeCancelledEnfc"`    // required, max 10000
	CancelledENCFQuantity int                    `json:"cancelledEncfQuantity"` // required
}

// CancellationHeader represents the cancellation header. Campo DGII <Encabezado>
type CancellationHeader struct {
	RNCSender             int    `json:"rncSender"`             // required
	CancelledENCFQuantity string `json:"cancelledEncfQuantity"` // required
}

// CreateCancellationRequest is the request body for creating a cancellation.
// POST /cancellations
type CreateCancellationRequest struct {
	Company       *DocumentCompany     `json:"company,omitempty"`
	Header        *CancellationHeader  `json:"header"`
	Cancellations []CancellationLine   `json:"cancellations"` // required, max 10
}

// CancellationResponse represents the response when creating a cancellation.
type CancellationResponse struct {
	ID                    string `json:"id"`
	StampDate             string `json:"stampDate,omitempty"`
	Status                string `json:"status"` // REGISTERED, TO_SEND, WAITING_RESPONSE, TO_NOTIFY, FINISHED
	CompanyIdentification string `json:"companyIdentification,omitempty"`
}

// CancellationGovernmentResponse represents the government response for a cancellation (flat structure).
type CancellationGovernmentResponse struct {
	Codigo string `json:"codigo"`
	Valor  string `json:"valor"`
}

// CancellationStatus represents the status of a cancellation (GET /cancellations/{id}).
type CancellationStatus struct {
	ID                    string                          `json:"id"`
	StampDate             string                          `json:"stampDate,omitempty"`
	Status                string                          `json:"status"`
	CompanyIdentification string                          `json:"companyIdentification,omitempty"`
	XML                   string                          `json:"xml,omitempty"`
	GovernmentResponse    *CancellationGovernmentResponse `json:"governmentResponse,omitempty"`
}
