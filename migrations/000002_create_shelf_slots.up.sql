CREATE TABLE shelf_slots (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             company_id UUID NOT NULL,
                             slot INT NOT NULL,
                             product_id UUID REFERENCES products(id) ON DELETE SET NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE UNIQUE INDEX idx_shelf_slots_company_slot
    ON shelf_slots(company_id, slot);


CREATE INDEX idx_shelf_slots_company_id
    ON shelf_slots(company_id);