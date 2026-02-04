package models

// BillHeaderAndFooterModel is a detailed model of a sales transaction.
type BillHeaderAndFooterModel struct {
	EPOSTransactionID              int            `json:"EPOSTransactionID" xml:"EPOSTransactionID"`
	SaleDate                       FlexibleTime   `json:"SaleDate" xml:"SaleDate"`
	WeekNo                         int            `json:"WeekNo" xml:"WeekNo"`
	CoverCount                     int            `json:"CoverCount" xml:"CoverCount"`
	TableNumber                    int            `json:"TableNumber" xml:"TableNumber"`
	Session                        string         `json:"Session" xml:"Session"`
	RevenueCenter                  string         `json:"RevenueCenter" xml:"RevenueCenter"`
	ActualGrossAmt                 float64        `json:"ActualGrossAmt" xml:"ActualGrossAmt"`
	ActualNetAmt                   float64        `json:"ActualNetAmt" xml:"ActualNetAmt"`
	TaxableNetAmt                  float64        `json:"TaxableNetAmt" xml:"TaxableNetAmt"`
	TaxAmt                         float64        `json:"TaxAmt" xml:"TaxAmt"`
	ServiceChargeAmt               float64        `json:"ServiceChargeAmt" xml:"ServiceChargeAmt"`
	GratuityAmt                    float64        `json:"GratuityAmt" xml:"GratuityAmt"`
	TotalDiscountAmount            float64        `json:"TotalDiscountAmount" xml:"TotalDiscountAmount"`
	OpenDateTime                   FlexibleTime   `json:"OpenDateTime" xml:"OpenDateTime"`
	ClosedDateTime                 FlexibleTime   `json:"ClosedDateTime" xml:"ClosedDateTime"`
	EmployeeOpenedBillID           int            `json:"EmployeeOpenedBillID" xml:"EmployeeOpenedBillID"`
	EmployeeOpenedBill             string         `json:"EmployeeOpenedBill" xml:"EmployeeOpenedBill"`
	EmployeeClosedBillID           int            `json:"EmployeeClosedBillID" xml:"EmployeeClosedBillID"`
	EmployeeClosedBill             string         `json:"EmployeeClosedBill" xml:"EmployeeClosedBill"`
	BillPaid                       bool           `json:"BillPaid" xml:"BillPaid"`
	BillTransferred                bool           `json:"BillTransferred" xml:"BillTransferred"`
	EmployeeWhoTransferredBillID   int            `json:"EmployeeWhoTransferredBillID" xml:"EmployeeWhoTransferredBillID"`
	EmployeeWhoTransferredBillName string         `json:"EmployeeWhoTransferredBillName" xml:"EmployeeWhoTransferredBillName"`
	TransferDateTime               FlexibleTime   `json:"TransferDateTime" xml:"TransferDateTime"`
	CRMMemberID                    int            `json:"CRMMemberID" xml:"CRMMemberID"`
	IsRefund                       bool           `json:"IsRefund" xml:"IsRefund"`
	Terminal                       string         `json:"Terminal" xml:"Terminal"`
	ExternalMembershipID           string         `json:"ExternalMembershipID" xml:"ExternalMembershipID"`
	BillItemInfos                  []BillItemInfo `json:"BillItemInfos" xml:"BillItemInfos>BillItemInfo"`
	PaymentInfos                   []PaymentInfo  `json:"PaymentInfos" xml:"PaymentInfos>PaymentInfo"`
}

