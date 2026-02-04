package tevalis

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/white-streak-technologies/tevalis-go/models"
)

const DefaultBaseURL = "https://api.tevalis.com"
const DefaultDateFormat = "2006-01-02"

// Client is the Tevalis API client.
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	headers     http.Header
	dateFormat  string
	contentType string
}

// Option configures the client.
type Option func(*Client) error

// NewClient creates a new Tevalis API client.
func NewClient(options ...Option) (*Client, error) {
	base, err := url.Parse(DefaultBaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:     base,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		headers:     make(http.Header),
		dateFormat:  DefaultDateFormat,
		contentType: DefaultContentType,
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithBaseURL sets a custom API base URL.
func WithBaseURL(rawURL string) Option {
	return func(c *Client) error {
		parsed, err := url.Parse(rawURL)
		if err != nil {
			return err
		}
		if parsed.Scheme == "" || parsed.Host == "" {
			return errors.New("invalid base URL")
		}
		c.baseURL = parsed
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.New("http client cannot be nil")
		}
		c.httpClient = httpClient
		return nil
	}
}

// WithHeader sets a default header for all requests.
func WithHeader(key, value string) Option {
	return func(c *Client) error {
		if strings.TrimSpace(key) == "" {
			return errors.New("header name cannot be empty")
		}
		c.headers.Set(key, value)
		return nil
	}
}

// WithDateFormat sets the date format used for StartDate and EndDate query values.
func WithDateFormat(layout string) Option {
	return func(c *Client) error {
		if strings.TrimSpace(layout) == "" {
			return errors.New("date format cannot be empty")
		}
		c.dateFormat = layout
		return nil
	}
}

// WithContentType sets the request/response format.
// Use application/xml (default) or application/json.
func WithContentType(contentType string) Option {
	return func(c *Client) error {
		trimmed := strings.TrimSpace(contentType)
		if trimmed == "" {
			return errors.New("content type cannot be empty")
		}
		c.contentType = trimmed
		return nil
	}
}

// WithCredentials sets the required authentication headers.
func WithCredentials(companyID int, guid, developerID string, guid2 ...string) Option {
	return func(c *Client) error {
		if companyID <= 0 {
			return errors.New("companyID must be greater than 0")
		}
		if strings.TrimSpace(guid) == "" {
			return errors.New("guid cannot be empty")
		}
		if strings.TrimSpace(developerID) == "" {
			return errors.New("developerID cannot be empty")
		}
		c.headers.Set(HeaderCompanyID, strconv.Itoa(companyID))
		c.headers.Set(HeaderGUID, guid)
		c.headers.Set(HeaderDeveloperID, developerID)
		if len(guid2) > 0 && strings.TrimSpace(guid2[0]) != "" {
			c.headers.Set(HeaderGUID2, guid2[0])
		}
		return nil
	}
}

func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values) (*http.Request, error) {
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}

	relative := &url.URL{Path: path, RawQuery: query.Encode()}
	u := c.baseURL.ResolveReference(relative)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.contentType != "" {
		req.Header.Set("Content-Type", c.contentType)
		req.Header.Set("Accept", c.contentType)
	}

	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

// GetSalesExport calls GET /{version}/Sales/GetSalesExport.
func (c *Client) GetSalesExport(ctx context.Context, version string, siteID int, startDate, endDate time.Time) (*models.SalesExportModel, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}
	if strings.TrimSpace(version) == "" {
		return nil, fmt.Errorf("version cannot be empty")
	}

	query := url.Values{}
	query.Set("StartDate", startDate.Format(c.dateFormat))
	query.Set("EndDate", endDate.Format(c.dateFormat))

	path := fmt.Sprintf("/%s/Sales/GetSalesExport/%s", strings.TrimPrefix(strings.TrimSpace(version), "/"), strconv.Itoa(siteID))
	request, err := c.newRequest(ctx, http.MethodGet, path, query)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, &APIError{StatusCode: response.StatusCode, Body: payload}
	}

	var salesResp models.SalesExportModel
	if c.contentType == "application/xml" {
		if err := xml.Unmarshal(payload, &salesResp); err != nil {
			return nil, err
		}
	} else {
		if err := json.Unmarshal(payload, &salesResp); err != nil {
			return nil, err
		}
	}

	return &salesResp, nil
}
