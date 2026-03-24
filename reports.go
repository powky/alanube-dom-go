package alanube

import (
	"context"
	"net/http"
	"net/url"
)

// reportQuery builds query params for report endpoints.
func reportQuery(opts *ReportQueryOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	if opts.LegalStatus != "" {
		q.Set("legalStatus", opts.LegalStatus)
	}
	if opts.DateFrom != "" {
		q.Set("dateFrom", opts.DateFrom)
	}
	if opts.DateUntil != "" {
		q.Set("dateUntil", opts.DateUntil)
	}
	return q
}

// GetUserDocumentTotals retrieves total emitted documents for the authenticated user.
// GET /reports/users/documents/total
func (c *Client) GetUserDocumentTotals(ctx context.Context, opts *ReportQueryOptions) (*UserDocumentTotalsResponse, error) {
	var resp UserDocumentTotalsResponse
	if err := c.do(ctx, http.MethodGet, "/reports/users/documents/total", reportQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCompanyDocumentTotals retrieves total emitted documents for a specific company.
// GET /reports/companies/{companyID}/documents/total
func (c *Client) GetCompanyDocumentTotals(ctx context.Context, companyID string, opts *ReportQueryOptions) (*CompanyDocumentTotalsResponse, error) {
	var resp CompanyDocumentTotalsResponse
	if err := c.do(ctx, http.MethodGet, "/reports/companies/"+companyID+"/documents/total", reportQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCompanyAcceptedDocuments retrieves the historical total of accepted documents for a company.
// Only includes documents with legalStatus ACCEPTED or ACCEPTED_WITH_OBSERVATIONS.
// GET /companies/{companyID}/accepted-documents
func (c *Client) GetCompanyAcceptedDocuments(ctx context.Context, companyID string) (*CompanyAcceptedDocumentTotals, error) {
	var resp CompanyAcceptedDocumentTotals
	if err := c.do(ctx, http.MethodGet, "/companies/"+companyID+"/accepted-documents", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCompanyEmittedDocuments retrieves the historical total of emitted documents for a company.
// No date or legal status filters - returns all-time totals.
// GET /companies/{companyID}/emitted-documents
func (c *Client) GetCompanyEmittedDocuments(ctx context.Context, companyID string) (*CompanyDocumentTotals, error) {
	var resp CompanyDocumentTotals
	if err := c.do(ctx, http.MethodGet, "/companies/"+companyID+"/emitted-documents", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
