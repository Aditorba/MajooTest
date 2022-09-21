package dto

type MerchantOmzetDTO struct {
	MerchantId   int     `json:"merchantId" form:"merchantId"`
	MerchantName string  `json:"merchantName" form:"merchantName"`
	Omzet        float64 `json:"omzet" form:"omzet"`
}
