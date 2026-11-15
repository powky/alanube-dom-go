package alanube

// --- Shared document types used across invoices, notes, purchases, etc. ---

// DocumentCompany identifies the company for document emission (when not using the main one).
type DocumentCompany struct {
	ID string `json:"id"`
}

// PaymentForm represents a payment form entry in the payment forms table.
type PaymentForm struct {
	PaymentMethod        *int    `json:"paymentMethod,omitempty"`        // 1-8
	PaymentAmount        *float64 `json:"paymentAmount,omitempty"`
	PaymentAccountType   string  `json:"paymentAccountType,omitempty"`   // CT, AH, OT
	PaymentAccountNumber string  `json:"paymentAccountNumber,omitempty"` // max 28
	BankPayment          string  `json:"bankPayment,omitempty"`          // max 75
}

// IdDoc represents document identification. Campo DGII <IdDoc>
type IdDoc struct {
	ENCF                      string        `json:"encf"`                                // required, 13 chars
	SequenceDueDate           string        `json:"sequenceDueDate,omitempty"`            // date
	CreditNoteIndicator       *int          `json:"creditNoteIndicator,omitempty"`        // 0 or 1, required for credit notes (34)
	DeferredDeliveryIndicator *int          `json:"deferredDeliveryIndicator,omitempty"`
	TaxAmountIndicator        *int          `json:"taxAmountIndicator,omitempty"`         // 1 or 2
	IncomeType                int           `json:"incomeType"`                           // required, 1-6
	PaymentType               int           `json:"paymentType"`                          // required, 1-3
	PaymentDeadline           string        `json:"paymentDeadline,omitempty"`            // date
	PaymentTerm               string        `json:"paymentTerm,omitempty"`                // max 15
	PaymentFormsTable         []PaymentForm `json:"paymentFormsTable,omitempty"`          // max 7
	DateFrom                  string        `json:"dateFrom,omitempty"`                   // date
	DateUntil                 string        `json:"dateUntil,omitempty"`                  // date
	TotalPages                *int          `json:"totalPages,omitempty"`
}

// Sender represents the document issuer. Campo DGII <Emisor>
type Sender struct {
	RNC                        string   `json:"rnc"`                                  // required, max 11
	CompanyName                string   `json:"companyName"`                          // required, max 150
	Tradename                  string   `json:"tradename,omitempty"`                  // max 150
	BranchOffice               string   `json:"branchOffice,omitempty"`               // max 20
	Address                    string   `json:"address"`                              // required, max 100
	Municipality               string   `json:"municipality,omitempty"`
	Province                   string   `json:"province,omitempty"`
	PhoneNumber                []string `json:"phoneNumber,omitempty"`                // max 3
	Mail                       string   `json:"mail,omitempty"`                       // max 80
	WebSite                    string   `json:"webSite,omitempty"`                    // max 50
	EconomicActivity           string   `json:"economicActivity,omitempty"`           // max 100
	SellerCode                 string   `json:"sellerCode,omitempty"`                 // max 60
	InternalInvoiceNumber      string   `json:"internalInvoiceNumber,omitempty"`      // max 20
	InternalOrderNumber        string   `json:"internalOrderNumber,omitempty"`        // max 20
	SaleArea                   string   `json:"saleArea,omitempty"`                   // max 20
	SaleRoute                  string   `json:"saleRoute,omitempty"`                  // max 20
	AdditionalInformationIssuer string  `json:"additionalInformationIssuer,omitempty"` // max 250
	StampDate                  string   `json:"stampDate"`                            // required, date
}

// Buyer represents the document recipient. Campo DGII <Comprador>
type Buyer struct {
	RNC                   string `json:"rnc,omitempty"`                   // max 11
	ForeignIdentifier     string `json:"foreignIdentifier,omitempty"`     // max 20
	CompanyName           string `json:"companyName,omitempty"`           // max 150
	Contact               string `json:"contact,omitempty"`               // max 80
	Mail                  string `json:"mail,omitempty"`                  // max 80
	Address               string `json:"address,omitempty"`               // max 100
	Municipality          string `json:"municipality,omitempty"`
	Province              string `json:"province,omitempty"`
	BuyerCountry          string `json:"buyerCountry,omitempty"`          // max 60, for export (46)
	DeliverDate           string `json:"deliverDate,omitempty"`           // date
	ContactDelivery       string `json:"contactDelivery,omitempty"`       // max 100
	DeliveryAddress       string `json:"deliveryAddress,omitempty"`       // max 100
	AdditionalPhone       string `json:"additionalPhone,omitempty"`       // max 12
	PurchaseOrderDate     string `json:"purchaseOrderDate,omitempty"`     // date
	PurchaseOrderNumber   string `json:"purchaseOrderNumber,omitempty"`   // max 20
	InternalCode          string `json:"internalCode,omitempty"`          // max 20
	ResponsibleForPayment string `json:"responsibleForPayment,omitempty"` // max 20
	AdditionalInformation string `json:"additionalInformation,omitempty"` // max 150
}

