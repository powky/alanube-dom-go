package alanube

// CompanyDocumentTotals represents document totals for a specific company.
type CompanyDocumentTotals struct {
	TotalEmittedDocuments int `json:"totalEmittedDocuments"`
	FiscalInvoices        int `json:"fiscalInvoices"`
	Invoices              int `json:"invoices"`
	CreditNotes           int `json:"creditNotes"`
	DebitNotes            int `json:"debitNotes"`
	ExportSupports        int `json:"exportSupports"`
	Gubernamentals        int `json:"gubernamentals"`
	MinorExpenses         int `json:"minorExpenses"`
	PaymentAbroadSupports int `json:"paymentAbroadSupports"`
	Purchases             int `json:"purchases"`
	SpecialRegimes        int `json:"specialRegimes"`
}

// UserDocumentTotalsData contains the report data for user document totals.
type UserDocumentTotalsData struct {
	TotalEmittedDocuments int                              `json:"totalEmittedDocuments"`
	Companies             map[string]CompanyDocumentTotals `json:"companies"`
}

// UserDocumentTotalsResponse is the response for the user document totals report.
type UserDocumentTotalsResponse struct {
	Data *UserDocumentTotalsData `json:"data"`
}

// CompanyDocumentTotalsResponse is the response for the company document totals report.
type CompanyDocumentTotalsResponse struct {
	Data *CompanyDocumentTotals `json:"data"`
}

// CompanyAcceptedDocumentTotals represents accepted document totals for a specific company.
type CompanyAcceptedDocumentTotals struct {
	TotalAcceptedDocuments int `json:"totalAcceptedDocumentsla"` // note: API field name includes trailing "la"
	FiscalInvoices         int `json:"fiscalInvoices"`
	Invoices               int `json:"invoices"`
	CreditNotes            int `json:"creditNotes"`
	DebitNotes             int `json:"debitNotes"`
	ExportSupports         int `json:"exportSupports"`
	Gubernamentals         int `json:"gubernamentals"`
	MinorExpenses          int `json:"minorExpenses"`
	PaymentAbroadSupports  int `json:"paymentAbroadSupports"`
	Purchases              int `json:"purchases"`
	SpecialRegimes         int `json:"specialRegimes"`
}

// ReportQueryOptions represents query parameters for report endpoints.
type ReportQueryOptions struct {
	LegalStatus string // required, comma-separated legal statuses
	DateFrom    string // YYYY-MM-DD
	DateUntil   string // YYYY-MM-DD
}
