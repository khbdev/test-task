CREATE TABLE products (
id UUID PRIMARY KEY,
company_id UUID NOT NULL,
name VARCHAR(255) NOT NULL,
sku VARCHAR(100) NOT NULL,
barcode TEXT[],
supply_price NUMERIC(12, 2) NOT NULL,
retail_price NUMERIC(12, 2) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP NULL
);


CREATE UNIQUE INDEX idx_products_company_sku
    ON products(company_id, sku)
    WHERE deleted_at IS NULL;


CREATE INDEX idx_products_company_id
    ON products(company_id);