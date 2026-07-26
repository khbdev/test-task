package service

import (
	"context"
	"test-task/internal/models/elasticc"
	"test-task/internal/models/repostory"
	"test-task/internal/models/service"
	"test-task/internal/repostory/database"
	"test-task/internal/repostory/elastic"

	_ "github.com/google/uuid"
)

type ProductService struct {
	productRepo *database.ProductRepository
	elasticRepo *elastic.ProductRepository
}

func NewProductService(productRepo *database.ProductRepository, elasticRepo *elastic.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		elasticRepo: elasticRepo,
	}
}

func (s *ProductService) BulkProductUpsert(ctx context.Context, products []service.Product) error {

	uniqueProducts := make(map[string]service.Product)

	for _, product := range products {

		key := product.CompanyID + "_" + product.SKU

		uniqueProducts[key] = product
	}

	cleanProducts := make([]service.Product, 0, len(uniqueProducts))

	for _, product := range uniqueProducts {
		cleanProducts = append(cleanProducts, product)
	}

	repoProducts := make([]repostory.Product, 0, len(cleanProducts))

	for _, product := range cleanProducts {

		repoProducts = append(repoProducts, repostory.Product{
			ID:          product.ID,
			CompanyID:   product.CompanyID,
			Name:        product.Name,
			SKU:         product.SKU,
			Barcode:     product.Barcode,
			SupplyPrice: product.SupplyPrice,
			RetailPrice: product.RetailPrice,
		})
	}

	err := s.productRepo.BulkProductUpsert(
		ctx,
		repoProducts,
	)

	if err != nil {
		return err
	}

	documents := make([]elasticc.ProductDocument, 0, len(cleanProducts))

	for _, product := range cleanProducts {

		documents = append(documents, elasticc.ProductDocument{
			ID:          product.ID,
			CompanyID:   product.CompanyID,
			Name:        product.Name,
			SKU:         product.SKU,
			Barcode:     product.Barcode,
			RetailPrice: product.RetailPrice,
			Slot:        "",
		})
	}

	err = s.elasticRepo.BulkProductUpsert(
		ctx,
		documents,
	)

	return err
}
