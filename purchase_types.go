package alanube

// CreatePurchaseRequest is the request body for creating a purchase document (eCF type 41).
// POST /purchases
type CreatePurchaseRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer,omitempty"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"`
	Pagination           []Pagination              `json:"pagination,omitempty"`
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}

// CreateMinorExpenseRequest is the request body for creating a minor expense document (eCF type 43).
// POST /minor-expenses
type CreateMinorExpenseRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer,omitempty"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"`
	Pagination           []Pagination              `json:"pagination,omitempty"`
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}

// CreatePaymentAbroadRequest is the request body for creating a payment abroad support document (eCF type 47).
// POST /payments-abroad
type CreatePaymentAbroadRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer,omitempty"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"`
	Pagination           []Pagination              `json:"pagination,omitempty"`
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}