// Transport represents transport information. Campo DGII <Transporte>
type Transport struct {
	MeansOfTransport                    *int   `json:"meansOfTransport,omitempty"`                    // 1=Terrestre, 2=Marítimo, 3=Aérea (export)
	OriginCountry                       string `json:"originCountry,omitempty"`                       // max 60 (export)
	DestinationAddress                  string `json:"destinationAddress,omitempty"`                  // max 100 (export)
	DestinationCountry                  string `json:"destinationCountry,omitempty"`                  // max 60 (export)
	RNCCarrierCompanyIdentification     string `json:"rncCarrierCompanyIdentification,omitempty"`     // max 20 (export)
	CarrierCompanyName                  string `json:"carrierCompanyName,omitempty"`                  // max 150 (export)
	TripNumber                          string `json:"tripNumber,omitempty"`                          // max 20 (export)
	Driver                              string `json:"driver,omitempty"`                              // max 20
	TransportDocument                   *int   `json:"transportDocument,omitempty"`
	File                                string `json:"file,omitempty"`                                // max 20
	LicensePlate                        string `json:"licensePlate,omitempty"`                        // max 7
	TransportationRoute                 string `json:"transportationRoute,omitempty"`                 // max 20
	TransportationZone                  string `json:"transportationZone,omitempty"`                  // max 20
	AlbaranNumber                       string `json:"albaranNumber,omitempty"`                       // max 20
}

// AdditionalInformationDoc represents additional document information. Campo DGII <InformacionesAdicionales>
type AdditionalInformationDoc struct {
	ShippingDate        string     `json:"shippingDate,omitempty"`        // date
	ShipmentNumber      string     `json:"shipmentNumber,omitempty"`      // max 25
	ContainerNumber     string     `json:"containerNumber,omitempty"`     // max 100
	ReferenceNumber     *int       `json:"referenceNumber,omitempty"`
	ShipmentPortName    string     `json:"shipmentPortName,omitempty"`    // max 40 (export)
	DeliveryConditions  string     `json:"deliveryConditions,omitempty"`  // max 3, CIF/FOB etc (export)
	TotalFob            *float64   `json:"totalFob,omitempty"`            // export
	Insurance           *float64   `json:"insurance,omitempty"`           // export
	Freight             *float64   `json:"freight,omitempty"`             // export
	OtherExpenses       *float64   `json:"otherExpenses,omitempty"`       // export
	TotalCif            *float64   `json:"totalCif,omitempty"`            // export
	CustomsRegime       string     `json:"customsRegime,omitempty"`       // max 35 (export)
	ExitPortName        string     `json:"exitPortName,omitempty"`        // max 40 (export)
	LandingPortName     string     `json:"landingPortName,omitempty"`     // max 40 (export)
	GrossWeight         *float64   `json:"grossWeight,omitempty"`
	NetWeight           *float64   `json:"netWeight,omitempty"`
	GrossWeightUnit     *int       `json:"grossWeightUnit,omitempty"`
	UnitNetWeight       *int       `json:"unitNetWeight,omitempty"`
	BulkQuantity        *float64   `json:"bulkQuantity,omitempty"`
	BulkUnit            *int       `json:"bulkUnit,omitempty"`
	BulkVolume          *float64   `json:"bulkVolume,omitempty"`
	UnitVolume          *int       `json:"unitVolume,omitempty"`
	Transport           *Transport `json:"transport,omitempty"`
}

// --- Totals ---

