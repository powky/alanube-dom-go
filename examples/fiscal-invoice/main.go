// Example: Fiscal credit invoice (31) with foreign currency USD.
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

	doc, err := client.CreateFiscalInvoice(ctx, &alanube.CreateFiscalInvoiceRequest{
		IdDoc: &alanube.IdDoc{
			ENCF:               "E310000000000",
			SequenceDueDate:    "2025-12-31",
			TaxAmountIndicator: intPtr(0),
			IncomeType:         1,
			PaymentType:        1,
		},
		Sender: &alanube.Sender{
			RNC:                         "132109122",
			CompanyName:                 "Mi Razón Social SRL",
			Tradename:                   "Mi Nombre Comercial",
			BranchOffice:                "Nombre Sucursal",
			Address:                     "C Luis Lembert Esq. Dr. Heriberto Pieter Plaza Hache. 10123",
			Municipality:                "010101",
			Province:                    string(alanube.ProvinceDistritoNacional),
			PhoneNumber:                 []string{"809-465-6799", "809-456-5646"},
			Mail:                        "correodelemisor@correo.com",
			WebSite:                     "www.sitiowebemisor.com",
			EconomicActivity:            "Descripción de la actividad comercial de mi empresa",
			SellerCode:                  "Carlos Segura ID458-457",
			InternalInvoiceNumber:       "456789",
			InternalOrderNumber:         "562344",
			SaleArea:                    "Santo Domingo Este",
			SaleRoute:                   "Ruta 1",
			AdditionalInformationIssuer: "Información adicional: Realizar pago en el Banco Popular",
			StampDate:                   "2023-09-05",
		},
		Buyer: &alanube.Buyer{
			RNC:         "132250011",
			CompanyName: "Adatum Corporation",
			Contact:     "Robert Townes",
			Mail:        "gerard.diaz@kcpdynamics.com",
			Address:     "Station Road, 21",
		},
		Totals: &alanube.Totals{
			TotalTaxedAmount: floatPtr(56.8),
			I1AmountTaxed:    floatPtr(56.8),
			ItbisS1:          intPtr(18),
			ItbisTotal:       floatPtr(10.22),
			Itbis1Total:      floatPtr(10.22),
			TotalAmount:      67.02,
			OtherCurrency: &alanube.OtherCurrency{
				CurrencyType:                  string(alanube.CurrencyUSD),
				ExchangeRate:                  56.8,
				TotalTaxedAmountOtherCurrency: floatPtr(1),
				AmountTaxed1OtherCurrency:     floatPtr(1),
				ItbisTotalOtherCurrency:       floatPtr(0.18),
				Itbis1TotalOtherCurrency:      floatPtr(0.18),
				TotalAmountOtherCurrency:      1.18,
			},
		},
		ItemDetails: []alanube.ItemDetail{
			{
				LineNumber:           1,
				BillingIndicator:     1,
				GoodServiceIndicator: 1,
				ItemCodeTable: []alanube.ItemCodeEntry{
					{CodeType: "Interna", ItemCode: "001"},
				},
				ItemName:     "ALMOHADILLA PARA ROLL-ON",
				QuantityItem: 1,
				UnitPriceItem: 56.8,
				ItemAmount:    56.8,
				OtherCurrencyDetail: &alanube.OtherCurrencyDetail{
					PriceOtherCurrency:      1,
					DiscountOtherCurrency:   floatPtr(0),
					SurchargeAnotherCurrency: floatPtr(0),
					AmountItemOtherCurrency: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("create fiscal invoice: %v", err)
	}

	fmt.Printf("Fiscal Invoice created!\n")
	fmt.Printf("  ID: %s\n", doc.ID)
	fmt.Printf("  ENCF: %s\n", doc.ENCF)
	fmt.Printf("  Status: %s\n", doc.Status)
	fmt.Printf("  Stamp Date: %s\n", doc.StampDate)
}

func intPtr(v int) *int       { return &v }
func floatPtr(v float64) *float64 { return &v }
