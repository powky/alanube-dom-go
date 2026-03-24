package alanube

// Certificate represents a digital certificate for request bodies.
type Certificate struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Content   string `json:"content"`
	Password  string `json:"password"`
}

// CertificateResponse represents certificate details returned by the API.
type CertificateResponse struct {
	Name       string `json:"name"`
	Extension  string `json:"extension"`
	IssuerName string `json:"issuerName"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

// --- Webhooks (nested structure matching API exactly) ---

// WebhookAuthCacheConfig configures caching for webhook auth responses.
type WebhookAuthCacheConfig struct {
	Status string `json:"status"`           // "active" | "inactive"
	TTL    int    `json:"ttl,omitempty"`     // cache TTL in seconds
}

// WebhookAuthHeadersResponse configures which headers from the auth response to use.
type WebhookAuthHeadersResponse struct {
	Fields []string                `json:"fields"`
	Cache  *WebhookAuthCacheConfig `json:"cache,omitempty"`
}

// WebhookAuth configures authentication for a webhook endpoint.
type WebhookAuth struct {
	Status          string                      `json:"status"`
	URL             string                      `json:"url,omitempty"`
	HeadersRequest  map[string]string           `json:"headersRequest,omitempty"`
	Body            map[string]interface{}       `json:"body,omitempty"`
	HeadersResponse *WebhookAuthHeadersResponse `json:"headersResponse,omitempty"`
}

// WebhookEndpoint configures a single webhook endpoint.
type WebhookEndpoint struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	Status  string            `json:"status,omitempty"` // "active" | "inactive"
	Auth    *WebhookAuth      `json:"auth,omitempty"`
}

// WebhooksGeneralConfig configures general webhooks.
type WebhooksGeneralConfig struct {
	GovernmentStatusChanged *WebhookEndpoint `json:"governmentStatusChanged,omitempty"`
}

// WebhooksDocumentsConfig configures document-related webhooks.
type WebhooksDocumentsConfig struct {
	EmissionFinished *WebhookEndpoint `json:"emissionFinished,omitempty"`
	Cancellations    *WebhookEndpoint `json:"cancellations,omitempty"`
	Reception        *WebhookEndpoint `json:"reception,omitempty"`
}

// WebhooksConfig configures all webhook endpoints for a company.
type WebhooksConfig struct {
	General   *WebhooksGeneralConfig   `json:"general,omitempty"`
	Documents *WebhooksDocumentsConfig `json:"documents,omitempty"`
}

// --- Notification by email ---

// NotificationByEmailConfig configures email notifications to document recipients.
type NotificationByEmailConfig struct {
	Enabled bool   `json:"enabled"`
	Message string `json:"message,omitempty"`
}

// --- Company URLs ---

// CompanyURLs contains the URLs assigned to a company for DGII certification.
type CompanyURLs struct {
	Reception      string `json:"reception"`
	Approval       string `json:"approval"`
	Authentication string `json:"authentication"`
}

// --- Company (response) ---

// Company represents a registered company in Alanube.
type Company struct {
	ID                  string                     `json:"id"`
	Name                string                     `json:"name"`
	TradeName           string                     `json:"tradeName"`
	Identification      string                     `json:"identification"`
	Type                string                     `json:"type"` // "main" | "associated"
	Address             string                     `json:"address"`
	Province            string                     `json:"province"`
	Municipality        string                     `json:"municipality"`
	Email               string                     `json:"email"`
	CompanyURLs         *CompanyURLs               `json:"companyUrls,omitempty"`
	Certificate         *CertificateResponse       `json:"certificate,omitempty"`
	Webhooks            *WebhooksConfig            `json:"webhooks,omitempty"`
	NotificationByEmail *NotificationByEmailConfig `json:"notificationByEmail,omitempty"`
}

// --- Requests ---

// CreateCompanyRequest is the request body for creating a company.
type CreateCompanyRequest struct {
	Name                string                     `json:"name"`
	TradeName           string                     `json:"tradeName"`
	Identification      string                     `json:"identification"`
	Address             string                     `json:"address"`
	Province            string                     `json:"province"`
	Municipality        string                     `json:"municipality"`
	Type                string                     `json:"type,omitempty"` // "main" | "associated"
	Email               string                     `json:"email,omitempty"`
	Certificate         *Certificate               `json:"certificate"`
	Webhooks            *WebhooksConfig            `json:"webhooks,omitempty"`
	NotificationByEmail *NotificationByEmailConfig `json:"notificationByEmail,omitempty"`
	Logo                string                     `json:"logo,omitempty"`
}

// UpdateCompanyRequest is the request body for updating a company (PATCH).
type UpdateCompanyRequest struct {
	Name                string                     `json:"name,omitempty"`
	TradeName           string                     `json:"tradeName,omitempty"`
	Identification      string                     `json:"identification,omitempty"`
	Address             string                     `json:"address,omitempty"`
	Province            string                     `json:"province,omitempty"`
	Municipality        string                     `json:"municipality,omitempty"`
	Type                string                     `json:"type,omitempty"`
	Email               string                     `json:"email,omitempty"`
	Certificate         *Certificate               `json:"certificate,omitempty"`
	Webhooks            *WebhooksConfig            `json:"webhooks,omitempty"`
	NotificationByEmail *NotificationByEmailConfig `json:"notificationByEmail,omitempty"`
	Logo                string                     `json:"logo,omitempty"`
}

// companyWrapper wraps the API response for company endpoints.
type companyWrapper struct {
	Company *Company `json:"company"`
}