// AdditionalTax represents an additional tax entry. Campo DGII <ImpuestosAdicionales>
type AdditionalTax struct {
	TaxType                                    int      `json:"taxType"`                                              // required, 1-39
	AdditionalTaxRate                          *float64 `json:"additionalTaxRate,omitempty"`                           // 0-999.99
	SelectiveTaxAmountSpecificConsumption      *float64 `json:"selectiveTaxAmountSpecificConsumption,omitempty"`
	AmountSelectiveConsumptionTaxAdvalorem     *float64 `json:"amountSelectiveConsumptionTaxAdvalorem,omitempty"`
	OtherAdditionalTaxes                       *float64 `json:"otherAdditionalTaxes,omitempty"`
}

// AdditionalTaxOtherCurrency represents an additional tax in another currency.
type AdditionalTaxOtherCurrency struct {
	TaxTypeOtherCurrency                                       int      `json:"taxTypeOtherCurrency"`                                       // required, 1-39
	AdditionalTaxRateOtherCurrency                             *float64 `json:"additionalTaxRateOtherCurrency,omitempty"`
	SelectiveTaxAmountSpecificConsumptionOtherCurrency         *float64 `json:"selectiveTaxAmountSpecificConsumptionOtherCurrency,omitempty"`
	AmountSelectiveConsumptionTaxAdvaloremOtherCurrency        *float64 `json:"amountSelectiveConsumptionTaxAdvaloremOtherCurrency,omitempty"`
	OtherAdditionalTaxesOtherCurrency                          *float64 `json:"otherAdditionalTaxesOtherCurrency,omitempty"`
}

// OtherCurrency represents totals in an alternative currency. Campo DGII <OtraMoneda>
type OtherCurrency struct {
	CurrencyType                         string                       `json:"currencyType"`                                   // required
	ExchangeRate                         float64                      `json:"exchangeRate"`                                   // required
	TotalTaxedAmountOtherCurrency        *float64                     `json:"totalTaxedAmountOtherCurrency,omitempty"`
	AmountTaxed1OtherCurrency            *float64                     `json:"amountTaxed1OtherCurrency,omitempty"`
	AmountTaxed2OtherCurrency            *float64                     `json:"amountTaxed2OtherCurrency,omitempty"`
	AmountTaxed3OtherCurrency            *float64                     `json:"amountTaxed3OtherCurrency,omitempty"`
	ExemptAmountOtherCurrency            *float64                     `json:"exemptAmountOtherCurrency,omitempty"`
	ItbisTotalOtherCurrency              *float64                     `json:"itbisTotalOtherCurrency,omitempty"`
	Itbis1TotalOtherCurrency             *float64                     `json:"itbis1TotalOtherCurrency,omitempty"`
	Itbis2TotalOtherCurrency             *float64                     `json:"itbis2TotalOtherCurrency,omitempty"`
	Itbis3TotalOtherCurrency             *float64                     `json:"itbis3TotalOtherCurrency,omitempty"`
	AdditionalTaxAmountOtherCurrency     *float64                     `json:"additionalTaxAmountOtherCurrency,omitempty"`
	AdditionalTaxesOtherCurrency         []AdditionalTaxOtherCurrency `json:"additionalTaxesOtherCurrency,omitempty"` // max 20
	TotalAmountOtherCurrency             float64                      `json:"totalAmountOtherCurrency"`               // required
}

// Totals represents the document totals. Campo DGII <Totales>
type Totals struct {
	TotalTaxedAmount      *float64         `json:"totalTaxedAmount,omitempty"`
	I1AmountTaxed         *float64         `json:"i1AmountTaxed,omitempty"`
	I2AmountTaxed         *float64         `json:"i2AmountTaxed,omitempty"`
	I3AmountTaxed         *float64         `json:"i3AmountTaxed,omitempty"`
	ExemptAmount          *float64         `json:"exemptAmount,omitempty"`
	ItbisS1               *int             `json:"itbisS1,omitempty"`
	ItbisS2               *int             `json:"itbisS2,omitempty"`
	ItbisS3               *int             `json:"itbisS3,omitempty"`
	ItbisTotal            *float64         `json:"itbisTotal,omitempty"`
	Itbis1Total           *float64         `json:"itbis1Total,omitempty"`
	Itbis2Total           *float64         `json:"itbis2Total,omitempty"`
	Itbis3Total           *float64         `json:"itbis3Total,omitempty"`
	AdditionalTaxAmount   *float64         `json:"additionalTaxAmount,omitempty"`
	AdditionalTaxes       []AdditionalTax  `json:"additionalTaxes,omitempty"` // max 20
	TotalAmount           float64          `json:"totalAmount"`               // required
	NonBillableAmount     *float64         `json:"nonBillableAmount,omitempty"`
	AmountPeriod          *float64         `json:"amountPeriod,omitempty"`
	PreviousBalance       *float64         `json:"previousBalance,omitempty"`
	AmountAdvancePayment  *float64         `json:"amountAdvancePayment,omitempty"`
	PayValue              *float64         `json:"payValue,omitempty"`
	ItbisTotalRetained    *float64         `json:"itbisTotalRetained,omitempty"`
	IsrTotalRetention     *float64         `json:"isrTotalRetention,omitempty"`
	ItbisTotalPerception  *float64         `json:"itbisTotalPerception,omitempty"`
	IsrTotalPerception    *float64         `json:"isrTotalPerception,omitempty"`
	OtherCurrency         *OtherCurrency   `json:"otherCurrency,omitempty"`
}

