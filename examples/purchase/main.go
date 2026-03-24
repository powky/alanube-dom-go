// Example: Purchase (41) with retention (ITBIS and ISR).
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

	doc, err := client.CreatePurchase(ctx, &alanube.CreatePurchaseRequest{
		IdDoc: &alanube.IdDoc{
			ENCF:            "E410000000000",
			SequenceDueDate: "2025-12-31",
			TaxAmountIndicator: intPtr(0),
			PaymentType:     2,
			PaymentDeadline: "2023-09-30",
			PaymentTerm:     "1  mes",
			PaymentFormsTable: []alanube.PaymentForm{
				{PaymentMethod: intPtr(1), PaymentAmount: floatPtr(50)},
				{PaymentMethod: intPtr(4), PaymentAmount: floatPtr(150)},
			},
		},
		Sender: &alanube.Sender{
			RNC:                         "132109122",
			CompanyName:                 "Mi Razón Social SRL",
			Tradename:                   "Mi Nombre Comercial",
			BranchOffice:                "Santo Domingo Este",
			Address:                     "C Luis Lembert Esq. Dr. Heriberto Pieter Plaza Hache. 10123",
			Municipality:                "010101",
			Province:                    string(alanube.ProvinceDistritoNacional),
			PhoneNumber:                 []string{"809-465-6799", "809-456-5646", "809-431-4489"},
			Mail:                        "correodecontacto@correo.com",
			WebSite:                     "www.empresadecontacto.com",
			EconomicActivity:            "Descripción de la actividad comercial de mi empresa",
			InternalInvoiceNumber:       "4548",
			InternalOrderNumber:         "1234548648",
			AdditionalInformationIssuer: "Costos, plazos de pago y limitaciones de responsabilidad se detallan en la factura.",
			StampDate:                   "2023-08-11",
		},
		Buyer: &alanube.Buyer{
			RNC:                   "00305985365",
			CompanyName:           "Proveedor Informal de Prueba",
			Contact:               "Carlos Gomez Estrada, Teléfono 809-456-87-98",
			Mail:                  "contacto@proveedor.com",
			Address:               "Av. Núñez de Cáceres 593, Santo Domingo 10133",
			Municipality:          "010101",
			Province:              string(alanube.ProvinceDistritoNacional),
			InternalCode:          "PI-545899",
			ResponsibleForPayment: "132109122",
			AdditionalInformation: "Información importante del proveedor.",
		},
		Totals: &alanube.Totals{
			TotalTaxedAmount:     floatPtr(100),
			I1AmountTaxed:        floatPtr(100),
			ItbisS1:              intPtr(18),
			ItbisTotal:           floatPtr(18),
			Itbis1Total:          floatPtr(18),
			TotalAmount:          118,
			AmountPeriod:         floatPtr(118),
			PreviousBalance:      floatPtr(100),
			AmountAdvancePayment: floatPtr(200),
			PayValue:             floatPtr(18),
			ItbisTotalRetained:   floatPtr(18),
			IsrTotalRetention:    floatPtr(2),
		},
		ItemDetails: []alanube.ItemDetail{
			{
				LineNumber:           1,
				BillingIndicator:     1,
				GoodServiceIndicator: 2,
				ItemName:             "Servicio técnico con impuesto ITBIS 18%",
				ItemDescription:      "Valores con ITBIS 18%",
				QuantityItem:         1,
				UnitPriceItem:        100,
				ItemAmount:           100,
				Retention: &alanube.Retention{
					IndicatorAgentWithholdingPerception: intPtr(1),
					ItbisAmountWithheld:                 floatPtr(18),
					IsrAmountWithheld:                   floatPtr(2),
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("create purchase: %v", err)
	}

	fmt.Printf("Purchase created!\n")
	fmt.Printf("  ID: %s\n", doc.ID)
	fmt.Printf("  ENCF: %s\n", doc.ENCF)
	fmt.Printf("  Status: %s\n", doc.Status)
}

func intPtr(v int) *int           { return &v }
func floatPtr(v float64) *float64 { return &v }
