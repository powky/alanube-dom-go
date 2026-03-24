package alanube

// CreateDebitNoteRequest is the request body for creating a debit note (eCF type 33).
// POST /debit-notes
type CreateDebitNoteRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"`
	Pagination           []Pagination              `json:"pagination,omitempty"`
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}

// CreateCreditNoteRequest is the request body for creating a credit note (eCF type 34).
// POST /credit-notes
type CreateCreditNoteRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"`
	Pagination           []Pagination              `json:"pagination,omitempty"`
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}