// --- Item Details ---

// ItemCodeEntry represents an item code in the item code table.
type ItemCodeEntry struct {
	CodeType string `json:"codeType,omitempty"` // max 14
	ItemCode string `json:"itemCode,omitempty"`
}

// Retention represents retention info on an item. Campo DGII <Retencion>
type Retention struct {
	IndicatorAgentWithholdingPerception *int     `json:"indicatorAgentWithholdingPerception,omitempty"` // 1=R, 2=P
	ItbisAmountWithheld                *float64 `json:"itbisAmountWithheld,omitempty"`
	IsrAmountWithheld                  *float64 `json:"isrAmountWithheld,omitempty"`
}

// SubquantityEntry represents a subquantity table entry.
type SubquantityEntry struct {
	Subquantity     *float64 `json:"subquantity,omitempty"`
	CodeSubquantity *int     `json:"codeSubquantity,omitempty"`
}

// SubDiscount represents a sub-discount on an item.
type SubDiscount struct {
	SubDiscountRate       string   `json:"subDiscountRate,omitempty"`       // "%" or "$"
	SubDiscountPercentage *float64 `json:"subDiscountPercentage,omitempty"`
	SubDiscountAmount     *float64 `json:"subDiscountAmount,omitempty"`
}

// SubSurcharge represents a sub-surcharge on an item.
type SubSurcharge struct {
	SubSurchargeType       string   `json:"subSurchargeType,omitempty"`       // "%" or "$"
	SubSurchargePercentage *float64 `json:"subSurchargePercentage,omitempty"`
	SubSurchargeAmount     *float64 `json:"subSurchargeAmount,omitempty"`
}

// ItemAdditionalTax represents an additional tax on an item line.
type ItemAdditionalTax struct {
	TaxType int `json:"taxType"` // required, 1-39
}

// OtherCurrencyDetail represents item-level amounts in another currency.
type OtherCurrencyDetail struct {
	PriceOtherCurrency           float64  `json:"priceOtherCurrency"`                     // required
	DiscountOtherCurrency        *float64 `json:"discountOtherCurrency,omitempty"`
	SurchargeAnotherCurrency     *float64 `json:"surchargeAnotherCurrency,omitempty"`
	AmountItemOtherCurrency      float64  `json:"amountItemOtherCurrency"`                // required
}

// Mining represents mining-related fields on an item. Campo DGII <Mineria>
type Mining struct {
	NetWeightKilogram *float64 `json:"netWeightKilogram,omitempty"`
	NetWeightMining   *float64 `json:"netWeightMining,omitempty"`
	AffiliationType   *int     `json:"affiliationType,omitempty"` // 1=Afiliada, 2=No afiliada
	Settlement        *int     `json:"settlement,omitempty"`      // 1=Provisional, 2=Final
}

