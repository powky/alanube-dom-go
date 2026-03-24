// Example: Minor expense (43) with exempt items.
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

	doc, err := client.CreateMinorExpense(ctx, &alanube.CreateMinorExpenseRequest{
		IdDoc: &alanube.IdDoc{
			ENCF:            "E430000000000",
			SequenceDueDate: "2025-12-31",
			PaymentType:     1,
		},
		Sender: &alanube.Sender{
			RNC:                         "132109122",
			CompanyName:                 "Mi Razón Social SRL",
			Tradename:                   "Mi Nombre Comercial",
			BranchOffice:                "Sucursal Name",
			Address:                     "C Luis Lembert Esq. Dr. Heriberto Pieter Plaza Hache. 10123",
			Municipality:                "010101",
			Province:                    string(alanube.ProvinceDistritoNacional),
			PhoneNumber:                 []string{"809-465-6799", "809-456-5646", "809-431-4489"},
			Mail:                        "micorreo@correoemisor.com",
			WebSite:                     "www.empresaEmisor.com",
			EconomicActivity:            "Descripción de la actividad comercial de mi empresa",
			InternalInvoiceNumber:       "4859",
			InternalOrderNumber:         "4546976",
			AdditionalInformationIssuer: "Términos y condiciones establecidos en este documento.",
			StampDate:                   "2023-09-04",
		},
		Totals: &alanube.Totals{
			ExemptAmount:         floatPtr(1000),
			TotalAmount:          1000,
			AmountPeriod:         floatPtr(1000),
			PreviousBalance:      floatPtr(100),
			AmountAdvancePayment: floatPtr(200),
			PayValue:             floatPtr(900),
		},
		ItemDetails: []alanube.ItemDetail{
			{
				LineNumber:           1,
				BillingIndicator:     4, // Exempt
				GoodServiceIndicator: 1,
				ItemName:             "Cambio periodico de gomas",
				ItemDescription:      "Se realiza el cambio de la 4 gomas vehiculo principal de entregas.",
				QuantityItem:         1,
				UnitMeasure:          intPtr(int(alanube.UnitUnidad)),
				UnitPriceItem:        1000,
				ItemAmount:           1000,
			},
		},
	})
	if err != nil {
		log.Fatalf("create minor expense: %v", err)
	}

	fmt.Printf("Minor Expense created!\n")
	fmt.Printf("  ID: %s\n", doc.ID)
	fmt.Printf("  ENCF: %s\n", doc.ENCF)
	fmt.Printf("  Status: %s\n", doc.Status)
}

func intPtr(v int) *int           { return &v }
func floatPtr(v float64) *float64 { return &v }
