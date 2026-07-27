# Electronic Shelf Price Service

Go backend servis. Do'kondagi elektron narx ko'rsatkichlari uchun
yozilgan.

## Tech Stack

-   Go 1.25
-   Gin
-   PostgreSQL
-   sqlx
-   Elasticsearch 8
-   golang-migrate
-   Docker Compose

## Run Project

``` bash
docker compose up -d
```

API: localhost:8083\
PostgreSQL: localhost:5432\
Elasticsearch: localhost:9200

## Migration

Database schema golang-migrate orqali boshqariladi.

``` bash
migrate -path migrations -database "postgres://postgres:secret@localhost:5432/test_task?sslmode=disable" up
```

## Architecture

Handler -\> Service -\> Repository -\> PostgreSQL / Elasticsearch

Handler HTTP request va response bilan ishlaydi.

Service business logic uchun javob beradi.

Repository database va Elasticsearch bilan ishlaydi.

## Database Decisions

`shelf_slots` jadvalida product ma'lumotlari saqlanmaydi. Faqat
product_id saqlanadi.

Sababi product ma'lumotlari products jadvalida mavjud. Duplicate data
saqlamaslik va normalizationni saqlash uchun JOIN ishlatiladi.

Barcha querylarda company_id filter ishlatiladi.

Sababi bitta kompaniya boshqa kompaniyaning ma'lumotlarini ko'rmasligi
kerak.

## Product Upsert

Productlar `(company_id, sku)` bo'yicha unique.

Bulk request ichida bir xil SKU bir necha marta kelsa oxirgi qiymat
olinadi.

## Elasticsearch

Product search Elasticsearch orqali ishlaydi.

Sababi fuzzy search, typo bilan qidirish, SKU, barcode va slot bo'yicha
qidiruv uchun Elasticsearch ishlatiladi.

Slot o'zgarganda faqat `slot` field update qilinadi.

## Indexes

Products uchun:

``` sql
CREATE UNIQUE INDEX idx_products_company_sku
ON products(company_id, sku)
WHERE deleted_at IS NULL;
```

Shelf slots uchun:

``` sql
CREATE UNIQUE INDEX idx_shelf_slots_company_slot
ON shelf_slots(company_id, slot);
```

Indexlar tez qidirish va uniqueness uchun qo'yilgan.

## Docker Compose

PostgreSQL va Elasticsearch Docker Compose orqali ko'tariladi.

Sababi projectni bitta buyruq bilan bir xil environmentda ishga
tushirish uchun.

## API

Products: - POST /products/bulk - GET /products/search?q= - DELETE
/products

Slots: - PUT /slots - GET /slots

Reports: - GET /reports/stock-value