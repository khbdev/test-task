package service

type Product struct {
	ID          string
	CompanyID   string
	Name        string
	SKU         string
	Barcode     []string
	SupplyPrice float64
	RetailPrice float64
}
