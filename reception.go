package alanube

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// receivedDocsQuery builds query params for received documents list endpoints.
func receivedDocsQuery(opts *ReceivedDocumentsQueryOptions, companyID string) url.Values {
	q := url.Values{}
	if companyID != "" {
		q.Set("idCompany", companyID)
	}
	if opts == nil {
		return q
	}
	if opts.Status != "" {
		q.Set("status", opts.Status)
	}
	if opts.CommercialResponse != "" {
		q.Set("commercialResponse", opts.CommercialResponse)
	}
	if opts.RNC != "" {
		q.Set("rnc", opts.RNC)
	}
	if opts.DocumentNumber != "" {
		q.Set("documentNumber", opts.DocumentNumber)
	}
	if opts.Limit > 0 {
		q.Set("limit", fmt.Sprintf("%d", opts.Limit))
	}
	if opts.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", opts.Page))
	}
	if opts.Start != "" {
		q.Set("start", opts.Start)
	}
	if opts.End != "" {
		q.Set("end", opts.End)
	}
	return q
}

// ReceiveDocumentFromDGII simulates receiving an electronic document from DGII (sandbox only).
// On success returns 201 with empty body. The received document will appear in the electronic mailbox.
// POST /reception/dgii/receive-document
func (c *Client) ReceiveDocumentFromDGII(ctx context.Context, req *ReceiveDocumentRequest) error {
	return c.do(ctx, http.MethodPost, "/reception/dgii/receive-document", nil, req, nil)
}

// GetReceivedDocuments lists received documents for the main company.
// GET /received-documents
func (c *Client) GetReceivedDocuments(ctx context.Context, opts *ReceivedDocumentsQueryOptions) (*ReceivedDocumentsResponse, error) {
	var resp ReceivedDocumentsResponse
	if err := c.do(ctx, http.MethodGet, "/received-documents", receivedDocsQuery(opts, ""), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedDocumentsByCompany lists received documents for a specific company.
// GET /received-documents/idCompany/{companyID}
func (c *Client) GetReceivedDocumentsByCompany(ctx context.Context, companyID string, opts *ReceivedDocumentsQueryOptions) (*ReceivedDocumentsResponse, error) {
	var resp ReceivedDocumentsResponse
	if err := c.do(ctx, http.MethodGet, "/received-documents/idCompany/"+companyID, receivedDocsQuery(opts, ""), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedDocument retrieves a received document by ID.
// GET /received-documents/{id}
func (c *Client) GetReceivedDocument(ctx context.Context, id string) (*ReceivedDocument, error) {
	var resp ReceivedDocument
	if err := c.do(ctx, http.MethodGet, "/received-documents/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedDocumentByCompany retrieves a received document by ID and company.
// GET /received-documents/{id}/idCompany/{companyID}
func (c *Client) GetReceivedDocumentByCompany(ctx context.Context, id, companyID string) (*ReceivedDocument, error) {
	var resp ReceivedDocument
	if err := c.do(ctx, http.MethodGet, "/received-documents/"+id+"/idCompany/"+companyID, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCommercialResponse generates a commercial response for a received document.
// POST /received-documents/{docID}/commercial-response
func (c *Client) CreateCommercialResponse(ctx context.Context, docID string, req *CommercialResponseRequest) (*CommercialResponseResult, error) {
	var resp CommercialResponseResult
	if err := c.do(ctx, http.MethodPost, "/received-documents/"+docID+"/commercial-response", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCommercialResponseByCompany generates a commercial response for a received document by company.
// POST /received-documents/{docID}/commercial-response/idCompany/{companyID}
func (c *Client) CreateCommercialResponseByCompany(ctx context.Context, docID, companyID string, req *CommercialResponseRequest) (*CommercialResponseResult, error) {
	var resp CommercialResponseResult
	if err := c.do(ctx, http.MethodPost, "/received-documents/"+docID+"/commercial-response/idCompany/"+companyID, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCommercialApproval retrieves a generated commercial approval by ID.
// GET /commercial-approvals/{id}
func (c *Client) GetCommercialApproval(ctx context.Context, id string) (*CommercialApproval, error) {
	var resp CommercialApproval
	if err := c.do(ctx, http.MethodGet, "/commercial-approvals/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCommercialApprovalByCompany retrieves a generated commercial approval by ID and company.
// GET /commercial-approvals/{id}?idCompany={companyID}
func (c *Client) GetCommercialApprovalByCompany(ctx context.Context, id, companyID string) (*CommercialApproval, error) {
	var resp CommercialApproval
	q := queryWith("idCompany", companyID)
	if err := c.do(ctx, http.MethodGet, "/commercial-approvals/"+id, q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// commercialApprovalsQuery builds query params for commercial approvals list endpoints.
func commercialApprovalsQuery(opts *CommercialApprovalsQueryOptions) url.Values {
	q := url.Values{}
	if opts != nil {
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Page > 0 {
			q.Set("page", fmt.Sprintf("%d", opts.Page))
		}
	}
	return q
}

// GetReceivedCommercialApprovals lists received commercial approvals for the main company.
// GET /received-commercial-approvals
func (c *Client) GetReceivedCommercialApprovals(ctx context.Context, opts *CommercialApprovalsQueryOptions) (*CommercialApprovalsResponse, error) {
	var resp CommercialApprovalsResponse
	if err := c.do(ctx, http.MethodGet, "/received-commercial-approvals", commercialApprovalsQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedCommercialApprovalsByCompany lists received commercial approvals for a specific company.
// GET /received-commercial-approvals/idCompany/{companyID}
func (c *Client) GetReceivedCommercialApprovalsByCompany(ctx context.Context, companyID string, opts *CommercialApprovalsQueryOptions) (*CommercialApprovalsResponse, error) {
	var resp CommercialApprovalsResponse
	if err := c.do(ctx, http.MethodGet, "/received-commercial-approvals/idCompany/"+companyID, commercialApprovalsQuery(opts), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedCommercialApproval retrieves a received commercial approval by ID.
// GET /received-commercial-approvals/{id}
func (c *Client) GetReceivedCommercialApproval(ctx context.Context, id string) (*CommercialApproval, error) {
	var resp CommercialApproval
	if err := c.do(ctx, http.MethodGet, "/received-commercial-approvals/"+id, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReceivedCommercialApprovalByCompany retrieves a received commercial approval by ID and company.
// GET /received-commercial-approvals/{id}/idCompany/{companyID}
func (c *Client) GetReceivedCommercialApprovalByCompany(ctx context.Context, id, companyID string) (*CommercialApproval, error) {
	var resp CommercialApproval
	if err := c.do(ctx, http.MethodGet, "/received-commercial-approvals/"+id+"/idCompany/"+companyID, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