// ItemDetail represents a line item in a document. Campo DGII <DetallesItem>
type ItemDetail struct {
	LineNumber          int                  `json:"lineNumber"`                    // required, 1-1000
	ItemCodeTable       []ItemCodeEntry      `json:"itemCodeTable,omitempty"`       // max 5
	BillingIndicator    int                  `json:"billingIndicator"`              // required, 0-4
	Retention           *Retention           `json:"retention,omitempty"`
	ItemName            string               `json:"itemName"`                      // required, max 80
	GoodServiceIndicator int                 `json:"goodServiceIndicator"`          // required, 1=Bien, 2=Servicio
	ItemDescription     string               `json:"itemDescription,omitempty"`     // max 1000
	QuantityItem        float64              `json:"quantityItem"`                  // required
	UnitMeasure         *int                 `json:"unitMeasure,omitempty"`         // 1-62
	QuantityReference   *float64             `json:"quantityReference,omitempty"`
	ReferenceUnit       *int                 `json:"referenceUnit,omitempty"`
	SubquantityTable    []SubquantityEntry   `json:"subquantityTable,omitempty"`    // max 5
	DegreesAlcohol      *float64             `json:"degreesAlcohol,omitempty"`      // 0-99
	UnitPriceReference  *float64             `json:"unitPriceReference,omitempty"`
	ElaborationDate     string               `json:"elaborationDate,omitempty"`     // date
	ExpirationDateItem  string               `json:"expirationDateItem,omitempty"`  // date
	Mining              *Mining              `json:"mining,omitempty"`
	UnitPriceItem       float64              `json:"unitPriceItem"`                 // required
	DiscountAmount      *float64             `json:"discountAmount,omitempty"`
	SubDiscounts        []SubDiscount        `json:"subDiscounts,omitempty"`        // max 12
	SurchargeAmount     *float64             `json:"surchargeAmount,omitempty"`
	SubSurcharge        []SubSurcharge       `json:"subSurcharge,omitempty"`        // max 12
	AdditionalTaxes     []ItemAdditionalTax  `json:"additionalTaxes,omitempty"`     // max 2
	OtherCurrencyDetail *OtherCurrencyDetail `json:"otherCurrencyDetail,omitempty"`
	ItemAmount          float64              `json:"itemAmount"`                    // required
}

// --- Subtotals ---

// Subtotal represents a subtotal group. Campo DGII <Subtotales>
type Subtotal struct {
	SubTotalNumber           *int     `json:"subTotalNumber,omitempty"`
	SubtotalDescription      string   `json:"subtotalDescription,omitempty"`      // max 40
	Order                    *int     `json:"order,omitempty"`
	SubTotalAmountTaxedTotal *float64 `json:"subTotalAmountTaxedTotal,omitempty"`
	SubTotalAmountTaxedI1    *float64 `json:"subTotalAmountTaxedI1,omitempty"`
	SubTotalAmountTaxedI2    *float64 `json:"subTotalAmountTaxedI2,omitempty"`
	SubTotalAmountTaxedI3    *float64 `json:"subTotalAmountTaxedI3,omitempty"`
	ItbisSubTotal            *float64 `json:"itbisSubTotal,omitempty"`
	Itbis1SubTotal           *float64 `json:"itbis1SubTotal,omitempty"`
	Itbis2SubTotal           *float64 `json:"itbis2SubTotal,omitempty"`
	Itbis3SubTotal           *float64 `json:"itbis3SubTotal,omitempty"`
	SubTotalAdditionalTax    *float64 `json:"subTotalAdditionalTax,omitempty"`
	SubTotalExempt           *float64 `json:"subTotalExempt,omitempty"`
	SubTotalAmount           *float64 `json:"subTotalAmount,omitempty"`
	Lines                    *int     `json:"lines,omitempty"`
}

// --- Discounts or Surcharges ---

// DiscountOrSurcharge represents a global discount or surcharge. Campo DGII <DescuentosORecargos>
type DiscountOrSurcharge struct {
	LineNumber                                 *int     `json:"lineNumber,omitempty"`
	FitType                                    string   `json:"fitType,omitempty"`                                    // "D" or "R"
	Norma1007Indicator                         *int     `json:"norma1007Indicator,omitempty"`
	DescriptionDiscountOrSurcharge             string   `json:"descriptionDiscountOrSurcharge,omitempty"`             // max 45
	TypeValue                                  string   `json:"typeValue,omitempty"`                                  // "%" or "$"
	DiscountValueOrSurcharge                   *float64 `json:"discountValueOrSurcharge,omitempty"`                   // 0-999
	DiscountAmountOrSurcharge                  *float64 `json:"discountAmountOrSurcharge,omitempty"`
	DiscountAmountOrSurchargeOtherCurrency     *float64 `json:"discountAmountOrSurchargeOtherCurrency,omitempty"`
	IndicatorBillingDiscountOrSurcharge        *int     `json:"indicatorBillingDiscountOrSurcharge,omitempty"`        // 0-4
}

// --- Pagination ---

