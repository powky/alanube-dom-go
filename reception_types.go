package alanube

// ReceivedDocumentsQueryOptions represents query parameters for listing received documents.
type ReceivedDocumentsQueryOptions struct {
	Status             string // RECEIVED, NOT_RECEIVED
	CommercialResponse string // NOT_DECLARED, ACCEPTED, REJECTED
	RNC                string // 9-11 chars
	DocumentNumber     string // max 13
	Limit              int    // 1-1000, default 25
	Page               int    // >= 1, default 1
	Start              string // YYYY-MM-DD
	End                string // YYYY-MM-DD
}

// PaginationMetadata represents pagination info in list responses.
type PaginationMetadata struct {
	CurrentPage int `json:"current_page"`
	Limit       int `json:"limit"`
	From        int `json:"from"`
	To          int `json:"to"`
}

// ReceivedDocument represents a document received through the electronic mailbox.
type ReceivedDocument struct {
	ID                   string  `json:"id"`
	IssuerIdentification string  `json:"issuerIdentification"`
	BuyerIdentification  string  `json:"buyerIdentification"`
	DocumentType         string  `json:"documentType"`
	DocumentNumber       string  `json:"documentNumber"`
	DocumentStampDate    string  `json:"documentStampDate"`
	SignatureDateTime    string  `json:"signatureDateTime"`
	TotalAmount          string  `json:"totalAmount"`
	Status               string  `json:"status"` // RECEIVED, NOT_RECEIVED
	ErrorMsg             *string `json:"errorMsg"`
	CommercialResponse   string  `json:"commercialResponse"` // ACCEPTED, REJECTED, NOT_DECLARED
	Timestamp            string  `json:"timestamp"`
	XML                  string  `json:"xml,omitempty"` // present when fetching by ID
}

// ReceivedDocumentsResponse is the paginated response for listing received documents.
type ReceivedDocumentsResponse struct {
	Metadata  *PaginationMetadata `json:"metadata"`
	Documents []ReceivedDocument  `json:"documents"`
}

// ReceivedDocReceipt is the receipt acknowledgment for a received document.
type ReceivedDocReceipt struct {
	ID             string  `json:"id"`
	DocumentNumber string  `json:"documentNumber"`
	Status         string  `json:"status"`
	ErrorMsg       *string `json:"errorMsg"`
	Timestamp      string  `json:"timestamp"`
}

// ReceiveDocumentRequest is the request body for simulating document reception from DGII (sandbox only).
// POST /reception/dgii/receive-document
type ReceiveDocumentRequest struct {
	IDCompany    string `json:"idCompany"`    // required
	DocumentType string `json:"documentType"` // required: "31", "33", "34", "44", "47"
}

// CommercialResponseRequest is the request body for generating a commercial response.
type CommercialResponseRequest struct {
	CommercialResponse string `json:"commercialResponse"` // required: "ACCEPTED" or "REJECTED"
	NotAcceptedDetail  string `json:"notAcceptedDetail,omitempty"`
}

// CommercialResponseResult is the response when generating a commercial response.
type CommercialResponseResult struct {
	CommercialApprovalInfo *CommercialApprovalInfo `json:"commercialApprovalInfo"`
}

// CommercialApprovalInfo contains the ID of the generated commercial approval.
type CommercialApprovalInfo struct {
	ID string `json:"id"`
}

// CommercialApproval represents a commercial approval record.
type CommercialApproval struct {
	ID                   string `json:"id"`
	DocumentNumber       string `json:"documentNumber"`
	BuyerIdentification  string `json:"buyerIdentification"`
	IssuerIdentification string `json:"issuerIdentification"`
	IssuerStatus         string `json:"issuerStatus,omitempty"`
	IssuerSendAttempts   int    `json:"issuerSendAttempts,omitempty"`
	DGIIStatus           string `json:"dgiiStatus,omitempty"`
	DGIINotifyAttempts   int    `json:"dgiiNotifyAttempts,omitempty"`
	Timestamp            string `json:"timestamp"`
	XML                  string `json:"xml,omitempty"` // present when fetching by ID
}

// CommercialApprovalsQueryOptions represents query parameters for listing commercial approvals.
type CommercialApprovalsQueryOptions struct {
	Limit int // required, 1-1000
	Page  int // required, >= 1
}

// CommercialApprovalsResponse is the paginated response for listing commercial approvals.
type CommercialApprovalsResponse struct {
	Metadata  *PaginationMetadata  `json:"metadata"`
	Documents []CommercialApproval `json:"documents"`
}
