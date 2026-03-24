package alanube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// dgiiQuery builds query params for DGII status endpoints.
func dgiiQuery(opts *DGIIQueryOptions) url.Values {
	if opts == nil {
		return nil
	}
	q := url.Values{}
	if opts.Environment != "" {
		q.Set("environment", string(opts.Environment))
	} else if opts.Maintenance {
		q.Set("maintenance", "yes")
	}
	return q
}

// CheckDGIIStatus queries the DGII service status.
// GET /check-dgii-status
//
// The response type depends on the query options:
//   - No options: returns *DGIIStatusResponse
//   - Environment: returns *EnvironmentStatus
//   - Maintenance: returns *MaintenanceResponse
//
// Use the typed helper methods for convenience.
func (c *Client) CheckDGIIStatus(ctx context.Context, opts *DGIIQueryOptions) (json.RawMessage, error) {
	var resp json.RawMessage
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status", dgiiQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckDGIIStatusByCompany queries the DGII service status for a specific company.
// GET /check-dgii-status/idCompany/{companyID}
func (c *Client) CheckDGIIStatusByCompany(ctx context.Context, companyID string, opts *DGIIQueryOptions) (json.RawMessage, error) {
	var resp json.RawMessage
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status/idCompany/"+companyID, dgiiQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Typed helper methods ---

// GetDGIIServiceStatus returns the status of all DGII services.
// GET /check-dgii-status (no query params)
func (c *Client) GetDGIIServiceStatus(ctx context.Context) (*DGIIStatusResponse, error) {
	var resp DGIIStatusResponse
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDGIIServiceStatusByCompany returns the status of all DGII services for a specific company.
// GET /check-dgii-status/idCompany/{companyID} (no query params)
func (c *Client) GetDGIIServiceStatusByCompany(ctx context.Context, companyID string) (*DGIIStatusResponse, error) {
	var resp DGIIStatusResponse
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status/idCompany/"+companyID, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDGIIEnvironmentStatus checks the health of a specific DGII environment.
// GET /check-dgii-status?environment={env}
func (c *Client) GetDGIIEnvironmentStatus(ctx context.Context, env DGIIEnvironment) (*EnvironmentStatus, error) {
	var resp EnvironmentStatus
	q := url.Values{}
	q.Set("environment", string(env))
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status", q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDGIIMaintenance checks the DGII maintenance schedule.
// GET /check-dgii-status?maintenance=yes
func (c *Client) GetDGIIMaintenance(ctx context.Context) (*MaintenanceResponse, error) {
	var resp MaintenanceResponse
	q := url.Values{}
	q.Set("maintenance", "yes")
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status", q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDGIIEnvironmentStatusByCompany checks the health of a specific DGII environment for a company.
// GET /check-dgii-status/idCompany/{companyID}?environment={env}
func (c *Client) GetDGIIEnvironmentStatusByCompany(ctx context.Context, companyID string, env DGIIEnvironment) (*EnvironmentStatus, error) {
	var resp EnvironmentStatus
	q := url.Values{}
	q.Set("environment", string(env))
	if err := c.do(ctx, http.MethodGet, "/check-dgii-status/idCompany/"+companyID, q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDGIIMaintenanceByCompany checks the DGII maintenance schedule for a specific company.
// GET /check-dgii-status/idCompany/{companyID}?maintenance=yes
func (c *Client) GetDGIIMaintenanceByCompany(ctx context.Context, companyID string) (*MaintenanceResponse, error) {
	var resp MaintenanceResponse
	q := url.Values{}
	q.Set("maintenance", "yes")
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("/check-dgii-status/idCompany/%s", companyID), q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
