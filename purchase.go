package alanube

import (
	"context"
	"net/http"
)

// --- Purchase (41) ---

// CreatePurchase creates a purchase document (eCF type 41).
// POST /purchases
func (c *Client) CreatePurchase(ctx context.Context, req *CreatePurchaseRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/purchases", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPurchase checks the status of a purchase document.
// GET /purchases/{id}
func (c *Client) GetPurchase(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/purchases/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPurchaseByCompany checks a purchase by ID and company.
// GET /purchases/{id}/{companyID}
func (c *Client) GetPurchaseByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/purchases/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyPurchase sends a purchase notification by email.
// POST /purchases/notify-by-email
func (c *Client) NotifyPurchase(ctx context.Context, req *NotifyByEmailRequest) (*NotifyByEmailResponse, error) {
	var resp NotifyByEmailResponse
	if err := c.do(ctx, http.MethodPost, "/purchases/notify-by-email", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Minor Expense (43) ---

// CreateMinorExpense creates a minor expense document (eCF type 43).
// POST /minor-expenses
func (c *Client) CreateMinorExpense(ctx context.Context, req *CreateMinorExpenseRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/minor-expenses", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMinorExpense checks the status of a minor expense document.
// GET /minor-expenses/{id}
func (c *Client) GetMinorExpense(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/minor-expenses/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMinorExpenseByCompany checks a minor expense by ID and company.
// GET /minor-expenses/{id}/{companyID}
func (c *Client) GetMinorExpenseByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/minor-expenses/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --- Payment Abroad (47) ---

// CreatePaymentAbroad creates a payment abroad support document (eCF type 47).
// POST /payment-abroad-supports
func (c *Client) CreatePaymentAbroad(ctx context.Context, req *CreatePaymentAbroadRequest) (*DocumentResponse, error) {
	var resp DocumentResponse
	if err := c.do(ctx, http.MethodPost, "/payment-abroad-supports", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPaymentAbroad checks the status of a payment abroad document.
// GET /payment-abroad-supports/{id}
func (c *Client) GetPaymentAbroad(ctx context.Context, id string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/payment-abroad-supports/"+id, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPaymentAbroadByCompany checks a payment abroad by ID and company.
// GET /payment-abroad-supports/{id}/{companyID}
func (c *Client) GetPaymentAbroadByCompany(ctx context.Context, id, companyID string, opts *DocumentQueryOptions) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/payment-abroad-supports/"+id+"/idCompany/"+companyID, documentQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
