package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"test-task/internal/models/elasticc"

	"github.com/elastic/go-elasticsearch/v8"
)

type ProductRepository struct {
	client *elasticsearch.Client
}

func NewProductRepository(client *elasticsearch.Client) *ProductRepository {
	return &ProductRepository{
		client: client,
	}
}

func (r *ProductRepository) BulkProductUpsert(ctx context.Context, products []elasticc.ProductDocument) error {

	var buf bytes.Buffer

	for _, product := range products {

		meta := []byte(
			fmt.Sprintf(
				`{"update":{"_index":"products","_id":"%s"}}`,
				product.ID,
			),
		)

		buf.Write(meta)
		buf.WriteByte('\n')

		doc := map[string]interface{}{
			"doc":           product,
			"doc_as_upsert": true,
		}

		data, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		buf.Write(data)
		buf.WriteByte('\n')
	}

	_, err := r.client.Bulk(
		bytes.NewReader(buf.Bytes()),
		r.client.Bulk.WithContext(ctx),
	)

	return err
}

func (r *ProductRepository) SearchProducts(ctx context.Context, companyID string, q string) ([]elasticc.ProductDocument, error) {

	query := fmt.Sprintf(`
	{
		"query": {
			"bool": {
				"must": {
					"multi_match": {
						"query": "%s",
						"fields": [
							"name",
							"sku",
							"barcode",
							"slot"
						],
						"fuzziness": "AUTO"
					}
				},
				"filter": {
					"term": {
						"company_id.keyword": "%s"
					}
				}
			}
		}
	}`, q, companyID)

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("products"),
		r.client.Search.WithBody(
			bytes.NewReader([]byte(query)),
		),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data struct {
		Hits struct {
			Hits []struct {
				Product elasticc.ProductDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	json.NewDecoder(res.Body).Decode(&data)

	products := make([]elasticc.ProductDocument, 0)

	for _, item := range data.Hits.Hits {
		products = append(products, item.Product)
	}

	return products, nil
}
