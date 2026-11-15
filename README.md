# alanube-dom-go

Go SDK for the [Alanube API](https://developer.alanube.co) - Dominican Republic electronic invoicing (DGII).

## Installation

```bash
go get github.com/powky/alanube-dom-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    alanube "github.com/powky/alanube-dom-go"
)

func main() {
    // Use NewClient for production, NewSandboxClient for testing
    client := alanube.NewSandboxClient("your-jwt-token")
    ctx := context.Background()

    // Create a company
    company, err := client.CreateCompany(ctx, &alanube.CreateCompanyRequest{
        Name:           "Mi Empresa SRL",
        TradeName:      "Mi Empresa",
        Identification: "123456789",
        Address:        "Calle Principal #1",
        Province:       string(alanube.ProvinceSantoDomingo),
        Municipality:   "320101",
        Certificate: &alanube.Certificate{
            Name:      "cert",
            Extension: "p12",
            Content:   "<base64-encoded-certificate>",
            Password:  "secret",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Company: %s (ID: %s)\n", company.Name, company.ID)
}
```

## Configuration

```go
// Production client
client := alanube.NewClient("your-jwt-token")

// Sandbox client
client := alanube.NewSandboxClient("your-jwt-token")

// With options
client := alanube.NewClient("your-jwt-token",
    alanube.WithTimeout(60 * time.Second),
    alanube.WithHTTPClient(customHTTPClient),
    alanube.WithBaseURL("https://custom-url.com/dom/v1"),
)
```

## API Coverage

### Companies (Empresas)

```go
company, err := client.CreateCompany(ctx, &alanube.CreateCompanyRequest{...})
company, err := client.GetSelfCompany(ctx)
company, err := client.UpdateSelfCompany(ctx, &alanube.UpdateCompanyRequest{...})
company, err := client.GetCompany(ctx, "companyID")
company, err := client.UpdateCompany(ctx, "companyID", &alanube.UpdateCompanyRequest{...})
```

### Certification (Certificación)

```go
setTest, err := client.CreateSetTest(ctx, &alanube.CreateSetTestRequest{...})
check, err := client.GetSetTest(ctx, "setTestID")
check, err := client.GetSetTestByCompany(ctx, "setTestID", "companyID")
signed, err := client.SignDocument(ctx, xmlFileReader, "filename.xml")
signed, err := client.SignDocumentByCompany(ctx, "companyID", xmlFileReader, "filename.xml")
info, err := client.GetProviderInfo(ctx)
```

### Electronic Invoices (Facturas Electrónicas)

All invoice types follow the same pattern: Create, Get, GetByCompany, Notify.

```go
// Fiscal Credit Invoice (31)
doc, err := client.CreateFiscalInvoice(ctx, &alanube.CreateFiscalInvoiceRequest{...})
status, err := client.GetFiscalInvoice(ctx, "id", nil)
status, err := client.GetFiscalInvoiceByCompany(ctx, "id", "companyID", &alanube.DocumentQueryOptions{PDFType: "generic"})
resp, err := client.NotifyFiscalInvoice(ctx, &alanube.NotifyByEmailRequest{ID: "id", Mail: "email@example.com"})

// Consumption Invoice (32)
doc, err := client.CreateConsumptionInvoice(ctx, &alanube.CreateConsumptionInvoiceRequest{...})

// Special Regimes (44)
doc, err := client.CreateSpecialRegimesInvoice(ctx, &alanube.CreateSpecialRegimesInvoiceRequest{...})

// Governmental (45)
doc, err := client.CreateGovernmentalInvoice(ctx, &alanube.CreateGovernmentalInvoiceRequest{...})

// Export (46)
doc, err := client.CreateExportInvoice(ctx, &alanube.CreateExportInvoiceRequest{...})
```

### Electronic Notes (Notas Electrónicas)

```go
// Debit Note (33)
doc, err := client.CreateDebitNote(ctx, &alanube.CreateDebitNoteRequest{...})

// Credit Note (34)
doc, err := client.CreateCreditNote(ctx, &alanube.CreateCreditNoteRequest{...})
```

### Purchases & Expenses (Compras y Gastos)

```go
// Purchase (41)
doc, err := client.CreatePurchase(ctx, &alanube.CreatePurchaseRequest{...})

// Minor Expense (43)
doc, err := client.CreateMinorExpense(ctx, &alanube.CreateMinorExpenseRequest{...})

// Payment Abroad (47)
doc, err := client.CreatePaymentAbroad(ctx, &alanube.CreatePaymentAbroadRequest{...})
```

### Cancellations (Anulaciones)

```go
resp, err := client.CreateCancellation(ctx, &alanube.CreateCancellationRequest{...})
status, err := client.GetCancellation(ctx, "id")
docStatus, err := client.GetCancellationByCompany(ctx, "id", "companyID")
```

### Received Documents (Documentos Recibidos)

```go
// List received documents (paginated)
docs, err := client.GetReceivedDocuments(ctx, &alanube.ReceivedDocumentsQueryOptions{
    Status: "RECEIVED",
    Limit:  25,
    Page:   1,
    Start:  "2024-01-01",
    End:    "2024-12-31",
})

docs, err := client.GetReceivedDocumentsByCompany(ctx, "companyID", opts)
doc, err := client.GetReceivedDocument(ctx, "id")
doc, err := client.GetReceivedDocumentByCompany(ctx, "id", "companyID")

// Simulate document reception (sandbox only)
err := client.ReceiveDocumentFromDGII(ctx, &alanube.ReceiveDocumentRequest{
    IDCompany:    "companyID",
    DocumentType: "31",
})
```

### Commercial Approvals (Aprobaciones Comerciales)

```go
// Generate commercial response
result, err := client.CreateCommercialResponse(ctx, "docID", &alanube.CommercialResponseRequest{
    CommercialResponse: "ACCEPTED",
})
result, err := client.CreateCommercialResponseByCompany(ctx, "docID", "companyID", &alanube.CommercialResponseRequest{
    CommercialResponse: "REJECTED",
    NotAcceptedDetail:  "Reason for rejection",
})

// Query received commercial approvals (paginated)
approvals, err := client.GetReceivedCommercialApprovals(ctx, &alanube.CommercialApprovalsQueryOptions{Limit: 25, Page: 1})
approval, err := client.GetReceivedCommercialApproval(ctx, "id")

// Query generated commercial approvals
approval, err := client.GetCommercialApproval(ctx, "id")
approval, err := client.GetCommercialApprovalByCompany(ctx, "id", "companyID")
```

### Acknowledgments (Acuses de Recibo)

```go
// External acknowledgments (as issuer)
ack, err := client.GetExternalAcknowledgment(ctx, "companyID", "id")
acks, err := client.GetExternalAcknowledgments(ctx, "companyID", &alanube.AcknowledgmentsQueryOptions{
    Status: "RECEIVED",
    Limit:  25,
})

// Internal receipts (as receiver)
receipt, err := client.GetInternalReceipt(ctx, "companyID", "id")
receipts, err := client.GetInternalReceipts(ctx, "companyID", opts)
```

### Directory (Directorio)

```go
entries, err := client.GetDirectory(ctx, &alanube.DirectoryQueryOptions{RNC: "123456789"})
entries, err := client.GetDirectoryByCompany(ctx, "companyID", nil)
```

### DGII Status

```go
// Service status
status, err := client.GetDGIIServiceStatus(ctx)
status, err := client.GetDGIIServiceStatusByCompany(ctx, "companyID")

// Environment status
envStatus, err := client.GetDGIIEnvironmentStatus(ctx, alanube.DGIIProduction)

// Maintenance windows
maint, err := client.GetDGIIMaintenance(ctx)
```

### Reports (Reportes)

```go
// User totals (all companies)
totals, err := client.GetUserDocumentTotals(ctx, &alanube.ReportQueryOptions{
    LegalStatus: "ACCEPTED,ACCEPTED_WITH_OBSERVATIONS",
    DateFrom:    "2024-01-01",
    DateUntil:   "2024-03-31",
})

// Company totals (filtered)
totals, err := client.GetCompanyDocumentTotals(ctx, "companyID", opts)

// Historical totals (all-time)
emitted, err := client.GetCompanyEmittedDocuments(ctx, "companyID")
accepted, err := client.GetCompanyAcceptedDocuments(ctx, "companyID")
```

### Webhooks

```go
http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
    event, err := alanube.ParseWebhook(r)
    if err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    switch {
    case event.GovernmentStatusChanged != nil:
        fmt.Printf("Company %s certified: %s\n",
            event.GovernmentStatusChanged.Company.ID,
            event.GovernmentStatusChanged.Company.HasBeenCertified)

    case event.EmissionFinished != nil:
        fmt.Printf("Document %s: %s (%s)\n",
            event.EmissionFinished.DocumentNumber,
            event.EmissionFinished.Status,
            event.EmissionFinished.LegalStatus)

    case event.DocumentReception != nil:
        fmt.Printf("Received %s: %s\n",
            event.DocumentReception.Type,
            event.DocumentReception.DocumentNumber)
    }

    w.WriteHeader(200)
})
```

## Error Handling

```go
doc, err := client.CreateFiscalInvoice(ctx, req)
if err != nil {
    // Check for specific error codes
    if alanube.IsErrorCode(err, alanube.ErrENCFUsed) {
        log.Println("ENCF already used, try a different sequence")
    } else if alanube.IsErrorCode(err, alanube.ErrDGIITimeoutError) {
        log.Println("DGII timeout, document will be processed async")
    }

    // Access full error details
    if apiErr, ok := err.(*alanube.Error); ok {
        fmt.Printf("Code: %s, Messages: %v\n", apiErr.Code, apiErr.Messages)
    }

    // The request succeeded but the response had an unexpected shape:
    // the document WAS created — recover it from the raw body, don't retry.
    var decErr *alanube.DecodeError
    if errors.As(err, &decErr) {
        log.Printf("unexpected response: %s", decErr.Body)
    }
}
```

DGII response codes (`governmentResponse.value[].codigo`) are typed as `FlexString`
because the API returns them as numbers even though the docs say string.

## Catalogs (Constants)

The SDK includes typed constants for all DGII catalogs:

```go
// Document types
alanube.DocumentTypeFiscalInvoice  // "fiscalInvoice"
alanube.DocumentTypeCreditNote     // "creditNote"

// Statuses
alanube.StatusFinished             // "FINISHED"
alanube.LegalStatusAccepted        // "ACCEPTED"

// Tax types (39 types)
alanube.TaxTypePropinaLegal        // 1 - Propina Legal (10%)
alanube.TaxTypeISCEspecificoCerveza // 6 - ISC Específico Cerveza

// Unit measures (62 units)
alanube.UnitUnidad                 // 43 - UND
alanube.UnitKilogramo              // 21 - KG

// Currencies (16 codes)
alanube.CurrencyUSD                // "USD"
alanube.CurrencyEUR                // "EUR"

// DGII Environments
alanube.DGIIProduction             // "2"
alanube.DGIICertification          // "3"

// Provinces & Municipalities (32 provinces, 155 municipalities)
alanube.ProvinceDistritoNacional   // "010000"
alanube.ProvinceSantoDomingo       // "320000"
alanube.MunicipalitySantoDomingoEste // "320100"

// Error codes (~170 codes)
alanube.ErrENCFUsed                // "AP3001"
alanube.ErrDGIITimeoutError        // "AEP2006"
```

## Examples

See the [examples](./examples) directory:

- [basic](./examples/basic) - Company creation and basic invoice
- [fiscal-invoice](./examples/fiscal-invoice) - Fiscal credit invoice (31) with USD
- [debit-note](./examples/debit-note) - Debit note (33) with document reference
- [purchase](./examples/purchase) - Purchase (41) with ITBIS and ISR retention
- [minor-expense](./examples/minor-expense) - Minor expense (43) with exempt items
- [payment-abroad](./examples/payment-abroad) - Payment abroad (47) with ISR retention

## API Reference

- [Alanube Developer Portal](https://developer.alanube.co)
- [API Reference](https://developer.alanube.co/reference)
- Base URLs:
  - Sandbox: `https://sandbox.alanube.co/dom/v1`
  - Production: `https://api.alanube.co/dom/v1`

## License

MIT
