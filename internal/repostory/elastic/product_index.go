package elastic

import (
	"bytes"
	"context"
	"encoding/json"
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
			`{"update":{"_index":"products","_id":"` +
				product.ID +
				`"}}`,
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
