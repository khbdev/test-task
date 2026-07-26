package database

import (
	"context"
	"test-task/internal/models/repostory"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func (r *ProductRepository) DeleteProducts(ctx context.Context, companyID string, ids []string) error {

	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	// slotlarni bo'shatish

	_, err = tx.ExecContext(
		ctx,
		`
		UPDATE shelf_slots
		SET product_id = NULL
		WHERE company_id = $1
		AND product_id = ANY($2)
		`,
		companyID,
		pq.Array(ids),
	)

	if err != nil {
		return err
	}

	// soft delete

	_, err = tx.ExecContext(
		ctx,
		`
		UPDATE products
		SET deleted_at = NOW()
		WHERE company_id = $1
		AND id = ANY($2)
		`,
		companyID,
		pq.Array(ids),
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}