type BillItemInfo struct {
	EPOSTransactionItemID int             `json:"EPOSTransactionItemID" xml:"EPOSTransactionItemID"`
	EPOSTransactionID     int             `json:"EPOSTransactionID" xml:"EPOSTransactionID"`
	QtySold               float64         `json:"QtySold" xml:"QtySold"`
	UnitQuantity          float64         `json:"UnitQuantity" xml:"UnitQuantity"`
	OfSku                 float64         `json:"OfSku" xml:"OfSku"`
	CostPrice             float64         `json:"CostPrice" xml:"CostPrice"`
	ExpectedGrossAmt      float64         `json:"ExpectedGrossAmt" xml:"ExpectedGrossAmt"`
	ActualGrossAmt        float64         `json:"ActualGrossAmt" xml:"ActualGrossAmt"`
	TaxableNetAmt         float64         `json:"TaxableNetAmt" xml:"TaxableNetAmt"`
	TaxAmt                float64         `json:"TaxAmt" xml:"TaxAmt"`
	VATRate               float64         `json:"VATRate" xml:"VATRate"`
	VATID                 int             `json:"VATID" xml:"VATID"`
	MeasureID             int             `json:"MeasureID" xml:"MeasureID"`
	MeasureName           string          `json:"MeasureName" xml:"MeasureName"`
	ProductID             int             `json:"ProductID" xml:"ProductID"`
	ProductName           string          `json:"ProductName" xml:"ProductName"`
	ProductGroupID        int             `json:"ProductGroupID" xml:"ProductGroupID"`
	ProductGroupName      string          `json:"ProductGroupName" xml:"ProductGroupName"`
	ProductTypeID         int             `json:"ProductTypeID" xml:"ProductTypeID"`
	ProductTypeName       string          `json:"ProductTypeName" xml:"ProductTypeName"`
	PriceGroupID          int             `json:"PriceGroupID" xml:"PriceGroupID"`
	PriceGroupName        string          `json:"PriceGroupName" xml:"PriceGroupName"`
	PLU                   string          `json:"PLU" xml:"PLU"`
	CostCenter            string          `json:"CostCenter" xml:"CostCenter"`
	SessionName           string          `json:"SessionName" xml:"SessionName"`
	EmployeeWhoSoldID     int             `json:"EmployeeWhoSoldID" xml:"EmployeeWhoSoldID"`
	EmployeeWhoSoldName   string          `json:"EmployeeWhoSoldName" xml:"EmployeeWhoSoldName"`
	EnterDateTime         FlexibleTime    `json:"EnterDateTime" xml:"EnterDateTime"`
	IsVoided              bool            `json:"IsVoided" xml:"IsVoided"`
	VoidReason            *string         `json:"VoidReason" xml:"VoidReason"`
	EmployeeWhoVoidedID   int             `json:"EmployeeWhoVoidedID" xml:"EmployeeWhoVoidedID"`
	EmployeeWhoVoidedName string          `json:"EmployeeWhoVoidedName" xml:"EmployeeWhoVoidedName"`
	VoidDateTime          FlexibleTime    `json:"VoidDateTime" xml:"VoidDateTime"`
	IsErrorCorrect        bool            `json:"IsErrorCorrect" xml:"IsErrorCorrect"`
	IsOption              bool            `json:"IsOption" xml:"IsOption"`
	IsSold                bool            `json:"IsSold" xml:"IsSold"`
	IsRefund              bool            `json:"IsRefund" xml:"IsRefund"`
	IsRefundReturnToStock bool            `json:"IsRefundReturnToStock" xml:"IsRefundReturnToStock"`
	DiscountInfos         []DiscountInfo  `json:"DiscountInfos" xml:"DiscountInfos>DiscountInfo"`
	PromotionInfos        []PromotionInfo `json:"PromotionInfos" xml:"PromotionInfos>PromotionInfo"`
}

type DiscountInfo struct {
	DiscountReasonID           int          `json:"DiscountReasonID" xml:"DiscountReasonID"`
	DiscountReasonDesc         string       `json:"DiscountReasonDesc" xml:"DiscountReasonDesc"`
	DiscountAmt                float64      `json:"DiscountAmt" xml:"DiscountAmt"`
	DiscountApprovedByEmployee bool         `json:"DiscountApprovedByEmployee" xml:"DiscountApprovedByEmployee"`
	DiscountApprovalDateTime   FlexibleTime `json:"DiscountApprovalDateTime" xml:"DiscountApprovalDateTime"`
}

type PromotionInfo struct {
	PromotionID           int     `json:"PromotionID" xml:"PromotionID"`
	PromotionName         string  `json:"PromotionName" xml:"PromotionName"`
	PromotionLossGrossAmt float64 `json:"PromotionLossGrossAmt" xml:"PromotionLossGrossAmt"`
	PromotionLossNetAmt   float64 `json:"PromotionLossNetAmt" xml:"PromotionLossNetAmt"`
}

type PaymentInfo struct {
	EPOSTransactionID   int          `json:"EPOSTransactionID" xml:"EPOSTransactionID"`
	TenderAmount        float64      `json:"TenderAmount" xml:"TenderAmount"`
	TenderTypeID        int          `json:"TenderTypeID" xml:"TenderTypeID"`
	TenderTypeName      string       `json:"TenderTypeName" xml:"TenderTypeName"`
	EmployeeID          int          `json:"EmployeeID" xml:"EmployeeID"`
	EmployeeName        string       `json:"EmployeeName" xml:"EmployeeName"`
	TenderedDateTime    FlexibleTime `json:"TenderedDateTime" xml:"TenderedDateTime"`
	TableNo             int          `json:"TableNo" xml:"TableNo"`
	TenderID            int          `json:"TenderID" xml:"TenderID"`
	ThirdPartyReference string       `json:"ThirdPartyReference" xml:"ThirdPartyReference"`
}
