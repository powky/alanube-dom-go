package alanube

// ItemExample defines example item data for test set generation.
type ItemExample struct {
	BillingIndicator     int    `json:"billingIndicator"`           // 1=ITBIS 18%, 2=ITBIS 16%, 3=ITBIS 0%, 4=Exento
	ItemName             string `json:"itemName,omitempty"`         // max 80 chars
	GoodServiceIndicator int    `json:"goodServiceIndicator"`      // 1=Bien, 2=Servicio
	ItemDescription      string `json:"itemDescription,omitempty"` // max 1000 chars
	UnitPriceItem        int    `json:"unitPriceItem"`             // 1 to 99999999
}

// CreateSetTestRequest is the request body for generating a certification test set.
type CreateSetTestRequest struct {
	IDCompany   string       `json:"idCompany,omitempty"`
	RetryNumber int          `json:"retryNumber,omitempty"` // 0 to 99
	ItemExample *ItemExample `json:"itemExample"`
}

// SetTestCreateResponse is the response when creating a test set (POST /set-tests).
type SetTestCreateResponse struct {
	ID                    string `json:"id"`
	Status                string `json:"status"`
	IDCompany             string `json:"idCompany"`
	CompanyIdentification string `json:"companyIdentification"`
	CreationDate          string `json:"creationDate"`
}

// SetTestDocumentInfo represents a document within a test set check response.
type SetTestDocumentInfo struct {
	DocumentType string `json:"documentType"`
	IDDocument   string `json:"idDocument"`
	Status       string `json:"status"` // ACCEPTED, REJECTED, IN_PROGRESS
	ENCF         string `json:"encf,omitempty"`
	XML          string `json:"xml,omitempty"`
	PDF          string `json:"pdf,omitempty"`
}

// SetTestResumeDocument represents a resume document within a test set check response.
type SetTestResumeDocument struct {
	DocumentType string `json:"documentType"`
	IDDocument   string `json:"idDocument"`
	Status       string `json:"status"` // ACCEPTED, REJECTED, IN_PROGRESS
	ENCF         string `json:"encf,omitempty"`
}

// SetTestResumesInfo contains resume documents info and a zip download URL.
type SetTestResumesInfo struct {
	ZipURL    string                  `json:"zipUrl,omitempty"`
	Documents []SetTestResumeDocument `json:"documents,omitempty"`
}

// SetTestCheckResponse is the response when checking a test set (GET /check-set-tests/{id}).
type SetTestCheckResponse struct {
	TestID                 string                `json:"testId"`
	IDCompany              string                `json:"idCompany"`
	Status                 string                `json:"status"`
	RetryNumber            int                   `json:"retryNumber"`
	TotalDocumentProcessed int                   `json:"totalDocumentProcessed"`
	DocumentsInfo          []SetTestDocumentInfo `json:"documentsInfo,omitempty"`
	ResumesInfo            *SetTestResumesInfo   `json:"resumesInfo,omitempty"`
}

// SignDocumentResponse is the response for a signed document.
type SignDocumentResponse struct {
	SignedDocumentURL string `json:"signedDocumentUrl"`
}

// ProviderInfoField represents a field with description and value in the provider info response.
type ProviderInfoField struct {
	Description string      `json:"description"`
	Value       interface{} `json:"value,omitempty"`
}

// ProviderDataRNC represents the RNC field within provider data.
type ProviderDataRNC struct {
	Description string `json:"description"`
	Value       int    `json:"value"`
}

// ProviderDataName represents a name field within provider data.
type ProviderDataName struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// ProviderData contains the provider company details.
type ProviderData struct {
	Description string            `json:"description"`
	RNC         *ProviderDataRNC  `json:"rnc"`
	Name        *ProviderDataName `json:"name"`
	TradeName   *ProviderDataName `json:"tradeName"`
}

// ProviderInfoData contains the full provider information.
type ProviderInfoData struct {
	SoftwareType    *ProviderInfoField `json:"softwareType"`
	SoftwareName    *ProviderInfoField `json:"softwareName"`
	SoftwareVersion *ProviderInfoField `json:"softwareVersion"`
	ProviderData    *ProviderData      `json:"providerData"`
}

// ProviderInfo is the response wrapper for the provider info endpoint.
type ProviderInfo struct {
	Data *ProviderInfoData `json:"data"`
}
