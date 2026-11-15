package alanube

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return NewClient("test-token", WithBaseURL(srv.URL))
}

func TestFlexStringDecodesStringNumberAndNull(t *testing.T) {
	cases := map[string]string{
		`"1934"`: "1934",
		`1934`:   "1934",
		`null`:   "",
		`"abc"`:  "abc",
	}
	for in, want := range cases {
		var f FlexString
		if err := json.Unmarshal([]byte(in), &f); err != nil {
			t.Fatalf("%s: unexpected error: %v", in, err)
		}
		if f.String() != want {
			t.Fatalf("%s: got %q, want %q", in, f, want)
		}
	}

	var f FlexString
	if err := json.Unmarshal([]byte(`{"x":1}`), &f); err == nil {
		t.Fatal("expected error for object input")
	}
}

// The Alanube docs declare governmentResponse.value[].codigo as a string but
// their own example (and DGII) return it as a number. A rejected document must
// still decode, otherwise the caller loses the document ID.
func TestCreateFiscalInvoiceDecodesNumericGovernmentCode(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/fiscal-invoices" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Errorf("authorization header = %q", got)
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{
			"id": "01G021Z3QSRZ58GSTBH7TPGD2J",
			"status": "FINISHED",
			"legalStatus": "REJECTED",
			"documentNumber": "E310000000001",
			"sequenceConsumed": true,
			"governmentResponse": {
				"value": [{"valor": "El campo MontoGravadoI1 no coincide", "codigo": 1934}],
				"code": 2
			}
		}`)
	})

	doc, err := c.CreateFiscalInvoice(context.Background(), &CreateFiscalInvoiceRequest{})
	if err != nil {
		t.Fatalf("CreateFiscalInvoice: %v", err)
	}
	if doc.ID != "01G021Z3QSRZ58GSTBH7TPGD2J" || doc.LegalStatus != string(LegalStatusRejected) {
		t.Fatalf("unexpected document: %+v", doc)
	}
	if got := doc.GovernmentResponse.Value[0].Code; got != "1934" {
		t.Fatalf("codigo = %q, want 1934", got)
	}
	if doc.GovernmentResponse.Code != 2 {
		t.Fatalf("code = %d, want 2", doc.GovernmentResponse.Code)
	}
}

func TestCreateConsumptionInvoiceSendsIdDocPaymentForms(t *testing.T) {
	var body map[string]json.RawMessage
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/invoices" {
			t.Errorf("path = %s, want /invoices", r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"id":"x","status":"FINISHED","legalStatus":"ACCEPTED"}`)
	})

	method := 1
	_, err := c.CreateConsumptionInvoice(context.Background(), &CreateConsumptionInvoiceRequest{
		IdDoc: &IdDoc{
			ENCF:              "E320000000001",
			IncomeType:        1,
			PaymentType:       1,
			PaymentFormsTable: []PaymentForm{{PaymentMethod: &method}},
		},
		Sender:      &Sender{RNC: "131234567", CompanyName: "X", Address: "Y", StampDate: "2026-01-01"},
		Totals:      &Totals{TotalAmount: 118},
		ItemDetails: []ItemDetail{{LineNumber: 1, BillingIndicator: 1, ItemName: "Clase", GoodServiceIndicator: 2, QuantityItem: 1, UnitPriceItem: 100, ItemAmount: 100}},
	})
	if err != nil {
		t.Fatalf("CreateConsumptionInvoice: %v", err)
	}
	if _, ok := body["paymentFormsTable"]; ok {
		t.Fatal("paymentFormsTable must be nested in idDoc, not top-level")
	}
	var idDoc map[string]json.RawMessage
	if err := json.Unmarshal(body["idDoc"], &idDoc); err != nil {
		t.Fatalf("idDoc: %v", err)
	}
	if _, ok := idDoc["paymentFormsTable"]; !ok {
		t.Fatal("idDoc.paymentFormsTable missing")
	}
	if _, ok := idDoc["sequenceDueDate"]; ok {
		t.Fatal("empty sequenceDueDate must be omitted (E32 does not accept it)")
	}
	if _, ok := body["buyer"]; ok {
		t.Fatal("nil buyer must be omitted for E32")
	}
}

