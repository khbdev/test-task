package database

import (
	"context"
	"test-task/internal/models/repostory"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) BulkProductUpsert(ctx context.Context, products []repostory.Product) error {

	query := `INSERT INTO products (
    id,
    company_id,
    name,
    sku,
    barcode,
    supply_price,
    retail_price
)
VALUES (
    :id,
    :company_id,
    :name,
    :sku,
    :barcode,
    :supply_price,
    :retail_price
)

ON CONFLICT (company_id, sku)
WHERE deleted_at IS NULL

DO UPDATE SET
name = EXCLUDED.name,
barcode = EXCLUDED.barcode,
supply_price = EXCLUDED.supply_price,
retail_price = EXCLUDED.retail_price;`

	_, err := r.db.NamedExecContext(
		ctx,
		query,
		products,
	)

	return err
}
