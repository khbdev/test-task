CREATE TABLE shelf_slots (
id UUID PRIMARY KEY,
 company_id UUID NOT NULL,
slot INT NOT NULL,
product_id UUID REFERENCES products(id) ON DELETE SET NULL,
 created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
