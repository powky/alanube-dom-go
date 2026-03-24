package alanube

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// Client is the Alanube API client.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithBaseURL overrides the default base URL.
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = u }
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.httpClient.Timeout = d }
}

// NewClient creates a new production API client.
func NewClient(token string, opts ...Option) *Client {
	c := &Client{
		baseURL:    ProductionBaseURL,
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// NewSandboxClient creates a new sandbox API client.
func NewSandboxClient(token string, opts ...Option) *Client {
	c := &Client{
		baseURL:    SandboxBaseURL,
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// do performs an HTTP request and decodes the JSON response.
// If body is nil, no request body is sent.
// If result is nil, the response body is discarded.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body, result interface{}) error {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("alanube: marshal request: %w", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, reqBody)
	if err != nil {
		return fmt.Errorf("alanube: create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("alanube: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("alanube: decode response: %w", err)
		}
	}

	return nil
}

// doNoAuth performs an HTTP request without authentication header.
func (c *Client) doNoAuth(ctx context.Context, method, fullURL string, query url.Values, result interface{}) error {
	u := fullURL
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return fmt.Errorf("alanube: create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("alanube: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("alanube: decode response: %w", err)
		}
	}

	return nil
}

// doMultipartFile performs a multipart/form-data POST with a single file field.
func (c *Client) doMultipartFile(ctx context.Context, path, fieldName, fileName string, fileContent io.Reader, result interface{}) error {
	u := c.baseURL + path

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return fmt.Errorf("alanube: create form file: %w", err)
	}
	if _, err := io.Copy(part, fileContent); err != nil {
		return fmt.Errorf("alanube: copy file content: %w", err)
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, &buf)
	if err != nil {
		return fmt.Errorf("alanube: create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("alanube: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("alanube: decode response: %w", err)
		}
	}

	return nil
}

func parseError(resp *http.Response) error {
	var apiErr Error
	apiErr.StatusCode = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Error{StatusCode: resp.StatusCode, Code: fmt.Sprintf("%d", resp.StatusCode)}
	}

	if err := json.Unmarshal(body, &apiErr); err != nil {
		return &Error{StatusCode: resp.StatusCode, Code: fmt.Sprintf("%d", resp.StatusCode), Messages: []string{string(body)}}
	}

	return &apiErr
}
