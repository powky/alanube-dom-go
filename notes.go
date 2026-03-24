package alanube

import (
	"context"
	"net/http"
)

// --- Debit Note (33) ---

// CreateDebitNote creates a debit note (eCF type 33).
// POST /debit-notes
func (c *Client) CreateDebitNote(ctx context.Context, req *CreateDebitNoteRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/debit-notes", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDebitNote checks the status of a debit note.
// GET /debit-notes/{id}
func (c *Client) GetDebitNote(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/debit-notes/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDebitNoteByCompany checks a debit note by ID and company.
// GET /debit-notes/{id}/idCompany/{companyID}
func (c *Client) GetDebitNoteByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/debit-notes/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyDebitNote sends a debit note notification by email.
// POST /debit-notes/notify-by-email
func (c *Client) NotifyDebitNote(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/debit-notes/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Credit Note (34) ---

// CreateCreditNote creates a credit note (eCF type 34).
// POST /credit-notes
func (c *Client) CreateCreditNote(ctx context.Context, req *CreateCreditNoteRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/credit-notes", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCreditNote checks the status of a credit note.
// GET /credit-notes/{id}
func (c *Client) GetCreditNote(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/credit-notes/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCreditNoteByCompany checks a credit note by ID and company.
// GET /credit-notes/{id}/idCompany/{companyID}
func (c *Client) GetCreditNoteByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/credit-notes/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyCreditNote sends a credit note notification by email.
// POST /credit-notes/notify-by-email
func (c *Client) NotifyCreditNote(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/credit-notes/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
