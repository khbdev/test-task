package database

import (
	"context"

	"test-task/internal/models/repostory"

	"github.com/jmoiron/sqlx"
)

type ShelfSlotRepository struct {
	db *sqlx.DB
}

func NewShelfSlotRepository(db *sqlx.DB) *ShelfSlotRepository {
	return &ShelfSlotRepository{
		db: db,
	}
}

func (r *ShelfSlotRepository) UpdateSlots(ctx context.Context, companyID string, slots []repostory.ShelfSlot) error {

	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO shelf_slots (
			company_id,
			slot,
			product_id
		)
		VALUES (
			$1,
			$2,
			$3
		)
		ON CONFLICT (company_id, slot)
		DO UPDATE SET
			product_id = EXCLUDED.product_id;
	`

	for _, slot := range slots {

		_, err = tx.ExecContext(
			ctx,
			query,
			companyID,
			slot.Slot,
			slot.ProductID,
		)

		if err != nil {
			return err
		}

	}

	return tx.Commit()
}

func (r *ShelfSlotRepository) GetSlots(ctx context.Context, companyID string, page int, limit int, search string) ([]repostory.ShelfSlotGET, int, error) {

	offset := (page - 1) * limit

	var total int

	err := r.db.GetContext(ctx, &total, `
		SELECT COUNT(*)
		FROM shelf_slots ss
		LEFT JOIN products p ON p.id = ss.product_id
		WHERE ss.company_id=$1
	`, companyID)

	if err != nil {
		return nil, 0, err
	}

	var slots []repostory.ShelfSlotGET

	err = r.db.SelectContext(ctx, &slots, `
		SELECT
			ss.slot,
			p.name,
			p.sku,
			p.retail_price
		FROM shelf_slots ss
		LEFT JOIN products p
			ON p.id = ss.product_id
			AND p.deleted_at IS NULL
		WHERE ss.company_id=$1
		AND (
			$2 = ''
			OR p.name ILIKE '%'||$2||'%'
			OR p.sku ILIKE '%'||$2||'%'
			OR ss.slot::text ILIKE '%'||$2||'%'
		)
		ORDER BY ss.slot
		LIMIT $3 OFFSET $4
	`,
		companyID,
		search,
		limit,
		offset,
	)

	return slots, total, err
}
