package alanube

import (
	"context"
	"net/http"
	"net/url"
)

// directoryQuery builds query params for directory endpoints.
func directoryQuery(opts *DirectoryQueryOptions) url.Values {
	if opts == nil || opts.RNC == "" {
		return nil
	}
	q := url.Values{}
	q.Set("rnc", opts.RNC)
	return q
}

// GetDirectory lists active companies in the electronic document directory.
// GET /check-directory
func (c *Client) GetDirectory(ctx context.Context, opts *DirectoryQueryOptions) ([]DirectoryEntry, error) {
	var resp []DirectoryEntry
	if err := c.do(ctx, http.MethodGet, "/check-directory", directoryQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDirectoryByCompany lists active companies in the directory for a specific company.
// GET /check-directory/idCompany/{companyID}
func (c *Client) GetDirectoryByCompany(ctx context.Context, companyID string, opts *DirectoryQueryOptions) ([]DirectoryEntry, error) {
	var resp []DirectoryEntry
	if err := c.do(ctx, http.MethodGet, "/check-directory/idCompany/"+companyID, directoryQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
