package repostory

import "github.com/lib/pq"

type ShelfSlot struct {
	ID        string  `db:"id"`
	CompanyID string  `db:"company_id"`
	Slot      int     `db:"slot"`
	ProductID *string `db:"product_id"`
}

type ShelfSlotGET struct {
	ID        string  `db:"id" json:"id"`
	Slot      int     `db:"slot" json:"slot"`
	ProductID *string `db:"product_id" json:"product_id"`

	Name    *string        `db:"name" json:"name"`
	SKU     *string        `db:"sku" json:"sku"`
	Barcode pq.StringArray `db:"barcode" json:"barcode"`

	SupplyPrice *float64 `db:"supply_price" json:"supply_price"`
	RetailPrice *float64 `db:"retail_price" json:"retail_price"`
}
type StockValueReport struct {
	SupplyTotal float64 `db:"supply_total" json:"supply_total"`
	Occupied    int     `db:"occupied" json:"occupied"`
	Empty       int     `db:"empty" json:"empty"`
}
