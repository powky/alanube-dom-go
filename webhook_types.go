package alanube

// WebhookEventType identifies the type of webhook event.
type WebhookEventType string

const (
	WebhookGovernmentStatusChanged WebhookEventType = "general.governmentStatusChanged"
	WebhookEmissionFinished        WebhookEventType = "documents.emissionFinished"
	WebhookDocumentReception       WebhookEventType = "documents.reception"
)

// GovernmentStatusChangedEvent is the payload for the general.governmentStatusChanged webhook.
type GovernmentStatusChangedEvent struct {
	Company GovernmentStatusCompany `json:"company"`
}

// GovernmentStatusCompany is the company data within a government status change event.
type GovernmentStatusCompany struct {
	ID                string `json:"id"`
	HasBeenCertified  string `json:"hasBeenCertified"` // "yes" | "no"
	CertificationStep int    `json:"certificationStep"`
	Type              string `json:"type"`
}

// EmissionFinishedEvent is the payload for the documents.emissionFinished webhook.
type EmissionFinishedEvent struct {
	Type                  string              `json:"type"`
	ID                    string              `json:"id"`
	StampDate             string              `json:"stampDate"`
	Status                string              `json:"status"`
	LegalStatus           string              `json:"legalStatus"`
	CompanyIdentification string              `json:"companyIdentification"`
	TrackID               string              `json:"trackId"`
	DocumentNumber        string              `json:"documentNumber"`
	SequenceConsumed      bool                `json:"sequenceConsumed"`
	SignatureDate         string              `json:"signatureDate"`
	SecurityCode          string              `json:"securityCode"`
	DocumentStampURL      string              `json:"documentStampUrl,omitempty"`
	XML                   string              `json:"xml,omitempty"`
	ResumeXML             string              `json:"resumeXml,omitempty"`
	PDF                   string              `json:"pdf,omitempty"`
	GovernmentResponse    *GovernmentResponse `json:"governmentResponse,omitempty"`
	Error                 *DocumentError      `json:"error,omitempty"`
}

// DocumentReceptionEvent is the payload for the documents.reception webhook.
// The Type field determines which variant this is.
type DocumentReceptionEvent struct {
	Type                 string              `json:"type"` // "receivedDocument", "receivedCommercialApproval", "sentCommercialApproval"
	ID                   string              `json:"id"`
	IssuerIdentification string              `json:"issuerIdentification,omitempty"`
	BuyerIdentification  string              `json:"buyerIdentification,omitempty"`
	DocumentType         string              `json:"documentType,omitempty"`
	DocumentNumber       string              `json:"documentNumber,omitempty"`
	DocumentStampDate    string              `json:"documentStampDate,omitempty"`
	SignatureDateTime    string              `json:"signatureDateTime,omitempty"`
	TotalAmount          string              `json:"totalAmount,omitempty"`
	Status               string              `json:"status,omitempty"`
	ErrorMsg             *string             `json:"errorMsg,omitempty"`
	CommercialResponse   string              `json:"commercialResponse,omitempty"`
	Timestamp            string              `json:"timestamp"`
	Receipt              *ReceivedDocReceipt `json:"receipt,omitempty"`
	// Fields for sentCommercialApproval
	IssuerStatus       string `json:"issuerStatus,omitempty"`
	IssuerSendAttempts int    `json:"issuerSendAttempts,omitempty"`
	DGIIStatus         string `json:"dgiiStatus,omitempty"`
	DGIINotifyAttempts int    `json:"dgiiNotifyAttempts,omitempty"`
}
