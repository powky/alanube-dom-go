package alanube

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// acknowledgmentsQuery builds query params for acknowledgments list endpoints.
func acknowledgmentsQuery(opts *AcknowledgmentsQueryOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	if opts.DocumentNumber != "" {
		q.Set("documentNumber", opts.DocumentNumber)
	}
	if opts.RNC != "" {
		q.Set("rnc", opts.RNC)
	}
	if opts.Status != "" {
		q.Set("status", opts.Status)
	}
	if opts.Start != "" {
		q.Set("start", opts.Start)
	}
	if opts.End != "" {
		q.Set("end", opts.End)
	}
	if opts.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", opts.Limit))
	}
	if opts.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", opts.Page))
	}
	return q
}

// GetExternalAcknowledgment retrieves an external reception acknowledgment by ID.
// GET /reception/acknowledgments/{companyID}/{id}
func (c *Client) GetExternalAcknowledgment(ctx context.Context, companyID, id string) (*Acknowledgment, error) {
	var resp Acknowledgment
	if err := c.do(ctx, http.MethodGet, "/reception/acknowledgments/"+companyID+"/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExternalAcknowledgments lists external reception acknowledgments for a company.
// GET /reception/acknowledgments/{companyID}
func (c *Client) GetExternalAcknowledgments(ctx context.Context, companyID string, opts *AcknowledgmentsQueryOptions) ([]Acknowledgment, error) {
	var resp []Acknowledgment
	if err := c.do(ctx, http.MethodGet, "/reception/acknowledgments/"+companyID, acknowledgmentsQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetInternalReceipt retrieves an internal reception receipt by ID.
// GET /reception/receipts/{companyID}/{id}
func (c *Client) GetInternalReceipt(ctx context.Context, companyID, id string) (*Acknowledgment, error) {
	var resp Acknowledgment
	if err := c.do(ctx, http.MethodGet, "/reception/receipts/"+companyID+"/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInternalReceipts lists internal reception receipts for a company.
// GET /reception/receipts/{companyID}
func (c *Client) GetInternalReceipts(ctx context.Context, companyID string, opts *AcknowledgmentsQueryOptions) ([]Acknowledgment, error) {
	var resp []Acknowledgment
	if err := c.do(ctx, http.MethodGet, "/reception/receipts/"+companyID, acknowledgmentsQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
