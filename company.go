package alanube

import (
	"context"
	"net/http"
)

// CreateCompany registers a new company.
// POST /company
func (c *Client) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*Company, error) {
	var resp companyWrapper
	if err := c.do(ctx, http.MethodPost, "/company", nil, req, &resp); err != nil {
		return nil, err
	}
	return resp.Company, nil
}

// GetSelfCompany retrieves the company associated with the current token.
// GET /company
func (c *Client) GetSelfCompany(ctx context.Context) (*Company, error) {
	var resp companyWrapper
	if err := c.do(ctx, http.MethodGet, "/company", nil, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Company, nil
}

// UpdateSelfCompany updates the company associated with the current token.
// PATCH /company
func (c *Client) UpdateSelfCompany(ctx context.Context, req *UpdateCompanyRequest) (*Company, error) {
	var resp companyWrapper
	if err := c.do(ctx, http.MethodPatch, "/company", nil, req, &resp); err != nil {
		return nil, err
	}
	return resp.Company, nil
}

// GetCompany retrieves a company by ID.
// GET /company/{id}
func (c *Client) GetCompany(ctx context.Context, id string) (*Company, error) {
	var resp companyWrapper
	if err := c.do(ctx, http.MethodGet, "/company/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Company, nil
}

// UpdateCompany updates a company by ID.
// PATCH /company/{id}
func (c *Client) UpdateCompany(ctx context.Context, id string, req *UpdateCompanyRequest) (*Company, error) {
	var resp companyWrapper
	if err := c.do(ctx, http.MethodPatch, "/company/"+id, nil, req, &resp); err != nil {
		return nil, err
	}
	return resp.Company, nil
}
