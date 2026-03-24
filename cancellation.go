package alanube

import (
	"context"
	"net/http"
)

// CreateCancellation creates a document cancellation (range annulment).
// POST /cancellations
func (c *Client) CreateCancellation(ctx context.Context, req *CreateCancellationRequest) (*CancellationResponse, error) {
	var resp CancellationResponse
	if err := c.do(ctx, http.MethodPost, "/cancellations", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCancellation checks the status of a cancellation.
// GET /cancellations/{id}
func (c *Client) GetCancellation(ctx context.Context, id string) (*CancellationStatus, error) {
	var resp CancellationStatus
	if err := c.do(ctx, http.MethodGet, "/cancellations/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCancellationByCompany checks a cancellation by ID and company.
// GET /cancellations/{id}/idCompany/{companyID}
func (c *Client) GetCancellationByCompany(ctx context.Context, id, companyID string) (*DocumentStatus, error) {
	var resp DocumentStatus
	if err := c.do(ctx, http.MethodGet, "/cancellations/"+id+"/idCompany/"+companyID, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
