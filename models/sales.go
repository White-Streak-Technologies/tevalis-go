package models

import (
	"net/http"
)

// SalesExportResponse contains raw response data from the sales export endpoint.
type SalesExportResponse struct {
	StatusCode  int
	ContentType string
	Headers     http.Header
	Body        []byte
}
