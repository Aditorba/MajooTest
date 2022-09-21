package dto

type OutletOmzetDTO struct {
	MerchantId   int     `json:"merchantId" form:"merchantId"`
	OutletId     int     `json:"outletId" form:"outletId"`
	MerchantName string  `json:"merchantName" form:"merchantName"`
	OutletName   string  `json:"outletName" form:"outletName"`
	Omzet        float64 `json:"omzet" form:"omzet"`
}
