package elasticc

type ProductDocument struct {
	ID          string   `json:"id"`
	CompanyID   string   `json:"company_id"`
	Name        string   `json:"name"`
	SKU         string   `json:"sku"`
	Barcode     []string `json:"barcode"`
	SupplyPrice float64  `json:"supply_price"`
	RetailPrice float64  `json:"retail_price"`
	Slot        string   `json:"slot"`
}