func TestDecodeErrorKeepsBody(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"id":"keep-me","status":{"unexpected":"shape"}}`)
	})

	_, err := c.CreateConsumptionInvoice(context.Background(), &CreateConsumptionInvoiceRequest{})
	var de *DecodeError
	if !errors.As(err, &de) {
		t.Fatalf("expected *DecodeError, got %T: %v", err, err)
	}
	if de.StatusCode != http.StatusCreated || !bytes.Contains(de.Body, []byte("keep-me")) {
		t.Fatalf("decode error lost the response: %+v", de)
	}
}

func TestAPIErrorFormats(t *testing.T) {
	cases := []struct {
		name     string
		status   int
		body     string
		wantCode string
		wantMsg  string
	}{
		{"400 array", 400, `{"message":["encf must be a string"],"code":400}`, "400", "encf must be a string"},
		{"404 string", 404, `{"message":"Not found","code":404}`, "404", "Not found"},
		{"500 coded", 500, `{"message":"boom","code":"EPR500"}`, "EPR500", "boom"},
		{"ap code", 400, `{"message":"ENCF document was used","code":"AP3001"}`, "AP3001", "ENCF document was used"},
		{"non-json", 502, `<html>bad gateway</html>`, "502", "<html>bad gateway</html>"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				io.WriteString(w, tc.body)
			})
			_, err := c.GetSelfCompany(context.Background())
			var apiErr *Error
			if !errors.As(err, &apiErr) {
				t.Fatalf("expected *Error, got %T: %v", err, err)
			}
			if apiErr.StatusCode != tc.status || apiErr.Code != tc.wantCode {
				t.Fatalf("got status=%d code=%q", apiErr.StatusCode, apiErr.Code)
			}
			if len(apiErr.Messages) == 0 || apiErr.Messages[0] != tc.wantMsg {
				t.Fatalf("messages = %v", apiErr.Messages)
			}
		})
	}

	if !IsErrorCode(&Error{Code: "AP3001"}, ErrENCFUsed) {
		t.Fatal("IsErrorCode should match AP3001")
	}
}

func TestGetSelfCompanyUnwrapsCompany(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"company":{"id":"C1","name":"Sculpted","identification":"131234567",
			"companyUrls":{"approval":"https://a","authentication":"https://b"},
			"certificate":{"name":"cert","extension":"p12","endDate":"2027-01-01 00:00:00"}}}`)
	})
	co, err := c.GetSelfCompany(context.Background())
	if err != nil {
		t.Fatalf("GetSelfCompany: %v", err)
	}
	if co.ID != "C1" || co.CompanyURLs.Approval != "https://a" || co.Certificate.EndDate == "" {
		t.Fatalf("unexpected company: %+v", co)
	}
}

func TestSignDocumentSendsMultipartXML(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			t.Errorf("content-type = %s", r.Header.Get("Content-Type"))
		}
		f, hdr, err := r.FormFile("xml")
		if err != nil {
			t.Fatalf("form file: %v", err)
		}
		data, _ := io.ReadAll(f)
		if hdr.Filename != "postulacion.xml" || string(data) != "<x/>" {
			t.Errorf("file = %s %q", hdr.Filename, data)
		}
		io.WriteString(w, `{"signedDocumentUrl":"https://signed"}`)
	})
	resp, err := c.SignDocument(context.Background(), strings.NewReader("<x/>"), "postulacion.xml")
	if err != nil {
		t.Fatalf("SignDocument: %v", err)
	}
	if resp.SignedDocumentURL != "https://signed" {
		t.Fatalf("url = %q", resp.SignedDocumentURL)
	}
}

func TestParseWebhookEvents(t *testing.T) {
	parse := func(body string) *WebhookEvent {
		t.Helper()
		r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(body))
		ev, err := ParseWebhook(r)
		if err != nil {
			t.Fatalf("ParseWebhook(%s): %v", body, err)
		}
		return ev
	}

	ev := parse(`{"company":{"id":"C1","hasBeenCertified":"yes","certificationStep":5}}`)
	if ev.GovernmentStatusChanged == nil || ev.GovernmentStatusChanged.Company.HasBeenCertified != "yes" {
		t.Fatalf("government status event not detected: %+v", ev)
	}

	ev = parse(`{"type":"fiscalInvoice","id":"D1","status":"FINISHED","legalStatus":"REJECTED",
		"documentNumber":"E310000000001","governmentResponse":{"value":[{"valor":"x","codigo":1934}],"code":2}}`)
	if ev.EmissionFinished == nil || ev.EmissionFinished.ID != "D1" {
		t.Fatalf("emission event not detected: %+v", ev)
	}
	if ev.EmissionFinished.GovernmentResponse.Value[0].Code != "1934" {
		t.Fatalf("emission codigo = %q", ev.EmissionFinished.GovernmentResponse.Value[0].Code)
	}

	ev = parse(`{"type":"receivedDocument","id":"R1","timestamp":"t"}`)
	if ev.DocumentReception == nil {
		t.Fatalf("reception event not detected: %+v", ev)
	}

	r := httptest.NewRequest(http.MethodPost, "/webhook", strings.NewReader(`{"foo":1}`))
	if _, err := ParseWebhook(r); err == nil {
		t.Fatal("expected error for unknown webhook format")
	}
}
