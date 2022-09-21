package domain

type DataReportTemplate struct {
	Id           int64
	MerchantId   int64
	MerchantName string
	OutletId     int64
	OutletName   string
	UserId       int64
	Name         string
	Omzet        float64
}
