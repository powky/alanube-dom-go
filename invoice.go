package alanube

import (
	"context"
	"net/http"
	"net/url"
)

// documentQuery builds query params from DocumentQueryOptions.
func documentQuery(opts *DocumentQueryOptions) url.Values {
	if opts == nil || opts.PDFType == "" {
		return nil
	}
	q := url.Values{}
	q.Set("pdfType", opts.PDFType)
	return q
}

// --- Fiscal Credit Invoice (31) ---

// CreateFiscalInvoice creates a fiscal credit invoice (eCF type 31).
// POST /fiscal-invoices
func (c *Client) CreateFiscalInvoice(ctx context.Context, req *CreateFiscalInvoiceRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/fiscal-invoices", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFiscalInvoice checks the status of a fiscal credit invoice.
// GET /fiscal-invoices/{id}
func (c *Client) GetFiscalInvoice(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/fiscal-invoices/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFiscalInvoiceByCompany checks a fiscal credit invoice by ID and company.
// GET /fiscal-invoices/{id}/idCompany/{idCompany}
func (c *Client) GetFiscalInvoiceByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/fiscal-invoices/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyFiscalInvoice sends a fiscal credit invoice notification by email.
// POST /fiscal-invoices/notify-by-email
func (c *Client) NotifyFiscalInvoice(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/fiscal-invoices/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Consumption Invoice (32) ---

// CreateConsumptionInvoice creates a consumption invoice (eCF type 32).
// POST /invoices
func (c *Client) CreateConsumptionInvoice(ctx context.Context, req *CreateConsumptionInvoiceRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/invoices", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConsumptionInvoice checks the status of a consumption invoice.
// GET /invoices/{id}
func (c *Client) GetConsumptionInvoice(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/invoices/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConsumptionInvoiceByCompany checks a consumption invoice by ID and company.
// GET /invoices/{id}/idCompany/{idCompany}
func (c *Client) GetConsumptionInvoiceByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/invoices/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyConsumptionInvoice sends a consumption invoice notification by email.
// POST /invoices/notify-by-email
func (c *Client) NotifyConsumptionInvoice(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/invoices/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Special Regimes Invoice (44) ---

// CreateSpecialRegimesInvoice creates a special regimes invoice (eCF type 44).
// POST /special-regimes
func (c *Client) CreateSpecialRegimesInvoice(ctx context.Context, req *CreateSpecialRegimesInvoiceRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/special-regimes", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSpecialRegimesInvoice checks the status of a special regimes invoice.
// GET /special-regimes/{id}
func (c *Client) GetSpecialRegimesInvoice(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/special-regimes/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSpecialRegimesInvoiceByCompany checks a special regimes invoice by ID and company.
// GET /special-regimes/{id}/idCompany/{idCompany}
func (c *Client) GetSpecialRegimesInvoiceByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/special-regimes/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifySpecialRegimesInvoice sends a special regimes invoice notification by email.
// POST /special-regimes/notify-by-email
func (c *Client) NotifySpecialRegimesInvoice(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/special-regimes/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Governmental Invoice (45) ---

// CreateGovernmentalInvoice creates a governmental invoice (eCF type 45).
// POST /gubernamentals
func (c *Client) CreateGovernmentalInvoice(ctx context.Context, req *CreateGovernmentalInvoiceRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/gubernamentals", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGovernmentalInvoice checks the status of a governmental invoice.
// GET /gubernamentals/{id}
func (c *Client) GetGovernmentalInvoice(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/gubernamentals/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGovernmentalInvoiceByCompany checks a governmental invoice by ID and company.
// GET /gubernamentals/{id}/idCompany/{idCompany}
func (c *Client) GetGovernmentalInvoiceByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/gubernamentals/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyGovernmentalInvoice sends a governmental invoice notification by email.
// POST /gubernamentals/notify-by-email
func (c *Client) NotifyGovernmentalInvoice(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/gubernamentals/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Export Invoice (46) ---

// CreateExportInvoice creates an export invoice (eCF type 46).
// POST /export-supports
func (c *Client) CreateExportInvoice(ctx context.Context, req *CreateExportInvoiceRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/export-supports", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExportInvoice checks the status of an export invoice.
// GET /export-supports/{id}
func (c *Client) GetExportInvoice(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/export-supports/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExportInvoiceByCompany checks an export invoice by ID and company.
// GET /export-supports/{id}/idCompany/{idCompany}
func (c *Client) GetExportInvoiceByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/export-supports/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyExportInvoice sends an export invoice notification by email.
// POST /export-supports/notify-by-email
func (c *Client) NotifyExportInvoice(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/export-supports/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
