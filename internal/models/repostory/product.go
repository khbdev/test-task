package repostory

type Product struct {
	ID          string   `db:"id"`
	CompanyID   string   `db:"company_id"`
	Name        string   `db:"name"`
	SKU         string   `db:"sku"`
	Barcode     []string `db:"barcode"`
	SupplyPrice float64  `db:"supply_price"`
	RetailPrice float64  `db:"retail_price"`
	CreatedAt   string   `db:"created_at"`
	DeletedAt   *string  `db:"deleted_at"`
}