// SubtotalAdditionalTaxPage represents page-level additional tax subtotals.
type SubtotalAdditionalTaxPage struct {
	SubtotalSelectiveTaxForSpecificConsumptionPage *float64 `json:"subtotalSelectiveTaxForSpecificConsumptionPage,omitempty"`
	SubtotalOtherTax                               *float64 `json:"subtotalOtherTax,omitempty"`
}

// Pagination represents page layout information. Campo DGII <Paginacion>
type Pagination struct {
	PageNo                          int                        `json:"pageNo"`
	NoLineFrom                      int                        `json:"noLineFrom"`
	NoLineUntil                     int                        `json:"noLineUntil"`
	SubtotalAmountTaxedPage         *float64                   `json:"subtotalAmountTaxedPage,omitempty"`
	SubtotalAmountTaxed1Page        *float64                   `json:"subtotalAmountTaxed1Page,omitempty"`
	SubtotalAmountTaxed2Page        *float64                   `json:"subtotalAmountTaxed2Page,omitempty"`
	SubtotalAmountTaxed3Page        *float64                   `json:"subtotalAmountTaxed3Page,omitempty"`
	ExemptSubtotalPage              *float64                   `json:"exemptSubtotalPage,omitempty"`
	ItbisSubtotalPage               *float64                   `json:"itbisSubtotalPage,omitempty"`
	Itbis1SubtotalPage              *float64                   `json:"itbis1SubtotalPage,omitempty"`
	Itbis2SubtotalPage              *float64                   `json:"itbis2SubtotalPage,omitempty"`
	Itbis3SubtotalPage              *float64                   `json:"itbis3SubtotalPage,omitempty"`
	SubtotalAdditionalTaxPage       *float64                   `json:"subtotalAdditionalTaxPage,omitempty"`
	SubtotalAdditionalTax           *SubtotalAdditionalTaxPage `json:"subtotalAdditionalTax,omitempty"`
	SubtotalAmountPage              *float64                   `json:"subtotalAmountPage,omitempty"`
	SubtotalNonBillableAmountPage   *float64                   `json:"subtotalNonBillableAmountPage,omitempty"`
}

// --- Information Reference ---

// InformationReference represents reference to a modified document. Campo DGII <InformacionReferencia>
type InformationReference struct {
	NCFModified            string `json:"ncfModified,omitempty"`            // max 19
	RNCOtherTaxpayer       string `json:"rncOtherTaxpayer,omitempty"`       // max 11
	NCFModifiedDate        string `json:"ncfModifiedDate,omitempty"`        // date
	ModificationCode       *int   `json:"modificationCode,omitempty"`       // 1-5
	ReasonForModification  string `json:"reasonForModification,omitempty"`  // max 90, for debit/credit notes
}

// --- Config ---

// PDFConfig configures PDF generation.
type PDFConfig struct {
	Type string `json:"type,omitempty"` // "generic"
	Note string `json:"note,omitempty"` // max 250
}

// DocumentConfig configures additional document options.
type DocumentConfig struct {
	PDF *PDFConfig `json:"pdf,omitempty"`
}

// --- Fiscal Invoice (31) Request ---

// CreateFiscalInvoiceRequest is the request body for creating a fiscal credit invoice (eCF type 31).
// POST /fiscal-invoices
type CreateFiscalInvoiceRequest struct {
	Company              *DocumentCompany          `json:"company,omitempty"`
	IdDoc                *IdDoc                    `json:"idDoc"`
	Sender               *Sender                   `json:"sender"`
	Buyer                *Buyer                    `json:"buyer"`
	AdditionalInfo       *AdditionalInformationDoc `json:"additionalInformation,omitempty"`
	Totals               *Totals                   `json:"totals"`
	ItemDetails          []ItemDetail              `json:"itemDetails"`                     // required, 1-1000
	Subtotals            []Subtotal                `json:"subtotals,omitempty"`             // max 20
	DiscountsOrSurcharges []DiscountOrSurcharge    `json:"discountsOrSurcharges,omitempty"` // max 20
	Pagination           []Pagination              `json:"pagination,omitempty"`            // 2-100
	InformationReference *InformationReference     `json:"informationReference,omitempty"`
	Config               *DocumentConfig           `json:"config,omitempty"`
}

// --- Document Response (shared across all document types) ---

