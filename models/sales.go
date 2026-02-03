package models

// SalesExportResponse contains raw response data from the sales export endpoint.
type SalesExportResponse struct {
	XMLName      struct{}      `json:"-" xml:"SalesExportResponse"`
	SiteID       int           `json:"SiteID" xml:"SiteID"`
	Transactions []Transaction `json:"transactions" xml:"transactions>transaction"`
}
