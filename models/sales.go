package models

// SalesExportModel contains raw response data from the sales export endpoint.
type SalesExportModel struct {
	XMLName      struct{}                   `json:"-" xml:"SalesExportModel"`
	SiteID       int                        `json:"SiteID" xml:"SiteID"`
	Transactions []BillHeaderAndFooterModel `json:"transactions" xml:"transactions>BillHeaderAndFooterModel"`
}