// ResponseMessage represents an API response message (e.g., async processing notifications).
type ResponseMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// DocumentResponse represents the response when creating/emitting a document.
type DocumentResponse struct {
	ID                    string              `json:"id"`
	StampDate             string              `json:"stampDate,omitempty"`
	Status                string              `json:"status"`                          // REGISTERED, TO_SEND, WAITING_RESPONSE, TO_NOTIFY, FINISHED
	LegalStatus           string              `json:"legalStatus,omitempty"`           // ACCEPTED, ACCEPTED_WITH_OBSERVATIONS, REJECTED
	CompanyIdentification string              `json:"companyIdentification,omitempty"`
	ENCF                  string              `json:"encf,omitempty"`
	DocumentNumber        string              `json:"documentNumber,omitempty"`
	XML                   string              `json:"xml,omitempty"`
	ResumeXML             string              `json:"resumeXml,omitempty"`
	PDF                   string              `json:"pdf,omitempty"`
	DocumentStampURL      string              `json:"documentStampUrl,omitempty"`
	SignatureDate         string              `json:"signatureDate,omitempty"`
	SecurityCode          string              `json:"securityCode,omitempty"`
	SequenceConsumed      bool                `json:"sequenceConsumed,omitempty"`
	GovernmentResponse    *GovernmentResponse `json:"governmentResponse,omitempty"`
	Response              []ResponseMessage   `json:"response,omitempty"`
}

// --- Other invoice type requests (reuse shared types) ---

// CreateConsumptionInvoiceRequest is the request body for creating a consumption invoice (eCF type 32).
// POST /consumption-invoices
type CreateConsumptionInvoiceRequest struct {
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

// CreateSpecialRegimesInvoiceRequest is the request body for creating a special regimes invoice (eCF type 44).
// POST /special-regime-invoices
type CreateSpecialRegimesInvoiceRequest struct {
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

// CreateGovernmentalInvoiceRequest is the request body for creating a governmental invoice (eCF type 45).
// POST /governmental-invoices
type CreateGovernmentalInvoiceRequest struct {
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

// CreateExportInvoiceRequest is the request body for creating an export invoice (eCF type 46).
// POST /export-invoices
type CreateExportInvoiceRequest struct {
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

// --- Query Options ---

// DocumentQueryOptions represents optional query parameters for document GET endpoints.
type DocumentQueryOptions struct {
	PDFType string // "generic" (default) or other allowed types
}

// --- Notify ---

// NotifyByEmailRequest is the request body for notifying a document by email.
type NotifyByEmailRequest struct {
	ID        string `json:"id"`
	IDCompany string `json:"idCompany,omitempty"`
	Mail      string `json:"mail,omitempty"`
	PDFType   string `json:"pdfType,omitempty"` // "generic" (default)
}

// NotifyByEmailResponse is the response from the notify-by-email endpoint.
type NotifyByEmailResponse struct {
	Message string `json:"message"`
}

// --- Document Status (used by GET endpoints) ---

// GovernmentResponseDetail represents a government validation error or warning.
type GovernmentResponseDetail struct {
	Value string     `json:"valor"`
	Code  FlexString `json:"codigo"`
}

// GovernmentResponse represents the government validation response.
type GovernmentResponse struct {
	Value []GovernmentResponseDetail `json:"value,omitempty"`
	Code  int                        `json:"code,omitempty"`
}

// DocumentError represents a processing error for a document.
type DocumentError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// DocumentStatus represents the full status of an emitted document (from webhook or GET).
type DocumentStatus struct {
	ID                    string              `json:"id"`
	Type                  string              `json:"type"`
	Status                string              `json:"status"`
	LegalStatus           string              `json:"legalStatus,omitempty"`
	CompanyIdentification string              `json:"companyIdentification,omitempty"`
	TrackID               string              `json:"trackId,omitempty"`
	DocumentNumber        string              `json:"documentNumber,omitempty"`
	SequenceConsumed      bool                `json:"sequenceConsumed,omitempty"`
	StampDate             string              `json:"stampDate,omitempty"`
	SignatureDate         string              `json:"signatureDate,omitempty"`
	SecurityCode          string              `json:"securityCode,omitempty"`
	DocumentStampURL      string              `json:"documentStampUrl,omitempty"`
	XML                   string              `json:"xml,omitempty"`
	ResumeXML             string              `json:"resumeXml,omitempty"`
	PDF                   string              `json:"pdf,omitempty"`
	GovernmentResponse    *GovernmentResponse `json:"governmentResponse,omitempty"`
	Error                 *DocumentError      `json:"error,omitempty"`
}
