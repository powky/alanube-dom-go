package main

import (
	"context"
	"fmt"
	"log"
	"os"

	alanube "github.com/powky/alanube-dom-go"
)

func main() {
	token := os.Getenv("ALANUBE_TOKEN")
	if token == "" {
		log.Fatal("ALANUBE_TOKEN environment variable is required")
	}

	client := alanube.NewSandboxClient(token)
	ctx := context.Background()

	// Create a company.
	company, err := client.CreateCompany(ctx, &alanube.CreateCompanyRequest{
		Name:           "Mi Empresa SRL",
		TradeName:      "Mi Empresa",
		Identification: "123456789",
		Address:        "Calle Principal #1",
		Province:       "Santo Domingo",
		Municipality:   "Distrito Nacional",
		Certificate: &alanube.Certificate{
			Name:      "cert",
			Extension: "p12",
			Content:   "<base64-encoded-certificate>",
			Password:  "secret",
		},
	})
	if err != nil {
		log.Fatalf("create company: %v", err)
	}
	fmt.Printf("Company created: %s (ID: %s)\n", company.Name, company.ID)

	// Create a fiscal credit invoice (type 31).
	doc, err := client.CreateFiscalInvoice(ctx, &alanube.CreateFiscalInvoiceRequest{
		Company: &alanube.DocumentCompany{ID: company.ID},
		IdDoc: &alanube.IdDoc{
			ENCF:            "E310000000001",
			SequenceDueDate: "2026-12-31",
			IncomeType:      1,
			PaymentType:     1,
		},
		Sender: &alanube.Sender{
			RNC:         "123456789",
			CompanyName: "Mi Empresa SRL",
			Address:     "Calle Principal #1",
			StampDate:   "2026-03-23",
		},
		Buyer: &alanube.Buyer{
			RNC:         "987654321",
			CompanyName: "Cliente Ejemplo SRL",
		},
		Totals: &alanube.Totals{
			TotalAmount: 5900,
		},
		ItemDetails: []alanube.ItemDetail{
			{
				LineNumber:           1,
				BillingIndicator:     1,
				GoodServiceIndicator: 2,
				ItemName:             "Servicio de consultoría",
				QuantityItem:         1,
				UnitPriceItem:        5000,
				ItemAmount:           5000,
			},
		},
	})
	if err != nil {
		log.Fatalf("create fiscal invoice: %v", err)
	}
	fmt.Printf("Invoice created: %s (Status: %s, ENCF: %s)\n", doc.ID, doc.Status, doc.ENCF)

	// Check DGII service status.
	dgiiStatus, err := client.GetDGIIServiceStatus(ctx)
	if err != nil {
		log.Fatalf("check dgii status: %v", err)
	}
	for _, group := range dgiiStatus.Response {
		for _, s := range group {
			fmt.Printf("DGII Service: %s - %s\n", s.Service, s.Status)
		}
	}
}
