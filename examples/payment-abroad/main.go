// Example: Payment abroad support (47) with ISR retention.
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

	doc, err := client.CreatePaymentAbroad(ctx, &alanube.CreatePaymentAbroadRequest{
		IdDoc: &alanube.IdDoc{
			ENCF:            "E470000000000",
			SequenceDueDate: "2025-12-31",
			PaymentType:     2,
			PaymentTerm:     "30 días",
			PaymentDeadline: "2024-12-31",
			PaymentFormsTable: []alanube.PaymentForm{
				{PaymentMethod: intPtr(1), PaymentAmount: floatPtr(1000)},
				{PaymentMethod: intPtr(2), PaymentAmount: floatPtr(1000)},
			},
			DateFrom:  "2023-08-01",
			DateUntil: "2023-08-31",
		},
		Sender: &alanube.Sender{
			RNC:                   "132109122",
			CompanyName:           "Razón Social del Emisor",
			Tradename:             "Nombre Comercial del Emisor",
			BranchOffice:          "Punto de Fabrica",
			Address:               "Av. México No. 48, Gazcue, Distrito Nacional, Santo Domingo",
			Province:              string(alanube.ProvinceDistritoNacional),
			Municipality:          "010101",
			PhoneNumber:           []string{"809-499-6894", "809-564-6894"},
			Mail:                  "contacto@emisordele-ncf.com",
			WebSite:               "www.emisordele-ncf.com",
			EconomicActivity:      "Descripción de la actividad que desarrolla la empresa",
			InternalInvoiceNumber: "1234",
			InternalOrderNumber:   "123456",
			StampDate:             "2023-08-01",
		},
		Buyer: &alanube.Buyer{
			ForeignIdentifier: "900900900",
			CompanyName:       "Empresa Colombiana de Comercio S.A.S",
		},
		AdditionalInfo: &alanube.AdditionalInformationDoc{
			Transport: &alanube.Transport{
				DestinationCountry: "Colombia",
			},
		},
		Totals: &alanube.Totals{
			ExemptAmount:         floatPtr(100),
			TotalAmount:          100,
			AmountPeriod:         floatPtr(100),
			PreviousBalance:      floatPtr(100),
			AmountAdvancePayment: floatPtr(50),
			PayValue:             floatPtr(150),
			IsrTotalRetention:    floatPtr(27),
		},
		ItemDetails: []alanube.ItemDetail{
			{
				LineNumber:           1,
				BillingIndicator:     4, // Exempt
				GoodServiceIndicator: 2, // Service
				ItemCodeTable: []alanube.ItemCodeEntry{
					{CodeType: "EAN", ItemCode: "1224545-4545"},
					{CodeType: "PLU", ItemCode: "799889457844884"},
				},
				ItemName:        "Nombre del producto o servicio.",
				ItemDescription: "Descripción adicional del ítem.",
				QuantityItem:    1,
				UnitMeasure:     intPtr(int(alanube.UnitUnidad)),
				UnitPriceItem:   100,
				ItemAmount:      100,
				Retention: &alanube.Retention{
					IndicatorAgentWithholdingPerception: intPtr(1),
					IsrAmountWithheld:                   floatPtr(27),
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("create payment abroad: %v", err)
	}

	fmt.Printf("Payment Abroad created!\n")
	fmt.Printf("  ID: %s\n", doc.ID)
	fmt.Printf("  ENCF: %s\n", doc.ENCF)
	fmt.Printf("  Status: %s\n", doc.Status)
}

func intPtr(v int) *int           { return &v }
func floatPtr(v float64) *float64 { return &v }
