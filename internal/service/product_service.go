package service

import (
	"context"
	"test-task/internal/models/elasticc"
	"test-task/internal/models/repostory"
	"test-task/internal/models/service"
	"test-task/internal/repostory/database"
	"test-task/internal/repostory/elastic"

	"github.com/google/uuid"
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

func (s *ProductService) BulkProductUpsert(ctx context.Context, companyID string, products []service.Product) error {

	uniqueProducts := make(map[string]service.Product)

	for _, product := range products {

		product.CompanyID = companyID

		if product.ID == "" {
			product.ID = uuid.New().String()
		}

		key := companyID + "_" + product.SKU

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

	return s.elasticRepo.BulkProductUpsert(
		ctx,
		documents,
	)
}

func (s *ProductService) SearchProducts(
	ctx context.Context,
	companyID string, q string,
) ([]elasticc.ProductDocument, error) {

	products, err := s.elasticRepo.SearchProducts(
		ctx,
		companyID,
		q,
	)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) DeleteProducts(
	ctx context.Context,
	companyID string, ids []string,
) error {

	err := s.productRepo.DeleteProducts(
		ctx,
		companyID,
		ids,
	)

	if err != nil {
		return err
	}

	err = s.elasticRepo.DeleteProducts(
		ctx,
		ids,
	)

	return err
}
