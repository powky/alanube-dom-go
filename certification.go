package alanube

import (
	"context"
	"io"
	"net/http"
)

// CreateSetTest generates the test set required for the certification process.
// POST /set-tests
func (c *Client) CreateSetTest(ctx context.Context, req *CreateSetTestRequest) (*SetTestCreateResponse, error) {
	var resp SetTestCreateResponse
	if err := c.do(ctx, http.MethodPost, "/set-tests", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSetTest checks the test set status for the main company.
// GET /check-set-tests/{id}
func (c *Client) GetSetTest(ctx context.Context, id string) (*SetTestCheckResponse, error) {
	var resp SetTestCheckResponse
	if err := c.do(ctx, http.MethodGet, "/check-set-tests/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSetTestByCompany checks the test set status for an associated company.
// GET /check-set-tests/{id}/idCompany/{companyID}
func (c *Client) GetSetTestByCompany(ctx context.Context, id, companyID string) (*SetTestCheckResponse, error) {
	var resp SetTestCheckResponse
	if err := c.do(ctx, http.MethodGet, "/check-set-tests/"+id+"/idCompany/"+companyID, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SignDocument signs an XML document for the certification process (main company).
// POST /sign-document (multipart/form-data with xml file)
func (c *Client) SignDocument(ctx context.Context, xmlFile io.Reader, fileName string) (*SignDocumentResponse, error) {
	var resp SignDocumentResponse
	if err := c.doMultipartFile(ctx, "/sign-document", "xml", fileName, xmlFile, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SignDocumentByCompany signs an XML document for the certification process for a specific company.
// POST /sign-document/idCompany/{companyID} (multipart/form-data with xml file)
func (c *Client) SignDocumentByCompany(ctx context.Context, companyID string, xmlFile io.Reader, fileName string) (*SignDocumentResponse, error) {
	var resp SignDocumentResponse
	if err := c.doMultipartFile(ctx, "/sign-document/idCompany/"+companyID, "xml", fileName, xmlFile, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProviderInfo retrieves Alanube provider information for DGII portal registration.
// GET /provider-info
func (c *Client) GetProviderInfo(ctx context.Context) (*ProviderInfo, error) {
	var resp ProviderInfo
	if err := c.do(ctx, http.MethodGet, "/provider-info", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
